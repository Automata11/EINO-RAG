package main

import (
	"context"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	InitClient()
	defer MilvusCli.Close()
	ctx := context.Background()
	embedder := NewChatEmbedder(ctx)
	// indexer := NewArkIndexer(ctx, embedder)
	// splitter := NewTranslator(ctx)
	retriever := NewArkRetriever(ctx, embedder)

	// // 准备文档
	// bs, err := os.ReadFile("./documents.md")
	// if err != nil {
	// 	panic(err)
	// }
	// docs := []*schema.Document{
	// 	{ID: "1",
	// 		Content: string(bs),
	// 	},
	// }

	// // 分割文档
	// results, err := splitter.Transform(ctx, docs)
	// if err != nil {
	// 	panic(err)
	// }
	// for i, doc := range results {
	// 	doc.ID = docs[0].ID + "_" + strconv.Itoa(i+1)
	// 	fmt.Printf("Document %d ID: %s\n", i+1, doc.ID)
	// }

	// // 保存分割结果
	// ids, err := indexer.Store(ctx, results)
	// fmt.Printf("文档保存成功,ids:%v\n", ids)

	documents, err := retriever.Retrieve(ctx, "打枪时常嘴微张，自嘲像植物大战")
	if err != nil {
		panic(err)
	}
	for i, doc := range documents {
		fmt.Printf("Document %d:\n", i+1)
		fmt.Printf("ID: %s\n", doc.ID)
		fmt.Printf("Content: %s\n", doc.Content)
		fmt.Printf("Metadata: %v\n", doc.MetaData)
		fmt.Println()
	}

}
