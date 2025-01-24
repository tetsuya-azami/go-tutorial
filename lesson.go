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
	var t, f bool = true, false
	fmt.Println("Hello world", "golang", "\n", time.Now(), "\n", "int: ", i, "\n", "bool: ", t, f)
	fmt.Println(user.Current())
	printTypeDefaults()
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
