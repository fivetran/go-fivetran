package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/connectors"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestConnectorDetailsMock(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/connectors/connector_id").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareConnectorDetailsResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewConnectorDetails().ConnectorID("connector_id").Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	// assert
	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)

	assertConnectorDetailsResponse(t, response)
}

func TestCustomConnectorDetailsMock(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/connectors/connector_id").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareConnectorDetailsResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewConnectorDetails().ConnectorID("connector_id").DoCustom(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	// assert
	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)

	assertCustomConnectorDetailsResponse(t, response)
}

func TestCustomMergedConnectorDetailsMock(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/connectors/connector_id").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareConnectorDetailsResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewConnectorDetails().ConnectorID("connector_id").DoCustomMerged(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	// assert
	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)

	assertCustomMergedConnectorDetailsResponse(t, response)
}

func prepareConnectorDetailsResponse() string {
	return `{
		"code": "Success",
		"data": {
			"id": "speak_inexpensive",
			"group_id": "projected_sickle",
			"service": "criteo",
			"service_version": 0,
			"schema": "criteo",
			"paused": true,
			"pause_after_trial": true,
			"connected_by": "interment_burdensome",
			"created_at": "2018-12-01T15:43:29.013729Z",
			"succeeded_at": null,
			"failed_at": null,
			"sync_frequency": 1440,
			"daily_sync_time": "03:00",
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
				"share_url": "share_url",
				"is_keypair": true,
				"secrets_list": [
					{
						"key": "key", 
						"value": "value"
					}
				],
				"fake_field": "unmapped-value"
			}
		}
	}`
}

func assertConnectorDetailsResponse(t *testing.T, response connectors.DetailsWithConfigNoTestsResponse) {

	assertEqual(t, response.Code, "Success")

	assertEqual(t, response.Data.Config.SecretsList[0].Key, "key")
	assertEqual(t, response.Data.Config.SecretsList[0].Value, "value")
	assertEqual(t, response.Data.Config.ShareURL, "share_url")
	assertEqual(t, *response.Data.Config.IsKeypair, true)
}

func assertCustomConnectorDetailsResponse(t *testing.T, response connectors.DetailsWithCustomConfigNoTestsResponse) {

	assertEqual(t, response.Code, "Success")

	assertKey(t, "share_url", response.Data.Config, "share_url")
	assertKey(t, "is_keypair", response.Data.Config, true)

	secretsList, ok := response.Data.Config["secrets_list"].([]interface{})

	assertEqual(t, ok, true)
	assertEqual(t, len(secretsList), 1)

	secret := secretsList[0].(map[string]interface{})

	assertKey(t, "key", secret, "key")
	assertKey(t, "value", secret, "value")
}

func assertCustomMergedConnectorDetailsResponse(t *testing.T, response connectors.DetailsWithCustomMergedConfigNoTestsResponse) {

	assertEqual(t, response.Code, "Success")

	assertHasNoKey(t, response.Data.CustomConfig, "share_url")
	assertHasNoKey(t, response.Data.CustomConfig, "is_keypair")
	assertHasNoKey(t, response.Data.CustomConfig, "secrets_list")

	assertKeyValue(t, response.Data.CustomConfig, "fake_field", "unmapped-value")

	assertEqual(t, response.Data.Config.SecretsList[0].Key, "key")
	assertEqual(t, response.Data.Config.SecretsList[0].Value, "value")
	assertEqual(t, response.Data.Config.ShareURL, "share_url")
	assertEqual(t, *response.Data.Config.IsKeypair, true)

}
