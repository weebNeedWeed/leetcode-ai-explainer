package storage

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/weebNeedWeed/leetcode-ai-explainer/internal/app_config"
)

type Problem struct {
	Name string
	Path string
}

type ProblemInDb struct {
	Pk   string `dynamodbav:"pk"`
	Sk   string `dynamodbav:"sk"`
	Path string `dynamodbav:"path"`
}

type Storage interface {
	GetProblem(zeroPaddedProblemId string) (Problem, error)
}

var ErrorNoProblemFound = errors.New("no problem found")

type DynamodbStorage struct {
	client *dynamodb.Client
}

func NewDynamodbStorage() *DynamodbStorage {
	conf, _ := config.LoadDefaultConfig(context.TODO())
	client := dynamodb.NewFromConfig(conf)

	return &DynamodbStorage{client}
}

func (d *DynamodbStorage) GetProblem(zeroPaddedProblemId string) (Problem, error) {
	keyCond := expression.Key("pk").Equal(expression.Value("PROBLEMS")).
		And(expression.Key("sk").BeginsWith(zeroPaddedProblemId))
	builder := expression.NewBuilder().WithKeyCondition(keyCond)
	exp, _ := builder.Build()

	o, err := d.client.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:                 aws.String(app_config.Config.DynamoDbTableName),
		KeyConditionExpression:    exp.KeyCondition(),
		ExpressionAttributeNames:  exp.Names(),
		ExpressionAttributeValues: exp.Values(),
	})
	if err != nil {
		return Problem{}, err
	}

	if o.Count == 0 {
		return Problem{}, ErrorNoProblemFound
	}

	pInDb := ProblemInDb{}
	raw := o.Items[0]

	attributevalue.UnmarshalMap(raw, &pInDb)

	ans := Problem{
		Name: pInDb.Sk,
		Path: pInDb.Path,
	}
	return ans, nil
}

func (p *ProblemInDb) GetKey() map[string]types.AttributeValue {
	pk, _ := attributevalue.Marshal("PROBLEMS")
	sk, _ := attributevalue.Marshal(p.Sk)
	return map[string]types.AttributeValue{
		"pk": pk,
		"sk": sk,
	}
}
