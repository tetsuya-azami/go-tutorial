package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	m   map[string]int
	mux sync.Mutex
}

func (c *Counter) Increase(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.m[key]++
}

func (c *Counter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.m[key]
}

func main() {
	// goroutineMain()
	// channel()
	// bufferdChannel()
	// bufferdChannel()
	// prodCons()
	// testChannelLoop()
	// pipeline()
	// selectChannel()
	syncMutex()
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

func pipeline() {
	first := make(chan int)
	second := make(chan int)
	third := make(chan int)

	go producer1(first)
	go multii2(first, second)
	go multi4(second, third)
	for result := range third {
		fmt.Println(result)
	}
}

func producer1(first chan int) {
	defer close(first)
	for i := 0; i < 10; i++ {
		first <- i
	}
}

func multii2(first <-chan int, second chan<- int) {
	defer close(second)
	for i := range first {
		second <- i * 2
	}
}

func multi4(second <-chan int, third chan<- int) {
	defer close(third)
	for i := range second {
		third <- i * 4
	}
}

func selectChannel() {
	c1 := make(chan string)
	c2 := make(chan string)
	go selectChannelGoroutine1(c1)
	go selectChannelGoroutine2(c2)

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

func selectChannelGoroutine1(c chan string) {
	for {
		c <- "packed from 1"
		time.Sleep(1 * time.Second)
	}
}

func selectChannelGoroutine2(c chan string) {
	for {
		c <- "packed from 2"
		time.Sleep(1 * time.Second)
	}
}

func syncMutex() {
	c := Counter{m: make(map[string]int)}

	go func() {
		for i := 0; i < 10; i++ {
			c.Increase("key")
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c.Increase("key")
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println(c.Value("key"))
}
