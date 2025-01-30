package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Person struct {
	Name      string   `json:"name"`
	Age       int      `json:"age,string,omitempty"` // stringを指定することで数値を文字列として扱う
	Nicknames []string `json:"nicknames"`
}

func (p Person) MarshalJSON() ([]byte, error) { // カスタムマーシャル
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr." + p.Name,
	})
	return v, err
}

func (p *Person) UnmarshalJSON(b []byte) error { // カスタムアンマーシャル
	type Person2 struct {
		Name string
	}
	var p2 Person2
	err := json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Println(err)
	}
	p.Name = p2.Name

	return err
}

func main() {
	// getRequestAndPrintResult("http://example.com")
	// parseUrl("http://example.com/geao", "/test?a=1&b=2") // geaoの部分は反映されない
	// httpNewGetRequest("http://example.com", "/test?a=1&b=2")
	// httpNewPostRequest("http://example.com", "/test", "password")
	unmarshalJson(`
	{
		"name":"Alice",
		"age":"25",
		"nicknames":["Ally","Alicia"]
	}`)
	// unmarshalJson(`{
	// 	"name":"Alice",
	// 	"nicknames":["Ally","Alicia"]
	// }`)
	// marshalJson()
}

func getRequestAndPrintResult(domain string) {
	resp, _ := http.Get(domain)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func parseUrl(baseUrl string, pathAndQuery string) {
	base, _ := url.Parse(baseUrl)
	reference, _ := url.Parse(pathAndQuery)
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)
	fmt.Println(reference.Query())
}

func httpNewGetRequest(baseUrl string, query string) {
	base, _ := url.Parse(baseUrl)
	reference, _ := url.Parse(query)

	// クエリの追加
	q := reference.Query()
	q.Add("c", "3&%")
	reference.RawQuery = q.Encode()

	endpoint := base.ResolveReference(reference).String()
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	fmt.Println("endpoint: ", endpoint)

	var client *http.Client = &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func httpNewPostRequest(baseUrl string, path string, reqBody string) {
	base, _ := url.Parse(baseUrl)
	parsedPath, _ := url.Parse(path)
	url := base.ResolveReference(parsedPath)

	req, _ := http.NewRequest("POST", url.String(), bytes.NewBuffer([]byte(reqBody)))

	var client *http.Client = &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println(resp)
	fmt.Println(string(respBody))
}

func unmarshalJson(targetJson string) {
	b := []byte(targetJson)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(p.Name)
	fmt.Println(p.Age)
	fmt.Println(p.Nicknames)
}

func marshalJson() {
	p := Person{
		Name:      "Alice",
		Age:       25,
		Nicknames: []string{"Ally", "Alicia"},
	}
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(string(b))
}
