package main

import (
	"context"
	"os"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
)

func NewChatEmbedder(ctx context.Context) *ark.Embedder {
	//初始化embedder
	embedder, err := ark.NewEmbedder(ctx, &ark.EmbeddingConfig{
		APIKey: os.Getenv("EMBEDDER_API_KEY"),
		Model:  os.Getenv("EMBEDDER"),
	})
	if err != nil {
		panic(err)
	}
	return embedder
}
