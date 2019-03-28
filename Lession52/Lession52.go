package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "mkeqmnk3ne3!@@(你好_eo2"

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("加密前：", data)
	fmt.Println("加密後：", sEnc)

	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println("解密後：", string(sDec))
	fmt.Println()

	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
