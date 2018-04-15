# dynamock
Simple DynamoDB mocking for testing

# How to use

The syntax for expectation store is similar to jmockit in Java. 

```
// mocked struct
var mocked = dynamock.New()

// Your db struct
var db = &dynamoDB{
	svc : mocked,
}

func TestValidGetByDb(t *testing.T) {
	mocked.ExpectGetItem().WithTable("TableName").WithKeys(key).ThenReturn(output)

	_, err := db.GetByID("validId")
	if err != nil {
		t.Fail()
	}
}

func TestInvalidGetByDb(t *testing.T) {
	mocked.ExpectGetItem().WithTable("TableName").WithKeys(key).ThenThrow(fmt.Errorf("Invalid Id"))

	_, err := db.GetByID("dummyId")
	if err == nil {
		t.Fail()
	}
}
```

# Supported operation

Currently, only the below methods of `dynamodbiface.DynamoDBAPI` are supported. However, more will come in later later

``` go
func (e *dynamodbiface.DynamoDBAPI) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error)
func (e *dynamodbiface.DynamoDBAPI) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) 
```