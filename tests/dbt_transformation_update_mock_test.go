package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestDbtTransformationUpdateMock(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(
		http.MethodPatch,
		"/v1/dbt/transformations/_moonbeam_no_schedule_model").
		ThenCall(

			func(req *http.Request) (*http.Response, error) {
				body := requestBodyToJson(t, req)
				assertDbtTransformationUpdateRequest(t, body)
				response := mock.NewResponse(
					req,
					http.StatusOK,
					prepareDbtTransformationUpdateResponse())
				return response, nil
			})

	// act
	response, err := ftClient.NewDbtTransformationModifyService().
		DbtTransformationId("_moonbeam_no_schedule_model").
		RunTests(false).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	// assert
	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)

	assertDbtTransformationUpdateResponse(t, response)
}

func assertDbtTransformationUpdateRequest(t *testing.T, request map[string]interface{}) {
	assertKeyValue(t, request, "run_tests", false)
}

func assertDbtTransformationUpdateResponse(t *testing.T, response fivetran.DbtTransformationModifyResponse) {
	assertEqual(t, response.Code, "Success")

	assertEqual(t, *response.Data.RunTests, false)
}

func prepareDbtTransformationUpdateResponse() string {
	return `{
		"code": "Success",
		"message": "Transformation has been updated",
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
					"FRIDAY"
				],
				"interval": null,
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
