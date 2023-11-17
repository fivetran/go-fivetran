package groups_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/groups"
	testutils "github.com/fivetran/go-fivetran/test_utils"
	"github.com/fivetran/go-fivetran/tests"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestGroupsListServiceDo(t *testing.T) {
	// arrange
	limit := 10
	cursor := "some_cursor"

	ftClient, mockClient := tests.CreateTestClient()
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
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
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

func assertGroupsListResponse(t *testing.T, response groups.GroupsListResponse) {
	testutils.AssertEqual(t, response.Code, "Success")

	testutils.AssertEqual(t, len(response.Data.Items), 2)
	testutils.AssertEqual(t, response.Data.Items[0].ID, "projected_sickle")
	testutils.AssertEqual(t, response.Data.Items[0].Name, "Staging")

	testutils.AssertEqual(t, response.Data.Items[1].ID, "schoolmaster_heedless")
	testutils.AssertEqual(t, response.Data.Items[1].Name, "Production")

	testutils.AssertEqual(t, response.Data.NextCursor, "eyJza2lwIjoyfQ")
}
