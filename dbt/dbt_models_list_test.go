package dbt_test

import (
    "context"
    "net/http"
    "testing"
    "strconv"

    
    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestDbtModelsListService(t *testing.T) {
	// arrange
	limit := 2
	cursor := "some_cursor"
	nextCursor := "next_cursor"
	projectId := "dbt_project_id"

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/dbt/models").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			var query = req.URL.Query()
			testutils.AssertEqual(t, query.Get("project_id"), projectId)
			testutils.AssertEqual(t, query.Get("cursor"), cursor)
			testutils.AssertEqual(t, query.Get("limit"), strconv.Itoa(limit))
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
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.NextCursor, nextCursor)
	testutils.AssertEqual(t, len(response.Data.Items), 2)
	testutils.AssertEqual(t, response.Data.Items[0].ID, "model_id")
	testutils.AssertEqual(t, response.Data.Items[0].ModelName, "model_name")
	testutils.AssertEqual(t, response.Data.Items[0].Scheduled, true)
	testutils.AssertEqual(t, response.Data.Items[1].ID, "model_id_1")
	testutils.AssertEqual(t, response.Data.Items[1].ModelName, "model_name_1")
	testutils.AssertEqual(t, response.Data.Items[1].Scheduled, false)
}
