package dynamock

import (
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// WithTable specifies table name
func (e *GetItemExpectation) WithTable(table string) *GetItemExpectation {
	e.table = &table
	return e
}

// WithKeys specifies key values
func (e *GetItemExpectation) WithKeys(keys map[string]*dynamodb.AttributeValue) *GetItemExpectation {
	e.key = keys
	return e
}

// ThenReturn specifies returned value
func (e *GetItemExpectation) ThenReturn(res dynamodb.GetItemOutput) *GetItemExpectation {
	e.output = &res
	return e
}

// ThenThrow specifies returned error
func (e *GetItemExpectation) ThenThrow(err error) *GetItemExpectation {
	e.error = err
	return e
}

// GetItem to satisfy GetItem function from dynamodb service
func (e *Mocked) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	for _, x := range e.store.getItemExpect {
		if *x.table == *input.TableName {
			if x.key != nil {
				if !reflect.DeepEqual(x.key, input.Key) {
					return nil, fmt.Errorf("expect key %+v, found key %+v", x.key, input.Key)
				}
			}
			// delete first element of expectation
			e.store.getItemExpect = append(e.store.getItemExpect[:0], e.store.getItemExpect[1:]...)
			return x.output, x.error
		}
	}

	return nil, fmt.Errorf("expectation store is empty")
}
