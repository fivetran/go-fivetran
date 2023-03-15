package tests

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestNewDbtTransformationCreateMock(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(
		http.MethodPost,
		"/v1/dbt/projects/_moonbeam_project/transformations").
		ThenCall(

			func(req *http.Request) (*http.Response, error) {
				body := requestBodyToJson(t, req)
				assertDbtTransformationRequest(t, body)
				response := mock.NewResponse(
					req,
					http.StatusCreated,
					prepareDbtTransformationResponse())
				return response, nil
			})

	// act
	response, err := ftClient.NewDbtTransformationCreateService().
		DbtProjectID("_moonbeam_project").
		DbtModelID("_moonbeam_no_schedule_model").
		Schedule(prepareDbtTransformationSchedule()).
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

	assertDbtTransformationResponse(t, response)
}

func assertDbtTransformationResponse(t *testing.T, response fivetran.DbtTransformationCreateResponse) {

	assertEqual(t, response.Code, "Success")
	assertNotEmpty(t, response.Message)

	assertEqual(t, response.Data.DbtModelID, "_moonbeam_no_schedule_model")
}

func assertDbtTransformationRequest(t *testing.T, request map[string]interface{}) {
	assertKeyValue(t, request, "dbt_model_id", "_moonbeam_no_schedule_model")
}

func prepareDbtTransformationSchedule() *fivetran.DbtTransformationSchedule {
	schedule := fivetran.NewDbtTransformationSchedule().
		ScheduleType("TIME_OF_DAY").
		DaysOfWeek([]string{"MONDAY", "SUNDAY"}).
		TimeOfDay(time.Now())
	return schedule
}

func prepareDbtTransformationResponse() string {
	return `{
		"code": "Success",
		"message": "Transformation has been created",
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
