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

const (
	TRANSFORMATION_ID = "123"
	STATUS            = "SUCCEEDED"
	DBT_MODEL_ID      = "model_1"
	RUN_TESTS         = false
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

var daysOfWeek = []string{
	"Monday",
}

func TestNewDbtTransformationCreateFullMappingMock(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/dbt/transformations").
		ThenCall(

			func(req *http.Request) (*http.Response, error) {
				body := requestBodyToJson(t, req)
				assertDbtTransformationRequest(t, body)
				response := mock.NewResponse(req, http.StatusCreated, prepareDbtTransformationResponse())
				return response, nil
			})

	// act
	response, err := ftClient.NewDbtTransformationCreateService().
		DbtModelId(DBT_MODEL_ID).
		RunTests(RUN_TESTS).
		ProjectId(DBT_PROJECT_ID).
		Schedule(prepareSchedule()).
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

func prepareDbtTransformationResponse() string {
	return fmt.Sprintf(
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
		CONNECTOR_ID,
	)
}

func prepareSchedule() *fivetran.DbtTransformationSchedule {
	schedule := fivetran.NewDbtTransformationSchedule().
		ScheduleType(SCHEDULE_TYPE).
		DaysOfWeek(daysOfWeek).
		Interval(INTERVAL).
		TimeOfDay(TIME_OF_DAY)
	return schedule
}

func assertDbtTransformationRequest(t *testing.T, request map[string]interface{}) {
	assertKey(t, "dbt_model_id", request, DBT_MODEL_ID)
	assertKey(t, "run_tests", request, RUN_TESTS)
	assertKey(t, "project_id", request, DBT_PROJECT_ID)

	c, ok := request["schedule"]
	assertEqual(t, ok, true)
	schedule, ok := c.(map[string]interface{})
	assertEqual(t, ok, true)

	assertKey(t, "schedule_type", schedule, SCHEDULE_TYPE)

	daysOfWeek, ok := schedule["days_of_week"]
	assertEqual(t, ok, true)
	assertEqual(t, daysOfWeek, daysOfWeek)
	assertKey(t, "interval", schedule, float64(INTERVAL))
	assertKey(t, "time_of_day", schedule, TIME_OF_DAY)
}

func assertDbtTransformationResponse(t *testing.T, response fivetran.DbtTransformationCreateResponse) {

	assertEqual(t, response.Code, "Created")
	assertNotEmpty(t, response.Message)

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
