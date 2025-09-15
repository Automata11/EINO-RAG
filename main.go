package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	// InitClient()
	// defer MilvusCli.Close()

	// embedder := NewChatEmbedder(ctx)
	// // indexer := NewArkIndexer(ctx, embedder)
	// // splitter := NewTranslator(ctx)
	// retriever := NewArkRetriever(ctx, embedder)

	// documents, err := retriever.Retrieve(ctx, "打枪时常嘴微张，自嘲像植物大战")
	// if err != nil {
	// 	panic(err)
	// }
	// for i, doc := range documents {
	// 	fmt.Printf("Document %d:\n", i+1)
	// 	fmt.Printf("ID: %s\n", doc.ID)
	// 	fmt.Printf("Content: %s\n", doc.Content)
	// 	fmt.Printf("Metadata: %v\n", doc.MetaData)
	// 	fmt.Println()
	// }

	timeout := 30 * time.Second
	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey:  os.Getenv("ARK_API_KEY"),
		Model:   os.Getenv("MODEL"),
		Timeout: &timeout,
	})
	if err != nil {
		panic(err)
	}

	// 用之前定义的工具初始化tools
	todoTools := []tool.BaseTool{
		CreateTool(),
	}

	// 获取工具信息并绑定到model
	toolInfos := make([]*schema.ToolInfo, len(todoTools))
	for i, tool := range todoTools {
		info, err := tool.Info(ctx)
		if err != nil {
			panic(err)
		}
		toolInfos[i] = info
	}
	err = model.BindTools(toolInfos)
	if err != nil {
		panic(err)
	}

	// 创建Tool节点
	toolNode, err := compose.NewToolNode(ctx, &compose.ToolsNodeConfig{
		Tools: todoTools,
	})
	if err != nil {
		panic(err)
	}

	// 构建完整处理链，创建链条，添加节点，编译运行
	chain := compose.NewChain[[]*schema.Message, []*schema.Message]()
	chain.
		AppendChatModel(model, compose.WithNodeName("ark model")).
		AppendToolsNode(toolNode, compose.WithNodeName("tools"))
	r, err := chain.Compile(ctx)
	if err != nil {
		panic(err)
	}
	resp, err := r.Invoke(ctx, []*schema.Message{
		{
			Role:    schema.User,
			Content: "用户需要Minecraft的信息.",
		},
	})
	if err != nil {
		panic(err)
	}
	for _, m := range resp {
		fmt.Println(m.Content)
	}

	// // 使用lamda工具将String数据转换为Message数组，与模型输入格式一致
	// lambda := compose.InvokableLambda(func(_ context.Context, input string) ([]*schema.Message, error) {
	// 	output := input + ",回答结尾加上desuwa"
	// 	results := []*schema.Message{
	// 		{
	// 			Role:    schema.User,
	// 			Content: output,
	// 		},
	// 	}
	// 	return results, nil
	// })

	// chain := compose.NewChain[string, *schema.Message]()
	// chain.AppendLambda(lambda).AppendChatModel(model)
	// r, err := chain.Compile(ctx)
	// if err != nil {
	// 	panic(err)
	// }
	// anser, err := r.Invoke(ctx, "你叫什么名字")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(anser.Content)

}
