package groups_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/tests"
	"github.com/fivetran/go-fivetran/tests/mock"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestGroupRemoveUserServiceDo(t *testing.T) {
	// arrange
	expectedGroupID := "projected_sickle"
	expectedUserID := "nozzle_eat"

	ftClient, mockClient := tests.CreateTestClient()
	handler := mockClient.When(http.MethodDelete, "/v1/groups/"+expectedGroupID+"/users/"+expectedUserID).
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, fmt.Sprintf(`
			{
				"code": "Success",
				"message": "User with id '%v' has been removed from the group"
			}`,
				expectedUserID))
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
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Message, fmt.Sprintf("User with id '%v' has been removed from the group", expectedUserID))
}
