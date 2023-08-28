package tests

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

const (
	GroupCreateExampleResponse = `
	{
		"code": "Success",
		"message": "Group has been created",
		"data": {
			"id": "decent_dropsy",
			"name": "Primary_Snowflake",
			"created_at": "2020-05-25T15:26:47.306509Z"
		}
	}`

	ExpectedGroupCreateCode      = "Success"
	ExpectedGroupCreateMessage   = "Group has been created"
	ExpectedGroupCreateID        = "decent_dropsy"
	ExpectedGroupCreateName      = "Primary_Snowflake"
	ExpectedGroupCreateCreatedAt = "2020-05-25T15:26:47.306509Z"
)

func TestGroupCreateServiceDo(t *testing.T) {
	// Arrange
	ftClient, mockClient := CreateTestClient()
	groupName := "NewGroup"

	handler := mockClient.When(http.MethodPost, "/v1/groups").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			body := requestBodyToJson(t, req)
			assertGroupCreateRequest(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareGroupCreateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewGroupCreate().
		Name(groupName).
		Do(context.Background())

	// assert
	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)

	assertGroupCreateResponse(t, response)
}

func prepareGroupCreateResponse() string {
	return GroupCreateExampleResponse
}

func assertGroupCreateResponse(t *testing.T, response fivetran.GroupCreateResponse) {
	assertEqual(t, response.Code, ExpectedGroupCreateCode)
	assertEqual(t, response.Message, ExpectedGroupCreateMessage)

	assertEqual(t, response.Data.ID, ExpectedGroupCreateID)
	assertEqual(t, response.Data.Name, ExpectedGroupCreateName)

	expectedCreatedAt, _ := time.Parse(time.RFC3339, ExpectedGroupCreateCreatedAt)
	assertEqual(t, response.Data.CreatedAt, expectedCreatedAt)
}

func assertGroupCreateRequest(t *testing.T, request map[string]interface{}) {
	assertKey(t, "name", request, "NewGroup")
}
