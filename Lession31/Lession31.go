//注意：管道需要用<-來拋值
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //將得到的值輸入到管道c中
}

func main() {
	s := []int{1, 2, 0, -4, 4, 6, 8, 2, 4, 3}

	c := make(chan int)
	//只能初始化創建標準類型map, slicem channel返回的是對像飲用
	//new 和常規知識一樣，返回的都是字定義的對象只針
	slice := s[:len(s)/2]
	for _, v := range slice {
		fmt.Printf("%v ", v)
	}
	fmt.Println("\n--------------------------")
	go sum(slice, c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c //從管道中獲取數值
	fmt.Println(x, y, x+y)
}
