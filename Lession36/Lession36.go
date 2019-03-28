package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		//構造一個新的錯誤並返回
		return -1, errors.New("Can't work with 42")
	}
	//無錯誤則回nil
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

//重載error的interface
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		//自定義錯誤，則這樣返回錯誤
		return -1, &argError{arg, "Can't work with it"}
	}
	return arg + 3, nil
}
func main() {
	//測試標準錯誤
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	//測試自定義錯誤
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	//注意下文的if使用
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
