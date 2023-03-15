package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestDbtTransformationDetailsMock(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(
		http.MethodGet,
		"/v1/dbt/transformations/_moonbeam_no_schedule_model").
		ThenCall(
			func(req *http.Request) (*http.Response, error) {
				response := mock.NewResponse(
					req,
					http.StatusOK,
					prepareDbtTransformationDetailsResponse())
				return response, nil
			})

	// act
	response, err := ftClient.NewDbtTransformationDetailsService().
		DbtTransformationID("_moonbeam_no_schedule_model").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	// asert
	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)

	assertDbtTransformationDetailsMock(t, response)
}

func assertDbtTransformationDetailsMock(t *testing.T, response fivetran.DbtTransformationDetailsResponse) {

	assertEqual(t, response.Code, "Success")

	assertEqual(t, response.Data.DbtModelId, "_moonbeam_no_schedule_model")
	assertEqual(t, response.Data.OutputModelName, "no_schedule_model")
	assertEqual(t, response.Data.DbtProjectId, "_moonbeam_project")
	assertEqual(t, response.Data.Status, "PENDING")
	assertEqual(t, response.Data.RunTests, false)
}

func prepareDbtTransformationDetailsResponse() string {
	return `{
		"code": "Success",
		"message": "Transformation has been fetched",
		"data": {
			"id": "_moonbeam_no_schedule_model",
			"dbt_model_id": "_moonbeam_no_schedule_model",
			"output_model_name": "no_schedule_model",
			"dbt_project_id": "_moonbeam_project",
			"last_run": "2022-04-29T11:24:41.312868Z",
			"next_run": "2022-04-30T10:00:00.000000Z",
			"status": "PENDING",
			"schedule": {
				"schedule_type": "TIME_OF_DAY",
				"days_of_week": [
					"MONDAY",
					"SUNDAY"
				],
				"days_of_week": null,
				"time_of_day": "10:00"
			},
			"run_tests": false,
			"connector_ids": [
				"_moonbeam_marketo"
			],
			"model_ids": []
		}
	 }`
}
