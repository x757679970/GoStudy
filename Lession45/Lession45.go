package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	var ops int64 = 0 //執行的操作數，原子計數

	//讀寫管道
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	//該協程負責管理狀態，也就是所謂的"狀態協程"
	//它會不停的嘗試從通道或許請求並處理，做出響應
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	//這裡啟動了100個協程
	//它們每次啟動讀取都需要向狀態協程提交請求
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	//創建10個寫入協程
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	//主線休眠1秒，等待協程自己處理
	time.Sleep(time.Second)

	//輸出執行次數
	opsFainl := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFainl)

}
