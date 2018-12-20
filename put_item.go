package dynamock

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// WithTable specifies table name
func (e *PutItemExpectation) WithTable(table string) *PutItemExpectation {
	e.table = &table
	return e
}

// WithItem specifies item value
func (e *PutItemExpectation) WithItem(item map[string]*dynamodb.AttributeValue) *PutItemExpectation {
	e.item = item
	return e
}

// ThenReturn specifies returned value
func (e *PutItemExpectation) ThenReturn(res dynamodb.PutItemOutput) *PutItemExpectation {
	e.output = &res
	return e
}

// ThenThrow specifies returned error
func (e *PutItemExpectation) ThenThrow(err error) *PutItemExpectation {
	e.error = err
	return e
}

// PutItem to satisfy PutItem function from dynamodb service
func (e *Mocked) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	for _, x := range e.store.putItemExpectation {
		if *x.table == *input.TableName {
			if x.item != nil {
				return x.output, nil
			}
			// delete first element of expectation
			e.store.putItemExpectation = append(e.store.putItemExpectation[:0], e.store.putItemExpectation[1:]...)
			return x.output, x.error
		}
	}
	return nil, fmt.Errorf("expectation store is empty")
}
