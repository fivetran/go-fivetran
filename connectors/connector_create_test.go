package connectors_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/connectors"

	"github.com/fivetran/go-fivetran/tests/mock"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

const (
	CONNECTOR_SERVICE = "test_service"
	SYNC_FREQUENCY    = 5
)

func TestNewConnectorSecretsListMappingMock(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	syncFrequency := SYNC_FREQUENCY
	handler := mockClient.When(http.MethodPost, "/v1/connectors").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertConnectorRequest(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareConnectorCreateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewConnectorCreate().
		Service(CONNECTOR_SERVICE).
		GroupID("test_id").
		SyncFrequency(&syncFrequency).
		LocalProcessingAgentId("lpa_id").
		ProxyAgentId("proxy_id").
		PrivateLinkId("private_link_id").
		NetworkingMethod("networking_method").
		Config(prepareConnectorConfig()).
		Auth(prepareConnectorAuth()).
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

	assertConnectorResponse(t, response)
}

func TestNewConnectorCustomSecretsListMappingMock(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	syncFrequency := SYNC_FREQUENCY
	handler := mockClient.When(http.MethodPost, "/v1/connectors").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertConnectorRequest(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareConnectorCreateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewConnectorCreate().
		Service(CONNECTOR_SERVICE).
		GroupID("test_id").
		SyncFrequency(&syncFrequency).
		LocalProcessingAgentId("lpa_id").
		ProxyAgentId("proxy_id").
		PrivateLinkId("private_link_id").
		NetworkingMethod("networking_method").
		ConfigCustom(prepareConnectorCustomConfig()).
		AuthCustom(prepareConnectorCustomAuth()).
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
	assertConnectorCustomResponse(t, response)
}

func TestNewConnectorCustomMergedMappingMock(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	syncFrequency := SYNC_FREQUENCY
	handler := mockClient.When(http.MethodPost, "/v1/connectors").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertConnectorRequest(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareConnectorCustomMergedCreateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewConnectorCreate().
		Service(CONNECTOR_SERVICE).
		GroupID("test_id").
		SyncFrequency(&syncFrequency).
		LocalProcessingAgentId("lpa_id").
		ProxyAgentId("proxy_id").
		PrivateLinkId("private_link_id").
		NetworkingMethod("networking_method").
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
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	assertConnectorCustomMergedResponse(t, response)
}

func TestNewConnectorWihtNilSyncFrequency(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/connectors").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertConnectorRequestWithNilSyncFrequency(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareConnectorCreateResponseWithNilSyncFrequency())
			return response, nil
		})

	// act
	response, err := ftClient.NewConnectorCreate().
		Service(CONNECTOR_SERVICE).
		GroupID("test_id").
		SyncFrequency(nil).
		LocalProcessingAgentId("lpa_id").
		ProxyAgentId("proxy_id").
		PrivateLinkId("private_link_id").
		NetworkingMethod("networking_method").
		Config(prepareConnectorConfig()).
		Auth(prepareConnectorAuth()).
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

	assertConnectorResponse(t, response)
	testutils.AssertIsNil(t, response.Data.SyncFrequency)
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
            "local_processing_agent_id": "lpa_id",
            "proxy_agent_id": "proxy_id",
            "private_link_id": "private_link_id",
            "networking_method": "networking_method",
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
            "local_processing_agent_id": "lpa_id",
            "proxy_agent_id": "proxy_id",
            "private_link_id": "private_link_id",
            "networking_method": "networking_method",
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

func prepareConnectorCreateResponseWithNilSyncFrequency() string {
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
            "sync_frequency": null,
            "daily_sync_time": "03:00",
            "local_processing_agent_id": "lpa_id",
            "proxy_agent_id": "proxy_id",
            "private_link_id": "private_link_id",
            "networking_method": "networking_method",
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
	testutils.AssertKey(t, "service", request, CONNECTOR_SERVICE)
	testutils.AssertKey(t, "group_id", request, "test_id")
	testutils.AssertKey(t, "sync_frequency", request, float64(5))
	testutils.AssertKey(t, "local_processing_agent_id", request, "lpa_id")
	testutils.AssertKey(t, "proxy_agent_id", request, "proxy_id")
	testutils.AssertKey(t, "private_link_id", request, "private_link_id")
	testutils.AssertKey(t, "networking_method", request, "networking_method")

	config, ok := request["config"].(map[string]interface{})
	testutils.AssertEqual(t, ok, true)

	secretsList, ok := config["secrets_list"].([]interface{})
	testutils.AssertEqual(t, ok, true)

	testutils.AssertEqual(t, len(secretsList), int(1))

	secret, ok := secretsList[0].(map[string]interface{})
	testutils.AssertEqual(t, ok, true)

	testutils.AssertKey(t, "key", secret, "key")
	testutils.AssertKey(t, "value", secret, "value")

	auth, ok := request["auth"].(map[string]interface{})
	testutils.AssertEqual(t, ok, true)

	clientAccess, ok := auth["client_access"].(map[string]interface{})
	testutils.AssertEqual(t, ok, true)

	testutils.AssertKey(t, "client_id", clientAccess, "client_id")
	testutils.AssertKey(t, "client_secret", clientAccess, "client_secret")
}

func assertConnectorRequestWithNilSyncFrequency(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "service", request, CONNECTOR_SERVICE)
	testutils.AssertKey(t, "group_id", request, "test_id")
	testutils.AssertKey(t, "local_processing_agent_id", request, "lpa_id")
	testutils.AssertKey(t, "proxy_agent_id", request, "proxy_id")
	testutils.AssertKey(t, "private_link_id", request, "private_link_id")
	testutils.AssertKey(t, "networking_method", request, "networking_method")

	config, ok := request["config"].(map[string]interface{})
	testutils.AssertEqual(t, ok, true)

	secretsList, ok := config["secrets_list"].([]interface{})
	testutils.AssertEqual(t, ok, true)

	testutils.AssertEqual(t, len(secretsList), int(1))

	secret, ok := secretsList[0].(map[string]interface{})
	testutils.AssertEqual(t, ok, true)

	testutils.AssertKey(t, "key", secret, "key")
	testutils.AssertKey(t, "value", secret, "value")

	auth, ok := request["auth"].(map[string]interface{})
	testutils.AssertEqual(t, ok, true)

	clientAccess, ok := auth["client_access"].(map[string]interface{})
	testutils.AssertEqual(t, ok, true)

	testutils.AssertKey(t, "client_id", clientAccess, "client_id")
	testutils.AssertKey(t, "client_secret", clientAccess, "client_secret")
}

func assertConnectorResponse(t *testing.T, response connectors.DetailsWithConfigResponse) {

	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertNotEmpty(t, response.Message)

	testutils.AssertEqual(t, response.Data.Config.SecretsList[0].Key, "key")
	testutils.AssertEqual(t, response.Data.Config.SecretsList[0].Value, "value")
}

func assertConnectorCustomResponse(t *testing.T, response connectors.DetailsWithCustomConfigResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertNotEmpty(t, response.Message)

	testutils.AssertKey(t, "key", response.Data.Config["secrets_list"].([]interface{})[0].(map[string]interface{}), "key")
	testutils.AssertKey(t, "value", response.Data.Config["secrets_list"].([]interface{})[0].(map[string]interface{}), "value")
}

func assertConnectorCustomMergedResponse(t *testing.T, response connectors.DetailsWithCustomMergedConfigResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertNotEmpty(t, response.Message)

	testutils.AssertEqual(t, response.Data.Config.SecretsList[0].Key, "key")
	testutils.AssertEqual(t, response.Data.Config.SecretsList[0].Value, "value")

	testutils.AssertKey(t, "user", response.Data.CustomConfig["fake_list"].([]interface{})[0].(map[string]interface{}), "user")
	testutils.AssertKey(t, "password", response.Data.CustomConfig["fake_list"].([]interface{})[0].(map[string]interface{}), "******")
	testutils.AssertKey(t, "fake_field", response.Data.CustomConfig, "unmapped-value")
}
