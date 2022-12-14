package tests

import (
	"testing"

	"github.com/fivetran/go-fivetran"
)

type testStruct struct {
	SomeField    *string `json:"some_field,omitempty"`
	AnotherField *string `json:"another_field,omitempty"`
}

func TestMergeIntoMap(t *testing.T) {
	someFieldValue := "someFieldValue"
	anotherFieldValue := "anotherFieldValue"
	testValue := &testStruct{
		SomeField:    &someFieldValue,
		AnotherField: &anotherFieldValue,
	}
	testMap := make(map[string]interface{})
	testMap["some_key"] = "someKeyValue"
	fivetran.MergeIntoMap(testValue, &testMap)
	assertKeyValue(t, testMap, "some_field", someFieldValue)
	assertKeyValue(t, testMap, "another_field", anotherFieldValue)
	assertKeyValue(t, testMap, "some_key", "someKeyValue")
}

func TestFetchFromMap(t *testing.T) {
	testMap := make(map[string]interface{})
	testMap["some_key"] = "someKeyValue"
	testMap["some_field"] = "someFieldValue"
	testMap["another_field"] = "anotherFieldValue"
	var testValue testStruct

	err := fivetran.FetchFromMap(&testMap, &testValue)

	assertIsNil(t, err)
	assertIsNotNil(t, testValue)

	assertHasNoKey(t, testMap, "some_field")
	assertHasNoKey(t, testMap, "another_field")
	assertHasKey(t, testMap, "some_key")

	assertEqual(t, *testValue.SomeField, "someFieldValue")
	assertEqual(t, *testValue.AnotherField, "anotherFieldValue")
}
