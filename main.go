package main

import (
	"context"
	"log"
	"os"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/cloudwego/eino-ext/components/indexer/milvus"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
	// "github.com/cloudwego/eino-ext/components/model/ark"
)

func main() {

	InitClient()
	defer MilvusCli.Close()

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

	var collection = collName

	var fields = []*entity.Field{
		{
			Name:     "id",
			DataType: entity.FieldTypeVarChar,
			TypeParams: map[string]string{
				"max_length": "256",
			},
			PrimaryKey: true,
			AutoID:     false,
		},
		{
			Name:     "vector", // 确保字段名匹配
			DataType: entity.FieldTypeBinaryVector,
			TypeParams: map[string]string{
				"dim": "81920",
			},
		},
		{
			Name:     "content",
			DataType: entity.FieldTypeVarChar,
			TypeParams: map[string]string{
				"max_length": "8192",
			},
		},
		{
			Name:     "metadata",
			DataType: entity.FieldTypeJSON,
		},
	}

	indexer, err := milvus.NewIndexer(ctx, &milvus.IndexerConfig{
		Client:     MilvusCli,
		Collection: collection,
		Fields:     fields,
		Embedding:  embedder,
	})

	if err != nil {
		log.Fatalf("Failed to create indexer: %v", err)
	}

	docs := []*schema.Document{
		{
			ID:      "1",                                // ID is the unique identifier of the document.
			Content: "Go 语言是一种开源的编程语言，具有简洁、高效和并发编程的特点。", // Content is the content of the document.
			MetaData: map[string]any{ // MetaData is the metadata of the document, can be used to store extra information.
				"author": "David Stone",
				"year":   "2025",
			},
		},
	}

	ids, err := indexer.Store(ctx, docs)
	if err != nil {
		log.Fatalf("Failed to store documents: %v", err)
	}
	log.Printf("Successfully stored documents with IDs: %v", ids)

}
