package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Entry struct {
	Title     string
	Author    string
	URL       string
	Permalink string
}

type Feed struct {
	Data struct {
		Children []struct {
			Data Entry
		}
	}
}

func main() {
	//請求的網站
	url := "https://www.reddit.com/r/golang/new.json"

	//獲取url
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error fetching", err)
	}

	//延遲關閉響應
	defer resp.Body.Close()

	//檢查我們是否得到OK
	if resp.StatusCode != http.StatusOK {
		log.Fatal("Error Status not OK:", resp.StatusCode)
	}

	//若OK，則讀取響應數據Body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Error reading body:", err)
	}

	//創建Feed對象，解析Json並寫入
	var entries Feed
	if err := json.Unmarshal(body, &entries); err != nil {
		log.Fatalln("Error decoing JSON", err)
	}

	//循環填充到子單元中
	for _, ed := range entries.Data.Children {
		Entry := ed.Data
		log.Println(">>>")
		log.Println("Title		:", Entry.Title)
		log.Println("Author		:", Entry.Author)
		log.Println("URL		:", Entry.URL)
		log.Println("Comments: http://reddit.com%s \n", Entry.Permalink)
	}
}
