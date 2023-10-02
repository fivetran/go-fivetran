package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/dbt"
	"github.com/fivetran/go-fivetran/tests/mock"
)

const (
	NEW_RUN_TESTS     = true
	NEW_PAUSED        = false
	NEW_SCHEDULE_TYPE = "schedule_type_2"
	NEW_INTERVAL      = 0
	NEW_TIME_OF_DAY   = "12:00"
)

var newDaysOfWeek = []string{
	"Tuesday",
}

func TestDbtTransformationModifyService(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/dbt/transformations/"+TRANSFORMATION_ID).
		ThenCall(
			func(req *http.Request) (*http.Response, error) {
				body := requestBodyToJson(t, req)
				assertTransformationUpdateRequest(t, body)
				response := mock.NewResponse(req, http.StatusOK, prepareDbtTransformationModifyResponse())
				return response, nil
			})

	newSchedule := fivetran.NewDbtTransformationSchedule().
		ScheduleType(NEW_SCHEDULE_TYPE).
		DaysOfWeek(newDaysOfWeek).
		Interval(NEW_INTERVAL).
		TimeOfDay(NEW_TIME_OF_DAY)

	service := ftClient.NewDbtTransformationModifyService().
		DbtTransformationId(TRANSFORMATION_ID).
		RunTests(NEW_RUN_TESTS).
		Paused(NEW_PAUSED).
		Schedule(newSchedule)

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

	assertDbtTransformationModifyResponse(t, response)
}

func assertTransformationUpdateRequest(t *testing.T, request map[string]interface{}) {
	assertKey(t, "run_tests", request, NEW_RUN_TESTS)
	assertKey(t, "paused", request, NEW_PAUSED)
	assertHasKey(t, request, "schedule")

	schedule := request["schedule"].(map[string]interface{})

	assertKey(t, "schedule_type", schedule, NEW_SCHEDULE_TYPE)
	assertKey(t, "interval", schedule, float64(NEW_INTERVAL))
	assertKey(t, "days_of_week", schedule, []interface{}{"Tuesday"})
	assertKey(t, "time_of_day", schedule, NEW_TIME_OF_DAY)
}

func prepareDbtTransformationModifyResponse() string {
	return fmt.Sprintf(
		`{
			"code": "Success",
			"message": "Dbt transformation has been updated",
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
		NEW_SCHEDULE_TYPE,
		newDaysOfWeek[0],
		NEW_INTERVAL,
		NEW_TIME_OF_DAY,
		CREATED_AT,
		OUTPUT_MODEL_NAME,
		DBT_PROJECT_ID,
		DBT_MODEL_ID,
		NEXT_RUN,
		CREATED_AT,
		NEW_RUN_TESTS,
		NEW_PAUSED,
		MODEL_ID,
		CONNECTOR_ID,
	)
}

func assertDbtTransformationModifyResponse(t *testing.T, response dbt.DbtTransformationResponse) {
	assertEqual(t, response.Code, "Success")
	assertEqual(t, response.Message, "Dbt transformation has been updated")

	assertEqual(t, response.Data.ID, TRANSFORMATION_ID)
	assertEqual(t, response.Data.Status, STATUS)
	assertEqual(t, response.Data.LastRun, CREATED_AT)
	assertEqual(t, response.Data.OutputModelName, OUTPUT_MODEL_NAME)
	assertEqual(t, response.Data.DbtProjectId, DBT_PROJECT_ID)
	assertEqual(t, response.Data.DbtModelId, DBT_MODEL_ID)
	assertEqual(t, response.Data.NextRun, NEXT_RUN)
	assertEqual(t, response.Data.CreatedAt, CREATED_AT)
	assertEqual(t, response.Data.ModelIds[0], MODEL_ID)
	assertEqual(t, response.Data.ConnectorIds[0], CONNECTOR_ID)
	assertEqual(t, response.Data.RunTests, NEW_RUN_TESTS)
	assertEqual(t, response.Data.Paused, NEW_PAUSED)
	assertEqual(t, response.Data.Schedule.ScheduleType, NEW_SCHEDULE_TYPE)
	assertEqual(t, response.Data.Schedule.DaysOfWeek, newDaysOfWeek)
	assertEqual(t, response.Data.Schedule.Interval, NEW_INTERVAL)
	assertEqual(t, response.Data.Schedule.TimeOfDay, NEW_TIME_OF_DAY)
}
