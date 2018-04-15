package dynamock

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type Mocked struct {
	dynamodbiface.DynamoDBAPI
	store *ExpectationStore
}

type ExpectationStore struct {
	GetItemExpect []GetItemExpectation
	ScanExpect    []ScanExpectation
}

type GetItemExpectation struct {
	table  *string
	key    map[string]*dynamodb.AttributeValue
	output *dynamodb.GetItemOutput
	error  error
}

type ScanExpectation struct {
	table  *string
	output *dynamodb.ScanOutput
	error  error
}

type AnyValue struct{}
