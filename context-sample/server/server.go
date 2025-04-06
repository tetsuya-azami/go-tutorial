package server

import (
	"context"
	"context-sample/auth"
	"context-sample/handlers"
	"context-sample/session"
	"fmt"
)

type MyServer struct {
	router map[string]handlers.MyHandleFunc
}

func NewMyServer() *MyServer {
	return &MyServer{
		router: map[string]handlers.MyHandleFunc{
			"/test": handlers.GetGreeting,
		},
	}
}

func (srv *MyServer) ListenAndServe() {
	for {
		var path, token string
		fmt.Scan(&path)
		fmt.Scan(&token)

		ctx := session.SetSessionID(context.Background())
		ctx = auth.SetAuthToken(ctx, token)

		go srv.Request(ctx, path)
	}
}
