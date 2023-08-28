package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

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

func prepareGroupListUsersResponse() string {
	return `{
		"code": "Success",
		"data": {
			"items": [
				{
					"id": "nozzle_eat",
					"email": "john@mycompany.com",
					"given_name": "John",
					"family_name": "White",
					"verified": true,
					"invited": false,
					"picture": null,
					"phone": null,
					"role": null,
					"logged_in_at": "2019-01-03T08:44:45.369Z",
					"created_at": "2018-01-15T11:00:27.329220Z",
					"active": true
				}
			],
			"next_cursor": "eyJza2lwIjoyfQ"
		}
	}`
}

func assertGroupListUsersResponse(t *testing.T, response fivetran.GroupListUsersResponse) {
	createdAt, _ := time.Parse(time.RFC3339Nano, "2018-01-15T11:00:27.329220Z")
	loggedInAt, _ := time.Parse(time.RFC3339Nano, "2019-01-03T08:44:45.369Z")

	assertEqual(t, response.Code, "Success")
	assertEqual(t, len(response.Data.Items), 1)
	item := response.Data.Items[0]

	assertEqual(t, item.ID, "nozzle_eat")
	assertEqual(t, item.Email, "john@mycompany.com")
	assertEqual(t, item.GivenName, "John")
	assertEqual(t, item.FamilyName, "White")
	assertEqual(t, *item.Verified, true)
	assertEqual(t, *item.Invited, false)
	assertEqual(t, item.Picture, "")
	assertEqual(t, item.Phone, "")
	assertEqual(t, item.Role, "")
	assertEqual(t, item.LoggedInAt, loggedInAt)
	assertEqual(t, item.CreatedAt, createdAt)
}
