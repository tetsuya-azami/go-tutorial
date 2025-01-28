package main

import (
	"fmt"
	"go-tutorial/chapter6/mylib"
	"go-tutorial/chapter6/mylib/under"
	"regexp"
	"time"

	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)

func main() {
	// useMylib()
	// useLib()
	// useTime()
	useRegexp()
}

func useMylib() {
	nums := []int{1, 2, 3, 4, 5}
	avg := mylib.Average(nums)
	fmt.Println(avg)
	under.Hello()
}

func useLib() {
	spy, _ := quote.NewQuoteFromYahoo("spy", "2016-01-01", "2016-04-01", quote.Daily, true)
	fmt.Println(spy.CSV())
	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2)
}

func useTime() {
	t := time.Now()
	fmt.Println("default: ", t)
	fmt.Println("RFC3339: ", t.Format(time.RFC3339))
	fmt.Println("年:", t.Year(), "月:", int(t.Month()), "日:", t.Day())
}

func useRegexp() {
	// // simple use
	// match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	// fmt.Println(match)

	// // usex mustCompile
	// r := regexp.MustCompile("a([a-z]+)e")
	// result := r.MatchString("apple")
	// fmt.Println(result)

	// // findString
	// r2 := regexp.MustCompile("/(view|edit|save)/([a-zA-Z0-9]+)$")
	// result := r2.FindString("/view/test")
	// fmt.Println(result)

	// findStringSubmatch
	r2 := regexp.MustCompile("/(view|edit|save)/([a-zA-Z0-9]+)$")
	result := r2.FindStringSubmatch("/view/test")
	fmt.Println(result)
	fmt.Println(len(result))
}
