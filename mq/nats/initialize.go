package nats

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/WayneShenHH/servermodule/config"
	"github.com/WayneShenHH/servermodule/flow"
	"github.com/WayneShenHH/servermodule/gracefulshutdown"
	"github.com/WayneShenHH/servermodule/logger"
	"github.com/WayneShenHH/servermodule/mq"
	"github.com/WayneShenHH/servermodule/protocol"
	"github.com/WayneShenHH/servermodule/protocol/constant"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

const (
	status_init    = 0
	status_running = 1
)

type handler struct {
	ctx             context.Context
	config          *config.NatsConfig
	connection      stan.Conn
	status          int                     // 標記此模組 運作中，關閉中，結束
	shutdownChan    chan struct{}           // graceful shutdown 信號
	subscriptionMap subscriptionMap         // map[RoomGuid]stan.Subscription 記住所有註冊
	subscriberMap   map[string]func() error // map[RoomGuid]func() error 記住所有註冊方法
	publishBackup   sync.Map
}

func New(cfg *config.NatsConfig) mq.Handler {
	ctx, shutdownChan := gracefulshutdown.GetContext(constant.Nats_Level)
	hdr := &handler{
		ctx:           ctx,
		config:        cfg,
		status:        status_init,
		shutdownChan:  shutdownChan,
		subscriberMap: make(map[string]func() error),
	}

	err := hdr.connect()
	if err != nil {
		return nil
	}

	go hdr.shutdownListener()

	return hdr
}

// 當所有模組都啟動後才能執行，避免太早收到訊號但無法處理
func (hdr *handler) Start() {
	hdr.status = status_running

	for _, subscribe := range hdr.subscriberMap {
		err := subscribe()
		if err != nil {
			logger.Errorf("nats/Start/subscribe failed, %v", err)
		}
	}
}

// connect 嘗試連線，成功會開啟 publish() 生命週期
func (hdr *handler) connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), hdr.config.ConnectTimeOut)
	defer cancel()
	ticker := time.NewTicker(hdr.config.ReconnInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			logger.Fatalf("nats/connect failed: %v, timeout: %v", ctx.Err(), hdr.config.ConnectTimeOut)
			hdr.shutdownChan <- struct{}{}
			return fmt.Errorf("nats/connect failed: %v, timeout: %v", ctx.Err(), hdr.config.ConnectTimeOut)
		case <-ticker.C:
			// 建立 Nats Streaming 連線
			sc, err := stan.Connect(
				hdr.config.ClusterID,
				fmt.Sprintf("GS-%v-%v", time.Now().UnixNano(), rand.Intn(10000)),
				stan.NatsURL(hdr.config.Addr),
				stan.NatsOptions(nats.UserInfo(hdr.config.Username, hdr.config.Password)),
				stan.Pings(hdr.config.StanPingsInterval, hdr.config.StanPingsMaxOut),
				stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
					logger.Errorf("nats/connect Connection lost, reason: %v", reason.Error())
					hdr.reconnect()
				}),
			)

			if err == nil {
				logger.Infof("nats/connect to %v success", hdr.config.Addr)
				hdr.connection = sc

				hdr.publishBackup.Range(func(key, value interface{}) bool {
					hdr.Publish(key.(string), value.(*flow.DataFlow))
					err := hdr.connection.Publish(key.(string), value.([]byte))
					if err == nil {
						hdr.publishBackup.Delete(key)
					}
					return true
				})

				return nil
			}

			logger.Warnf("nats/connect: %v, retry in %v", err, hdr.config.ReconnInterval)
		}
	}
}

func (hdr *handler) reconnect() {
	hdr.disconnect()
	err := hdr.connect()
	if err != nil {
		return
	}
	for _, subscribe := range hdr.subscriberMap {
		err := subscribe()
		if err != nil {
			logger.Errorf("nats/reconnect/subscribe failed, %v", err)
		}
	}
}

// disconnect 使 publish() 能夠正常結束，取消訂閱，關閉連線
func (hdr *handler) disconnect() {
	hdr.unsubscribe()

	err := hdr.connection.Close()
	if err != nil {
		logger.Errorf("nats/disconnect error: %v", err)
		return
	}
	logger.Infof("nats/disconnect closed")
}

// unsubscribe 使所有已經訂閱的房間取消訂閱主題
func (hdr *handler) unsubscribe() {
	hdr.subscriptionMap.Range(func(key string, sub stan.Subscription) bool {
		err := sub.Unsubscribe()
		if err != nil {
			logger.Errorf("nats/unsubscribe error: %v", err)
		}
		hdr.subscriptionMap.Delete(key)
		return true
	})

	logger.Infof("nats/unsubscribe ok")
}

func (hdr *handler) unsubscribeByKey(key string) {
	subs, exist := hdr.subscriptionMap.Load(key)
	if exist {
		err := subs.Unsubscribe()
		if err != nil {
			logger.Errorf("nats/unsubscribe %v error: %v", key, err)
			return
		}
	}

	delete(hdr.subscriberMap, key)
	hdr.subscriptionMap.Delete(key)

	logger.Infof("nats/unsubscribe %v ok", key)
}

func (hdr *handler) shutdownListener() {
	<-hdr.ctx.Done()
	// 先將所有等待處理的訊息處理完再做關閉
	// 只處理現有剩餘訊息，保持 hdr.writeFlow 開著
	hdr.disconnect()
	hdr.shutdownChan <- struct{}{}
	logger.Warn("nats shutdown")
}

