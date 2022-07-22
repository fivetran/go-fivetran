package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

var (
	TEST_KEY    = "test_key"
	TEST_SECRET = "test_secret"

	TEST_CONNECTOR_ID = "test_connector_id"
	TEST_HASH         = "test_hash"
	TEST_PUBLIC_KEY   = "test_public_key"
)

func CreateTestClient() (*fivetran.Client, *mock.HttpClient) {
	ftClient := fivetran.New(TEST_KEY, TEST_SECRET)
	mockClient := mock.NewHttpClient()
	ftClient.SetHttpClient(mockClient)
	return ftClient, mockClient
}

func requestBodyToJson(t *testing.T, req *http.Request) map[string]interface{} {
	t.Helper()

	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		t.Errorf("requestBodyToJson, cannot read request body: %s", err)
	}
	req.Body.Close()

	result := map[string]interface{}{}
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		t.Errorf("requestBodyToJson, cannot parse request body: %s", err)
	}

	return result
}

func printError(t *testing.T, actual interface{}, expected interface{}) {
	t.Helper()
	t.Errorf("Expected: %s"+
		"\n     but: <%s>\n",
		fmt.Sprintf("value equal to <%v>", expected),
		fmt.Sprintf("%v", actual),
	)
}

func isEmpty(actual interface{}) bool {
	var isEmpty bool = false

	if actual == nil {
		isEmpty = true
	} else if actualValue, ok := actual.(string); ok {
		isEmpty = actualValue == ""
	} else if reflect.ValueOf(actual).Len() == 0 {
		isEmpty = true
	}
	return isEmpty
}

func assertNotEmpty(t *testing.T, actual interface{}) {
	t.Helper()

	var isEmpty bool = isEmpty(actual)

	if isEmpty {
		printError(t, actual, "none-empty value")
	}
}

func assertEqual(t *testing.T, actual interface{}, expected interface{}) {
	t.Helper()

	if !reflect.DeepEqual(expected, actual) {
		printError(t, actual, expected)
	}
}
