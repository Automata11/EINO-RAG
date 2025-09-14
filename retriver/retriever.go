package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/cloudwego/eino-ext/components/retriever/milvus"
	"github.com/joho/godotenv"
	// "github.com/cloudwego/eino-ext/components/model/ark"
)

func main_retriever() {

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

	// var collection = collName

	// var fields = []*entity.Field{
	// 	{
	// 		Name:     "id",
	// 		DataType: entity.FieldTypeVarChar,
	// 		TypeParams: map[string]string{
	// 			"max_length": "256",
	// 		},
	// 		PrimaryKey: true,
	// 		AutoID:     false,
	// 	},
	// 	{
	// 		Name:     "vector", // 确保字段名匹配
	// 		DataType: entity.FieldTypeBinaryVector,
	// 		TypeParams: map[string]string{
	// 			"dim": "81920",
	// 		},
	// 	},
	// 	{
	// 		Name:     "content",
	// 		DataType: entity.FieldTypeVarChar,
	// 		TypeParams: map[string]string{
	// 			"max_length": "8192",
	// 		},
	// 	},
	// 	{
	// 		Name:     "metadata",
	// 		DataType: entity.FieldTypeJSON,
	// 	},
	// }

	// indexer, err := milvus.NewIndexer(ctx, &milvus.IndexerConfig{
	// 	Client:     MilvusCli,
	// 	Collection: collection,
	// 	Fields:     fields,
	// 	Embedding:  embedder,
	// })

	// if err != nil {
	// 	log.Fatalf("Failed to create indexer: %v", err)
	// }

	// docs := []*schema.Document{
	// 	{
	// 		ID:      "3",                              // ID is the unique identifier of the document.
	// 		Content: "uestc是电子科技大学，是一所位于四川省成都市的985高校", // Content is the content of the document.
	// 		MetaData: map[string]any{ // MetaData is the metadata of the document, can be used to store extra information.
	// 			"author": "David Stone",
	// 			"year":   "2025",
	// 		},
	// 	},
	// }

	// ids, err := indexer.Store(ctx, docs)
	// if err != nil {
	// 	log.Fatalf("Failed to store documents: %v", err)
	// }
	// log.Printf("Successfully stored documents with IDs: %v", ids)

	retriever, err := milvus.NewRetriever(ctx, &milvus.RetrieverConfig{
		Client:      MilvusCli,
		Collection:  "test",
		VectorField: "vector",
		Embedding:   embedder,
		TopK:        2,
		OutputFields: []string{
			"id",
			"content",
			"metadata",
		},
	})
	if err != nil {
		panic(err)
	}
	documents, err := retriever.Retrieve(ctx, "uestc是哪里")
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
