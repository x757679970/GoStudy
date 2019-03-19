package main

import "fmt"

//定義一個struct
//裡面有名字跟顏色
type Dog struct {
	Name  string
	Color string
}

//定義一個呼叫結構並且印出
func (d Dog) Call() {
	fmt.Printf("Come here %s dog, %s \n", d.Name, d.Color)
}

func main() {
	//定義一個對象實例，並賦與屬性
	SpotDogInstance := Dog{Name: "Spor", Color: "brown"}
	//調用
	SpotDogInstance.Call()
}
