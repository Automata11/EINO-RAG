package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
)

func main0() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	timeout := 30 * time.Second
	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey:  os.Getenv("ARK_API_KEY"),
		Model:   os.Getenv("MODEL"),
		Timeout: &timeout,
	})
	if err != nil {
		panic(err)
	}
	messages := []*schema.Message{
		schema.SystemMessage("你是一个助手"),
		schema.UserMessage("你好"),
		schema.UserMessage("简单说明如何判断涨潮时间"),
	}

	reader, err := model.Stream(ctx, messages)
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	for {
		chunk, err := reader.Recv()
		if err != nil {
			break
		}
		fmt.Print(chunk.Content)
	}
	// response, err := model.Generate(ctx, input)
	// if err != nil {
	// 	panic(err)
	// }
	// println(response.Content)

}
