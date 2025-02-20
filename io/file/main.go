package main

import (
	"fmt"
	"os"
)

func main() {
	writeText()
}

func readText() {
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	data := make([]byte, 1024)
	count, err := f.Read(data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("read %d bytes:\n", count)
	fmt.Println(string(data))
}

func writeText() {
	f, err := os.Create("write.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	s := "write from golang changed"
	b := []byte(s)
	count, err := f.Write(b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("write %d bytes\n", count)
}
