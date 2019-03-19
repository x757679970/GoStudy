package main

import "fmt"
import "time"

func main() {
	num := 2
	num2 := 3
	if num%2 == 1 {
		fmt.Println("odd")
	} else {
		fmt.Println("偶數")
	}
	if num2%2 == 1 {
		fmt.Println("odd")
	} else {
		fmt.Println("偶數")
	}

	//switch case是循序往下判斷，並且預設都會break，無須另外寫
	//這個是無條件的switch
	switch {
	case num == 2:
		fmt.Println("I got the number 2")
	case num2 == 3:
		fmt.Println("I got the number 3")
	default:
		fmt.Println("others....")
	}

	//一般的switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("normal")
	}
}
