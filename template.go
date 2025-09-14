package main

import (
	"context"

	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

// 创建模板，使用 FString 格式

func t() {
	template := prompt.FromMessages(schema.FString,
		schema.SystemMessage("你是一个{role}，你需要用{language}回答问题，你的目标是帮助程序员提供技术建议"),
		// 插入历史对话
		schema.MessagesPlaceholder("{history}", true),
		// 插入用户问题
		schema.UserMessage("问题：{question}"),
	)

	messages, err := template.Format(context.Background(), map[string]any{
		"role":     "助手",
		"language": "中文",
		"question": "如何使用 Go 语言读取文件？",
		"history": []*schema.Message{
			schema.UserMessage("你是谁？"),
			schema.AssistantMessage("我是助手兼鼓励师", nil),
		},
	})
	if err != nil {
		panic(err)
	}
	_ = messages // Use messages as needed
}
