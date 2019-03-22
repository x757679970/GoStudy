package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	i, j := 42, 2222
	p := &i
	fmt.Println("變更前址：", p)
	fmt.Println("變更前值：", *p)
	fmt.Println(*p)
	*p = 21 //修改只有指針
	fmt.Println(i)

	//打印指針
	fmt.Println("變更後址：", *p)
	fmt.Println("變更後值：", &p)

	p = &j
	*p = *p / 37
	fmt.Println(j)

	v := Vertex{1, 2}
	q := &v
	q.X = 1e9 //類型被強制轉換
	fmt.Println(v)
}
