package main

import (
	"fmt"
	"os/user"
	"time"
)

func init() {
	fmt.Println("init")
}

func main() {
	bazz()
	var i int = 1
	t, f := true, false
	fmt.Println("Hello world", "golang", "\n", time.Now(), "\n", "int: ", i, "\n", "bool: ", t, f)
	fmt.Println(user.Current())
	printTypeDefaults()
	array()
	slice()
}

func bazz() {
	fmt.Println("bazz")
}

func printTypeDefaults() {
	var (
		i    int
		f64  float64
		s    string
		t, f bool
	)
	fmt.Printf("型のデフォルト値: %v %v %q %v %v\n", i, f64, s, t, f)
}

func array() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a) // [100 200]

	b := [2]int{100} // 要素が足りない場合はデフォルト値(この場合は0)が入る
	fmt.Println(b)   // [100 0]
}

func slice() {
	a := []int{1, 2, 3}
	fmt.Println(a) // [1 2 3]
	a = append(a, 4)
	fmt.Println(a) // [1 2 3 4]
	a[2] = 100
	fmt.Println(a) // [1 2 100 4]
}
