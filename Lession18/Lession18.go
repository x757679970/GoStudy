package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//隨機產生一個[0,100)的半開半閉區間的int值
	//注意，如果沒有給初始化種子，前三個數字永遠是81、87、47
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println("----------------")

	//添加隨機種子，開始產生隨機數字
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println("----------------")

	//其他一些隨機方法，例如隨機變更type
	//Perm將0~N的數隨機分布為一個數組
	fmt.Println("Random Int:", rand.Int())
	fmt.Println("Random float:", rand.Float32())
	fmt.Println("Random Permutation of N ints:", rand.Perm(31))

	//該NormFloat64隨機數調用多次的話會生成一個取縣，不算標準隨機
	for i := 0; i < 10; i++ {
		fmt.Println("Normalized", rand.NormFloat64())
	}
	fmt.Println("----------------")

	//隨機整數
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	//隨機float
	fmt.Println(rand.Float64())
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	//設置種子
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	//設置相同的種子之後，隨機數是完全相同的
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)

	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)

	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
	fmt.Println()
}
