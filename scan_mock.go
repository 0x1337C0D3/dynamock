package dynamock

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// WithTable specifies table name
func (e *ScanExpectation) WithTable(table string) *ScanExpectation {
	e.table = &table
	return e
}

// ThenReturn specifies returned value
func (e *ScanExpectation) ThenReturn(res dynamodb.ScanOutput) *ScanExpectation {
	e.output = &res
	return e
}

// ThenThrow specifies returned error
func (e *ScanExpectation) ThenThrow(err error) *ScanExpectation {
	e.error = err
	return e
}

// Scan to satisfy Scan function from dynamodb service
func (e *Mocked) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	for _, x := range e.store.scanExpect {
		if *x.table == *input.TableName {
			// delete first element of expectation
			e.store.scanExpect = append(e.store.scanExpect[:0], e.store.scanExpect[1:]...)
			return x.output, x.error
		}
	}

	return nil, fmt.Errorf("expectation store is empty")
}
