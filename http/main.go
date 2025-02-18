package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/products/37", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Host = "www.example.com"

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}
