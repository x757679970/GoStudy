package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signl.Notify(sigs, syscall.SIGINT.syscall.SIGTERM)
	//啟動協程，等待系統信號
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()
	fmt.Println("awaiting singnal")
	<-done
	fmt.Println("exiting")

}
