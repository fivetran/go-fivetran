package externallogging_test

import (
    "context"
    "fmt"
    "net/http"
    "testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/external_logging"
    
    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestExternalLoggingUpdateService(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/external-logging/log_id").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertExternalLoggingUpdateRequest(t, body)
			response := mock.NewResponse(req, http.StatusOK, prepareExternalLoggingUpdateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewExternalLoggingUpdate().
		ExternalLoggingId("log_id").
		Enabled(EXTLOG_ENABLED).
		Config(prepareExternalLoggingUpdateConfig()).
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

	assertExternalLoggingUpdateResponse(t, response)
}

func TestExternalLoggingCustomUpdateService(t *testing.T) {
	// arrange
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/external-logging/log_id").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertExternalLoggingUpdateCustomRequest(t, body)
			response := mock.NewResponse(req, http.StatusOK, prepareExternalLoggingUpdateMergedResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewExternalLoggingUpdate().
		ExternalLoggingId("log_id").
		Enabled(EXTLOG_ENABLED).
		ConfigCustom(prepareExternalLoggingCustomConfig()).
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

	assertExternalLoggingUpdateCustomResponse(t, response)
}

func TestExternalLoggingCustomMergedUpdateService(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/external-logging/log_id").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertExternalLoggingUpdateCustomMergedRequest(t, body)
			response := mock.NewResponse(req, http.StatusOK, prepareExternalLoggingUpdateMergedResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewExternalLoggingUpdate().
		ExternalLoggingId("log_id").
		Enabled(EXTLOG_ENABLED).
		Config(prepareExternalLoggingUpdateConfig()).
		ConfigCustom(prepareExternalLoggingCustomMergedConfig()).
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

	assertExternalLoggingUpdateCustomMergedResponse(t, response)
}

func prepareExternalLoggingUpdateResponse() string {
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

func prepareExternalLoggingUpdateMergedResponse() string {
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

func prepareExternalLoggingUpdateConfig() *externallogging.ExternalLoggingConfig {
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
func assertExternalLoggingUpdateRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "enabled", request, EXTLOG_ENABLED)

	config, ok := request["config"].(map[string]interface{})
	testutils.AssertEqual(t, ok, true)

	testutils.AssertKey(t, "workspace_id", config, EXTLOG_WORKSPACEID)
	testutils.AssertKey(t, "primary_key", config, EXTLOG_PRIMARYKEY)
}

func assertExternalLoggingUpdateCustomRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "enabled", request, EXTLOG_ENABLED)

	config, ok := request["config"].(map[string]interface{})

	testutils.AssertEqual(t, ok, true)

	testutils.AssertKey(t, "fake_field", config, "unmapped-value")
}

func assertExternalLoggingUpdateCustomMergedRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "enabled", request, EXTLOG_ENABLED)

	config, ok := request["config"].(map[string]interface{})

	testutils.AssertEqual(t, ok, true)

	testutils.AssertKey(t, "workspace_id", config, EXTLOG_WORKSPACEID)
	testutils.AssertKey(t, "primary_key", config, EXTLOG_PRIMARYKEY)
	testutils.AssertKey(t, "fake_field", config, "unmapped-value")
}

// assert Response
func assertExternalLoggingUpdateResponse(t *testing.T, response externallogging.ExternalLoggingResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Message, "External logging service has been updated")
	testutils.AssertEqual(t, response.Data.Id, EXTLOG_GROUPID)
	testutils.AssertEqual(t, response.Data.Service, EXTLOG_SERVICE)
	testutils.AssertEqual(t, response.Data.Enabled, EXTLOG_ENABLED)

	testutils.AssertEqual(t, response.Data.Config.WorkspaceId, EXTLOG_WORKSPACEID)
	testutils.AssertEqual(t, response.Data.Config.PrimaryKey, EXTLOG_PRIMARYKEY)
}

func assertExternalLoggingUpdateCustomResponse(t *testing.T, response externallogging.ExternalLoggingCustomResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Message, "External logging service has been updated")
	testutils.AssertEqual(t, response.Data.Id, EXTLOG_GROUPID)
	testutils.AssertEqual(t, response.Data.Service, EXTLOG_SERVICE)
	testutils.AssertEqual(t, response.Data.Enabled, EXTLOG_ENABLED)

	testutils.AssertKey(t, "fake_field", response.Data.Config, "unmapped-value")
}

func assertExternalLoggingUpdateCustomMergedResponse(t *testing.T, response externallogging.ExternalLoggingCustomMergedResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Message, "External logging service has been updated")
	testutils.AssertEqual(t, response.Data.Id, EXTLOG_GROUPID)
	testutils.AssertEqual(t, response.Data.Service, EXTLOG_SERVICE)
	testutils.AssertEqual(t, response.Data.Enabled, EXTLOG_ENABLED)

	testutils.AssertEqual(t, response.Data.Config.WorkspaceId, EXTLOG_WORKSPACEID)
	testutils.AssertEqual(t, response.Data.Config.PrimaryKey, EXTLOG_PRIMARYKEY)
	testutils.AssertKey(t, "fake_field", response.Data.CustomConfig, "unmapped-value")
}
