package dbt_test

import (
    "context"
    "net/http"
    "testing"
    "strconv"

	
    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestDbtProjectsListService(t *testing.T) {
	// arrange
	limit := 2
	cursor := "some_cursor"
	nextCursor := "next_cursor"

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/dbt/projects").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			var query = req.URL.Query()
			testutils.AssertEqual(t, query.Get("cursor"), cursor)
			testutils.AssertEqual(t, query.Get("limit"), strconv.Itoa(limit))
			response := mock.NewResponse(req, http.StatusOK, `{
				"code": "Success",
				"data": {
					"items": [
						{
							"id": "dbt_project_id",
							"group_id": "group_id",
							"created_at": "created_at",
							"created_by_id": "created_by_id"
						},
						{
							"id": "dbt_project_id_2",
							"group_id": "group_id_2",
							"created_at": "created_at_2",
							"created_by_id": "created_by_id_2"
						}
					],
					"next_cursor": "`+nextCursor+`"
				}
			}`)
			return response, nil
		})

	// act
	response, err := ftClient.NewDbtProjectsList().
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
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.NextCursor, nextCursor)
	testutils.AssertEqual(t, len(response.Data.Items), 2)
	testutils.AssertEqual(t, response.Data.Items[0].ID, "dbt_project_id")
	testutils.AssertEqual(t, response.Data.Items[0].GroupId, "group_id")
	testutils.AssertEqual(t, response.Data.Items[0].CreatedAt, "created_at")
	testutils.AssertEqual(t, response.Data.Items[0].CreatedById, "created_by_id")
	testutils.AssertEqual(t, response.Data.Items[1].ID, "dbt_project_id_2")
	testutils.AssertEqual(t, response.Data.Items[1].GroupId, "group_id_2")
	testutils.AssertEqual(t, response.Data.Items[1].CreatedAt, "created_at_2")
	testutils.AssertEqual(t, response.Data.Items[1].CreatedById, "created_by_id_2")
}
