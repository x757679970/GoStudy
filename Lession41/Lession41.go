package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	//創建三個協程
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	//主線程添加工作
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs) //全部工作完成

	//從協程中獲取工作結果
	for a := 1; a <= 9; a++ {
		fmt.Printf("- job %d done\n", <-results/2)
	}
}
