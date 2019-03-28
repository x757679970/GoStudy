package main

import (
	"fmt"
	"time"
)

//c和quit都是int類型的channel
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		//select和switch非常相似，但是select強制要求其case必須是I/O操作
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	//創建一個協程，讓其去輸出c channel 的數據10次
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		//輸出結束，給quit發個信號
		quit <- 0
	}()
	fibonacci(c, quit)

	fmt.Println("====================")
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	//注意for循環的次數用來檢察接收次數
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
