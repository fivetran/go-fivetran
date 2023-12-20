package groups_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/groups"
	
	"github.com/fivetran/go-fivetran/tests/mock"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

const (
	ExpectedGroupCreateCode      = "Success"
	ExpectedGroupCreateMessage   = "Group has been created"
	ExpectedGroupCreateID        = "decent_dropsy"
	ExpectedGroupCreateName      = "Primary_Snowflake"
	ExpectedGroupCreateCreatedAt = "2020-05-25T15:26:47.306509Z"
)

func TestGroupCreateServiceDo(t *testing.T) {
	// Arrange
	ftClient, mockClient := testutils.CreateTestClient()
	groupName := "NewGroup"

	handler := mockClient.When(http.MethodPost, "/v1/groups").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
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
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)

	assertGroupCreateResponse(t, response)
}

func prepareGroupCreateResponse() string {
	return fmt.Sprintf(`{
		"code": "%v",
		"message": "%v",
		"data": {
			"id": "%v",
			"name": "%v",
			"created_at": "%v"
		}
	}`,
		ExpectedGroupCreateCode,
		ExpectedGroupCreateMessage,
		ExpectedGroupCreateID,
		ExpectedGroupCreateName,
		ExpectedGroupCreateCreatedAt,
	)
}

func assertGroupCreateResponse(t *testing.T, response groups.GroupDetailsResponse) {
	testutils.AssertEqual(t, response.Code, ExpectedGroupCreateCode)
	testutils.AssertEqual(t, response.Message, ExpectedGroupCreateMessage)

	testutils.AssertEqual(t, response.Data.ID, ExpectedGroupCreateID)
	testutils.AssertEqual(t, response.Data.Name, ExpectedGroupCreateName)

	testutils.AssertTimeEqual(t, response.Data.CreatedAt, ExpectedGroupCreateCreatedAt)
}

func assertGroupCreateRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "name", request, "NewGroup")
}
