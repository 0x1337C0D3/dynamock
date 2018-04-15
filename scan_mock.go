package dynamock

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (e *ScanExpectation) WithTable(table string) *ScanExpectation {
	e.table = &table
	return e
}

func (e *ScanExpectation) ThenReturns(res dynamodb.ScanOutput) *ScanExpectation {
	e.output = &res
	return e
}

func (e *ScanExpectation) ThenThrow(err error) *ScanExpectation {
	e.error = err
	return e
}

func (e *Mocked) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	if len(e.store.ScanExpect) > 0 {
		x := e.store.ScanExpect[0] //get first element of expectation

		if x.table != nil {
			if *x.table != *input.TableName {
				return nil, fmt.Errorf("expect table %s but found table %s", *x.table, *input.TableName)
			}
		}

		// delete first element of expectation
		e.store.ScanExpect = append(e.store.ScanExpect[:0], e.store.ScanExpect[1:]...)
		return x.output, x.error
	}

	return nil, fmt.Errorf("expectation store is empty")
}
