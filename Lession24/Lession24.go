package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	//簡易版
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	newInt := intSeq()
	fmt.Println(newInt())
	fmt.Println("====================")

	//複雜版
	pos, neq := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neq(-2*i),
		)
	}
}
