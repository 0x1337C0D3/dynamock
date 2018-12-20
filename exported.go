package dynamock

import "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

var mock *Mocked

// New to create new mock instance
func New() (dynamodbiface.DynamoDBAPI, *ExpectationStore) {
	mock = new(Mocked)
	mock.store = new(ExpectationStore)
	return mock, mock.store
}

// ExpectGetItem retrieves GetItemExpectation store
func (m *ExpectationStore) ExpectGetItem() *GetItemExpectation {
	getItemExpect := GetItemExpectation{table: nil, key: nil}
	m.getItemExpect = append(m.getItemExpect, getItemExpect)
	return &m.getItemExpect[len(m.getItemExpect)-1]
}

// ExpectScan retrieves ScanExpectation store
func (m *ExpectationStore) ExpectScan() *ScanExpectation {
	ScanExpect := ScanExpectation{table: nil}
	m.scanExpect = append(m.scanExpect, ScanExpect)
	return &m.scanExpect[len(m.scanExpect)-1]
}

// ExpectQuery retrieves ScanExpectation store
func (m *ExpectationStore) ExpectQuery() *QueryExpectation {
	queryExpect := QueryExpectation{table: nil}
	m.queryExpectation = append(m.queryExpectation, queryExpect)
	return &m.queryExpectation[len(m.queryExpectation)-1]
}

// ExpectPutItem retrieves putItemExpectation store
func (m *ExpectationStore) ExpectPutItem() *PutItemExpectation {
	putItemExpect := PutItemExpectation{table: nil, item: nil}
	m.putItemExpectation = append(m.putItemExpectation, putItemExpect)
	return &m.putItemExpectation[len(m.putItemExpectation)-1]
}

// ExpectDescribeTable retrieves descTableExpectation store
func (m *ExpectationStore) ExpectDescribeTable() *DescribeTableExpectation {
	item := DescribeTableExpectation{table: nil}
	m.descTableExpectation = append(m.descTableExpectation, item)
	return &m.descTableExpectation[len(m.descTableExpectation)-1]
}
