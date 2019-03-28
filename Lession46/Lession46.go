package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func main() {
	_, err := os.Create(getCurrentDirectory() + "/text.txt")
	if err != nil {
		panic(err)
	}
	//這裡會當掉，並控制台輸出堆棧
	panic("a problem")
}
