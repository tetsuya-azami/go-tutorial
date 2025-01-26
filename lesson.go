package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func init() {
	fmt.Println("init")
}

func main() {
	// bazz()
	// var i int = 1
	// t, f := true, false
	// fmt.Println("Hello world", "golang", "\n", time.Now(), "\n", "int: ", i, "\n", "bool: ", t, f)
	// fmt.Println(user.Current())
	// printTypeDefaults()
	// array()
	// slice()
	// byteString()
	// mapFunc()
	// added, subtracted := calc(1, 2)
	// fmt.Println(added, subtracted)
	// funcLiteral()
	// cloasure()
	// forStatement()
	// deferPractice()
	// stackingDefer()
	// fileLoad()
	// logFatalError()
	// loggingSettings("test.log")
	// logFatalError()
	save()
	fmt.Println("ok?")
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
	fmt.Println(a)      // [1 2 100 4]
	fmt.Println(a[1:3]) // [2 100]
	fmt.Println(a[:])   // [1 2 100 4]
}

func byteString() {
	a := []byte{97, 98, 99}
	fmt.Println(a)         // [97 98 99]
	fmt.Println(string(a)) // abc
}

func mapFunc() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m) // map[apple:100 banana:200]
	m["apple"] = 150
	fmt.Println(m["apple"])
	v, ok := m["apple"]
	fmt.Println(v, ok) // 150 true
	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2) // 0 false
}

func calc(x int, y int) (added int, subtracted int) {
	added = x + y
	subtracted = x - y
	return
}

func funcLiteral() {
	f := func(x int, y int) (added int) {
		added = x + y
		return
	}
	fmt.Println(f(1, 2))
}

func cloasure() {
	x := 0
	incrementCallByValue := func(x int) int { // 値渡し
		x++
		return x
	}

	fmt.Println(incrementCallByValue(x)) // 1
	fmt.Println(incrementCallByValue(x)) // 1
	fmt.Println(incrementCallByValue(x)) // 1
	fmt.Println(x)                       // 0

	x = 0
	incrementCallByReference := func(x *int) int { // 参照渡し
		*x++
		return *x
	}

	fmt.Println(incrementCallByReference(&x)) // 1
	fmt.Println(incrementCallByReference(&x)) // 2
	fmt.Println(incrementCallByReference(&x)) // 3
	fmt.Println(x)                            // 3
}

func forStatement() {
	// lenでsliceをループ
	// s := []string{"java", "python", "c"}
	// for i := 0; i < len(s); i++ {
	// 	fmt.Println(i, s[i])
	// }

	// rangeでスライスをループ
	// s := []string{"java", "python", "c"}
	// for i, v := range s {
	// 	fmt.Println(i, v)
	// }

	// mapをループ
	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func deferPractice() {
	defer fmt.Println("deferPractice start") // deferは関数の最後に実行される
	deferPracticeInner := func() {
		defer fmt.Println("deferPracticeInner start")
		fmt.Println("deferPracticeInner end")
	}
	deferPracticeInner()
	fmt.Println("deferPractice end")
	// deferPracticeInner end
	// deferPracticeInner start
	// deferPractice end
	// deferPractice start
}

func stackingDefer() {
	defer fmt.Println("stackingDefer 1")
	defer fmt.Println("stackingDefer 2")
	defer fmt.Println("stackingDefer 3")
}

func fileLoad() {
	f, _ := os.Open("lesson.go")
	defer f.Close()
	data := make([]byte, 100)
	f.Read(data)
	fmt.Println(string(data))
}

func logFatalError() {
	f, err := os.Open("notExistsFile")
	defer f.Close()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("ok")
}

func loggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

func thirdPartyConnectDB() {
	panic("Unable to connect database")
}

func save() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}
