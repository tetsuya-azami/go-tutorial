package controller

import (
	"context"
	api "mvc-api/openapi"
)

type PingController struct{}

func (pc *PingController) GetPing(ctx context.Context, request api.GetPingRequestObject) (api.GetPingResponseObject, error) {
	return api.GetPing200JSONResponse{Message: "ok 200"}, nil
}
