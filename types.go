package dynamock

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// Mocked struct
type Mocked struct {
	dynamodbiface.DynamoDBAPI
	store *ExpectationStore
}

// ExpectationStore contains all expectations for all dynamodb operation (e.g. getItem, scan, etc)
type ExpectationStore struct {
	getItemExpect        []GetItemExpectation
	scanExpect           []ScanExpectation
	queryExpectation     []QueryExpectation
	putItemExpectation   []PutItemExpectation
	descTableExpectation []DescribeTableExpectation
}

// GetItemExpectation contains expectation for getItem operation
type GetItemExpectation struct {
	table  *string
	key    map[string]*dynamodb.AttributeValue
	output *dynamodb.GetItemOutput
	error  error
}

// ScanExpectation contains expectation for scan operation
type ScanExpectation struct {
	table  *string
	output *dynamodb.ScanOutput
	error  error
}

// QueryExpectation contains expectation for query operation
type QueryExpectation struct {
	table      *string
	conditions map[string]*dynamodb.Condition
	output     *dynamodb.QueryOutput
	error      error
}

// PutItemExpectation contains expectation for putItem operation
type PutItemExpectation struct {
	table  *string
	item   map[string]*dynamodb.AttributeValue
	output *dynamodb.PutItemOutput
	error  error
}

// DescribeTableExpectation contains expectation for describeTable operation
type DescribeTableExpectation struct {
	table  *string
	output *dynamodb.DescribeTableOutput
	error  error
}

// AnyValue is used as any value in expectation matcher
type AnyValue struct{}
