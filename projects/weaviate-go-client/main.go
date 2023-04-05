package main

import (
	"context"
	"fmt"

	"github.com/weaviate/weaviate-go-client/v4/weaviate"
)

func GetSchema() {
	cfg := weaviate.Config{
		Host:   "localhost:8080",
		Scheme: "http",
	}
	client := weaviate.New(cfg)

	schema, err := client.Schema().Getter().Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", schema)

	//models
	//client.Schema().ClassCreator().WithClass()

	model, err := client.Schema().ClassGetter().Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", model)
}

func main() {
	GetSchema()
}
