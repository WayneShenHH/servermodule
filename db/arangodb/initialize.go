package arangodb

import (
	"context"
	"crypto/tls"
	"errors"
	"net"
	defaulthttp "net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/WayneShenHH/servermodule/config"
	"github.com/WayneShenHH/servermodule/db"
	"github.com/WayneShenHH/servermodule/gracefulshutdown"
	"github.com/WayneShenHH/servermodule/logger"
	"github.com/WayneShenHH/servermodule/protocol/constant"
	"github.com/WayneShenHH/servermodule/util"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"go.uber.org/atomic"
	"golang.org/x/net/http2"
)

var (
	Error_Arangodb_Connection       = errors.New("arangodb connection error")
	Error_Arangodb_NewClient        = errors.New("arangodb newClient error")
	Error_Arangodb_Connect_Database = errors.New("arangodb connect database error")
	Error_Arangodb_Ping             = errors.New("arangodb ping error")
)

var handlerInstances []*arangoHdr // Client pool
var handlerIdx int                // 本次取用的 arangodb client index
var once sync.Once
var initialized bool
var mu sync.Mutex

type arangoHdr struct {
	db    driver.Database
	ctx   context.Context
	retry *util.Retry
}

// NewArango 建立 Arango 連線
func NewArango(cfg *config.ArangoDBConfig) db.NoSQL {
	initialize(cfg)
	return GetArango()
}

func GetArango() db.NoSQL {
	if !initialized {
		logger.Fatalf("GetArango instance is nil, have to initialize arango")
		return nil
	}

	mu.Lock()
	defer mu.Unlock()

	handlerIdx++
	if handlerIdx >= len(handlerInstances) {
		handlerIdx = 0
	}

	return handlerInstances[handlerIdx]
}

func retryLoop(ctx context.Context, shutdownChan, reTry chan struct{}, reTryCount *atomic.Uint32, handler *arangoHdr, config *config.ArangoDBConfig) {
	for {
		select {
		case <-ctx.Done():
			// 斷掉 redis。
			shutdownChan <- struct{}{}
			return
		case <-reTry:
		ReTryFlag:
			if reTryCount.Inc() > uint32(config.RetryCount) {
				logger.Fatalf("arangodb max retry count: %d", config.RetryCount)
				// 等待 shutdown 指令過來，然後等待關機。
				<-ctx.Done()
				shutdownChan <- struct{}{}
				return
			}
			aClient, err := connect(ctx, config, reTry, shutdownChan)
			// 連線完成。
			if err == nil {
				reTryCount.Store(0)
				handler.db = aClient
			} else {
				logger.Warnf("arangodb retry warn: %v", err.Error())
				goto ReTryFlag
			}
		}
	}
}

// 初始化 ArangoDB client instance
// 參數說明:
// HttpProtocol(1.1/2)
//  - 1.1: 初始化 1 個 instance, 並建立 ConnLimit 個連線
//  - 2:   初始化 ConnLimit 個 instance pool, 每個 instance 建立 1 個連線
//
// @param config ArangoDB 參數設定檔
func initialize(config *config.ArangoDBConfig) {
	once.Do(func() {
		ctx, shutdownChan := gracefulshutdown.GetContext(constant.Arangodb_Level)

		var hs []*arangoHdr
		var reTry []chan struct{}
		reTryCount := atomic.NewUint32(0)

		// 根據 HTTP protocol 初始化 client pool
		switch config.HttpProtocol {
		case "1.1":
			hs = make([]*arangoHdr, 1)
			reTry = make([]chan struct{}, 1)
			reTry[0] = make(chan struct{}, 1)

		case "2":
			hs = make([]*arangoHdr, config.Connlimit)
			reTry = make([]chan struct{}, config.Connlimit)
			for i := 0; i < config.Connlimit; i++ {
				reTry[i] = make(chan struct{}, 1)
			}
		}

		// 初始化 retry工具
		retry := util.NewRetry(config.RetryCount, config.RetryInterval)
		for i := 0; i < len(hs); i++ {
			hs[i] = &arangoHdr{db: nil, ctx: ctx, retry: retry}
		}

		for i := 0; i < len(hs); i++ {
			arangoClient, err := connect(ctx, config, reTry[i], shutdownChan)
			if err != nil {
				logger.Warnf("arangodb retry warn: %v", err.Error())
				i -= 1
				continue
			}

			hs[i].db = arangoClient
			go retryLoop(ctx, shutdownChan, reTry[i], reTryCount, hs[i], config)
		}

		logger.Info("arangodb Initialize Done")
		handlerInstances = hs

		initialized = true
	})
}

func connect(ctx context.Context, config *config.ArangoDBConfig, reTry, shutdownChan chan struct{}) (driver.Database, error) {
	dbIPs := strings.Split(config.Addr, ",")
	for _, u := range dbIPs {
		_, err := url.Parse("http://" + u)
		if err != nil {
			logger.Fatalf("arangodb IP url.Parse error: %v", err)
			// 等待 shutdown 指令過來，然後等待關機。
			<-ctx.Done()
			shutdownChan <- struct{}{}
			return nil, err
		}
	}

	var conn driver.Connection
	var err error

	// 根據 HTTP protocol 決定如何初始化 transport
	switch config.HttpProtocol {
	case "1.1":
		transport := &defaulthttp.Transport{
			DialContext: (&net.Dialer{
				KeepAlive: 60 * time.Second,
				DualStack: true}).DialContext,
			MaxIdleConns:          0,
			IdleConnTimeout:       30 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		}
		conn, err = http.NewConnection(http.ConnectionConfig{
			Endpoints: dbIPs,
			ConnLimit: config.Connlimit,
			Transport: transport,
		})

	case "2":
		transport := &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		}
		conn, err = http.NewConnection(http.ConnectionConfig{
			Endpoints: dbIPs,
			ConnLimit: config.Connlimit,
			Transport: transport,
		})
	}

	if err != nil {
		logger.Fatalf("arangodb connection error: %v", err)
		<-ctx.Done()
		shutdownChan <- struct{}{}
		return nil, Error_Arangodb_Connection
	}

	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(config.Username, config.Password),
	})
	if err != nil {
		logger.Errorf("arangodb NewClient error: %v", err)
		reTry <- struct{}{}
		return nil, Error_Arangodb_NewClient
	}
	if _, err := c.Version(ctx); err != nil {
		logger.Warnf("arangodb Ping filed, err: %v", err.Error())
		reTry <- struct{}{}
		return nil, Error_Arangodb_Ping
	}

	db, err := c.Database(ctx, config.Database)
	if err != nil {
		logger.Errorf("arangodb connect database error: %v", err)
		reTry <- struct{}{}
		return nil, Error_Arangodb_Connect_Database
	}
	go pingLoop(ctx, c, reTry, shutdownChan)
	return db, nil
}

func pingLoop(ctx context.Context, c driver.Client, reTry, shutdownChan chan struct{}) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if _, err := c.Version(ctx); err != nil {
				logger.Warnf("arangodb Ping filed, err: %v", err.Error())
				reTry <- struct{}{}
				return
			}
		case <-ctx.Done():
			// 等待 shutdown 指令過來，然後等待關機。
			shutdownChan <- struct{}{}
			return
		}
	}
}
