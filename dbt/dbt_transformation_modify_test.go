package dbt_test

import (
    "context"
    "fmt"
    "net/http"
    "testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/dbt"
	
    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
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
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/dbt/transformations/"+TRANSFORMATION_ID).
		ThenCall(
			func(req *http.Request) (*http.Response, error) {
				body := testutils.RequestBodyToJson(t, req)
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
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)

	assertDbtTransformationModifyResponse(t, response)
}

func assertTransformationUpdateRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "run_tests", request, NEW_RUN_TESTS)
	testutils.AssertKey(t, "paused", request, NEW_PAUSED)
	testutils.AssertHasKey(t, request, "schedule")

	schedule := request["schedule"].(map[string]interface{})

	testutils.AssertKey(t, "schedule_type", schedule, NEW_SCHEDULE_TYPE)
	testutils.AssertKey(t, "interval", schedule, float64(NEW_INTERVAL))
	testutils.AssertKey(t, "days_of_week", schedule, []interface{}{"Tuesday"})
	testutils.AssertKey(t, "time_of_day", schedule, NEW_TIME_OF_DAY)
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
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Message, "Dbt transformation has been updated")

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
	testutils.AssertEqual(t, response.Data.RunTests, NEW_RUN_TESTS)
	testutils.AssertEqual(t, response.Data.Paused, NEW_PAUSED)
	testutils.AssertEqual(t, response.Data.Schedule.ScheduleType, NEW_SCHEDULE_TYPE)
	testutils.AssertEqual(t, response.Data.Schedule.DaysOfWeek, newDaysOfWeek)
	testutils.AssertEqual(t, response.Data.Schedule.Interval, NEW_INTERVAL)
	testutils.AssertEqual(t, response.Data.Schedule.TimeOfDay, NEW_TIME_OF_DAY)
}
