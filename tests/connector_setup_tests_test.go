package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestConnectorSetupTestsService_Do(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/connectors/connector_id/test").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareConnectorSetupTestsResponse())
			return response, nil
		})
	service := ftClient.NewConnectorSetupTests().ConnectorID("connector_id").TrustCertificates(true).TrustFingerprints(true)

	// act
	response, err := service.Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)

	// assert
	assertConnectorSetupTestsResponse(t, response)
}

func assertConnectorSetupTestsResponse(t *testing.T, response fivetran.ConnectorSetupTestsResponse) {
	assertEqual(t, response.Code, "Success")
	assertNotEmpty(t, response.Message)

	// Assert data fields
	assertNotEmpty(t, response.Data.ID)
	assertNotEmpty(t, response.Data.GroupID)
	assertNotEmpty(t, response.Data.Service)

	// Assert status fields
	assertNotEmpty(t, response.Data.Status.SetupState)
	assertNotEmpty(t, response.Data.Status.SyncState)
	assertNotEmpty(t, response.Data.Status.UpdateState)

	// Assert setup tests
	for _, test := range response.Data.SetupTests {
		assertNotEmpty(t, test.Title)
		assertNotEmpty(t, test.Status)
		assertNotEmpty(t, test.Message)
	}

	// Add assertions for specific fields in the ConnectorConfigResponse struct
}

func prepareConnectorSetupTestsResponse() string {
	return `{
		"code": "Success",
		"message": "Setup tests were completed",
		"data": {
			"id": "speak_inexpensive",
			"group_id": "projected_sickle",
			"service": "criteo",
			"service_version": 0,
			"schema": "criteo",
			"paused": false,
			"pause_after_trial": true,
			"connected_by": "interment_burdensome",
			"created_at": "2018-12-01T15:43:29.013729Z",
			"succeeded_at": null,
			"failed_at": null,
			"sync_frequency":60,
			"status": {
				"setup_state": "incomplete",
				"sync_state": "scheduled",
				"update_state": "on_schedule",
				"is_historical_sync": true,
				"tasks": [],
				"warnings": []
			},
			"setup_tests": [{
				"title": "Validate Login",
				"status": "FAILED",
				"message": "Invalid login credentials"
			}],
			"config": {
				"username": "newuser",
				"password": "******",
				"api_token": "******",
				"service_version": "0"
			}
		}
	}
	`
}
