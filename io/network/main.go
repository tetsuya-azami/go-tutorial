package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	go listenServer()
	time.Sleep(1 * time.Second)
	receiveData()
}

func listenServer() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}
	defer ln.Close()

	conn, err := ln.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	fmt.Println("client connected")

	sendDataFromServer(conn)
}

func sendDataFromServer(conn net.Conn) {
	str := "Hello from server"
	data := []byte(str)
	_, err := conn.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func receiveData() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Dial failed :", err)
		return
	}

	data := make([]byte, 1024)
	count, err := conn.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Received from server: ", string(data[:count]))
}
