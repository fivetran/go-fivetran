package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/common"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestGroupRemoveUserServiceDo(t *testing.T) {
	// arrange
	expectedGroupID := "projected_sickle"
	expectedUserID := "nozzle_eat"

	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodDelete, "/v1/groups/"+expectedGroupID+"/users/"+expectedUserID).
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareGroupRemoveUserResponse(expectedUserID))
			return response, nil
		})

	// act
	response, err := ftClient.NewGroupRemoveUser().
		GroupID(expectedGroupID).
		UserID(expectedUserID).
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)
	assertGroupRemoveUserResponse(t, response, expectedUserID)
}

func prepareGroupRemoveUserResponse(expectedUserID string) string {
	return fmt.Sprintf(`{
		"code": "Success",
		"message": "User with id '%v' has been removed from the group"
	}`, expectedUserID)
}

func assertGroupRemoveUserResponse(t *testing.T, response common.CommonResponse, expectedUserID string) {
	assertEqual(t, response.Code, "Success")
	expectedMessage := fmt.Sprintf("User with id '%v' has been removed from the group", expectedUserID)
	assertEqual(t, response.Message, expectedMessage)
}
