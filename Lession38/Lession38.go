package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool) //通知關閉的信號通道

	//開了個協程，再等待通道
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received, job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	//主進程插入了三個通到數據
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs) //主進程主動關閉通道，如果不關閉，則協程還會再等待主線程發送通道，結果會互相等待，死鎖
	fmt.Println("sent all jobs")

	//主進程掛起，等待協程通知關閉
	fmt.Println(<-done)
}
