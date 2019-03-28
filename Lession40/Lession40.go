package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500) //500毫秒一次觸發
	go func() {
		for t := range ticker.C { //循環處發
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 1600) //1600毫秒後，關閉觸發器
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
