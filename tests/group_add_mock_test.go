package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/common"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestGroupAddUserServiceDo(t *testing.T) {
	// arrange
	groupID := "projected_sickle"
	email := "john.white@mycompany.com"
	role := "Account Administrator"

	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodPost, fmt.Sprintf("/v1/groups/%s/users", groupID)).
		ThenCall(func(req *http.Request) (*http.Response, error) {
			requestBody := RequestBodyToJson(t, req)
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
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)
	assertGroupAddUserResponse(t, response)
}

func assertGroupAddUserRequest(t *testing.T,
	body map[string]interface{},
	expectedEmail string,
	expectedRole string) {
	assertKey(t, "email", body, expectedEmail)
	assertKey(t, "role", body, expectedRole)
}

func assertGroupAddUserResponse(t *testing.T, response common.CommonResponse) {
	assertEqual(t, response.Code, "Success")
	assertEqual(t, response.Message, "User has been added to the group")
}
