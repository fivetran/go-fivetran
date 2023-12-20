package dbt_test

import (
    "context"
    "net/http"
    "testing"

    
    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestDbtModelDetailsService(t *testing.T) {
	// arrange
	modelId := "model_id"

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/dbt/models/"+modelId).
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, `{
				"code": "Success",
				"data": {
					"id": "model_id",
					"model_name": "model_name",
					"scheduled": true
				}
			}`)
			return response, nil
		})

	// act
	response, err := ftClient.NewDbtModelDetails().
		ModelId(modelId).
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
	testutils.AssertEqual(t, response.Data.ID, modelId)
	testutils.AssertEqual(t, response.Data.ModelName, "model_name")
	testutils.AssertEqual(t, response.Data.Scheduled, true)
}
