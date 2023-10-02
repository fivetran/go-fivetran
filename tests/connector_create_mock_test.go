package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/connectors"
	"github.com/fivetran/go-fivetran/tests/mock"
)

const (
	CONNECTOR_SERVICE = "test_service"
	SYNC_FREQUENCY    = 5
)

func TestNewConnectorSecretsListMappingMock(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/connectors").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := RequestBodyToJson(t, req)
			assertConnectorRequest(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareConnectorCreateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewConnectorCreate().
		Service(CONNECTOR_SERVICE).
		GroupID(ID).
		SyncFrequency(SYNC_FREQUENCY).
		Config(prepareConnectorConfig()).
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

	assertConnectorResponse(t, response)
}

func TestNewConnectorCustomSecretsListMappingMock(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/connectors").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := RequestBodyToJson(t, req)
			assertConnectorRequest(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareConnectorCreateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewConnectorCreate().
		Service(CONNECTOR_SERVICE).
		GroupID(ID).
		SyncFrequency(SYNC_FREQUENCY).
		ConfigCustom(prepareConnectorCustomConfig()).
		AuthCustom(prepareConnectorCustomAuth()).
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
	assertConnectorCustomResponse(t, response)
}

func TestNewConnectorCustomMergedMappingMock(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/connectors").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := RequestBodyToJson(t, req)
			assertConnectorRequest(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareConnectorCustomMergedCreateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewConnectorCreate().
		Service(CONNECTOR_SERVICE).
		GroupID(ID).
		SyncFrequency(SYNC_FREQUENCY).
		Config(prepareConnectorConfig()).
		ConfigCustom(prepareConnectorCustomMergedConfig()).
		AuthCustom(prepareConnectorCustomAuth()).
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
	assertConnectorCustomMergedResponse(t, response)
}

func prepareConnectorCreateResponse() string {
	return `{
		"code": "Success",
		"message": "Connector has been created",
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
				"secrets_list": [
					{
						"key": "key", 
						"value": "value"
					}
				]
			}
		}
	}`
}

func prepareConnectorCustomMergedCreateResponse() string {
	return `{
		"code": "Success",
		"message": "Connector has been created",
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
				"secrets_list": [
					{
						"key": "key", 
						"value": "value"
					}
				],
				"fake_field": "unmapped-value",
				"fake_list": [
					{
						"user": "user",
						"password": "******"
					}
				]
			}
		}
	}`
}

func prepareConnectorConfig() *connectors.ConnectorConfig {
	config := fivetran.NewConnectorConfig()
	secretsList := make([]*connectors.FunctionSecret, 0)
	secretsList = append(secretsList, fivetran.NewFunctionSecret().Key("key").Value("value"))
	config.SecretsList(secretsList)
	return config
}

func prepareConnectorAuth() *connectors.ConnectorAuth {
	auth := fivetran.NewConnectorAuth()

	clientAccess := fivetran.NewConnectorAuthClientAccess().ClientID("client_id").ClientSecret("client_secret")
	auth.ClientAccess(clientAccess)

	return auth
}

func prepareConnectorCustomMergedConfig() *map[string]interface{} {
	config := make(map[string]interface{})
	fakeList := make([]interface{}, 0)

	user := make(map[string]interface{})
	user["user"] = "user"
	user["password"] = "password"

	fakeList = append(fakeList, user)

	config["fake_list"] = fakeList
	config["fake_field"] = "unmapped-value"

	return &config
}

func prepareConnectorCustomConfig() *map[string]interface{} {
	config := make(map[string]interface{})
	secretsList := make([]interface{}, 0)

	secret := make(map[string]interface{})
	secret["key"] = "key"
	secret["value"] = "value"

	secretsList = append(secretsList, secret)

	config["secrets_list"] = secretsList

	return &config
}

func prepareConnectorCustomAuth() *map[string]interface{} {
	auth := make(map[string]interface{})
	clientAccess := make(map[string]interface{})

	clientAccess["client_id"] = "client_id"
	clientAccess["client_secret"] = "client_secret"

	auth["client_access"] = clientAccess

	return &auth
}

func assertConnectorRequest(t *testing.T, request map[string]interface{}) {
	assertKey(t, "service", request, CONNECTOR_SERVICE)
	assertKey(t, "group_id", request, ID)
	assertKey(t, "sync_frequency", request, float64(5))

	config, ok := request["config"].(map[string]interface{})
	assertEqual(t, ok, true)

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

func assertConnectorResponse(t *testing.T, response connectors.DetailsWithConfigResponse) {

	assertEqual(t, response.Code, "Success")
	assertNotEmpty(t, response.Message)

	assertEqual(t, response.Data.Config.SecretsList[0].Key, "key")
	assertEqual(t, response.Data.Config.SecretsList[0].Value, "value")
}

func assertConnectorCustomResponse(t *testing.T, response connectors.DetailsWithCustomConfigResponse) {
	assertEqual(t, response.Code, "Success")
	assertNotEmpty(t, response.Message)

	assertKey(t, "key", response.Data.Config["secrets_list"].([]interface{})[0].(map[string]interface{}), "key")
	assertKey(t, "value", response.Data.Config["secrets_list"].([]interface{})[0].(map[string]interface{}), "value")
}

func assertConnectorCustomMergedResponse(t *testing.T, response connectors.DetailsWithCustomMergedConfigResponse) {
	assertEqual(t, response.Code, "Success")
	assertNotEmpty(t, response.Message)

	assertEqual(t, response.Data.Config.SecretsList[0].Key, "key")
	assertEqual(t, response.Data.Config.SecretsList[0].Value, "value")

	assertKey(t, "user", response.Data.CustomConfig["fake_list"].([]interface{})[0].(map[string]interface{}), "user")
	assertKey(t, "password", response.Data.CustomConfig["fake_list"].([]interface{})[0].(map[string]interface{}), "******")
	assertKey(t, "fake_field", response.Data.CustomConfig, "unmapped-value")
}
