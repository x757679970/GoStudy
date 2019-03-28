package main

import (
	"fmt"
	"os"
)

func main() {
	//這個defer將永遠無法被執行到
	defer fmt.Println("!")

	//退出，返回值3。返回值只有0才認為是正常退出
	os.Exit(0)
}
