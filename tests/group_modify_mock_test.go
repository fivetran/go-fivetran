package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

const (
	GROUP_MODIFY_GROUP_ID            = "decent_dropsy"
	GROUP_MODIFY_EXPECTED_GROUP_NAME = "New_Group_Name"
	GROUP_MODIFY_CREATED_TIME        = "2020-05-25T15:26:47.306509Z"
)

func TestGroupModifyServiceDo(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/groups/"+ExpectedGroupID).
		ThenCall(func(req *http.Request) (*http.Response, error) {
			body := requestBodyToJson(t, req)
			assertGroupModifyRequest(t, body)
			response := mock.NewResponse(req, http.StatusOK, prepareGroupModifyResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewGroupModify().
		GroupID(ExpectedGroupID).
		Name(GROUP_MODIFY_EXPECTED_GROUP_NAME).
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)
	assertGroupModifyResponse(t, response)
}

func prepareGroupModifyResponse() string {
	return fmt.Sprintf(`{
		"code": "Success",
		"message": "Group has been updated",
		"data": {
			"id": "%s",
			"name": "%s",
			"created_at": "%s"
		}
	}`,
		GROUP_MODIFY_GROUP_ID,
		GROUP_MODIFY_EXPECTED_GROUP_NAME,
		GROUP_MODIFY_CREATED_TIME)
}

func assertGroupModifyRequest(t *testing.T, request map[string]interface{}) {
	assertKey(t, "name", request, GROUP_MODIFY_EXPECTED_GROUP_NAME)
}

func assertGroupModifyResponse(t *testing.T, response fivetran.GroupModifyResponse) {
	assertEqual(t, response.Code, "Success")
	assertEqual(t, response.Message, "Group has been updated")
	assertEqual(t, response.Data.ID, GROUP_MODIFY_GROUP_ID)
	assertEqual(t, response.Data.Name, GROUP_MODIFY_EXPECTED_GROUP_NAME)
	assertTimeEqual(t, response.Data.CreatedAt, GROUP_MODIFY_CREATED_TIME)
}
