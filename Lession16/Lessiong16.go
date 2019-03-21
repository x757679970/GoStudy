package main

import (
	"fmt"
	"strings"
)

func main() {
	//創建數組，聲明長度，無初始化
	var array1 = [5]string{}
	PrintStrArray(array1[:])

	//創建數組，聲明長度，局部順序初始化
	var array2 = [5]string{"Fri", "Thu"}
	PrintStrArray(array2[:])

	//創建數組，聲明長度，指定元素初始化
	var array3 = [5]string{1: "chris", 4: "Kris"}
	PrintStrArray(array3[:])

	//創建數組，不聲明長度，元素初始化
	var array4 = [...]string{"chris", "Kris"}
	PrintStrArray(array4[:])

	//創建數組，不聲明長度，指定元素初始化
	var array5 = [...]string{1: "chris", 7: "Kris"}
	PrintStrArray(array5[:])

	//創建數組切片
	var array6 = []string{"lala", "coco"}
	PrintStrArray(array6)

	//聲明新數組切片
	var array7 = [...]string{1: "chris", 4: "python", 7: "jquery"}
	var slicel = array7[2:5]
	PrintStrArray(slicel)

	var slicel2 = array7[2 : len(array7)-1]
	PrintStrArray(slicel2)

	//證明切片是值，而非只針
	AddSlice(slicel2)
	PrintStrArray(slicel2)
	PrintStrArray(array7[:])

	fmt.Println("============================")

	//注意，切片是地址，值。連copy都不是
	names := [4]string{
		"John",
		"Pau",
		"George",
		"Ring",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	fmt.Println(names)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
	///////////////////
	//注意切片的默認值
	s := []int{2, 3, 5, 11, 13}
	fmt.Println(s)

	s1 := s[:2]
	fmt.Println(s1)
	s2 := s[1:]
	fmt.Println(s2)
	s3 := s[:]
	fmt.Println(s3)
	fmt.Println("==========================")

	//空切片檢查
	var emptySlice []int
	fmt.Println(emptySlice, len(emptySlice), cap(emptySlice))
	if emptySlice == nil {
		fmt.Println("nil!")
	}
	fmt.Println("==========================")

	//make函數聲明動態切片，append添加切片
	//make只能初始化創建標準類型map,slice,channel，返回的是對象的飲用
	//new和常規認知一樣，返回的是字定義對象的指針
	aSlice := make([]int, 5)
	PrintSliceInfoWithStr("a", aSlice)

	bSlice := make([]int, 0, 5)
	PrintSliceInfoWithStr("b", bSlice)

	cSlice := bSlice[:2]
	PrintSliceInfoWithStr("c", cSlice)

	dSlice := bSlice[2:5]
	PrintSliceInfoWithStr("d", dSlice)

	eSlice := append(dSlice, 2, 3, 4)
	PrintSliceInfoWithStr("e", eSlice)
	fmt.Println("==========================")

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	//切片的切片組
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

//輸出字符串數組
func PrintStrArray(array []string) {
	for i, v := range array {
		fmt.Printf("index: %d	value:%s\t|", i, v)
	}
	fmt.Println()
}

//變更切換內容
func AddSlice(slice []string) {
	for i := 0; i < len(slice); i++ {
		slice[i] += "_l"
	}
}

//輸出切片訊息
func PrintSliceInfoWithStr(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
