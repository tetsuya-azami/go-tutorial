package handlers

import (
	"context"
	"context-sample/auth"
	"context-sample/db"
	"fmt"
	"time"
)

type MyRequest struct {
	path string
}

func (req *MyRequest) SetPath(path string) {
	req.path = path
}

func (req *MyRequest) GetPath() string {
	return req.path
}

type MyResponse struct {
	Code int
	Body string
	Err  error
}

type MyHandleFunc func(context.Context, MyRequest)

var GetGreeting MyHandleFunc = func(ctx context.Context, req MyRequest) {
	var res MyResponse

	userID, err := auth.VerifyAuthToken(ctx)
	if err != nil {
		res.Code = 401
		res.Body = "Unauthorized"
		fmt.Println(res)
		return
	}

	dbReqCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	rcvChan := db.DefaultDB.Search(dbReqCtx, userID)

	data := <-rcvChan
	if data.Err != nil {
		res.Code = 500
		res.Body = "Internal Server Error"
		fmt.Println(res)
		cancel()
		return
	}

	res.Code = 200
	res.Body = fmt.Sprintf("DB Access Succeed!. userID: %d, Detail: %s", userID, data.Data)
	cancel()
	fmt.Println(res)
}

func NotFoundHandler(ctx context.Context, req MyRequest) {
	var res MyResponse
	res.Code = 404
	res.Body = "Not Found"
	fmt.Println(res)

	return
}
