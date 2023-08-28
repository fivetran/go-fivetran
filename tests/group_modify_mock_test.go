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
	GroupID           = "decent_dropsy"
	ExpectedGroupName = "New_Group_Name"
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
		Name(ExpectedGroupName).
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
			"created_at": "2020-05-25T15:26:47.306509Z"
		}
	}`, GroupID, ExpectedGroupName)
}

func assertGroupModifyRequest(t *testing.T, request map[string]interface{}) {
	assertKey(t, "name", request, ExpectedGroupName)
}

func assertGroupModifyResponse(t *testing.T, response fivetran.GroupModifyResponse) {
	assertEqual(t, response.Code, "Success")
	assertEqual(t, response.Message, "Group has been updated")
	assertEqual(t, response.Data.ID, GroupID)
	assertEqual(t, response.Data.Name, ExpectedGroupName)
}
