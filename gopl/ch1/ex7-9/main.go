package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	urls := os.Args[1:]
	for _, url := range urls {
		fetch(os.Stdout, url)
	}
}

func fetch(dst io.Writer, url string) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(dst, resp.Body)
	dst.Write([]byte(resp.Status))
	// res, err := io.ReadAll(resp.Body)
	// fmt.Println(string(res))
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
}
