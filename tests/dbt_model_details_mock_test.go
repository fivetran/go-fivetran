package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestDbtModelDetailsService(t *testing.T) {
	// arrange
	modelId := "model_id"

	ftClient, mockClient := CreateTestClient()
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
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)
	assertEqual(t, response.Code, "Success")
	assertEqual(t, response.Data.ID, modelId)
	assertEqual(t, response.Data.ModelName, "model_name")
	assertEqual(t, response.Data.Scheduled, true)
}
