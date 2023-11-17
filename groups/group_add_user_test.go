package groups_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/common"
	"github.com/fivetran/go-fivetran/tests"
	"github.com/fivetran/go-fivetran/tests/mock"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestGroupAddUserServiceDo(t *testing.T) {
	// arrange
	groupID := "projected_sickle"
	email := "john.white@mycompany.com"
	role := "Account Administrator"

	ftClient, mockClient := tests.CreateTestClient()
	handler := mockClient.When(http.MethodPost, fmt.Sprintf("/v1/groups/%s/users", groupID)).
		ThenCall(func(req *http.Request) (*http.Response, error) {
			requestBody := tests.RequestBodyToJson(t, req)
			assertGroupAddUserRequest(t, requestBody, email, role)

			response := mock.NewResponse(req, http.StatusOK, `{
				"code": "Success",
				"message": "User has been added to the group"
			}`)
			return response, nil
		})

	// act
	response, err := ftClient.NewGroupAddUser().
		GroupID(groupID).
		Email(email).
		Role(role).
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	assertGroupAddUserResponse(t, response)
}

func assertGroupAddUserRequest(t *testing.T,
	body map[string]interface{},
	expectedEmail string,
	expectedRole string) {
	testutils.AssertKey(t, "email", body, expectedEmail)
	testutils.AssertKey(t, "role", body, expectedRole)
}

func assertGroupAddUserResponse(t *testing.T, response common.CommonResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Message, "User has been added to the group")
}
