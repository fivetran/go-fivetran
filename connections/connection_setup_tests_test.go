package connections_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/connections"
	
	"github.com/fivetran/go-fivetran/tests/mock"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestConnectionSetupTestsDo(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/connections/connection_id/test").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareConnectionSetupTestsResponse())
			return response, nil
		})
	service := ftClient.NewConnectionSetupTests().ConnectionID("connection_id").TrustCertificates(true).TrustFingerprints(true)

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

	assertConnectionSetupTestsResponse(t, response)
}

func TestConnectionSetupTestsServiceBadRequest(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/connections/connection_id/test").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusBadRequest, prepareConnectionReSyncTableResponse("BadRequest", "Invalid request"))
			return response, nil
		})

	service := ftClient.NewConnectionSetupTests().ConnectionID("connection_id").TrustCertificates(true).TrustFingerprints(true)

	// act
	response, err := service.Do(context.Background())

	if err == nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	// assert
	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	testutils.AssertEqual(t, response.Code, "BadRequest")
}

func assertConnectionSetupTestsResponse(t *testing.T, response connections.DetailsWithConfigResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertNotEmpty(t, response.Message)

	// Assert data fields
	testutils.AssertNotEmpty(t, response.Data.ID)
	testutils.AssertNotEmpty(t, response.Data.GroupID)
	testutils.AssertNotEmpty(t, response.Data.Service)

	// Assert status fields
	testutils.AssertNotEmpty(t, response.Data.Status.SetupState)
	testutils.AssertNotEmpty(t, response.Data.Status.SyncState)
	testutils.AssertNotEmpty(t, response.Data.Status.UpdateState)

	// Assert setup tests
	for _, test := range response.Data.SetupTests {
		testutils.AssertNotEmpty(t, test.Title)
		testutils.AssertNotEmpty(t, test.Status)
		testutils.AssertNotEmpty(t, test.Message)
	}

	testutils.AssertEqual(t, response.Data.Config.Username, "newuser")
	testutils.AssertEqual(t, response.Data.Config.Password, "******")
	testutils.AssertEqual(t, response.Data.Config.APIToken, "******")
	testutils.AssertEqual(t, response.Data.Config.ServiceVersion, "0")
}

func prepareConnectionSetupTestsResponse() string {
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
