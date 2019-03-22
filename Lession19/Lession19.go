package main

import (
	"fmt"
	"regexp"
)

func basicRegexes() {
	//創建一個茶找數字的正則表達式
	pattern := "[0-9]+"
	str := "The 12 monkeys have 48 bananas qwe99#EQWE"

	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex", err)
	}

	if re.MatchString(str) {
		fmt.Println("Yes. matched a number")
	} else {
		fmt.Println("No, not match")
	}

	//單次查詢
	result_two := re.FindString(str)
	fmt.Println("Number matched:", result_two)

	//全體查詢印出前兩個
	result_three := re.FindAllString(str, 2)
	for i, v := range result_three {
		fmt.Printf("Match %d: %s\n", i, v)
	}

	fmt.Println("================================")
	//全體查詢印出
	result_all := re.FindAllString(str, -1)
	for i, v := range result_all {
		fmt.Printf("Match %d: %s\n", i, v)
	}

	fmt.Println("================================")
	//正則替換
	result_four := re.ReplaceAllString(str, "xx")
	fmt.Println("Result:", result_four)
}

func case_insensitive() {
	ptn := `(?i)^t.`
	str := "To be or not to be"
	re, err := regexp.Compile(ptn)
	if err != nil {
		fmt.Println("Errpr compiling regex", err)
	}
	result := re.FindString(str)
	fmt.Println("Result:", result)
}

func main() {
	basicRegexes()
	fmt.Println("-------------------------")
	case_insensitive()
}
