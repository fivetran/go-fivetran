package transformations_test

import (
    "context"
    "net/http"
    "testing"

    "github.com/fivetran/go-fivetran/transformations"
    
    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestTransformationDetailsServiceDo(t *testing.T) {
    // arrange

    ftClient, mockClient := testutils.CreateTestClient()
    handler := mockClient.When(http.MethodGet, "/v1/transformations/id").
        ThenCall(func(req *http.Request) (*http.Response, error) {
            response := mock.NewResponse(req, http.StatusOK, prepareTransformationDetailsResponse())
            return response, nil
        })

    // act
    response, err := ftClient.NewTransformationDetails().
        TransformationId("id").
        Do(context.Background())

    // assert
    if err != nil {
        t.Error(err)
    }

    interactions := mockClient.Interactions()
    testutils.AssertEqual(t, len(interactions), 1)
    testutils.AssertEqual(t, interactions[0].Handler, handler)
    testutils.AssertEqual(t, handler.Interactions, 1)
    assertTransformationDetailsResponse(t, response)
}

func prepareTransformationDetailsResponse() string {
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
      "time_of_day": "14:00"
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
      ]
    }
  }
}`
}

func assertTransformationDetailsResponse(t *testing.T, response transformations.TransformationResponse) {
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