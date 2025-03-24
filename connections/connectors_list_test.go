package connectors_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/connectors"
	testutils "github.com/fivetran/go-fivetran/test_utils"
	
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestConnectorsListServiceDo(t *testing.T) {
	// arrange
	limit := 10
	cursor := "some_cursor"

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/connectors").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareConnectorsListResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewConnectorsList().
		Limit(limit).
		Cursor(cursor).
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	assertConnectorsListResponse(t, response)
}

func prepareConnectorsListResponse() string {
	return `{
  "code": "Success",
  "message": "Operation performed.",
  "data": {
    "items": [
      {
        "id": "connector_id",
        "service": "string",
        "schema": "gsheets.table",
        "paused": false,
        "status": {
          "tasks": [
            {
              "code": "resync_table_warning",
              "message": "Resync Table Warning",
              "details": "string"
            }
          ],
          "warnings": [
            {
              "code": "resync_table_warning",
              "message": "Resync Table Warning",
              "details": "string"
            }
          ],
          "schema_status": "ready",
          "update_state": "delayed",
          "setup_state": "connected",
          "sync_state": "scheduled",
          "is_historical_sync": false,
          "rescheduled_for": "2024-12-01T15:43:29.013729Z"
        },
        "config": {
          "property1": {},
          "property2": {}
        },
        "daily_sync_time": "14:00",
        "succeeded_at": "2024-12-01T15:43:29.013729Z",
        "sync_frequency": 360,
        "group_id": "group_id",
        "connected_by": "user_id",
        "setup_tests": [
          {
            "title": "Test Title",
            "status": "FAILED",
            "message": "Error message",
            "details": "Error details"
          }
        ],
        "source_sync_details": {},
        "service_version": 0,
        "created_at": "2024-12-01T15:43:29.013729Z",
        "failed_at": "2024-12-01T15:43:29.013729Z",
        "private_link_id": "string",
        "proxy_agent_id": "string",
        "networking_method": "Directly",
        "connect_card": {
          "token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJkIjp7ImxvZ2luIjp0cnVlLCJ1c2VyIjoiX2FjY291bnR3b3J0aHkiLCJhY2NvdW50IjoiX21vb25iZWFtX2FjYyIsImdyb3VwIjoiX21vb25iZWFtIiwiY29ubmVjdG9yIjoiY29iYWx0X2VsZXZhdGlvbiIsIm1ldGhvZCI6IlBiZkNhcmQiLCJpZGVudGl0eSI6ZmFsc2V9LCJpYXQiOjE2Njc4MzA2MzZ9.YUMGUbzxW96xsKJLo4bTorqzx8Q19GTrUi3WFRFM8BU",
          "uri": "https://fivetran.com/connect-card/setup?auth=eyJ0eXAiOiJKV1QiLCJh..."
        },
        "pause_after_trial": false,
        "data_delay_threshold": 0,
        "data_delay_sensitivity": "LOW",
        "schedule_type": "auto",
        "local_processing_agent_id": "string",
        "connect_card_config": {
          "redirect_uri": "https://your.site/path",
          "hide_setup_guide": true
        },
        "hybrid_deployment_agent_id": "string"
      }
    ],
    "next_cursor": "cursor_value"
  }
}`
}

func assertConnectorsListResponse(t *testing.T, response connectors.ConnectorsListResponse) {
	testutils.AssertEqual(t, response.Code, "Success")

	testutils.AssertEqual(t, len(response.Data.Items), 1)
	testutils.AssertEqual(t, response.Data.Items[0].ID, "connector_id")
	testutils.AssertEqual(t, response.Data.Items[0].Schema, "gsheets.table")
	testutils.AssertEqual(t, response.Data.Items[0].Service, "string")
	testutils.AssertEqual(t, response.Data.Items[0].ConnectedBy, "user_id")
	testutils.AssertEqual(t, response.Data.Items[0].ScheduleType, "auto")
	testutils.AssertEqual(t, response.Data.Items[0].DailySyncTime, "14:00")
	testutils.AssertEqual(t, response.Data.Items[0].HybridDeploymentAgentId, "string")
	testutils.AssertEqual(t, response.Data.Items[0].ProxyAgentId, "string")
	testutils.AssertEqual(t, response.Data.Items[0].PrivateLinkId, "string")
	testutils.AssertEqual(t, response.Data.Items[0].NetworkingMethod, "Directly")
	testutils.AssertEqual(t, response.Data.Items[0].DataDelaySensitivity, "LOW")

	testutils.AssertEqual(t, response.Data.NextCursor, "cursor_value")
}
