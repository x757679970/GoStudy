//聲明本包的名
package main

//加入使用的第三方套件
//like import ""
import "fmt"

func main() {
	//印出資料
	fmt.Println("This is my first Go"
	//此處將印出這個事情宣告給一個變數p使用
	//然後直接使用p就可以做印出的動作
	p := fmt.Println
	p("This is special use")
}
