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

func TestConnectionUpdateMock(t *testing.T) {
    // arrange
    ftClient, mockClient := testutils.CreateTestClient()
    dataDelayThreshold := DATA_DELAY_THRESHOLD
    handler := mockClient.When(http.MethodPatch, "/v1/connections/connection_id").ThenCall(

        func(req *http.Request) (*http.Response, error) {
            body := testutils.RequestBodyToJson(t, req)
            assertConnectionUpdateRequest(t, body)
            response := mock.NewResponse(req, http.StatusOK, prepareConnectionUpdateResponse())
            return response, nil
        })

    // act
    response, err := ftClient.NewConnectionUpdate().
        ConnectionID("connection_id").
        HybridDeploymentAgentId("hybrid_deployment_agent_id").
        ProxyAgentId("proxy_id").
        PrivateLinkId("private_link_id").
        NetworkingMethod("networking_method").
        DataDelayThreshold(&dataDelayThreshold).
        DataDelaySensitivity("CUSTOM").
        Paused(false).
        Config(prepareConfigUpdate()).
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

    assertConnectionUpdateResponse(t, response)
}

func TestCustomConnectionUpdateMock(t *testing.T) {
    // arrange
    ftClient, mockClient := testutils.CreateTestClient()
    dataDelayThreshold := DATA_DELAY_THRESHOLD
    handler := mockClient.When(http.MethodPatch, "/v1/connections/connection_id").ThenCall(

        func(req *http.Request) (*http.Response, error) {
            body := testutils.RequestBodyToJson(t, req)
            assertCustomConnectionUpdateRequest(t, body)
            response := mock.NewResponse(req, http.StatusOK, prepareConnectionUpdateResponse())
            return response, nil
        })

    // act
    response, err := ftClient.NewConnectionUpdate().
        ConnectionID("connection_id").
        HybridDeploymentAgentId("hybrid_deployment_agent_id").
        ProxyAgentId("proxy_id").
        PrivateLinkId("private_link_id").
        NetworkingMethod("networking_method").
        DataDelayThreshold(&dataDelayThreshold).
        DataDelaySensitivity("CUSTOM").
        Paused(false).
        ConfigCustom(prepareCustomUpdateConfig()).
        AuthCustom(prepareConnectionCustomAuthUpdate()).
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

    assertCustomConnectionUpdateResponse(t, response)
}

func prepareConnectionCustomAuthUpdate() *map[string]interface{} {
    auth := make(map[string]interface{})
    clientAccess := make(map[string]interface{})

    clientAccess["client_id"] = "client_id"
    clientAccess["client_secret"] = "client_secret"

    auth["client_access"] = clientAccess
    auth["custom_auth"] = "custom_auth"

    return &auth
}

func TestCustomMergedConnectionUpdateMock(t *testing.T) {
    // arrange
    ftClient, mockClient := testutils.CreateTestClient()
    dataDelayThreshold := DATA_DELAY_THRESHOLD
    handler := mockClient.When(http.MethodPatch, "/v1/connections/connection_id").ThenCall(

        func(req *http.Request) (*http.Response, error) {
            body := testutils.RequestBodyToJson(t, req)
            assertCustomConnectionUpdateRequest(t, body)
            response := mock.NewResponse(req, http.StatusOK, prepareConnectionUpdateResponse())
            return response, nil
        })

    // act
    response, err := ftClient.NewConnectionUpdate().
        ConnectionID("connection_id").
        Paused(false).
        HybridDeploymentAgentId("hybrid_deployment_agent_id").
        ProxyAgentId("proxy_id").
        PrivateLinkId("private_link_id").
        NetworkingMethod("networking_method").
        DataDelayThreshold(&dataDelayThreshold).
        DataDelaySensitivity("CUSTOM").
        ConfigCustom(prepareCustomMergedUpdateConfigMap()).
        Config(prepareCustomMergedConfigUpdate()).
        Auth(prepareCustomMergedAuth()).
        AuthCustom(prepareConnectionCustomMergedAuthMap()).
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

    assertCustomMergedConnectionUpdateResponse(t, response)
}

