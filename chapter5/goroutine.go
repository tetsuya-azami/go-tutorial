package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// goroutineMain()
	// channel()
	// bufferdChannel()
	// bufferdChannel()
	// prodCons()
	testChannelLoop()
}

func goroutineMain() {
	var wg sync.WaitGroup
	wg.Add(1)
	go goroutine("world", &wg)
	normal("hello")
	wg.Wait()
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func goroutine(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func channel() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int, 2)
	go goroutine1(s, c)
	for i := range c {
		fmt.Println(i)
	}
}

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum
	}
	close(c)
}

func goroutine2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum
	}
	close(c)
}

func bufferdChannel() {
	c := make(chan int, 2)
	c <- 100
	c <- 200
	close(c)

	for v := range c {
		fmt.Println(v)
	}
}

func prodCons() {
	var wg sync.WaitGroup
	ch := make(chan int)
	// producer
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}
	fmt.Println("producer done")
	// consumer
	go consumer(ch, &wg)
	wg.Wait() // Addした数だけDoneするまで待つ
	// close(ch)
	time.Sleep(2 * time.Second)
	fmt.Println("main done")
}

func producer(c chan int, i int) {
	c <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Println("process", i*1000)
		wg.Done()
	}
	fmt.Println("###########")
}

func testChannelLoop() {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("done")
}
