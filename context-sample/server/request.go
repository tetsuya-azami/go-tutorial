package server

import (
	"context"
	"context-sample/handlers"
)

func (srv *MyServer) Request(ctx context.Context, path string) {
	var req handlers.MyRequest
	req.SetPath(path)

	if handler, ok := srv.router[path]; ok {
		handler(ctx, req)
	} else {
		handlers.NotFoundHandler(ctx, req)
	}
}
