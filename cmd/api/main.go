package main

import (
	"log"

	"github.com/weebNeedWeed/leetcode-ai-explainer/internal/storage"
)

func main() {
	storage := storage.NewDynamodbStorage()
	app := application{storage}

	router := app.mount()

	log.Fatal(app.run(router))
}
