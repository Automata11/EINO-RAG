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
