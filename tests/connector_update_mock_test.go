package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/connectors"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestConnectorUpdateMock(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/connectors/connector_id").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := requestBodyToJson(t, req)
			assertConnectorUpdateRequest(t, body)
			response := mock.NewResponse(req, http.StatusOK, prepareConnectorUpdateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewConnectorModify().
		ConnectorID("connector_id").
		Paused(false).
		Config(prepareConfigUpdate()).
		Auth(prepareConnectorAuth()).
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

	assertConnectorUpdateResponse(t, response)
}

func TestCustomConnectorUpdateMock(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/connectors/connector_id").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := requestBodyToJson(t, req)
			assertCustomConnectorUpdateRequest(t, body)
			response := mock.NewResponse(req, http.StatusOK, prepareConnectorUpdateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewConnectorModify().
		ConnectorID("connector_id").
		Paused(false).
		ConfigCustom(prepareCustomUpdateConfig()).
		AuthCustom(prepareConnectorCustomAuthUpdate()).
		DoCustom(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	// assert
	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)

	assertCustomConnectorUpdateResponse(t, response)
}

func prepareConnectorCustomAuthUpdate() *map[string]interface{} {
	auth := make(map[string]interface{})
	clientAccess := make(map[string]interface{})

	clientAccess["client_id"] = "client_id"
	clientAccess["client_secret"] = "client_secret"

	auth["client_access"] = clientAccess
	auth["custom_auth"] = "custom_auth"

	return &auth
}

func TestCustomMergedConnectorUpdateMock(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/connectors/connector_id").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := requestBodyToJson(t, req)
			assertCustomConnectorUpdateRequest(t, body)
			response := mock.NewResponse(req, http.StatusOK, prepareConnectorUpdateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewConnectorModify().
		ConnectorID("connector_id").
		Paused(false).
		ConfigCustom(prepareCustomMergedUpdateConfigMap()).
		Config(prepareCustomMergedConfigUpdate()).
		Auth(prepareCustomMergedAuth()).
		AuthCustom(prepareConnectorCustomMergedAuthMap()).
		DoCustomMerged(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	// assert
	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)

	assertCustomMergedConnectorUpdateResponse(t, response)
}

func prepareConnectorCustomMergedAuthMap() *map[string]interface{} {
	auth := make(map[string]interface{})

	auth["custom_auth"] = "custom_auth"

	return &auth
}

func prepareCustomUpdateConfig() *map[string]interface{} {
	config := make(map[string]interface{})
	secretsList := make([]interface{}, 0)

	secret := make(map[string]interface{})
	secret["key"] = "key"
	secret["value"] = "value"

	secretsList = append(secretsList, secret)

	config["secrets_list"] = secretsList
	config["share_url"] = "share_url"
	config["is_keypair"] = true
	config["fake_field"] = "unmapped-value"

	return &config
}

func prepareCustomMergedAuth() *connectors.ConnectorAuth {
	auth := fivetran.NewConnectorAuth()

	clientAccess := fivetran.NewConnectorAuthClientAccess().ClientID("client_id").ClientSecret("client_secret")
	auth.ClientAccess(clientAccess)

	return auth
}

func prepareCustomMergedUpdateConfigMap() *map[string]interface{} {
	config := make(map[string]interface{})

	config["share_url"] = "share_url"
	config["fake_field"] = "unmapped-value"

	return &config
}

func prepareCustomMergedConfigUpdate() *connectors.ConnectorConfig {
	config := fivetran.NewConnectorConfig()
	secretsList := make([]*connectors.FunctionSecret, 0)
	secretsList = append(secretsList, fivetran.NewFunctionSecret().Key("key").Value("value"))
	config.
		SecretsList(secretsList).
		IsKeypair(true)

	return config
}

func prepareConfigUpdate() *connectors.ConnectorConfig {
	config := fivetran.NewConnectorConfig()
	secretsList := make([]*connectors.FunctionSecret, 0)
	secretsList = append(secretsList, fivetran.NewFunctionSecret().Key("key").Value("value"))
	config.
		SecretsList(secretsList).
		IsKeypair(true).
		ShareURL("share_url")

	return config
}

func prepareConnectorUpdateResponse() string {
	return `{
		"code": "Success",
		"data": {
			"id": "connector_id",
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

func assertConnectorConfig(t *testing.T, config connectors.ConnectorConfigResponse) {
	assertEqual(t, config.SecretsList[0].Key, "key")
	assertEqual(t, config.SecretsList[0].Value, "value")
	assertEqual(t, config.ShareURL, "share_url")
	assertEqual(t, *config.IsKeypair, true)
}

func assertConnectorUpdateResponse(t *testing.T, response fivetran.ConnectorModifyResponse) {
	assertEqual(t, response.Code, "Success")

	assertEqual(t, *response.Data.Paused, false)
	assertConnectorConfig(t, response.Data.Config)
}

func assertCustomConnectorUpdateResponse(t *testing.T, response fivetran.ConnectorCustomModifyResponse) {
	assertEqual(t, response.Code, "Success")

	assertEqual(t, *response.Data.Paused, false)

	assertKey(t, "share_url", response.Data.Config, "share_url")
	assertKey(t, "is_keypair", response.Data.Config, true)

	secretsList, ok := response.Data.Config["secrets_list"].([]interface{})

	assertEqual(t, ok, true)
	assertEqual(t, len(secretsList), 1)

	secret := secretsList[0].(map[string]interface{})

	assertKey(t, "key", secret, "key")
	assertKey(t, "value", secret, "value")
}

func assertCustomMergedConnectorUpdateResponse(t *testing.T, response fivetran.ConnectorCustomMergedModifyResponse) {
	assertEqual(t, response.Code, "Success")

	assertEqual(t, *response.Data.Paused, false)
	assertConnectorConfig(t, response.Data.Config)

	assertKey(t, "fake_field", response.Data.CustomConfig, "unmapped-value")

}

func assertConnectorUpdateRequest(t *testing.T, request map[string]interface{}) {
	assertKeyValue(t, request, "paused", false)
	config, ok := request["config"].(map[string]interface{})
	assertEqual(t, ok, true)

	assertKeyValue(t, config, "is_keypair", true)
	assertKeyValue(t, config, "share_url", "share_url")

	secretsList, ok := config["secrets_list"].([]interface{})
	assertEqual(t, ok, true)

	assertEqual(t, len(secretsList), int(1))

	secret, ok := secretsList[0].(map[string]interface{})
	assertEqual(t, ok, true)

	assertKey(t, "key", secret, "key")
	assertKey(t, "value", secret, "value")

	auth, ok := request["auth"].(map[string]interface{})
	assertEqual(t, ok, true)

	clientAccess, ok := auth["client_access"].(map[string]interface{})
	assertEqual(t, ok, true)

	assertKey(t, "client_id", clientAccess, "client_id")
	assertKey(t, "client_secret", clientAccess, "client_secret")
}

func assertCustomConnectorUpdateRequest(t *testing.T, request map[string]interface{}) {
	assertConnectorUpdateRequest(t, request)
	config := request["config"].(map[string]interface{})
	assertKey(t, "fake_field", config, "unmapped-value")
	auth := request["auth"].(map[string]interface{})
	assertKey(t, "custom_auth", auth, "custom_auth")
}
