package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	urls := os.Args[1:]
	fetchAll(urls)
}

func fetchAll(urls []string) {
	start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Println(time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	if !(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
		url = "http://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
	}
	defer resp.Body.Close()

	// f, err := os.Create(sanitizeUrl(url))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer func() {
	// 	err := f.Close()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }()

	// nbytes, err := io.Copy(f, resp.Body)
	nbytes, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("error occured %v", err)
	}

	ch <- fmt.Sprintf("%d bytes written", nbytes)
}

// func sanitizeUrl(url string) string {
// 	url = strings.TrimPrefix(url, "https://")
// 	url = strings.TrimPrefix(url, "http://")
// 	url = strings.ReplaceAll(url, "/", "_")

// 	return url
// }
