package controller

import (
	"context"
	api "mvc-api/openapi"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (Server) GetPing(ctx context.Context, request api.GetPingRequestObject) (api.GetPingResponseObject, error) {
	return api.GetPing200JSONResponse{Ping: "ok 200"}, nil
}
