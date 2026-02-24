package connections_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/connections"

	"github.com/fivetran/go-fivetran/tests/mock"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

const (
	CONNECTION_SERVICE = "test_service"
	SYNC_FREQUENCY    = 5
	DATA_DELAY_THRESHOLD = 1
)

func TestNewConnectionSecretsListMappingMock(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	syncFrequency := SYNC_FREQUENCY
	dataDelayThreshold := DATA_DELAY_THRESHOLD
	handler := mockClient.When(http.MethodPost, "/v1/connections").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertConnectionRequest(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareConnectionCreateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewConnectionCreate().
		Service(CONNECTION_SERVICE).
		GroupID("test_id").
		SyncFrequency(&syncFrequency).
		HybridDeploymentAgentId("lpa_id").
		ProxyAgentId("proxy_id").
		PrivateLinkId("private_link_id").
		NetworkingMethod("networking_method").
        DataDelayThreshold(&dataDelayThreshold).
        DataDelaySensitivity("CUSTOM").
		DestinationSchemaNames("FIVETRAN_NAMING").
		Config(prepareConnectionConfig()).
		Auth(prepareConnectionAuth()).
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

	assertConnectionResponse(t, response)
}

func TestNewConnectionCustomSecretsListMappingMock(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	syncFrequency := SYNC_FREQUENCY
	dataDelayThreshold := DATA_DELAY_THRESHOLD
	handler := mockClient.When(http.MethodPost, "/v1/connections").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertConnectionRequest(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareConnectionCreateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewConnectionCreate().
		Service(CONNECTION_SERVICE).
		GroupID("test_id").
		SyncFrequency(&syncFrequency).
		HybridDeploymentAgentId("lpa_id").
		ProxyAgentId("proxy_id").
		PrivateLinkId("private_link_id").
		NetworkingMethod("networking_method").
        DataDelayThreshold(&dataDelayThreshold).
        DataDelaySensitivity("CUSTOM").
		DestinationSchemaNames("FIVETRAN_NAMING").
		ConfigCustom(prepareConnectionCustomConfig()).
		AuthCustom(prepareConnectionCustomAuth()).
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
	assertConnectionCustomResponse(t, response)
}

func TestNewConnectionCustomMergedMappingMock(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	syncFrequency := SYNC_FREQUENCY
	dataDelayThreshold := DATA_DELAY_THRESHOLD
	handler := mockClient.When(http.MethodPost, "/v1/connections").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertConnectionRequest(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareConnectionCustomMergedCreateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewConnectionCreate().
		Service(CONNECTION_SERVICE).
		GroupID("test_id").
		SyncFrequency(&syncFrequency).
		HybridDeploymentAgentId("lpa_id").
		ProxyAgentId("proxy_id").
		PrivateLinkId("private_link_id").
		NetworkingMethod("networking_method").
        DataDelayThreshold(&dataDelayThreshold).
        DataDelaySensitivity("CUSTOM").
		DestinationSchemaNames("FIVETRAN_NAMING").
		Config(prepareConnectionConfig()).
		ConfigCustom(prepareConnectionCustomMergedConfig()).
		AuthCustom(prepareConnectionCustomAuth()).
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
	assertConnectionCustomMergedResponse(t, response)
}

func TestNewConnectionWihtNilSyncFrequency(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/connections").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertConnectionRequestWithNilSyncFrequency(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareConnectionCreateResponseWithNilSyncFrequency())
			return response, nil
		})

	// act
	response, err := ftClient.NewConnectionCreate().
		Service(CONNECTION_SERVICE).
		GroupID("test_id").
		SyncFrequency(nil).
		HybridDeploymentAgentId("lpa_id").
		ProxyAgentId("proxy_id").
		PrivateLinkId("private_link_id").
		NetworkingMethod("networking_method").
		DestinationSchemaNames("FIVETRAN_NAMING").
		Config(prepareConnectionConfig()).
		Auth(prepareConnectionAuth()).
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

	assertConnectionResponse(t, response)
	testutils.AssertIsNil(t, response.Data.SyncFrequency)
}

func prepareConnectionCreateResponse() string {
	return `{
        "code": "Success",
        "message": "Connection has been created",
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
            "hybrid_deployment_agent_id": "lpa_id",
            "proxy_agent_id": "proxy_id",
            "private_link_id": "private_link_id",
            "networking_method": "networking_method",
            "data_delay_threshold": 1,
            "data_delay_sensitivity": "CUSTOM",
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

func prepareConnectionCustomMergedCreateResponse() string {
	return `{
        "code": "Success",
        "message": "Connection has been created",
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
            "hybrid_deployment_agent_id": "lpa_id",
            "proxy_agent_id": "proxy_id",
            "private_link_id": "private_link_id",
            "networking_method": "networking_method",
            "data_delay_threshold": 1,
            "data_delay_sensitivity": "CUSTOM",
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

func prepareConnectionCreateResponseWithNilSyncFrequency() string {
	return `{
        "code": "Success",
        "message": "Connection has been created",
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
            "hybrid_deployment_agent_id": "lpa_id",
            "proxy_agent_id": "proxy_id",
            "private_link_id": "private_link_id",
            "networking_method": "networking_method",
            "data_delay_threshold": 1,
            "data_delay_sensitivity": "CUSTOM",
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

func prepareConnectionConfig() *connections.ConnectionConfig {
	config := fivetran.NewConnectionConfig()
	secretsList := make([]*connections.FunctionSecret, 0)
	secretsList = append(secretsList, fivetran.NewFunctionSecret().Key("key").Value("value"))
	config.SecretsList(secretsList)
	return config
}

func prepareConnectionAuth() *connections.ConnectionAuth {
	auth := fivetran.NewConnectionAuth()

	clientAccess := fivetran.NewConnectionAuthClientAccess().ClientID("client_id").ClientSecret("client_secret")
	auth.ClientAccess(clientAccess)

	return auth
}

func prepareConnectionCustomMergedConfig() *map[string]interface{} {
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

func prepareConnectionCustomConfig() *map[string]interface{} {
	config := make(map[string]interface{})
	secretsList := make([]interface{}, 0)

	secret := make(map[string]interface{})
	secret["key"] = "key"
	secret["value"] = "value"

	secretsList = append(secretsList, secret)

	config["secrets_list"] = secretsList

	return &config
}

func prepareConnectionCustomAuth() *map[string]interface{} {
	auth := make(map[string]interface{})
	clientAccess := make(map[string]interface{})

	clientAccess["client_id"] = "client_id"
	clientAccess["client_secret"] = "client_secret"

	auth["client_access"] = clientAccess

	return &auth
}

func assertConnectionRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "service", request, CONNECTION_SERVICE)
	testutils.AssertKey(t, "group_id", request, "test_id")
	testutils.AssertKey(t, "sync_frequency", request, float64(5))
	testutils.AssertKey(t, "hybrid_deployment_agent_id", request, "lpa_id")
	testutils.AssertKey(t, "proxy_agent_id", request, "proxy_id")
	testutils.AssertKey(t, "private_link_id", request, "private_link_id")
	testutils.AssertKey(t, "networking_method", request, "networking_method")
	testutils.AssertKey(t, "data_delay_sensitivity", request, "CUSTOM")
	testutils.AssertKey(t, "data_delay_threshold", request, float64(1))

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

func assertConnectionRequestWithNilSyncFrequency(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "service", request, CONNECTION_SERVICE)
	testutils.AssertKey(t, "group_id", request, "test_id")
	testutils.AssertKey(t, "hybrid_deployment_agent_id", request, "lpa_id")
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

func assertConnectionResponse(t *testing.T, response connections.DetailsWithConfigResponse) {

	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertNotEmpty(t, response.Message)

	testutils.AssertEqual(t, response.Data.Config.SecretsList[0].Key, "key")
	testutils.AssertEqual(t, response.Data.Config.SecretsList[0].Value, "value")
}

func assertConnectionCustomResponse(t *testing.T, response connections.DetailsWithCustomConfigResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertNotEmpty(t, response.Message)

	testutils.AssertKey(t, "key", response.Data.Config["secrets_list"].([]interface{})[0].(map[string]interface{}), "key")
	testutils.AssertKey(t, "value", response.Data.Config["secrets_list"].([]interface{})[0].(map[string]interface{}), "value")
}

func assertConnectionCustomMergedResponse(t *testing.T, response connections.DetailsWithCustomMergedConfigResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertNotEmpty(t, response.Message)

	testutils.AssertEqual(t, response.Data.Config.SecretsList[0].Key, "key")
	testutils.AssertEqual(t, response.Data.Config.SecretsList[0].Value, "value")

	testutils.AssertKey(t, "user", response.Data.CustomConfig["fake_list"].([]interface{})[0].(map[string]interface{}), "user")
	testutils.AssertKey(t, "password", response.Data.CustomConfig["fake_list"].([]interface{})[0].(map[string]interface{}), "******")
	testutils.AssertKey(t, "fake_field", response.Data.CustomConfig, "unmapped-value")
}
