package main

import (
	"fmt"
	"time"
)

func main() {
	//我們提交一些工作請求
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	//limiter通道每200ms會接受一次值
	limiter := time.Tick(time.Millisecond * 200)

	//我們通過從limiter通道中取值，來限制200毫秒才能進行一次消息發送
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	fmt.Println("=======================")

	//我們許可有三個突發事件通道隊列
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	//每200毫秒我們添加一個突發事件到通道隊列中
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	//我們當前模擬5個請求。前三個會在我們的緊急突發通道列表中
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests) //關閉通道
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