// 註冊要訂閱的主題，收到訊息後執行預先定義好要處理的事項
// @param key 要訂閱的主題
// @param callback 預先定義好要處理的事項，收到訊息後傳入(raw data, []byte)並執行
// @return error
func (hdr *handler) register(key string, callback func(data []byte, sequenceId protocol.SequenceId)) error {
	logger.Infof("nats/register key: [%v]", key)

	hdr.subscriberMap[key] = func() error {
		sub, err := hdr.connection.Subscribe(
			key,
			func(msg *stan.Msg) {
				callback(msg.Data, msg.Sequence)
				err := msg.Ack()
				if err != nil {
					logger.Errorf("nats/register msg.Ack error: %v", err)
				}
			},
			stan.SetManualAckMode(),
			stan.DurableName(key),
			stan.MaxInflight(1), // MaxInflight 固定1，讓每則訊息都能藉由nats resend
		)

		if err != nil {
			logger.Errorf("nats/register error: %v", err)
			return err
		}

		hdr.subscriptionMap.Store(key, sub)
		return nil
	}

	switch hdr.status {
	case status_init:
		// 先將訂閱訊息存放在 hdr.subscriberCallbackMap
		// 由 hdr.Start() 訂閱，避免太早接收訊息但是訂閱者尚未準備好
	case status_running:
		subscribe := hdr.subscriberMap[key]
		err := subscribe()
		if err != nil {
			return err
		}
	}

	return nil
}

// 註冊要訂閱的主題，收到訊息後執行預先定義好要處理的事項
// @param key 要訂閱的主題
// @param callback 預先定義好要處理的事項，收到訊息後傳入(raw data, []byte)並執行
// @return error
func (hdr *handler) registerBySequence(key string, sequenceId protocol.SequenceId, callback func(data []byte, sequenceId protocol.SequenceId)) error {
	logger.Infof("nats/register key: [%v]", key)

	hdr.subscriberMap[key] = func() error {
		sub, err := hdr.connection.Subscribe(
			key,
			func(msg *stan.Msg) {
				// 每次接收到訊息的Sequence都是當下訊號的序列號，則下一個序列號則是需要++後的號碼
				callback(msg.Data, msg.Sequence)
				err := msg.Ack()
				if err != nil {
					logger.Errorf("nats/registerBySequence msg.Ack error: %v", err)
				}
			},
			// 需注意每次設定都需以上一次序列號的下一碼為主，不然會多發送一次上次得最後一個訊息(Sequence會從1開始)
			// 當序列號設定0或1結果會是相同的，如果設定的序列號是比上次還多(第一次也依樣)，並不會影響序列號本身存在NATS的順序
			// 範例：
			// Data: [Seq:1 Msg:A, Seq:2 Msg:B, Seq:3 Msg:C]
			// 當我StartAtSequence設定1時，會把以上所以訊息都會重新接收到一遍Msg: [A, B, C](此時會照Seq順序發送)
			// 當我StartAtSequence設定3時，只會跑Seq為3的訊息 Msg:[C]
			// 當我StartAtSequence設定4時，則不會接收到任何訊息 Msg:[C]
			stan.StartAtSequence(sequenceId),
			stan.SetManualAckMode(),
			stan.DurableName(key),
			stan.MaxInflight(1), // MaxInflight 固定1，讓每則訊息都能藉由nats resend
		)

		if err != nil {
			logger.Errorf("nats/registerBySequence error: %v", err)
			return err
		}

		hdr.subscriptionMap.Store(key, sub)
		return nil
	}

	switch hdr.status {
	case status_init:
		// 先將訂閱訊息存放在 hdr.subscriberCallbackMap
		// 由 hdr.Start() 訂閱，避免太早接收訊息但是訂閱者尚未準備好
	case status_running:
		subscribe := hdr.subscriberMap[key]
		err := subscribe()
		if err != nil {
			return err
		}
	}

	return nil
}

// registerQueueGroupByCallback 註冊要訂閱的主題，收到訊息後執行預先定義好要處理的事項
// @param key 要訂閱的主題
// @param qgroup 要訂閱的主題 Group
// @param callback 預先定義好要處理的事項，收到訊息後傳入(raw data, []byte)並執行
// @return error
func (hdr *handler) registerQueueGroup(key, qgroup string, callback func(data []byte)) error {
	logger.Infof("nats/registerQueueGroup key: [%v]", key)

	hdr.subscriberMap[key] = func() error {
		sub, err := hdr.connection.QueueSubscribe(
			key,
			qgroup,
			func(msg *stan.Msg) {
				callback(msg.Data)
				err := msg.Ack()
				if err != nil {
					logger.Errorf("nats/registerQueueGroup msg.Ack error: %v", err)
				}
			},

			stan.SetManualAckMode(),
			stan.DurableName(key),
			stan.MaxInflight(1), // MaxInflight 固定1，讓每則訊息都能藉由nats resend
		)

		if err != nil {
			logger.Errorf("nats/registerQueueGroup error: %v", err)
			return err
		}

		hdr.subscriptionMap.Store(key, sub)
		return nil
	}

	switch hdr.status {
	case status_init:
		// 先將訂閱訊息存放在 hdr.subscriberCallbackMap
		// 由 hdr.Start() 訂閱，避免太早接收訊息但是訂閱者尚未準備好
	case status_running:
		subscribe := hdr.subscriberMap[key]
		err := subscribe()
		if err != nil {
			return err
		}
	}

	return nil
}
