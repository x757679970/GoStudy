//文件讀寫
package main

import "fmt"
import "io/ioutil"
import "log"

func main() {
	filename := "./../../res/testfile.txt"

	//讀取文件
	//有錯誤會返回Err； 若沒有錯誤Err為nil
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Error readfing file", filename)
	}

	//讀取完的資料content內容是字節數組[] byte，要轉型態為string再輸出
	fmt.Println(string(content))

	outPutFile := "./../../res/outPutFile.txt"

	err = ioutil.WriteFile(outPutFile, content, 0644)
	if err != nil {
		log.Fatal("Error writing file:", err)
	}
}
