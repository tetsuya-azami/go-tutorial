package main

import (
	"fmt"
	"go-tutorial/chapter6/mylib"
	"go-tutorial/chapter6/mylib/under"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	avg := mylib.Average(nums)
	fmt.Println(avg)
	under.Hello()
}
