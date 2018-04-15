package dynamock

import "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

var mock *Mocked

func New() (dynamodbiface.DynamoDBAPI, *ExpectationStore) {
	mock = new(Mocked)
	mock.store = new(ExpectationStore)
	return mock, mock.store
}

func (m *ExpectationStore) ExpectGetItem() *GetItemExpectation {
	getItemExpect := GetItemExpectation{table: nil, key: nil}
	m.GetItemExpect = append(m.GetItemExpect, getItemExpect)
	return &m.GetItemExpect[len(m.GetItemExpect)-1]
}

func (m *ExpectationStore) ExpectScan() *ScanExpectation {
	ScanExpect := ScanExpectation{table: nil}
	m.ScanExpect = append(m.ScanExpect, ScanExpect)
	return &m.ScanExpect[len(m.ScanExpect)-1]
}
