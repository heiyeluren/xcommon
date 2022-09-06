package common

import (
	"sync/atomic"
	"time"
)

var timeNow int64

func init() {
	Schedule(loadTime, time.Duration(1*1e9)) //每秒更新timeNow
}

func Schedule(f func(), delay time.Duration) chan<- bool {
	stop := make(chan bool)
	go func() {
	OuterLoop:
		for {
			select {
			case <-time.After(delay):
				f()
			case <-stop:
				break OuterLoop
			}
		}
	}()

	return stop
}

func loadTime() {
	for {
		oldTime := atomic.LoadInt64(&timeNow)
		if atomic.CompareAndSwapInt64(&timeNow, oldTime, time.Now().Unix()) {
			break
		}
	}
}

//获取本地时间戳，误差最多1秒
func GetTimeNow() int64 {
	return atomic.LoadInt64(&timeNow)
}
