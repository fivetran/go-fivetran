package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"testing"
	"time"

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

func RequestBodyToJson(t *testing.T, req *http.Request) map[string]interface{} {
	t.Helper()

	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		t.Errorf("requestBodyToJson, cannot read request body: %s", err)
	}
	req.Body.Close()
	req.Body = io.NopCloser(bytes.NewReader(bodyBytes))

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

func assertIsNil(t *testing.T, value interface{}) {
	t.Helper()

	if value != nil {
		printError(t, value, "nil")
	}
}

func assertIsNotNil(t *testing.T, value interface{}) {
	t.Helper()

	if value == nil {
		printError(t, value, "non-nil value")
	}
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

func assertKey(t *testing.T, key string, requestPart map[string]interface{}, expectedValue interface{}) {
	v, ok := requestPart[key]
	assertEqual(t, ok, true)
	assertEqual(t, v, expectedValue)
}

func assertHasKey(t *testing.T, source map[string]interface{}, key string) {
	t.Helper()
	_, ok := source[key]
	if !ok {
		t.Errorf("Expected Key not found in map: %s", key)
	}
}

func assertTimeEqual(t *testing.T, actualTime time.Time, expectedTime string) {
	ex, _ := time.Parse(time.RFC3339, expectedTime)
	assertEqual(t, ex, actualTime)
}

func assertHasNoKey(t *testing.T, source map[string]interface{}, key string) {
	t.Helper()
	_, ok := source[key]
	if ok {
		t.Errorf("Unexpected Key found in map: %s", key)
	}
}

func assertKeyValue(t *testing.T, source map[string]interface{}, key string, expected interface{}) {
	t.Helper()
	assertHasKey(t, source, key)
	actual := source[key]
	if !reflect.DeepEqual(actual, expected) {
		printError(t, actual, expected)
	}
}

func boolToStr(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
