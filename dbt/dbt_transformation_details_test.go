package dbt_test

import (
    "context"
    "fmt"
    "net/http"
    "testing"

	"github.com/fivetran/go-fivetran/dbt"
	
    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

var transformtationDetailsResponse = fmt.Sprintf(
	`{
		"code": "Success",
		"data": {
			"id": "%v",
			"status": "%v",
			"schedule": {
				"schedule_type": "%v",
				"days_of_week": [
					"%v"
				],
				"interval": %v,
				"time_of_day": "%v"
			},
			"last_run": "%v",
			"output_model_name": "%v",
			"dbt_project_id": "%v",
			"dbt_model_id": "%v",
			"next_run": "%v",
			"created_at": "%v",
			"run_tests": %v,
			"paused": %v,
			"model_ids": [
				"%v"
			],
			"connector_ids": [
				"%v"
			]
			}
		}`,
	TRANSFORMATION_ID,
	STATUS,
	SCHEDULE_TYPE,
	daysOfWeek[0],
	INTERVAL,
	TIME_OF_DAY,
	CREATED_AT,
	OUTPUT_MODEL_NAME,
	DBT_PROJECT_ID,
	DBT_MODEL_ID,
	NEXT_RUN,
	CREATED_AT,
	RUN_TESTS,
	PAUSED,
	MODEL_ID,
	CONNECTOR_ID)

func TestDbtTransformationDetailsService(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/dbt/transformations/"+TRANSFORMATION_ID).
		ThenCall(
			func(req *http.Request) (*http.Response, error) {
				response := mock.NewResponse(req, http.StatusOK, transformtationDetailsResponse)
				return response, nil
			})

	service := ftClient.NewDbtTransformationDetailsService().TransformationId(TRANSFORMATION_ID)

	// act
	response, err := service.Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	// assert
	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)

	assertDbtTransformationDetailsResponse(t, response)
}

func assertDbtTransformationDetailsResponse(t *testing.T, response dbt.DbtTransformationResponse) {
	testutils.AssertEqual(t, response.Code, "Success")

	testutils.AssertEqual(t, response.Data.ID, TRANSFORMATION_ID)
	testutils.AssertEqual(t, response.Data.Status, STATUS)
	testutils.AssertEqual(t, response.Data.LastRun, CREATED_AT)
	testutils.AssertEqual(t, response.Data.OutputModelName, OUTPUT_MODEL_NAME)
	testutils.AssertEqual(t, response.Data.DbtProjectId, DBT_PROJECT_ID)
	testutils.AssertEqual(t, response.Data.DbtModelId, DBT_MODEL_ID)
	testutils.AssertEqual(t, response.Data.NextRun, NEXT_RUN)
	testutils.AssertEqual(t, response.Data.CreatedAt, CREATED_AT)
	testutils.AssertEqual(t, response.Data.ModelIds[0], MODEL_ID)
	testutils.AssertEqual(t, response.Data.ConnectorIds[0], CONNECTOR_ID)
	testutils.AssertEqual(t, response.Data.RunTests, RUN_TESTS)
	testutils.AssertEqual(t, response.Data.Paused, PAUSED)
	testutils.AssertEqual(t, response.Data.Schedule.ScheduleType, SCHEDULE_TYPE)
	testutils.AssertEqual(t, response.Data.Schedule.DaysOfWeek, daysOfWeek)
	testutils.AssertEqual(t, response.Data.Schedule.Interval, INTERVAL)
	testutils.AssertEqual(t, response.Data.Schedule.TimeOfDay, TIME_OF_DAY)
}
