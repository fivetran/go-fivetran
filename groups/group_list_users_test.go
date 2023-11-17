package groups_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/groups"
	"github.com/fivetran/go-fivetran/tests"
	"github.com/fivetran/go-fivetran/tests/mock"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

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

func TestGroupListUsersServiceDo(t *testing.T) {
	// arrange
	groupID := "projected_sickle"
	limit := 10
	cursor := "eyJza2lwIjoxfQ"

	ftClient, mockClient := tests.CreateTestClient()
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
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	assertGroupListUsersResponse(t, response)
}

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

func assertGroupListUsersResponse(t *testing.T, response groups.GroupListUsersResponse) {

	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, len(response.Data.Items), 1)
	item := response.Data.Items[0]

	testutils.AssertEqual(t, item.ID, GROUP_LIST_USER_ID)
	testutils.AssertEqual(t, item.Email, GROUP_LIST_USER_EMAIL)
	testutils.AssertEqual(t, item.GivenName, GROUP_LIST_USER_GIVEN_NAME)
	testutils.AssertEqual(t, item.FamilyName, GROUP_LIST_USER_FAMILY_NAME)
	testutils.AssertEqual(t, *item.Verified, GROUP_LIST_USER_VERIFIED)
	testutils.AssertEqual(t, *item.Invited, GROUP_LIST_USER_INVITED)
	testutils.AssertEqual(t, item.Picture, GROUP_LIST_USER_PICTURE)
	testutils.AssertEqual(t, item.Phone, GROUP_LIST_USER_PHONE)
	testutils.AssertEqual(t, item.Role, GROUP_LIST_USER_ROLE)
	testutils.AssertTimeEqual(t, item.LoggedInAt, GROUP_LIST_USER_LOGGED_IN_AT)
	testutils.AssertTimeEqual(t, item.CreatedAt, GROUP_LIST_USER_CREATED_AT)
	testutils.AssertEqual(t, *item.Active, GROUP_LIST_USER_ACTIVE)
}
