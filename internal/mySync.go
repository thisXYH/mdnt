package internal

import (
	"sync/atomic"
	"time"
)

// MySync 利用channal的缓冲区实现一个限流器。
// 用来限制 gorotine 并发数量，并且等待所有 gorotine 结束
type MySync struct {
	innerCh chan bool
	working int32
}

func NewMySync(maxConcurrent int) *MySync {
	return &MySync{innerCh: make(chan bool, maxConcurrent), working: 0}
}

func (my *MySync) Add(num int) {
	for i := 0; i < num; i++ {
		atomic.AddInt32(&my.working, 1) // 先+
		my.innerCh <- true
	}
}

func (my *MySync) Done(num int) {
	for i := 0; i < num; i++ {
		<-my.innerCh
		atomic.AddInt32(&my.working, -1) // 后-
	}
}

// Wait 通过定时 loop 是否结束
func (my *MySync) Wait() {
	ch := time.Tick(time.Millisecond * 1)
	for {
		<-ch
		if my.working <= 0 {
			return
		}
	}
}
