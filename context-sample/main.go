package main

import "context-sample/server"

func main() {
	srv := server.NewMyServer()
	srv.ListenAndServe()
}
