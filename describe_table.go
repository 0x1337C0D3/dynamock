package dynamock

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// WithTable specifies table name
func (e *DescribeTableExpectation) WithTable(table string) *DescribeTableExpectation {
	e.table = &table
	return e
}

// ThenReturn specifies returned value
func (e *DescribeTableExpectation) ThenReturn(res dynamodb.DescribeTableOutput) *DescribeTableExpectation {
	e.output = &res
	return e
}

// ThenThrow specifies returned error
func (e *DescribeTableExpectation) ThenThrow(err error) *DescribeTableExpectation {
	e.error = err
	return e
}

// DescribeTable to satisfy DescribeTable function from dynamodb service
func (e *Mocked) DescribeTable(input *dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error) {
	for _, x := range e.store.descTableExpectation {
		if *x.table == *input.TableName {
			e.store.descTableExpectation = append(e.store.descTableExpectation[:0], e.store.descTableExpectation[1:]...)
			return x.output, x.error
		}
	}
	return nil, fmt.Errorf("expectation store is empty")
}
