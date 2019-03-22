package main

import "fmt"

type SFather struct {
	Name string
	Age  int
}

func (s *SFather) SetName(name string) {
	s.Name = name
}

func (s *SFather) GetName() string {
	return s.Name
}

func (s *SFather) SetAge(age int) {
	s.Age = age
}

func (s *SFather) GetAge() int {
	return s.Age
}

type SChildren struct {
	SFather //繼承
	nIndex  int
}

//多態
type IPeople interface {
	GetID() int
	SetID(int)
}

func (s *SFather) SetID(ID int) {
	s.Age = ID
}

func (s *SFather) GetID() int {
	return s.Age
}

func (s *SChildren) SetID(ID int) {
	s.nIndex = ID
}

func (s *SChildren) GetID() int {
	return s.nIndex
}

func main() {
	//封裝
	pFather := new(SFather)
	pFather.SetName("Jerry")
	fmt.Println(pFather.Name)

	//繼承
	pChildren := new(SChildren)
	pChildren.SetName("Ivy")
	fmt.Println(pChildren.Name)

	//多態
	pChildren.SetID(10)
	pFather.SetID(10)

	fmt.Printf("%v .. %T .. age=%d\n", pFather, pFather, pFather.GetAge())
	fmt.Printf("%v .. %T .. age=%d\n", pChildren, pChildren, pChildren.GetAge())
}
