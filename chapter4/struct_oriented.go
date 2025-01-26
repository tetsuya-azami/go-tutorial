package main

import "fmt"

type Vertex struct {
	X, Y int
}

type MyInt int

type Human interface {
	Say() string
}

type Person struct {
	Name string
}

func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	return p.Name
}

func New(x, y int) *Vertex {
	return &Vertex{x, y}
}

func (v Vertex) Add() int {
	return v.X + v.Y
}

func (v *Vertex) ChangeX(newx int) {
	v.X = newx
}

func main() {
	// add()
	// changeX()
	// newVertex()
	// originalType()
	// say()
	// DriveCar(&Person{"Mike"})
	// do(10)
	do2(10)
	do2("Mike")
	do2(true)
}

func add() {
	v := Vertex{3, 4}
	res := v.Add()
	fmt.Println(res)
}

func changeX() {
	v := Vertex{3, 4}
	v.ChangeX(10)
	fmt.Println(v)
}

func newVertex() {
	v := New(3, 4)
	v.ChangeX(10)
	fmt.Println(*v)
}

func (i MyInt) Double() int {
	return int(i) * 2
}

func originalType() {
	myInt := MyInt(10)
	fmt.Println(myInt.Double())
}

func say() {
	var mike Human = &Person{"Mike"}
	mike.Say()
}

func DriveCar(h Human) {
	if h.Say() == "Mr.Mike" {
		fmt.Println("Run")
	} else {
		fmt.Println("Get out")
	}
}

func do(i interface{}) {
	ii := i.(int) * 2
	fmt.Println(ii)
}

func do2(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	}
}
