package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Timer 1 創建")
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C //執行等待，此處會因通道訊息等帶導致卡死
	fmt.Println("Timer 1 過期")

	fmt.Println("Timer 2 創建")
	timer2 := time.NewTimer(time.Second * 2)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 過期") //這裡永遠都不會被調用，因為主線程已經先退出了，協程被殺
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 停止")
	}
}
