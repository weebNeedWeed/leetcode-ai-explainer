package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/generative-ai-go/genai"
	"github.com/weebNeedWeed/leetcode-ai-explainer/internal/app_config"
	"github.com/weebNeedWeed/leetcode-ai-explainer/internal/gemini"
	"github.com/weebNeedWeed/leetcode-ai-explainer/internal/storage"
)

func (a *application) getExplanationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	problemIdAsStr := chi.URLParam(r, "id")
	problemId, err := strconv.Atoi(problemIdAsStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		}{http.StatusBadRequest, "invalid problem id"})
		return
	}

	zeroPadded := fmt.Sprintf("%04d", problemId)

	problem, err := a.storage.GetProblem(zeroPadded)
	if err != nil && err == storage.ErrorNoProblemFound {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		}{http.StatusBadRequest, "problem with given id not found"})
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		}{http.StatusInternalServerError, "internal server error"})
		log.Print(err)
		return
	}

	// get solution of problem
	baseUrl := "https://api.github.com/repos/doocs/leetcode/contents/" + problem.Path + "/README_EN.md"

	req, _ := http.NewRequest(http.MethodGet, baseUrl, nil)
	req.Header.Set("Accept", "application/vnd.github.v3.raw")
	req.Header.Set("Authorization", "Bearer "+app_config.Config.GithubToken)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		}{http.StatusInternalServerError, "internal server error"})
		log.Print(err)
		return
	}

	bodyRaw, _ := io.ReadAll(response.Body)
	solution := string(bodyRaw)

	aiClient := gemini.NewClient()
	model := aiClient.GenerativeModel("gemini-2.0-flash")
	model.SetTemperature(0.9)
	model.SetTopP(0.5)
	model.ResponseMIMEType = "text/plain"
	model.SystemInstruction = genai.NewUserContent(genai.Text(gemini.GetSystemInstruction()))

	prompt := gemini.GetPrompt(solution)
	aiResp, err := model.GenerateContent(context.TODO(), genai.Text(prompt))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		}{http.StatusInternalServerError, "internal server error"})
		log.Print(err)
		return
	}

	raw := aiResp.Candidates[0].Content.Parts[0]
	txt := raw.(genai.Text)
	txtAsStr := string(txt)
	txtAsStr = strings.TrimPrefix(txtAsStr, "```html")
	txtAsStr = strings.TrimSuffix(txtAsStr, "```")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data string `json:"data"`
	}{http.StatusOK, "success", txtAsStr})
}
