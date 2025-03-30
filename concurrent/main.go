package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func getLuckyNum(c chan<- int) {
	fmt.Println("...")
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	num := rand.Intn(10)
	c <- num
}

func getLuckyNumAndPrint() {
	c := make(chan int)
	defer close(c)

	go getLuckyNum(c)

	fmt.Printf("Today's your lucky number is %d!\n", <-c)
}

func race() {
	src := []int{1, 2, 3, 4, 5}
	dst := []int{}

	c := make(chan int)

	for _, s := range src {
		go func(s int) {
			num := s * 2
			c <- num
		}(s)
	}

	for range src {
		num := <-c
		dst = append(dst, num)
	}

	fmt.Print(dst)
}

func race2() {
	src := []int{1, 2, 3, 4, 5}
	dst := []int{}

	var mu sync.Mutex

	for _, s := range src {
		go func(s int) {
			num := s * 2
			mu.Lock()
			dst = append(dst, num)
			mu.Unlock()
		}(s)
	}

	time.Sleep(1 * time.Second)

	fmt.Print(dst)
}

func main() {
	// getLuckyNumAndPrint()
	// race()
	// race2()
}
