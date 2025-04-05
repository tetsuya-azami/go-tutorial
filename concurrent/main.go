package main

import (
	"context"
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

func chanelWithTimeout(c chan int) {
	timeout := time.After(1 * time.Second)
	for {
		select {
		case num := <-c:
			fmt.Println(num)
		case <-timeout:
			fmt.Println("timeout")
			return
		}
	}
}

var wg sync.WaitGroup

func generatorWithContext(ctx context.Context, num int) <-chan int {
	out := make(chan int)
	go func() {
		defer wg.Done()

	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			case out <- num:
			}
		}

		close(out)
		fmt.Println("generator closed")
	}()

	return out
}

func cancelUsingContext() {
	ctx, cancel := context.WithCancel(context.Background())
	gen := generatorWithContext(ctx, 1)

	wg.Add(1)

	for i := 0; i < 5; i++ {
		fmt.Println(<-gen)
	}

	cancel()
	wg.Wait()
}

func cancelPropagationSeries() {
	ctx0 := context.Background()
	ctx1, _ := context.WithCancel(ctx0)
	// G1
	go func(ctx1 context.Context) {
		ctx2, cancel2 := context.WithCancel(ctx1)

		// G2-1
		go func(ctx2 context.Context) {
			// G2-2
			go func(ctx2 context.Context) {
				select {
				case <-ctx2.Done():
					fmt.Println("G2-2: canceled")
				case <-time.After(time.Duration(1) * time.Second):
					fmt.Println("G2-2: timeout")
				}
			}(ctx2)

			select {
			case <-ctx2.Done():
				fmt.Println("G2-1: canceled")
			}
		}(ctx2)

		time.Sleep(time.Duration(2) * time.Second)
		cancel2()

		select {
		case <-ctx1.Done():
			fmt.Println("G1: canceled")
		}
	}(ctx1)

	time.Sleep(time.Second * 3)
}

func cancelPropagationParalell() {
	ctx0 := context.Background()
	ctx1, cancel1 := context.WithCancel(ctx0)
	// G1-1
	go func(ctx1 context.Context) {
		select {
		case <-ctx1.Done():
			fmt.Println("G1-1: canceled")
		}
	}(ctx1)

	// G1-2
	go func(ctx1 context.Context) {
		select {
		case <-ctx1.Done():
			fmt.Println("G1-2: canceled")
		}
	}(ctx1)

	cancel1()
	time.Sleep(time.Second)
}

func cancelPropagationSibling() {
	ctx0 := context.Background()
	ctx1, cancel1 := context.WithCancel(ctx0)
	// G1
	go func(ctx1 context.Context) {
		select {
		case <-ctx1.Done():
			fmt.Println("G1: canceled")
		}
	}(ctx1)

	ctx2, _ := context.WithCancel(ctx0)
	// G2
	go func(ctx2 context.Context) {
		select {
		case <-ctx2.Done():
			fmt.Println("G2: canceled")
		}
	}(ctx2)

	cancel1()
	time.Sleep(time.Second)
}

func main() {
	// getLuckyNumAndPrint()
	// race()
	// race2()
	// practiceSelect()
	// callGenerator()
	// useFanIn()

	// c := make(chan int)
	// go chanelWithTimeout(c)
	// c <- 1
	// c <- 2
	// c <- 3
	// c <- 4
	// c <- 5
	// time.Sleep(2 * time.Second)

	// cancelUsingContext()
	// cancelPropagationSeries()
	// cancelPropagationParalell()
	cancelPropagationSibling()
}
