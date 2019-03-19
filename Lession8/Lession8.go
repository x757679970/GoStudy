package main

import "fmt"

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())

	//defer其實是一種壓棧行為
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("END")
}

func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f2() (result int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (result int) {
	t := 5
	defer func() {
		t = t + 5
		result += 1
	}()
	return t
}

func f4() (result int) {
	defer func(result int) {
		result = result + 5
	}(result)
	return 1
}
