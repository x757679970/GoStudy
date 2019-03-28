package main

import (
	"fmt"
	"strconv"
)

func main() {
	//float64 轉換
	f64, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f64)
	f32, _ := strconv.ParseFloat("1.234", 32)
	fmt.Println(f32)

	//int 轉換
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)
	//int 16位轉換
	d, _ := strconv.ParseInt("ox1c8", 0, 64)
	fmt.Println(d)

	//uint轉換
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	//簡便接口
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	//非法轉換時的處理
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
