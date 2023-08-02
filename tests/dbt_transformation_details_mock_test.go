package tests

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestDbtTransformationDetailsService(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/dbt/transformations/"+TRANSFORMATION_ID).
		ThenCall(
			func(req *http.Request) (*http.Response, error) {
				response := mock.NewResponse(req, http.StatusOK, prepareDbtTransformationDetailsResponse())
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
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)

	assertDbtTransformationDetailsResponse(t, response)
}

func prepareDbtTransformationDetailsResponse() string {
	return fmt.Sprintf(
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
					"interval": "%v",
					"time_of_day": "%v"
				},
				"last_run": "%v",
				"output_model_name": "%v",
				"dbt_project_id": "%v",
				"dbt_model_id": "%v",
				"next_run": "%v",
				"created_at": "%v",
				"run_tests": "%v",
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
		MODEL_ID,
		CONNECTOR_ID)
}

func assertDbtTransformationDetailsResponse(t *testing.T, response fivetran.DbtTransformationDetailsResponse) {
	assertEqual(t, response.Code, "Success")

	assertEqual(t, response.Data.ID, TRANSFORMATION_ID)
	assertEqual(t, response.Data.Status, STATUS)
	assertEqual(t, response.Data.LastRun, CREATED_AT)
	assertEqual(t, response.Data.OutputModelName, OUTPUT_MODEL_NAME)
	assertEqual(t, response.Data.DbtProjectId, DBT_PROJECT_ID)
	assertEqual(t, response.Data.DbtModelId, DBT_MODEL_ID)
	assertEqual(t, response.Data.NextRun, NEXT_RUN)
	assertEqual(t, response.Data.CreatedAt, CREATED_AT)
	assertEqual(t, response.Data.RunTests, boolToStr(RUN_TESTS))
	assertEqual(t, response.Data.ModelIds[0], MODEL_ID)
	assertEqual(t, response.Data.ConnectorIds[0], CONNECTOR_ID)

	assertEqual(t, response.Data.Schedule.ScheduleType, SCHEDULE_TYPE)
	assertEqual(t, response.Data.Schedule.DaysOfWeek, daysOfWeek)
	assertEqual(t, response.Data.Schedule.Interval, strconv.Itoa(INTERVAL))
	assertEqual(t, response.Data.Schedule.TimeOfDay, TIME_OF_DAY)
}
