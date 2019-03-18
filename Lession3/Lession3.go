package main

import "fmt"
import "strings"

func main() {
	p := fmt.Println
	strBig := "HI THIS IS BIG WORD TO SMAILL TEST =p"
	//大寫轉小寫
	lower := strings.ToLower(strBig)
	p(lower)

	//檢查字串內是否有包夸
	if strings.Contains(lower, "mail") {
		p("yes got the word")
	}

	//字符串切割
	p("Characters 3-9:" + strBig[3:9])
	p("First Five:" + strBig[:5])

	//字符串拆離
	sentence := "I'm a sentence made up if words"
	words := strings.Split(sentence, "")

	// 格式化 %v 是指輸出數據組, 注意每個print的功能有點不一樣
	fmt.Printf("%v\n", words)

	//如果是要用空格來切割的話，使用Fields函式更好
	//可以使用空格拆來分析，還支持所有空白字符拆分
	fields := strings.Fields(sentence)
	fmt.Printf("%v\n", fields)

	p("==========我是分隔線==========")
	p("Contains:  ", strings.Contains("test", "es"))
	p("Count:     ", strings.Count("test", "t"))
	p("HasPrefix: ", strings.HasPrefix("test", "te")) //確認某字串開頭字串
	p("HasSuffix: ", strings.HasSuffix("test", "st")) //確認字串結尾字串
	p("Index:     ", strings.Index("test", "st"))
	p("Join       ", strings.Join([]string{"a", "b"}, "-"))
	p("Repeat     ", strings.Repeat("a", 5))
	p("Replace:   ", strings.Replace("aoole", "o", "p", -1))
	p("Replace:   ", strings.Replace("aoole", "o", "p", 1))
	p("Split:	  ", strings.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", strings.ToLower("BIG"))
	p("ToUpper:   ", strings.ToUpper("smalle"))
	p("Len:       ", len("Lenght"))
	p("Char:      ", "hello"[1])
	p()
}
