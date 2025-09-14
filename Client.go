package main

import (
	"context"
	"fmt"
	"log"

	"github.com/milvus-io/milvus-sdk-go/v2/client"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
)

var (
	milvusAddr = "localhost:19530"
	dbName     = "AwesomeEino"
	collName   = "test"
)

func clientInit() {
	ctx := context.Background()

	// 1. 连接到 Milvus
	c, err := client.NewClient(ctx, client.Config{
		Address: milvusAddr,
	})
	if err != nil {
		log.Fatalf("无法连接到 Milvus: %v", err)
	}
	defer c.Close()
	fmt.Printf("成功连接到 Milvus: %s\n", milvusAddr)

	// 2. 创建数据库
	err = c.CreateDatabase(ctx, dbName)
	if err != nil {
		// 如果数据库已存在，这可能会返回错误，我们可以根据需要处理
		fmt.Printf("创建数据库 '%s' 失败 (可能已存在): %v\n", dbName, err)
	} else {
		fmt.Printf("成功创建数据库: %s\n", dbName)
	}

	// 切换到新创建的数据库上下文
	c.UsingDatabase(ctx, dbName)
	fmt.Printf("切换到数据库: %s\n", dbName)

	// 3. 定义并创建 collection
	schema := &entity.Schema{
		CollectionName: collName,
		Description:    "Test collection for AwesomeEino",
		Fields: []*entity.Field{
			{
				Name:       "id",
				DataType:   entity.FieldTypeVarChar,
				PrimaryKey: true,
				AutoID:     false,
				TypeParams: map[string]string{
					entity.TypeParamMaxLength: "256",
				},
			},
			{
				Name:     "vector",
				DataType: entity.FieldTypeBinaryVector,
				TypeParams: map[string]string{
					entity.TypeParamDim: "81920",
				},
			},
			{
				Name:     "content",
				DataType: entity.FieldTypeVarChar,
				TypeParams: map[string]string{
					entity.TypeParamMaxLength: "8192",
				},
			},
			{
				Name:     "metadata",
				DataType: entity.FieldTypeJSON,
			},
		},
	}

	err = c.CreateCollection(ctx, schema, entity.DefaultShardNumber)
	if err != nil {
		log.Fatalf("创建 collection '%s' 失败: %v", collName, err)
	}

	fmt.Printf("成功创建 collection: %s\n", collName)
}
