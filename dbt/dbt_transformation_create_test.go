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
	TRANSFORMATION_ID = "123"
	STATUS            = "SUCCEEDED"
	DBT_MODEL_ID      = "model_1"
	RUN_TESTS         = false
	PAUSED            = true
	DBT_PROJECT_ID    = "project_1"
	SCHEDULE_TYPE     = "schedule_type_1"
	INTERVAL          = 1
	TIME_OF_DAY       = "time_of_day_1"
	OUTPUT_MODEL_NAME = "output_model_name_1"
	MODEL_ID          = "model_id_1"
	CONNECTOR_ID      = "connector_id_1"
	CREATED_AT        = "2023-07-24T14:15:22Z"
	NEXT_RUN          = "2023-08-24T14:15:22Z"
)

var (
	daysOfWeek = []string{
		"Monday",
	}
	createResponse = fmt.Sprintf(
		`{
			"code": "Created",
			"message": "Dbt transformation has been created",
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
		CONNECTOR_ID,
	)
)

func TestNewDbtTransformationCreateFullMappingMock(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/dbt/transformations").
		ThenCall(

			func(req *http.Request) (*http.Response, error) {
				body := testutils.RequestBodyToJson(t, req)
				assertDbtTransformationRequest(t, body)
				response := mock.NewResponse(req, http.StatusCreated, createResponse)
				return response, nil
			})

	schedule := fivetran.NewDbtTransformationSchedule().
		ScheduleType(SCHEDULE_TYPE).
		DaysOfWeek(daysOfWeek).
		Interval(INTERVAL).
		TimeOfDay(TIME_OF_DAY)

	service := ftClient.NewDbtTransformationCreateService().
		DbtModelId(DBT_MODEL_ID).
		RunTests(RUN_TESTS).
		Paused(PAUSED).
		Schedule(schedule)

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
	assertDbtTransformationResponse(t, response)
}

func assertDbtTransformationRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "dbt_model_id", request, DBT_MODEL_ID)
	testutils.AssertKey(t, "run_tests", request, RUN_TESTS)
	testutils.AssertKey(t, "paused", request, true)

	c, ok := request["schedule"]
	testutils.AssertEqual(t, ok, true)
	schedule, ok := c.(map[string]interface{})
	testutils.AssertEqual(t, ok, true)

	testutils.AssertKey(t, "schedule_type", schedule, SCHEDULE_TYPE)

	daysOfWeek, ok := schedule["days_of_week"]
	testutils.AssertEqual(t, ok, true)
	testutils.AssertEqual(t, daysOfWeek, daysOfWeek)
	testutils.AssertKey(t, "interval", schedule, float64(INTERVAL))
	testutils.AssertKey(t, "time_of_day", schedule, TIME_OF_DAY)
}

func assertDbtTransformationResponse(t *testing.T, response dbt.DbtTransformationResponse) {

	testutils.AssertEqual(t, response.Code, "Created")
	testutils.AssertNotEmpty(t, response.Message)

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
