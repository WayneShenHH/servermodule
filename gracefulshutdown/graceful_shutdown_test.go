package gracefulshutdown

import (
	"context"
	"log"
	"os"
	"sync"
	"testing"
	"time"
)

// 請注意 nowSentTime 的時間，因為 fmt 印出來的 log 沒有線程安全，不保證順序。
func Test_Gracefulshutdown(t *testing.T) {
	signalChan := Start()
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			ctx, ch := GetContext(uint8(i))
			go func(c context.Context, cha chan struct{}, ii int) {
				for {
					<-c.Done()
					cha <- struct{}{}
					nowSentTime := time.Now().Nanosecond()
					log.Printf("Gracefulshutdown Level: %d Send at %d", ii, nowSentTime)
					return
				}
			}(ctx, ch, i)
		}
	}
	signalChan <- os.Interrupt
	select {}
}

func Test_ContextDone(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wgg *sync.WaitGroup, c context.Context, ii int) {
			log.Printf("go I: %d", ii)
			<-c.Done()
			log.Printf("go Done I: %d", ii)
			wgg.Done()
		}(&wg, ctx, i)
	}
	time.Sleep(4 * time.Second)
	cancel()
	log.Printf("cancel")
	wg.Wait()
	log.Printf("Wait")
	select {}
}
