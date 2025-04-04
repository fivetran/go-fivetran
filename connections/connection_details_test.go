package connections_test

import (
    "context"
    "net/http"
    "testing"

    "github.com/fivetran/go-fivetran/connections"
    
    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestConnectionDetailsMock(t *testing.T) {
    // arrange
    ftClient, mockClient := testutils.CreateTestClient()
    handler := mockClient.When(http.MethodGet, "/v1/connections/connection_id").ThenCall(

        func(req *http.Request) (*http.Response, error) {
            response := mock.NewResponse(req, http.StatusOK, prepareConnectionDetailsResponse())
            return response, nil
        })

    // act
    response, err := ftClient.NewConnectionDetails().ConnectionID("connection_id").Do(context.Background())

    if err != nil {
        t.Logf("%+v\n", response)
        t.Error(err)
    }

    // assert
    interactions := mockClient.Interactions()
    testutils.AssertEqual(t, len(interactions), 1)
    testutils.AssertEqual(t, interactions[0].Handler, handler)
    testutils.AssertEqual(t, handler.Interactions, 1)

    assertConnectionDetailsResponse(t, response)
}

func TestCustomConnectionDetailsMock(t *testing.T) {
    // arrange
    ftClient, mockClient := testutils.CreateTestClient()
    handler := mockClient.When(http.MethodGet, "/v1/connections/connection_id").ThenCall(

        func(req *http.Request) (*http.Response, error) {
            response := mock.NewResponse(req, http.StatusOK, prepareConnectionDetailsResponse())
            return response, nil
        })

    // act
    response, err := ftClient.NewConnectionDetails().ConnectionID("connection_id").DoCustom(context.Background())

    if err != nil {
        t.Logf("%+v\n", response)
        t.Error(err)
    }

    // assert
    interactions := mockClient.Interactions()
    testutils.AssertEqual(t, len(interactions), 1)
    testutils.AssertEqual(t, interactions[0].Handler, handler)
    testutils.AssertEqual(t, handler.Interactions, 1)

    assertCustomConnectionDetailsResponse(t, response)
}

func TestCustomMergedConnectionDetailsMock(t *testing.T) {
    // arrange
    ftClient, mockClient := testutils.CreateTestClient()
    handler := mockClient.When(http.MethodGet, "/v1/connections/connection_id").ThenCall(

        func(req *http.Request) (*http.Response, error) {
            response := mock.NewResponse(req, http.StatusOK, prepareConnectionDetailsResponse())
            return response, nil
        })

    // act
    response, err := ftClient.NewConnectionDetails().ConnectionID("connection_id").DoCustomMerged(context.Background())

    if err != nil {
        t.Logf("%+v\n", response)
        t.Error(err)
    }

    // assert
    interactions := mockClient.Interactions()
    testutils.AssertEqual(t, len(interactions), 1)
    testutils.AssertEqual(t, interactions[0].Handler, handler)
    testutils.AssertEqual(t, handler.Interactions, 1)

    assertCustomMergedConnectionDetailsResponse(t, response)
}

func prepareConnectionDetailsResponse() string {
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
            "hybrid_deployment_agent_id": "hybrid_deployment_agent_id",
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

func assertConnectionDetailsResponse(t *testing.T, response connections.DetailsWithConfigNoTestsResponse) {

    testutils.AssertEqual(t, response.Code, "Success")

    testutils.AssertEqual(t, response.Data.Config.SecretsList[0].Key, "key")
    testutils.AssertEqual(t, response.Data.Config.SecretsList[0].Value, "value")
    testutils.AssertEqual(t, response.Data.Config.ShareURL, "share_url")
    testutils.AssertEqual(t, *response.Data.Config.IsKeypair, true)
}

func assertCustomConnectionDetailsResponse(t *testing.T, response connections.DetailsWithCustomConfigNoTestsResponse) {

    testutils.AssertEqual(t, response.Code, "Success")

    testutils.AssertKey(t, "share_url", response.Data.Config, "share_url")
    testutils.AssertKey(t, "is_keypair", response.Data.Config, true)

    secretsList, ok := response.Data.Config["secrets_list"].([]interface{})

    testutils.AssertEqual(t, ok, true)
    testutils.AssertEqual(t, len(secretsList), 1)

    secret := secretsList[0].(map[string]interface{})

    testutils.AssertKey(t, "key", secret, "key")
    testutils.AssertKey(t, "value", secret, "value")
}

func assertCustomMergedConnectionDetailsResponse(t *testing.T, response connections.DetailsWithCustomMergedConfigNoTestsResponse) {

    testutils.AssertEqual(t, response.Code, "Success")

    testutils.AssertHasNoKey(t, response.Data.CustomConfig, "share_url")
    testutils.AssertHasNoKey(t, response.Data.CustomConfig, "is_keypair")
    testutils.AssertHasNoKey(t, response.Data.CustomConfig, "secrets_list")

    testutils.AssertKeyValue(t, response.Data.CustomConfig, "fake_field", "unmapped-value")

    testutils.AssertEqual(t, response.Data.Config.SecretsList[0].Key, "key")
    testutils.AssertEqual(t, response.Data.Config.SecretsList[0].Value, "value")
    testutils.AssertEqual(t, response.Data.Config.ShareURL, "share_url")
    testutils.AssertEqual(t, *response.Data.Config.IsKeypair, true)

}
