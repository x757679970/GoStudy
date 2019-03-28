package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

//類內部配有異步互斥鎖
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock() //如果去除這兩個lock行為，將會報錯題是寫入衝突
	c.v[key]++
	c.mux.Unlock() //如果去除這兩個lock行為，將會報錯題是寫入衝突
}

func (c *SafeCounter) value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("someKey") //開啟1000個進程，進行inc疊加
	}

	time.Sleep(time.Second)         //主線程休眠一秒，如果不休眠，則主進程輸出value時，可能部分協程還未進行完畢
	fmt.Println(c.value("somekey")) //輸出該數值

	fmt.Println("========================================")

	var state = make(map[int]int) //int,int 的map
	var mutex = &sync.Mutex{}     //鎖對象
	var ops int64 = 0             //原子計數，表示我們對狀態進行了多少次操作
	var create int64 = 0

	for r := 0; r < 100; r++ {
		//創建100個協程
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				//fmt.Println("Total:", total)
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}
	for w := 0; w < 10; w++ {
		//開啟10個協程，隨機產生map<int, int>表
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&create, 1)

				runtime.Gosched() //釋放時間切片
			}
		}()
	}
	time.Sleep(time.Second) //主線程休眠一秒，此時讓協程去執行一秒

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)
	createFinal := atomic.LoadInt64(&create)
	fmt.Println("Create map times:", createFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