func prepareConnectionCustomMergedAuthMap() *map[string]interface{} {
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

func prepareCustomMergedAuth() *connections.ConnectionAuth {
    auth := fivetran.NewConnectionAuth()

    clientAccess := fivetran.NewConnectionAuthClientAccess().ClientID("client_id").ClientSecret("client_secret")
    auth.ClientAccess(clientAccess)

    return auth
}

func prepareCustomMergedUpdateConfigMap() *map[string]interface{} {
    config := make(map[string]interface{})

    config["share_url"] = "share_url"
    config["fake_field"] = "unmapped-value"

    return &config
}

func prepareCustomMergedConfigUpdate() *connections.ConnectionConfig {
    config := fivetran.NewConnectionConfig()
    secretsList := make([]*connections.FunctionSecret, 0)
    secretsList = append(secretsList, fivetran.NewFunctionSecret().Key("key").Value("value"))
    config.
        SecretsList(secretsList).
        IsKeypair(true)

    return config
}

func prepareConfigUpdate() *connections.ConnectionConfig {
    config := fivetran.NewConnectionConfig()
    secretsList := make([]*connections.FunctionSecret, 0)
    secretsList = append(secretsList, fivetran.NewFunctionSecret().Key("key").Value("value"))
    config.
        SecretsList(secretsList).
        IsKeypair(true).
        ShareURL("share_url")

    return config
}

func prepareConnectionUpdateResponse() string {
    return `{
        "code": "Success",
        "data": {
            "id": "connection_id",
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
            "hybrid_deployment_agent_id": "hybrid_deployment_agent_id",
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

func assertConnectionConfig(t *testing.T, config connections.ConnectionConfigResponse) {
    testutils.AssertEqual(t, config.SecretsList[0].Key, "key")
    testutils.AssertEqual(t, config.SecretsList[0].Value, "value")
    testutils.AssertEqual(t, config.ShareURL, "share_url")
    testutils.AssertEqual(t, *config.IsKeypair, true)
}

func assertConnectionUpdateResponse(t *testing.T, response connections.DetailsWithConfigResponse) {
    testutils.AssertEqual(t, response.Code, "Success")

    testutils.AssertEqual(t, *response.Data.Paused, false)
    testutils.AssertEqual(t, response.Data.HybridDeploymentAgentId, "hybrid_deployment_agent_id")
    testutils.AssertEqual(t, response.Data.ProxyAgentId, "proxy_id")
    testutils.AssertEqual(t, response.Data.PrivateLinkId, "private_link_id")
    testutils.AssertEqual(t, response.Data.NetworkingMethod, "networking_method")
    testutils.AssertEqual(t, *response.Data.DataDelayThreshold, 1)
    testutils.AssertEqual(t, response.Data.DataDelaySensitivity, "CUSTOM")

    assertConnectionConfig(t, response.Data.Config)
}

func assertCustomConnectionUpdateResponse(t *testing.T, response connections.DetailsWithCustomConfigResponse) {
    testutils.AssertEqual(t, response.Code, "Success")

    testutils.AssertEqual(t, *response.Data.Paused, false)
    testutils.AssertEqual(t, response.Data.HybridDeploymentAgentId, "hybrid_deployment_agent_id")
    testutils.AssertEqual(t, response.Data.ProxyAgentId, "proxy_id")
    testutils.AssertEqual(t, response.Data.PrivateLinkId, "private_link_id")
    testutils.AssertEqual(t, response.Data.NetworkingMethod, "networking_method")
    testutils.AssertEqual(t, *response.Data.DataDelayThreshold, 1)
    testutils.AssertEqual(t, response.Data.DataDelaySensitivity, "CUSTOM")

    testutils.AssertKey(t, "share_url", response.Data.Config, "share_url")
    testutils.AssertKey(t, "is_keypair", response.Data.Config, true)

    secretsList, ok := response.Data.Config["secrets_list"].([]interface{})

    testutils.AssertEqual(t, ok, true)
    testutils.AssertEqual(t, len(secretsList), 1)

    secret := secretsList[0].(map[string]interface{})

    testutils.AssertKey(t, "key", secret, "key")
    testutils.AssertKey(t, "value", secret, "value")
}

func assertCustomMergedConnectionUpdateResponse(t *testing.T, response connections.DetailsWithCustomMergedConfigResponse) {
    testutils.AssertEqual(t, response.Code, "Success")

    testutils.AssertEqual(t, *response.Data.Paused, false)
    testutils.AssertEqual(t, response.Data.HybridDeploymentAgentId, "hybrid_deployment_agent_id")
    testutils.AssertEqual(t, response.Data.ProxyAgentId, "proxy_id")
    testutils.AssertEqual(t, response.Data.PrivateLinkId, "private_link_id")
    testutils.AssertEqual(t, response.Data.NetworkingMethod, "networking_method")
    testutils.AssertEqual(t, *response.Data.DataDelayThreshold, 1)
    testutils.AssertEqual(t, response.Data.DataDelaySensitivity, "CUSTOM")

    assertConnectionConfig(t, response.Data.Config)

    testutils.AssertKey(t, "fake_field", response.Data.CustomConfig, "unmapped-value")
}

func assertConnectionUpdateRequest(t *testing.T, request map[string]interface{}) {
    testutils.AssertKeyValue(t, request, "paused", false)
    testutils.AssertKeyValue(t, request, "hybrid_deployment_agent_id", "hybrid_deployment_agent_id")
    testutils.AssertKeyValue(t, request, "proxy_agent_id", "proxy_id")
    testutils.AssertKeyValue(t, request, "private_link_id", "private_link_id")
    testutils.AssertKeyValue(t, request, "networking_method", "networking_method")
    
    config, ok := request["config"].(map[string]interface{})
    testutils.AssertEqual(t, ok, true)

    testutils.AssertKeyValue(t, config, "is_keypair", true)
    testutils.AssertKeyValue(t, config, "share_url", "share_url")

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

func assertCustomConnectionUpdateRequest(t *testing.T, request map[string]interface{}) {
    assertConnectionUpdateRequest(t, request)
    config := request["config"].(map[string]interface{})
    testutils.AssertKey(t, "fake_field", config, "unmapped-value")
    auth := request["auth"].(map[string]interface{})
    testutils.AssertKey(t, "custom_auth", auth, "custom_auth")
}
