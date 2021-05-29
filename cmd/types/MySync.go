package types

import (
	"sync/atomic"
	"time"
)

type MySync struct {
	innerCh chan bool
	count   int32
}

func NewMySync() *MySync {
	//给定一个稍微大一点的容量，避免阻塞处理线程
	return &MySync{innerCh: make(chan bool, 1024), count: 0}
}

func (my *MySync) Add(num int) {
	for i := 0; i < num; i++ {
		atomic.AddInt32(&my.count, 1) // 先+
		my.innerCh <- true
	}
}

func (my *MySync) Done(num int) {
	for i := 0; i < num; i++ {
		<-my.innerCh
		atomic.AddInt32(&my.count, -1) // 后-
	}
}

func (my *MySync) Wait() {
	ch := time.Tick(time.Millisecond * 1)
	for {
		<-ch
		if my.count <= 0 {
			return
		}
	}

}
