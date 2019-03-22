package main

import (
	"fmt"
	"math"
)

//接口
//實現多態的核心
type Abser interface {
	Abs() float64
}

type MyFloat float64

//實現MyFloat的ABS
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

//實現Vertex的ABS
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//實現MyFloat的Abs

func main() {
	var a Abser //聲明一個Abser格式的a

	f := MyFloat(-2.0)
	v := Vertex{3, 4}

	a = f //MyFloat類型的Abs實現
	fmt.Println(a.Abs())

	a = &v //*Vertex類型的Abs實現
	fmt.Println(a.Abs())

}
