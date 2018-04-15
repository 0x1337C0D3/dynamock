package dynamock

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"reflect"
)

func (e *GetItemExpectation) WithTable(table string) *GetItemExpectation {
	e.table = &table
	return e
}

func (e *GetItemExpectation) WithKeys(keys map[string]*dynamodb.AttributeValue) *GetItemExpectation {
	e.key = keys
	return e
}

func (e *GetItemExpectation) ThenReturn(res dynamodb.GetItemOutput) *GetItemExpectation {
	e.output = &res
	return e
}

func (e *GetItemExpectation) ThenThrow(err error) *GetItemExpectation {
	e.error = err
	return e
}

func (e *Mocked) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	//TODO implement the input matcher instead of checking the first expectation
	if len(e.store.GetItemExpect) > 0 {
		x := e.store.GetItemExpect[0]

		if x.table != nil {
			if *x.table != *input.TableName {
				return nil, fmt.Errorf("expected table %s, found table : %s", *x.table, *input.TableName)
			}
		}

		if x.key != nil {
			if !reflect.DeepEqual(x.key, input.Key) {
				return nil, fmt.Errorf("expect key %+v, found key %+v", x.key, input.Key)
			}
		}

		// delete first element of expectation
		e.store.GetItemExpect = append(e.store.GetItemExpect[:0], e.store.GetItemExpect[1:]...)

		return x.output, x.error
	}

	return nil, fmt.Errorf("expectation store is empty")
}
