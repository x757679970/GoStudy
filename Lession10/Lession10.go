package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	path := "../../"
	//調用系統函數
	//該系統函數的第二個參數必須為指定參數和返回值得函數指針
	filepath.Walk(path, MyWalker)
}

//該函數會被系統調用多次，fi是文件的訊息
func MyWalker(fn string, fi os.FileInfo, err error) error {
	if err != nil {
		fmt.Println("Walker Error", err)
		return nil
	}

	if fi.IsDir() {
		fmt.Println("Directory: ", fn)
	} else {
		fmt.Println("File: ", fn)
	}
	return nil
}
