package transformations_test

import (
    "context"
    "net/http"
    "testing"

    "github.com/fivetran/go-fivetran"
    "github.com/fivetran/go-fivetran/transformations"
    
    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewTransformationCreateFullMappingMock(t *testing.T) {
  // arrange
  ftClient, mockClient := testutils.CreateTestClient()
  handler := mockClient.When(http.MethodPost, "/v1/transformations").ThenCall(

    func(req *http.Request) (*http.Response, error) {
      body := testutils.RequestBodyToJson(t, req)
      assertTransformationFullRequest(t, body)
      response := mock.NewResponse(req, http.StatusCreated, prepareTransformationResponse())
      return response, nil
    })

  // act
  response, err := ftClient.NewTransformationCreate().
    ProjectType("DBT_GIT").
    Paused(true).
    TransformationConfig(prepareTransformationConfig()).
    TransformationSchedule(prepareTransformationSchedule()).
    Do(context.Background())

  if err != nil {
    t.Logf("%+v\n", response)
    t.Error(err)
  }

  // assert
  interactions := mockClient.Interactions()
  testutils.AssertEqual(t, len(interactions), 1)
  testutils.AssertEqual(t, interactions[0].Handler, handler)
  testutils.AssertEqual(t, handler.Interactions, 1)

  assertTransformationResponse(t, response)
}

func TestNewTransformationCustomMappingMock(t *testing.T) {
  // arrange
  ftClient, mockClient := testutils.CreateTestClient()
  handler := mockClient.When(http.MethodPost, "/v1/transformations").ThenCall(

    func(req *http.Request) (*http.Response, error) {
      body := testutils.RequestBodyToJson(t, req)
      assertTransformationCustomRequest(t, body)
      response := mock.NewResponse(req, http.StatusCreated, prepareTransformationResponse())
      return response, nil
    })

  // act
  response, err := ftClient.NewTransformationCreate().
    ProjectType("DBT_GIT").
    Paused(true).
    TransformationConfigCustom(prepareTransformationCustomMergedConfig()).
    TransformationScheduleCustom(prepareTransformationCustomMergedSchedule()).
    DoCustom(context.Background())

  if err != nil {
    t.Logf("%+v\n", response)
    t.Error(err)
  }

  // assert
  interactions := mockClient.Interactions()
  testutils.AssertEqual(t, len(interactions), 1)
  testutils.AssertEqual(t, interactions[0].Handler, handler)
  testutils.AssertEqual(t, handler.Interactions, 1)

  assertTransformationCustomResponse(t, response)
}

func TestNewTransformationCustomMergedMappingMock(t *testing.T) {
  // arrange
  ftClient, mockClient := testutils.CreateTestClient()
  handler := mockClient.When(http.MethodPost, "/v1/transformations").ThenCall(

    func(req *http.Request) (*http.Response, error) {
      body := testutils.RequestBodyToJson(t, req)
      assertTransformationCustomMergedRequest(t, body)
      response := mock.NewResponse(req, http.StatusCreated, prepareTransformationResponse())
      return response, nil
    })

  // act
  response, err := ftClient.NewTransformationCreate().
    ProjectType("DBT_GIT").
    Paused(true).
    TransformationConfig(prepareTransformationConfig()).
    TransformationConfigCustom(prepareTransformationCustomMergedConfig()).
    TransformationSchedule(prepareTransformationSchedule()).
    TransformationScheduleCustom(prepareTransformationCustomMergedSchedule()).
    DoCustomMerged(context.Background())

  if err != nil {
    t.Logf("%+v\n", response)
    t.Error(err)
  }

  // assert
  interactions := mockClient.Interactions()
  testutils.AssertEqual(t, len(interactions), 1)
  testutils.AssertEqual(t, interactions[0].Handler, handler)
  testutils.AssertEqual(t, handler.Interactions, 1)
  assertTransformationCustomMergedResponse(t, response)
}

func prepareTransformationResponse() string {
  return `{
  "code": "Success",
  "message": "Operation performed.",
  "data": {
    "id": "transformation_id",
    "status": "RUNNING",
    "schedule": {
      "cron": [
        "0 */1 * * *"
      ],
      "interval": 60,
      "smart_syncing": true,
      "connection_ids": [
        "connection_id"
      ],
      "schedule_type": "TIME_OF_DAY",
      "days_of_week": [
        "MONDAY",
        "FRIDAY"
      ],
      "time_of_day": "14:00",
      "fake_field": "unmapped-value"
    },
    "type": "DBT_CORE",
    "paused": true,
    "created_at": "2024-01-02T00:00:00.743708Z",
    "output_model_names": [
      "string"
    ],
    "created_by_id": "user_id",
    "transformation_config": {
      "project_id": "string",
      "name": "string",
      "steps": [
        {
          "name": "string",
          "command": "string"
        }
      ],
      "fake_field": "unmapped-value"
    }
  }
}`
}

func prepareTransformationConfig() *transformations.TransformationConfig {
  config := fivetran.NewTransformationConfig()
  config.ProjectId("string")
  config.Name("string")
  config.Steps([]transformations.TransformationStep{{Name: "string", Command: "string"}})

  return config
}

func prepareTransformationCustomMergedConfig() *map[string]interface{} {
  config := make(map[string]interface{})

  config["fake_field"] = "unmapped-value"

  return &config
}

func prepareTransformationSchedule() *transformations.TransformationSchedule {
  config := fivetran.NewTransformationSchedule()
  config.Cron([]string{"0 */1 * * *"})
  config.ConnectionIds([]string{"connection_ids"})
  config.DaysOfWeek([]string{"MONDAY", "FRIDAY"})
  config.TimeOfDay("14:00")
  config.ScheduleType("TIME_OF_DAY")
  config.Interval(60)
  config.SmartSyncing(true)

  return config
}

func prepareTransformationCustomMergedSchedule() *map[string]interface{} {
  config := make(map[string]interface{})

  config["fake_field"] = "unmapped-value"

  return &config
}

func assertTransformationFullRequest(t *testing.T, request map[string]interface{}) {
    testutils.AssertKey(t, "type", request, "DBT_GIT")
    testutils.AssertKey(t, "paused", request, true)

    schedule, ok := request["schedule"].(map[string]interface{})
    testutils.AssertEqual(t, ok, true)
    testutils.AssertHasKey(t, schedule, "cron")
    testutils.AssertHasKey(t, schedule, "connection_ids")
    testutils.AssertHasKey(t, schedule, "days_of_week")
    testutils.AssertKey(t, "time_of_day", schedule, "14:00")
    testutils.AssertKey(t, "schedule_type", schedule, "TIME_OF_DAY")
    testutils.AssertKey(t, "interval", schedule, float64(60))
    testutils.AssertKey(t, "smart_syncing", schedule, true)

    config, ok := request["transformation_config"].(map[string]interface{})
    testutils.AssertEqual(t, ok, true)
    testutils.AssertKey(t, "project_id", config, "string")
    testutils.AssertKey(t, "name", config, "string")
    testutils.AssertHasKey(t, config, "steps")
}

func assertTransformationCustomRequest(t *testing.T, request map[string]interface{}) {
    testutils.AssertKey(t, "type", request, "DBT_GIT")
    testutils.AssertKey(t, "paused", request, true)

    config, ok := request["transformation_config"].(map[string]interface{})
    testutils.AssertEqual(t, ok, true)
    testutils.AssertKey(t, "fake_field", config, "unmapped-value")

    schedule, ok := request["schedule"].(map[string]interface{})
    testutils.AssertEqual(t, ok, true)
    testutils.AssertKey(t, "fake_field", schedule, "unmapped-value")
}

func assertTransformationCustomMergedRequest(t *testing.T, request map[string]interface{}) {
    testutils.AssertKey(t, "type", request, "DBT_GIT")
    testutils.AssertKey(t, "paused", request, true)

    schedule, ok := request["schedule"].(map[string]interface{})
    testutils.AssertEqual(t, ok, true)
    testutils.AssertHasKey(t, schedule, "cron")
    testutils.AssertHasKey(t, schedule, "connection_ids")
    testutils.AssertHasKey(t, schedule, "days_of_week")
    testutils.AssertKey(t, "time_of_day", schedule, "14:00")
    testutils.AssertKey(t, "schedule_type", schedule, "TIME_OF_DAY")
    testutils.AssertKey(t, "interval", schedule, float64(60))
    testutils.AssertKey(t, "smart_syncing", schedule, true)
    testutils.AssertKey(t, "fake_field", schedule, "unmapped-value")

    config, ok := request["transformation_config"].(map[string]interface{})
    testutils.AssertEqual(t, ok, true)
    testutils.AssertKey(t, "project_id", config, "string")
    testutils.AssertKey(t, "name", config, "string")
    testutils.AssertHasKey(t, config, "steps")
    testutils.AssertKey(t, "fake_field", config, "unmapped-value")
}

func assertTransformationResponse(t *testing.T, response transformations.TransformationResponse) {
    testutils.AssertEqual(t, response.Code, "Success")
    testutils.AssertNotEmpty(t, response.Message)
    testutils.AssertEqual(t, response.Data.Id, "transformation_id")
    testutils.AssertEqual(t, response.Data.ProjectType, "DBT_CORE")
    testutils.AssertEqual(t, response.Data.CreatedAt, "2024-01-02T00:00:00.743708Z")
    testutils.AssertEqual(t, response.Data.CreatedById, "user_id")
    testutils.AssertEqual(t, response.Data.OutputModelNames[0], "string")
    testutils.AssertEqual(t, response.Data.Paused, true)
    testutils.AssertEqual(t, response.Data.Status, "RUNNING")

    testutils.AssertEqual(t, response.Data.TransformationSchedule.Cron[0], "0 */1 * * *")
    testutils.AssertEqual(t, response.Data.TransformationSchedule.ConnectionIds[0], "connection_id")
    testutils.AssertEqual(t, response.Data.TransformationSchedule.DaysOfWeek[0], "MONDAY")
    testutils.AssertEqual(t, response.Data.TransformationSchedule.DaysOfWeek[1], "FRIDAY")
    testutils.AssertEqual(t, response.Data.TransformationSchedule.TimeOfDay, "14:00")
    testutils.AssertEqual(t, response.Data.TransformationSchedule.ScheduleType, "TIME_OF_DAY")
    testutils.AssertEqual(t, response.Data.TransformationSchedule.Interval, 60)
    testutils.AssertEqual(t, response.Data.TransformationSchedule.SmartSyncing, true)

    testutils.AssertEqual(t, response.Data.TransformationConfig.ProjectId, "string")
    testutils.AssertEqual(t, response.Data.TransformationConfig.Name, "string")
    testutils.AssertEqual(t, response.Data.TransformationConfig.Steps[0].Name, "string")
    testutils.AssertEqual(t, response.Data.TransformationConfig.Steps[0].Command, "string")
}

func assertTransformationCustomResponse(t *testing.T, response transformations.TransformationCustomResponse) {
    testutils.AssertEqual(t, response.Code, "Success")
    testutils.AssertNotEmpty(t, response.Message)
    testutils.AssertEqual(t, response.Data.Id, "transformation_id")
    testutils.AssertEqual(t, response.Data.ProjectType, "DBT_CORE")
    testutils.AssertEqual(t, response.Data.CreatedAt, "2024-01-02T00:00:00.743708Z")
    testutils.AssertEqual(t, response.Data.CreatedById, "user_id")
    testutils.AssertEqual(t, response.Data.OutputModelNames[0], "string")
    testutils.AssertEqual(t, response.Data.Paused, true)
    testutils.AssertEqual(t, response.Data.Status, "RUNNING")

    testutils.AssertHasKey(t, response.Data.TransformationSchedule, "cron")
    testutils.AssertHasKey(t, response.Data.TransformationSchedule, "connection_ids")
    testutils.AssertHasKey(t, response.Data.TransformationSchedule, "days_of_week")
    testutils.AssertEqual(t, response.Data.TransformationSchedule["time_of_day"], "14:00")
    testutils.AssertEqual(t, response.Data.TransformationSchedule["schedule_type"], "TIME_OF_DAY")
    testutils.AssertEqual(t, response.Data.TransformationSchedule["interval"], float64(60))
    testutils.AssertEqual(t, response.Data.TransformationSchedule["smart_syncing"], true)

    testutils.AssertEqual(t, response.Data.TransformationConfig["project_id"], "string")
    testutils.AssertEqual(t, response.Data.TransformationConfig["name"], "string")
    testutils.AssertHasKey(t, response.Data.TransformationConfig, "steps")
}

func assertTransformationCustomMergedResponse(t *testing.T, response transformations.TransformationCustomMergedResponse) {
    testutils.AssertEqual(t, response.Code, "Success")
    testutils.AssertNotEmpty(t, response.Message)
    testutils.AssertEqual(t, response.Data.Id, "transformation_id")
    testutils.AssertEqual(t, response.Data.ProjectType, "DBT_CORE")
    testutils.AssertEqual(t, response.Data.CreatedAt, "2024-01-02T00:00:00.743708Z")
    testutils.AssertEqual(t, response.Data.CreatedById, "user_id")
    testutils.AssertEqual(t, response.Data.OutputModelNames[0], "string")
    testutils.AssertEqual(t, response.Data.Paused, true)
    testutils.AssertEqual(t, response.Data.Status, "RUNNING")

    testutils.AssertEqual(t, response.Data.TransformationSchedule.Cron[0], "0 */1 * * *")
    testutils.AssertEqual(t, response.Data.TransformationSchedule.ConnectionIds[0], "connection_id")
    testutils.AssertEqual(t, response.Data.TransformationSchedule.DaysOfWeek[0], "MONDAY")
    testutils.AssertEqual(t, response.Data.TransformationSchedule.DaysOfWeek[1], "FRIDAY")
    testutils.AssertEqual(t, response.Data.TransformationSchedule.TimeOfDay, "14:00")
    testutils.AssertEqual(t, response.Data.TransformationSchedule.ScheduleType, "TIME_OF_DAY")
    testutils.AssertEqual(t, response.Data.TransformationSchedule.Interval, 60)
    testutils.AssertEqual(t, response.Data.TransformationSchedule.SmartSyncing, true)

    testutils.AssertKey(t, "fake_field", response.Data.TransformationScheduleCustom, "unmapped-value")

    testutils.AssertEqual(t, response.Data.TransformationConfig.ProjectId, "string")
    testutils.AssertEqual(t, response.Data.TransformationConfig.Name, "string")
    testutils.AssertEqual(t, response.Data.TransformationConfig.Steps[0].Name, "string")
    testutils.AssertEqual(t, response.Data.TransformationConfig.Steps[0].Command, "string")

    testutils.AssertKey(t, "fake_field", response.Data.TransformationConfigCustom, "unmapped-value")
}
