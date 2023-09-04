package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestExternalLoggingModifyService(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/external-logging/log_id").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			body := requestBodyToJson(t, req)
			assertExternalLoggingModifyRequest(t, body)
			response := mock.NewResponse(req, http.StatusOK, prepareExternalLoggingModifyResponse())
			return response, nil
		})

    // act
    response, err := ftClient.NewExternalLoggingModify().
        ExternalLoggingId("log_id").
        Enabled(EXTLOG_ENABLED).
        Config(prepareExternalLoggingModifyConfig()).
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

	assertExternalLoggingModifyResponse(t, response)
}


func TestExternalLoggingCustomModifyService(t *testing.T) {
    // arrange
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/external-logging/log_id").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			body := requestBodyToJson(t, req)
			assertExternalLoggingModifyCustomRequest(t, body)
			response := mock.NewResponse(req, http.StatusOK, prepareExternalLoggingModifyMergedResponse())
			return response, nil
		})

    // act
    response, err := ftClient.NewExternalLoggingModify().
        ExternalLoggingId("log_id").
        Enabled(EXTLOG_ENABLED).
        ConfigCustom(prepareExternalLoggingCustomMergedConfig()).
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

    assertExternalLoggingModifyCustomResponse(t, response)
}

func TestExternalLoggingCustomMergedModifyService(t *testing.T) {
    // arrange
    ftClient, mockClient := CreateTestClient()
    handler := mockClient.When(http.MethodPatch, "/v1/external-logging/log_id").ThenCall(

        func(req *http.Request) (*http.Response, error) {
            body := requestBodyToJson(t, req)
            assertExternalLoggingModifyCustomMergedRequest(t, body)
            response := mock.NewResponse(req, http.StatusOK, prepareExternalLoggingModifyMergedResponse())
            return response, nil
        })

    // act
    response, err := ftClient.NewExternalLoggingModify().
        ExternalLoggingId("log_id").
        Enabled(EXTLOG_ENABLED).
        Config(prepareExternalLoggingModifyConfig()).
        ConfigCustom(prepareExternalLoggingCustomMergedConfig()).
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

    assertExternalLoggingModifyCustomMergedResponse(t, response)
}

func prepareExternalLoggingModifyResponse() string {
    return fmt.Sprintf(
        `{
			"code": "Success",
			"message": "External logging service has been updated",
            "data":{
                "id": "%v",
                "service": "%v",
                "enabled": %v,
                "config": {
                    "workspace_id": "%v",
                    "primary_key": "%v"
                }
            }
        }`,
        EXTLOG_GROUPID,
        EXTLOG_SERVICE,
        EXTLOG_ENABLED,
        EXTLOG_WORKSPACEID,
        EXTLOG_PRIMARYKEY,
    )
}

func prepareExternalLoggingModifyMergedResponse() string {
    return fmt.Sprintf(
        `{
			"code": "Success",
			"message": "External logging service has been updated",
            "data":{
                "id": "%v",
                "service": "%v",
                "enabled": %v,
                "config": {
                    "workspace_id": "%v",
                    "primary_key": "%v",
                    "fake_field": "unmapped-value"
                }
            }
        }`,
        EXTLOG_GROUPID,
        EXTLOG_SERVICE,
        EXTLOG_ENABLED,
        EXTLOG_WORKSPACEID,
        EXTLOG_PRIMARYKEY,
    )
}

func prepareExternalLoggingModifyConfig() *fivetran.ExternalLoggingConfig {
    config := fivetran.NewExternalLoggingConfig()
    config.WorkspaceId(EXTLOG_WORKSPACEID)
    config.PrimaryKey(EXTLOG_PRIMARYKEY)

    return config
}

func prepareExternalLoggingCustomConfig() *map[string]interface{} {
    config := make(map[string]interface{})

    config["fake_field"] = "unmapped-value"

    return &config
}

// assert Requests
func assertExternalLoggingModifyRequest(t *testing.T, request map[string]interface{}) {
    assertKey(t, "enabled", request, EXTLOG_ENABLED)

    config, ok := request["config"].(map[string]interface{})
    assertEqual(t, ok, true)

    assertKey(t, "workspace_id", config, EXTLOG_WORKSPACEID)
    assertKey(t, "primary_key", config, EXTLOG_PRIMARYKEY)     
}

func assertExternalLoggingModifyCustomRequest(t *testing.T, request map[string]interface{}) {
    assertKey(t, "enabled", request, EXTLOG_ENABLED)

    config, ok := request["config"].(map[string]interface{})
    
    assertEqual(t, ok, true)

    assertKey(t, "fake_field", config, "unmapped-value")
}

func assertExternalLoggingModifyCustomMergedRequest(t *testing.T, request map[string]interface{}) {
    assertKey(t, "enabled", request, EXTLOG_ENABLED)

    config, ok := request["config"].(map[string]interface{})
    
    assertEqual(t, ok, true)

    assertKey(t, "workspace_id", config, EXTLOG_WORKSPACEID)
    assertKey(t, "primary_key", config, EXTLOG_PRIMARYKEY)     
    assertKey(t, "fake_field", config, "unmapped-value")
}


// assert Response
func assertExternalLoggingModifyResponse(t *testing.T, response fivetran.ExternalLoggingModifyResponse) {
	assertEqual(t, response.Code, "Success")
	assertEqual(t, response.Message, "External logging service has been updated")
    assertEqual(t, response.Data.Id, EXTLOG_GROUPID)
    assertEqual(t, response.Data.Service, EXTLOG_SERVICE)
    assertEqual(t, response.Data.Enabled, EXTLOG_ENABLED)

    assertEqual(t, response.Data.Config.WorkspaceId, EXTLOG_WORKSPACEID)
    assertEqual(t, response.Data.Config.PrimaryKey, EXTLOG_PRIMARYKEY)
}

func assertExternalLoggingModifyCustomResponse(t *testing.T, response fivetran.ExternalLoggingModifyCustomResponse) {
	assertEqual(t, response.Code, "Success")
	assertEqual(t, response.Message, "External logging service has been updated")
    assertEqual(t, response.Data.Id, EXTLOG_GROUPID)
    assertEqual(t, response.Data.Service, EXTLOG_SERVICE)
    assertEqual(t, response.Data.Enabled, EXTLOG_ENABLED)

    assertKey(t, "fake_field", response.Data.Config, "unmapped-value")
}

func assertExternalLoggingModifyCustomMergedResponse(t *testing.T, response fivetran.ExternalLoggingModifyCustomMergedResponse) {
	assertEqual(t, response.Code, "Success")
	assertEqual(t, response.Message, "External logging service has been updated")
    assertEqual(t, response.Data.Id, EXTLOG_GROUPID)
    assertEqual(t, response.Data.Service, EXTLOG_SERVICE)
    assertEqual(t, response.Data.Enabled, EXTLOG_ENABLED)

    assertEqual(t, response.Data.Config.WorkspaceId, EXTLOG_WORKSPACEID)
    assertEqual(t, response.Data.Config.PrimaryKey, EXTLOG_PRIMARYKEY)
    assertKey(t, "fake_field", response.Data.CustomConfig, "unmapped-value")
}