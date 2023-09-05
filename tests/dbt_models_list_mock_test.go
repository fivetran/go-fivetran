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
	projectId := "dbt_project_id"

	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/dbt/models").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			var query = req.URL.Query()
			assertEqual(t, query.Get("project_id"), projectId)
			assertEqual(t, query.Get("cursor"), cursor)
			assertEqual(t, query.Get("limit"), strconv.Itoa(limit))
			response := mock.NewResponse(req, http.StatusOK, `{
				"code": "Success",
				"data": {
					"items": [
						{
							"id": "model_id",
							"model_name": "model_name",
							"scheduled": true
						},
						{
							"id": "model_id_1",
							"model_name": "model_name_1",
							"scheduled": false
						}
					],
					"next_cursor": "`+nextCursor+`"
				}
			}`)
			return response, nil
		})

	// act
	response, err := ftClient.NewDbtModelsList().
		ProjectId(projectId).
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
	assertEqual(t, response.Data.Items[0].ID, "model_id")
	assertEqual(t, response.Data.Items[0].ModelName, "model_name")
	assertEqual(t, response.Data.Items[0].Scheduled, true)
	assertEqual(t, response.Data.Items[1].ID, "model_id_1")
	assertEqual(t, response.Data.Items[1].ModelName, "model_name_1")
	assertEqual(t, response.Data.Items[1].Scheduled, false)
}
