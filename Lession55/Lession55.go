package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	//--------------------------------------
	//啟動第三方程序，使用exec.Command
	//--------------------------------------
	//不需要參數的第三方程序
	ipconfigCmd := exec.Command("ipconfig")
	ipconfigOut, err := ipconfigCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ipconfig")
	fmt.Println(string(ipconfigOut))

	//給予輸入和輸出的地方三進程
	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	//顯示的接管Pipe，接下來指定其輸入輸出
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	//注意這裡的參數給予
	lsCmd := exec.Command("bash", "-c", "ls -a- -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
	fmt.Println("=========================")
	//--------------------------------------
	//啟動第三方程序，並關閉自身進程，使用syscall.Exec
	//--------------------------------------
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}
	//一些參數，注意，第一個參數應該是進程名字
	args := []string{"ls", "-a", "-l", "-h"}
	//獲取當前的環境變量
	env := os.Environ()
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
