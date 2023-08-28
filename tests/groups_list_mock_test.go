package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestGroupsListService_Do(t *testing.T) {
	// arrange
	limit := 10
	cursor := "some_cursor"

	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/groups").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareGroupsListResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewGroupsList().
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
	assertGroupsListResponse(t, response)
}

func prepareGroupsListResponse() string {
	return `{
		"code": "Success",
		"data": {
			"items": [
				{
					"id": "projected_sickle",
					"name": "Staging",
					"created_at": "2018-12-20T11:59:35.089589Z"
				},
				{
					"id": "schoolmaster_heedless",
					"name": "Production",
					"created_at": "2019-01-08T19:53:52.185146Z"
				}
			],
			"next_cursor": "eyJza2lwIjoyfQ"
		}
	}`
}

func assertGroupsListResponse(t *testing.T, response fivetran.GroupsListResponse) {
	assertEqual(t, response.Code, "Success")

	assertEqual(t, len(response.Data.Items), 2)
	assertEqual(t, response.Data.Items[0].ID, "projected_sickle")
	assertEqual(t, response.Data.Items[0].Name, "Staging")

	assertEqual(t, response.Data.Items[1].ID, "schoolmaster_heedless")
	assertEqual(t, response.Data.Items[1].Name, "Production")

	assertEqual(t, response.Data.NextCursor, "eyJza2lwIjoyfQ")
}
