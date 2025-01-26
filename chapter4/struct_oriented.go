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
	Age  int
}

type UserNotFound struct {
	Username string
}

func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	return p.Name
}

func (p Person) String() string {
	return fmt.Sprintf("My name is %v", p.Name)
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
	// do2(10)
	// do2("Mike")
	// do2(true)
	// stringer()
	customError()
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
	var mike Human = &Person{"Mike", 11}
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

func stringer() {
	mike := Person{"Mike", 22}
	fmt.Println(mike)
}

func myFunc() error {
	// something wrong
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{Username: "mike"}
}

func customError() {
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}
}

func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not found: %v", e.Username)
}
