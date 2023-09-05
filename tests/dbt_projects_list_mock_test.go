package tests

import (
	"context"
	"net/http"
	"strconv"
	"testing"

	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestDbtModelsListService(t *testing.T) {
	// arrange
	limit := 2
	cursor := "some_cursor"
	nextCursor := "next_cursor"

	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/dbt/projects").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			var query = req.URL.Query()
			assertEqual(t, query.Get("cursor"), cursor)
			assertEqual(t, query.Get("limit"), strconv.Itoa(limit))
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
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)
	assertEqual(t, response.Code, "Success")
	assertEqual(t, response.Data.NextCursor, nextCursor)
	assertEqual(t, len(response.Data.Items), 2)
	assertEqual(t, response.Data.Items[0].ID, "dbt_project_id")
	assertEqual(t, response.Data.Items[0].GroupId, "group_id")
	assertEqual(t, response.Data.Items[0].CreatedAt, "created_at")
	assertEqual(t, response.Data.Items[0].CreatedById, "created_by_id")
	assertEqual(t, response.Data.Items[1].ID, "dbt_project_id_2")
	assertEqual(t, response.Data.Items[1].GroupId, "group_id_2")
	assertEqual(t, response.Data.Items[1].CreatedAt, "created_at_2")
	assertEqual(t, response.Data.Items[1].CreatedById, "created_by_id_2")
}
