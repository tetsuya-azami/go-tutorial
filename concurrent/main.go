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

func restFunc() <-chan int {
	result := make(chan int)

	go func() {
		defer close(result)
		for i := 0; i < 5; i++ {
			result <- 1
		}
	}()

	return result
}

func practiceSelect() {
	gen1, gen2 := make(chan int), make(chan int)

	// if n1, ok := <-gen1; ok {
	// 	fmt.Println(n1)
	// } else if n2, ok := <-gen2; ok {
	// 	fmt.Println(n2)
	// } else {
	// 	fmt.Println("neither cannot use")
	// }
	select {
	case num := <-gen1:
		fmt.Println(num)
	case num := <-gen2:
		fmt.Println(num)
	default:
		fmt.Println("neither cannot use")
	}
}

func generator(done chan struct{}, i int) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)

		for {
			select {
			case <-done:
				fmt.Println("generator: doneに入りました")
				return
			default:
				fmt.Println("generator: defaultに入りました")
				result <- i
			}
		}
	}()
	return result
}

func fanIn1(done chan struct{}, c1, c2 <-chan int) <-chan int {
	result := make(chan int)

	go func() {
		defer fmt.Println("closed fanin")
		defer close(result)
		for {
			// caseはfor文で回せないので(=可変長は無理)
			// 統合元のチャネルがスライスでくるとかだとこれはできない
			// →応用編に続く
			select {
			case <-done:
				fmt.Println("done")
				return
			case num := <-c1:
				fmt.Println("send 1")
				result <- num
			case num := <-c2:
				fmt.Println("send 2")
				result <- num
			default:
				fmt.Println("continue")
				continue
			}
		}
	}()

	return result
}

func useFanIn() {
	done := make(chan struct{})

	gen1 := generator(done, 1) // int 1をひたすら送信するチャネル(doneで止める)
	gen2 := generator(done, 2) // int 2をひたすら送信するチャネル(doneで止める)

	result := fanIn1(done, gen1, gen2) // 1か2を受け取り続けるチャネル
	for i := 0; i < 5; i++ {
		<-result
	}
	close(done)
	fmt.Println("main close done")

	// これを使って、main関数でcloseしている間に送信された値を受信しないと
	// チャネルがブロックされてしまってゴールーチンリークになってしまう恐れがある
	for {
		if _, ok := <-result; !ok {
			break
		}
	}
}

func main() {
	// getLuckyNumAndPrint()
	// race()
	// race2()
	// practiceSelect()
	// callGenerator()
	useFanIn()
}
