package main

import (
	"fmt"
	"runtime"
	"time"
)

var quit chan int

func foo(id int) {
	fmt.Println(id)
	time.Sleep(time.Second)//停頓一秒
	quit <- 0				//發消息；執行完畢
}

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func main() {
	fmt.Println("當前CPU有的核心數：", runtime.NumCPU())
	runtime.GOMAXPROCS(4) //設置GO可用的核心數
	begin := time.Now()
	count := 1000
	quit = make(chan int, count)	//開1000個goroutine

	for i := 0; i < count; i++ {
		go foo(i)
	}

	for i := 0; i < count; i++ {	//等待所有完成消息發送完畢
		<-quit
	}
	end := time.Now()
	fmt.Println("累積消耗時間：",
		end.Second()-begin.Second(), "秒",
		end.Nanosecond()-begin.Nanosecond(), "納秒")

	//如果我們runtime.GOMAXPROCS(1)設置為1，則不會有任何輸出
	//因為其實當前GO協程式再進行單CPU的切片模擬多線程而已
	//如果我們runtime.GOMAXPROCS(4)，則可以正常輸出5次world
	//因為此時我們確實多CPU並行
	
	//我們要明確要求CPU並行，則需要
	// 1:可以runtime.GOMAXPROCS(x)指定多CPU
	// 2:可以runtime.GOMAXPROCS調適用表示讓出CPU切片
	go say("world")	//開一個新的Goroutines執行
	for {
	}
}
