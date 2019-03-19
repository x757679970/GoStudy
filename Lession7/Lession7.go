package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filePath := "./../../res/testfile.txt"
	//使用系統緩衝家載一個文件
	f, err := os.Open(filePath)
	fmt.Println(f)
	fmt.Println(err)
	//延遲釋放
	f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		//輸出文字
		line := scanner.Text()
		fmt.Println(line)
	}

	//如果有錯誤則輸出
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//寫文件
	f1, err := os.Create("./../../res/outPutFile.txt")
	if err != nil {
		log.Fatalln("Error create files ", err)
	}
	defer f1.Close()

	//寫入字符串數組
	for _, str := range []string{"apple", "banna", "carrot"} {
		bytes, err := f.WriteString(str)
		if err != nil {
			log.Fatal("Error writing string", err)
		}
		fmt.Printf("wrote %d nytes to file\n", bytes)
	}
}
