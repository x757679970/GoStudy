package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Person struct {
	Name, City string
}

var person Person

var people []Person

type Response1 struct {
	Page   int
	Fruits []string
}

//type Response2 struct{
//	Page	int			'json:"page"'
//	Fruits	[]string	'json:"fruits"'
//}

func main() {
	//json字符串
	json_str := "{\"Name\": \"Marcus\", \"City\": \"San Joses\"}"

	//Unmarshal函數參數要求是 []byte，所以我們將string轉換過去，並填充給person對象
	if err := json.Unmarshal([]byte(json_str), &person); err != nil {
		fmt.Println("Error parsing JSON: ", err)
	}

	fmt.Println("Name: %v, City: %v\n", person.Name, person.City)

	//------------------------
	//批量家載
	//------------------------
	//加載一個jason文件
	file, err := ioutil.ReadFile("../../res/Names.json")
	if err != nil {
		fmt.Println("Error reading file")
	}

	//將json中的數據填充給people的對象數據
	if err := json.Unmarshal(file, &people); err != nil {
		fmt.Println("Error parsing JSON", err)
	}
	fmt.Println(people)

	//使用Marshal函數向Json文件中寫入一個對象數據
	json, err := json.Marshal(people)
	if err != nil {
		fmt.Println("JSON Encoding Error", err)
	}
	fmt.Println(string(json))

	fmt.Println("===================================")

}
