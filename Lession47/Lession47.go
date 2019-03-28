package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("%v\n", p)    //值
	fmt.Printf("%+v\n", p)   //包跨key
	fmt.Printf("%#v\n", p)   //包跨類型，key和值
	fmt.Printf("%T\n", p)    //類型
	fmt.Printf("%t\n", true) //bool類型
	fmt.Printf("%d\n", 123)  //int類型
	fmt.Printf("%b\n", 14)   //二進制輸出
	fmt.Printf("%c\n", 33)
	fmt.Printf("%x\n", 456)
	fmt.Printf("%f\n", 78.09)
	fmt.Printf("%e\n", 123400000.0)  //科學計號表示
	fmt.Printf("%E\n", 123400000.0)  // 用E表示
	fmt.Printf("%s\n", "\"string\"") //字符串輸出
	fmt.Printf("%q\n", "\"string\"") //不做解析的字符串輸出
	fmt.Printf("%x\n", "hex this")   //hex處理

	fmt.Printf("%p\n", &p) //指定複雜格式輸出
	fmt.Printf("|%6d|%6d|\n", 12, 345)
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	//格式化
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
