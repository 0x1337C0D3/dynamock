package dynamock

import (
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// WithTable specifies table name
func (e *QueryExpectation) WithTable(table string) *QueryExpectation {
	e.table = &table
	return e
}

// WithKeyCondition specifies key conditions
func (e *QueryExpectation) WithKeyCondition(conditions map[string]*dynamodb.Condition) *QueryExpectation {
	e.conditions = conditions
	return e
}

// ThenReturn specifies returned value
func (e *QueryExpectation) ThenReturn(res dynamodb.QueryOutput) *QueryExpectation {
	e.output = &res
	return e
}

// ThenThrow specifies returned error
func (e *QueryExpectation) ThenThrow(err error) *QueryExpectation {
	e.error = err
	return e
}

// Query to satisfy Query function from dynamodb service
func (e *Mocked) Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	for _, x := range e.store.queryExpectation {
		if *x.table == *input.TableName {
			if x.conditions != nil {
				if !reflect.DeepEqual(x.conditions, input.KeyConditions) {
					return nil, fmt.Errorf("expect key %+v, found key %+v", x.conditions, input.KeyConditions)
				}
			}
			// delete first element of expectation
			e.store.queryExpectation = append(e.store.queryExpectation[:0], e.store.queryExpectation[1:]...)
			return x.output, x.error
		}
	}

	return nil, fmt.Errorf("expectation store is empty")
}
