package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func BenchmarkPlusStringEveryTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		plusStringEveryTime(io.Discard)
	}
}

func BenchmarkStringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringsJoin(io.Discard)
	}
}

func plusStringEveryTime(w io.Writer) {
	var s, sep string
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Fprintln(w, s)
}

func stringsJoin(w io.Writer) {
	s := strings.Join(os.Args, " ")
	fmt.Fprintln(w, s)
}
