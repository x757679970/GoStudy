package main

import (
	"fmt"
	"sort"
)

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	strEnglist := []string{"好", "不好", "可能好", "good", "不可能", "not good"}
	nums := []int{4, 2, 12, 5, 1, 3}
	fmt.Println("Raw ABC : ", strEnglist)
	fmt.Println("Raw Nums: ", nums)

	//string排序
	sort.Strings(strEnglist)
	fmt.Println("Sorted ABC:", strEnglist)

	//int 排序
	sort.Ints(nums)
	fmt.Println("Sorted Nums:", nums)

	//反向排序
	sort.Sort(sort.Reverse(sort.StringSlice(strEnglist)))
	fmt.Println("Reverse ABC:", strEnglist)

	//--------------------------------------
	//map排序
	//--------------------------------------
	hash := map[string]int{
		"c": 3,
		"a": 6,
		"b": 2,
		"e": 5,
		"d": 4,
	}

	//此時的Map是完全亂序，每次運行結果都不同
	for k, v := range hash {
		fmt.Printf("%s => %v\t", k, v)
	}
	fmt.Println()

	//根據Key排序
	keys := make([]string, 0, len(hash))
	for k := range hash {
		keys = append(keys, k)
	}
	//僅僅對Key做了排序
	sort.Strings(keys)
	//按照Key順序進行輸出，但注意，hash內部還是亂序
	for i := range keys {
		fmt.Printf("%s => %v\t", keys[i], hash[keys[i]])
	}

	fmt.Println("")
	fmt.Println("=================================")

	//--------------------------------------
	//字定義函數
	//--------------------------------------
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}
