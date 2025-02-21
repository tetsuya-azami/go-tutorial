package main

import (
	"bufio"
	"os"
	"testing"
)

func BenchmarkReadX(b *testing.B) {
	f, _ := os.Open("read.txt")
	defer f.Close()

	d := make([]byte, 40*1024*1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Read(d)
	}
}

func BenchmarkReadY(b *testing.B) {
	f, _ := os.Open("read.txt")
	defer f.Close()
	bf := bufio.NewReader(f)

	d := make([]byte, 4*1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bf.Read(d)
	}
}
