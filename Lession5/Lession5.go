package main

import "fmt"

func main() {
	var twoDot [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDot[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoDot)

	//字符串數組
	alphas := []string{"abc", "def", "ghi"}

	//以array的長度當條件
	for i := 0; i < len(alphas); i++ {
		fmt.Printf("%d: %s \n", i, alphas[i])
	}

	//以range來做回傳處理
	//此處回傳兩個數值i=所引編號，val=值
	for i, val := range alphas {
		fmt.Printf("%d:%s\t", i, val)
	}

	fmt.Println()
	for i := range alphas {
		fmt.Println(i)
	}

	//如果不關心所引編號，則用 _ 來取代，這表示，不保留該數值
	for _, val := range alphas {
		fmt.Println(val)
	}

	//類似while的方式
	//go 沒有while
	x := 0
	for x < 5 {
		fmt.Println(x)
		x++
	}

	//將go字符串視為數據，range取值
	for i, c := range "goo" {
		fmt.Println(i, c)
	}

	//無限循環
	//for {
	//	fmt.Println(".")
	//}
}
