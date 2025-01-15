package destinations_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/destinations"
	testutils "github.com/fivetran/go-fivetran/test_utils"
	
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestDestinationsListServiceDo(t *testing.T) {
	// arrange
	limit := 10
	cursor := "some_cursor"

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/destinations").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareDestinationsListResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewDestinationsList().
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
	assertDestinationsListResponse(t, response)
}

func prepareDestinationsListResponse() string {
	return `{
  "code": "Success",
  "message": "Operation performed.",
  "data": {
    "items": [
      {
        "id": "destination_id",
        "service": "string",
        "region": "GCP_US_EAST4",
        "networking_method": "Directly",
        "setup_status": "CONNECTED",
        "daylight_saving_time_enabled": true,
        "local_processing_agent_id": "local_processing_agent_id",
        "private_link_id": "private_link_id",
        "group_id": "group_id",
        "time_zone_offset": "+3",
        "hybrid_deployment_agent_id": "hybrid_deployment_agent_id"
      }
    ],
    "next_cursor": "cursor_value"
  }
}`
}

func assertDestinationsListResponse(t *testing.T, response destinations.DestinationsListResponse) {
	testutils.AssertEqual(t, response.Code, "Success")

  testutils.AssertEqual(t, response.Code, "Success")
  testutils.AssertEqual(t, response.Data.Items[0].ID, "destination_id")
  testutils.AssertEqual(t, response.Data.Items[0].GroupID, "group_id")
  testutils.AssertEqual(t, response.Data.Items[0].Service, "string")
  testutils.AssertEqual(t, response.Data.Items[0].DaylightSavingTimeEnabled, true)
  testutils.AssertEqual(t, response.Data.Items[0].HybridDeploymentAgentId, "hybrid_deployment_agent_id")
  testutils.AssertEqual(t, response.Data.Items[0].PrivateLinkId, "private_link_id")
  testutils.AssertEqual(t, response.Data.Items[0].NetworkingMethod, "Directly")
  testutils.AssertEqual(t, response.Data.Items[0].Region, "GCP_US_EAST4")
  testutils.AssertEqual(t, response.Data.Items[0].TimeZoneOffset, "+3")
  testutils.AssertEqual(t, response.Data.Items[0].SetupStatus, "CONNECTED")

	testutils.AssertEqual(t, response.Data.NextCursor, "cursor_value")
}
