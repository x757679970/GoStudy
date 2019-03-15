package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var x, y, z int
var check bool

var (
	index   int
	getdata string
)

//常量
const sConst string = "正在等待輸入..."

//額外說明，在go裡面function
//一定要是 funtionName{
//}
//不可以用functionName
//{
//}
//會編譯錯誤

func main() {
	fmt.Println("hello word")
	fmt.Println(x, y, z, index, getdata, check)
	var a, b int = 1, 2
	fmt.Println("1+2=", a+b)
	fmt.Println("7/3=", 7/3)
	//在go裡面，是float一定要用.0或.x顯示，不然會視為整數
	fmt.Println("7.0/3.0=", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)

	//等待輸入
	var inputString string
	fmt.Scanln(&inputString)
	fmt.Println("=====================================")

	//重製輸入緩衝區
	scanner := bufio.NewScanner(os.Stdin)

	//處理三次
	nTimes := 0
	for scanner.Scan() {
		//將輸入數據轉為大寫
		ucl := strings.ToUpper(scanner.Text())
		//輸出
		fmt.Println(ucl)
		nTimes++
		if nTimes >= 3 {
			break
		}

	}

	//如果有錯誤，則輸出錯誤
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

}
