package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
	// "github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/joho/godotenv"
)

func main_embedded() {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	// 初始化embedder
	embedder, err := ark.NewEmbedder(ctx, &ark.EmbeddingConfig{
		APIKey: os.Getenv("EMBEDDER_API_KEY"),
		Model:  os.Getenv("EMBEDDER"),
	})

	if err != nil {
		panic(err)
	}
	input := []string{
		"如何使用 Go 语言读取文件？",
		"Go 语言的切片和数组有什么区别？",
		"如何在 Go 中处理并发？",
	}
	embeddings, err := embedder.EmbedStrings(ctx, input)
	if err != nil {
		panic(err)
	}
	for i, embed := range embeddings {
		fmt.Printf("Input: %s\nEmbedding的长度: %v\n\n", input[i], len(embed))
	}

}
