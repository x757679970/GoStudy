package main

import (
	"fmt"
	"time"
)

func main() {
	//創建一個channel
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(2 * 1e9)
		timeout <- true
	}()
	ch := make(chan int)
	select {
	case <-ch:
		fmt.Println("ch pop one element")
	case <-timeout:
		fmt.Println("timeout!")
	}

	fmt.Println("------------------------------")
	select {
	case ch <- 2:
		//當ch滿了，此時向ch中插入2會失敗，則會執行dewfault
		//此時會實現對channel是否已滿的檢查
		fmt.Println("ccc")
	default:
		fmt.Println("channel is full !")
	}
}
