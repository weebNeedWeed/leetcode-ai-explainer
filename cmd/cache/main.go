package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/weebNeedWeed/leetcode-ai-explainer/internal/app_config"
	"github.com/weebNeedWeed/leetcode-ai-explainer/internal/storage"
)

type item struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	TheType string `json:"type"`
}

func main() {
	githubToken := app_config.Config.GithubToken

	listOfFolders := getAllProblemsFolder(githubToken)

	c := make(chan []item, 2)
	wg := &sync.WaitGroup{}
	for _, folder := range listOfFolders {
		wg.Add(1)
		go getProblemsInsideAFolder(githubToken, folder, c, wg)
	}

	// monitor, wait until all goroutines execute successfully, close the channel
	go func() {
		wg.Wait()
		close(c)
	}()

	items := []item{}
	for i := range c {
		items = append(items, i...)
	}
	fmt.Printf("total items: %d\n", len(items))

	totalInserted := insertToDdbTable(items)
	fmt.Printf("total inserted: %d\n", totalInserted)
}

func getAllProblemsFolder(token string) []item {
	// get all folders that store solution
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/repos/doocs/leetcode/contents/solution", nil)
	if err != nil {
		log.Fatal(err)
	}

	// edit the header, add token and other params
	req.Header.Add("Authorization", "Bearer "+token)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// get and parse body to a list of folders that contain solutions
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	listOfFolders := []item{}
	json.Unmarshal(body, &listOfFolders)

	// remove item with type != dir
	ans := []item{}
	for _, item := range listOfFolders {
		if item.TheType == "dir" {
			ans = append(ans, item)
		}
	}

	return ans
}

func getProblemsInsideAFolder(token string, folder item, c chan []item, wg *sync.WaitGroup) {
	defer wg.Done()
	baseUrl := "https://api.github.com/repos/doocs/leetcode/contents/solution/"

	// make a req to get all problems inside a folder
	req, err := http.NewRequest(http.MethodGet, baseUrl+folder.Name, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	problems := []item{}
	json.Unmarshal(body, &problems)

	fmt.Printf("fetched for - %s - status: %d\n", folder.Name, len(problems))

	c <- problems
}

func insertToDdbTable(problems []item) int {
	conf, _ := config.LoadDefaultConfig(context.TODO())
	client := dynamodb.NewFromConfig(conf)

	totalInserted := 0

	batchSize := 25
	start, end := 0, batchSize

	for start < len(problems) {
		writeReqs := []types.WriteRequest{}

		if end > len(problems) {
			end = len(problems)
		}

		for _, i := range problems[start:end] {
			tempStruct := storage.ProblemInDb{
				Pk:   "PROBLEMS",
				Sk:   i.Name,
				Path: i.Path,
			}
			a, _ := attributevalue.MarshalMap(tempStruct)
			req := types.WriteRequest{
				PutRequest: &types.PutRequest{
					Item: a,
				},
			}
			writeReqs = append(writeReqs, req)
		}

		_, err := client.BatchWriteItem(context.TODO(), &dynamodb.BatchWriteItemInput{
			RequestItems: map[string][]types.WriteRequest{
				app_config.Config.DynamoDbTableName: writeReqs,
			},
		})

		if err != nil {
			log.Fatal(err)
		}

		start = end
		end += batchSize

		totalInserted += len(writeReqs)
	}

	return totalInserted
}
