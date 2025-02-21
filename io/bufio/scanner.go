package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("read.txt")
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		fmt.Println(line)
	}
}
