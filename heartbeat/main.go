package main

import (
	"fmt"
	"time"
)

type Result struct {
	Data string
	Err  error
}

func NewResult(data string, err error) Result {
	return Result{
		Data: data,
		Err:  err,
	}
}

func doWork(done chan struct{}, pulseInterval time.Duration) (<-chan struct{}, <-chan Result) {
	heartBeat := make(chan struct{})
	result := make(chan Result)

	go func() {
		defer close(heartBeat)
		defer close(result)

		pulse := time.Tick(pulseInterval)
		stopHeartBeat := time.After(2 * time.Second) // heart beatを止める用
		workGen := time.Tick(2 * pulseInterval)

		sendResult := func(r time.Time) {
			for {
				select {
				case <-done:
					return
				case <-pulse:
					heartBeat <- struct{}{}
				case result <- NewResult("Ok work done!", nil):
					return
				}
			}
		}

		for {
			select {
			case <-pulse:
				select {
				case heartBeat <- struct{}{}:
				default: // heart beatを受け取る人がいない場合にブロックするのを防ぐために必要
				}
			case <-stopHeartBeat:
				time.Sleep(10 * time.Second) // heart beatを止める用
			case r := <-workGen:
				sendResult(r)
			case <-done:
				return
			}
		}
	}()

	return heartBeat, result
}

func main() {
	done := make(chan struct{})
	heartBeat, result := doWork(done, time.Duration(time.Second))
	time.AfterFunc(5*time.Second, func() { close(done) })

	isHeartBeatClosed := false
	isResultClosed := false

	for !isHeartBeatClosed || !isResultClosed {
		select {
		case _, ok := <-heartBeat:
			if !ok {
				fmt.Println("heart bead closed")
				isHeartBeatClosed = true
				heartBeat = nil
			} else {
				fmt.Println("receive heart beat")
			}
		case r, ok := <-result:
			if !ok {
				fmt.Println("result closed")
				isResultClosed = true
				result = nil
			} else {
				fmt.Printf("receive result: %v\n", r)
			}
		case <-time.After(1500 * time.Millisecond):
			fmt.Println("unhealthy worker")
			return
		}
	}

	fmt.Println("finish!")
}
