package main

import (
	"context"
	"os"
	"time"

	"github.com/cloudwego/eino-ext/components/model/ark"
)

func NewChatModel(ctx context.Context) *ark.ChatModel {
	timeout := 30 * time.Second
	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey:  os.Getenv("ARK_API_KEY"),
		Model:   os.Getenv("MODEL"),
		Timeout: &timeout,
	})
	if err != nil {
		panic(err)
	}
	return model
}
