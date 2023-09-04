package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestGroupListUsersServiceDo(t *testing.T) {
	// arrange
	groupID := "projected_sickle"
	limit := 10
	cursor := "eyJza2lwIjoxfQ"

	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodGet, fmt.Sprintf("/v1/groups/%s/users", groupID)).
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareGroupListUsersResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewGroupListUsers().
		GroupID(groupID).
		Limit(limit).
		Cursor(cursor).
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)
	assertGroupListUsersResponse(t, response)
}

const (
	GROUP_LIST_USER_ID           = "nozzle_eat"
	GROUP_LIST_USER_EMAIL        = "john@mycompany.com"
	GROUP_LIST_USER_GIVEN_NAME   = "John"
	GROUP_LIST_USER_FAMILY_NAME  = "White"
	GROUP_LIST_USER_VERIFIED     = true
	GROUP_LIST_USER_INVITED      = false
	GROUP_LIST_USER_PICTURE      = ""
	GROUP_LIST_USER_PHONE        = ""
	GROUP_LIST_USER_ROLE         = ""
	GROUP_LIST_USER_LOGGED_IN_AT = "2019-01-03T08:44:45.369Z"
	GROUP_LIST_USER_CREATED_AT   = "2018-01-15T11:00:27.329220Z"
	GROUP_LIST_USER_ACTIVE       = true
)

func prepareGroupListUsersResponse() string {
	value := fmt.Sprintf(`{
		"code": "Success",
		"data": {
			"items": [
				{
					"id": "%v",
					"email": "%v",
					"given_name": "%v",
					"family_name": "%v",
					"verified": %t,
					"invited": %t,
					"picture": "%v",
					"phone": "%v",
					"role": "%v",
					"logged_in_at": "%v",
					"created_at": "%v",
					"active": %t
				}
			],
			"next_cursor": "eyJza2lwIjoyfQ"
		}
	}`,
		GROUP_LIST_USER_ID,
		GROUP_LIST_USER_EMAIL,
		GROUP_LIST_USER_GIVEN_NAME,
		GROUP_LIST_USER_FAMILY_NAME,
		GROUP_LIST_USER_VERIFIED,
		GROUP_LIST_USER_INVITED,
		GROUP_LIST_USER_PICTURE,
		GROUP_LIST_USER_PHONE,
		GROUP_LIST_USER_ROLE,
		GROUP_LIST_USER_LOGGED_IN_AT,
		GROUP_LIST_USER_CREATED_AT,
		GROUP_LIST_USER_ACTIVE)
	return value
}

func assertGroupListUsersResponse(t *testing.T, response fivetran.GroupListUsersResponse) {

	assertEqual(t, response.Code, "Success")
	assertEqual(t, len(response.Data.Items), 1)
	item := response.Data.Items[0]

	assertEqual(t, item.ID, GROUP_LIST_USER_ID)
	assertEqual(t, item.Email, GROUP_LIST_USER_EMAIL)
	assertEqual(t, item.GivenName, GROUP_LIST_USER_GIVEN_NAME)
	assertEqual(t, item.FamilyName, GROUP_LIST_USER_FAMILY_NAME)
	assertEqual(t, *item.Verified, GROUP_LIST_USER_VERIFIED)
	assertEqual(t, *item.Invited, GROUP_LIST_USER_INVITED)
	assertEqual(t, item.Picture, GROUP_LIST_USER_PICTURE)
	assertEqual(t, item.Phone, GROUP_LIST_USER_PHONE)
	assertEqual(t, item.Role, GROUP_LIST_USER_ROLE)
	assertTimeEqual(t, item.LoggedInAt, GROUP_LIST_USER_LOGGED_IN_AT)
	assertTimeEqual(t, item.CreatedAt, GROUP_LIST_USER_CREATED_AT)
	assertEqual(t, *item.Active, GROUP_LIST_USER_ACTIVE)
}
