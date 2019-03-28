package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	//簡單的檢查自符串是否符合正則
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	//其他複雜的正則表達式需要先進行編譯
	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))                   //檢查是否符合正則
	fmt.Println(r.FindString("peach punch"))              //檢查符合正則的部分
	fmt.Println(r.FindStringIndex("peach punch"))         //檢查符合政則部分的字符索引
	fmt.Println(r.FindStringSubmatch("peach punch"))      //查找符合正則的子字串
	fmt.Println(r.FindStringSubmatchIndex("peach punch")) //查找符合正則的子字符串索引

	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	fmt.Println(r.FindAllString("peach punch pinch", 2))
	fmt.Println(r.Match([]byte("peach")))

	//進行正則表達式的編譯時，我們可以使用MustComile替代Compile，它只有一個返回值
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))

}
