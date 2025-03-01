package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

// func main() {
// 	// request()
// 	newConn
// }

// func request() {
// 	client := &http.Client{}
// 	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/products/37", nil)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	req.Host = "www.example.com"

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(string(body))
// }

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //オプションを解析します。デフォルトでは解析しません。
	fmt.Println(r.Form) //このデータはサーバのプリント情報に出力されます。
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //ここでwに入るものがクライアントに出力されます。
}

func main() {
	http.HandleFunc("/", sayhelloName)       //アクセスのルーティングを設定します。
	err := http.ListenAndServe(":9090", nil) //監視するポートを設定します。
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer l.Close()
	fmt.Printf("Listening on :8080, %T, %v\n", l, l)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Client connected:", conn.RemoteAddr())
	}
}
