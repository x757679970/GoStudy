package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -7439967,
	}
	m["Test Labs"] = Vertex{
		10.24, -9.23512,
	}

	var m2 = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"BAIDU":     {66.6666, -94.87},
	}

	//數組添加
	m2["Amazon"] = Vertex{1.2, -2.3}
	fmt.Println(m)
	fmt.Println(m2)
}
