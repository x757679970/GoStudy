package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Http負責處理的請求
	http.HandleFunc("/hello", helloRequest)
	http.HandleFunc("/", getRequest)

	//啟動Web 服務器
	log.Println("Listening on http://localhost:9999")
	log.Fatal(http.ListenAndServe(":9999", nil))

}

//處理/hello請求
func helloRequest(w http.ResponseWriter, r *http.Request) {
	//注意這裡用的是Fprint，而不是Printf
	fmt.Fprint(w, "Hello Welt")
	return
}

//this function simply serves up the file requested, or
//an index list if a directory is requesterd
func getRequest(w http.ResponseWriter, r *http.Request) {
	file_requested := "./" + r.URL.Path
	http.ServeFile(w, r, file_requested)
	return
}
