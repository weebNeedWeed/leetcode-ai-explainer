package app_config

import (
	"log"

	"github.com/weebNeedWeed/leetcode-ai-explainer/internal/env"
)

type Configuration struct {
	Addr              string
	GithubToken       string
	DynamoDbTableName string
	GeminiApiKey      string
}

var Config = getConfig()

func getConfig() Configuration {
	c := Configuration{}

	c.Addr = env.GetString("ADDR", ":9090")
	if c.Addr == "" {
		log.Fatalf("no addr found")
	}

	c.DynamoDbTableName = env.GetString("DDB_TABLE_NAME", "")
	if c.DynamoDbTableName == "" {
		log.Fatalf("no table name found")
	}

	c.GithubToken = env.GetString("GITHUB_TOKEN", "")
	if c.GithubToken == "" {
		log.Fatalf("no github token found")
	}

	c.GeminiApiKey = env.GetString("GEMINI_APIKEY", "")
	if c.GithubToken == "" {
		log.Fatalf("no gemini api key found")
	}

	return c
}
