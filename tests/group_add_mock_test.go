package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
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
			var requestBody fivetran.GroupAddUserRequest
			err := json.NewDecoder(req.Body).Decode(&requestBody)
			if err != nil {
				return nil, err
			}
			assertEqual(t, *requestBody.Email, email)
			assertEqual(t, *requestBody.Role, role)

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

func assertGroupAddUserResponse(t *testing.T, response fivetran.GroupAddUserResponse) {
	assertEqual(t, response.Code, "Success")
	assertEqual(t, response.Message, "User has been added to the group")
}
