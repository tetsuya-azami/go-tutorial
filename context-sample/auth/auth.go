package auth

import (
	"context"
	"errors"
	"fmt"
)

type ctxKey int

const (
	authToken ctxKey = iota
)

func SetAuthToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, authToken, token)
}

func getAuthToken(ctx context.Context) (string, error) {
	if token, ok := ctx.Value(authToken).(string); ok {
		return token, nil
	} else {
		fmt.Println("getAuthToken: no auth token in context")
		return "", errors.New("cannot get auth token")
	}
}

func VerifyAuthToken(ctx context.Context) (int, error) {
	token, err := getAuthToken(ctx)
	if err != nil {
		return 0, err
	}

	userID := len(token)
	if userID == 0 {
		return 0, errors.New("invalid auth token")
	}

	return userID, nil
}
