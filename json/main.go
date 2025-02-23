package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age"`
}

func main() {
	decode(encode())
}

func encode() string {
	p := Person{Name: "", Age: 27}

	bytes, _ := json.Marshal(p)
	// json.NewEncoder(os.Stdout).Encode(p)
	return string(bytes)
}

func decode(target string) {
	person := &Person{}
	// if err := json.Unmarshal([]byte(jsonString), person); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(*person)
	if err := json.NewDecoder(strings.NewReader(target)).Decode(person); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(person.Name)
}
