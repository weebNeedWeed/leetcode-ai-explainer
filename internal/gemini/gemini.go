package gemini

import (
	"bytes"
	"context"
	_ "embed"
	"log"
	"text/template"

	"github.com/google/generative-ai-go/genai"
	"github.com/weebNeedWeed/leetcode-ai-explainer/internal/app_config"
	"google.golang.org/api/option"
)

//go:embed prompt.tmpl
var rawPrompt string

//go:embed system_instruction.tmpl
var systemInstruction string

func NewClient() *genai.Client {
	ctx := context.TODO()
	client, err := genai.NewClient(ctx, option.WithAPIKey(app_config.Config.GeminiApiKey))
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func GetPrompt(solution string) string {
	t1 := template.New("t1")
	t1 = template.Must(t1.Parse(rawPrompt))

	b := bytes.NewBuffer([]byte{})
	t1.Execute(b, solution)

	return b.String()
}

func GetSystemInstruction() string {
	return systemInstruction
}
