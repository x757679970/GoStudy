package main

import (
	"fmt"
	"time"
)

func worker(done chan string) {
	fmt.Println("working...")
	time.Sleep(time.Second * 2)
	fmt.Println("work done")
	time.Sleep(time.Second)
	done <- "ok"       //通過這個給值，導致下面的fmt.Println卡死等待
	fmt.Println("die") //此處的die和主線程中的ok順序是不穩定的
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	done := make(chan string, 1)
	go worker(done)
	fmt.Println(<-done) //這裡會無限卡死等待done管道中有數據引發同步

	fmt.Println("======================")

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs) //這裡的數據兩次被管道傳輸
}
