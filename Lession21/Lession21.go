package main

import "fmt"

type Dog struct {
	Name  string
	Color string
}

func main() {
	//創建類對象，並賦予屬性
	Spot := Dog{Name: "Spot", Color: "bowrn"}

	fmt.Println("1: ", Spot.Color)

	//獲取對象只針
	SpotPointer := &Spot
	SpotPointer.Color = "black"

	fmt.Println("2: ", Spot.Color)
}
