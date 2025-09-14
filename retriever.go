package main

import (
	"context"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/cloudwego/eino-ext/components/retriever/milvus"
)

func NewArkRetriever(ctx context.Context, embedder *ark.Embedder) *milvus.Retriever {

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
	return retriever
}
