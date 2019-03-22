package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertex) Abs2() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Scale2(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//雖然不能對標準類型進行自定義類函數，但是可以這樣做
type MyFloat float64

func (f MyFloat) AbsFloat() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
	fmt.Println(v)
	Scale(&v, 10)
	fmt.Println(v)
	v.Scale2(10)
	fmt.Println(v.Abs2())

	fmt.Println("================")

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.AbsFloat())

	fmt.Println("================")

	//指針調用
	v2 := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, ABS: %v\n", v2, v2.Abs2())
	v2.Scale2(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v2, v2.Abs2())
}
