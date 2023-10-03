package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	externallogging "github.com/fivetran/go-fivetran/external_logging"
	"github.com/fivetran/go-fivetran/tests/mock"
)

const (
	EXTLOG_GROUPID = "group_id"
	EXTLOG_SERVICE = "service"
	EXTLOG_ENABLED = true

	EXTLOG_WORKSPACEID  = "workspace_id"
	EXTLOG_PRIMARYKEY   = "primary_key"
	EXTLOG_LOGGROUPNAME = "log_group_name"
	EXTLOG_ROLEARN      = "role_arn"
	EXTLOG_EXTERNALID   = "external_id"
	EXTLOG_REGION       = "region"
	EXTLOG_APIKEY       = "api_key"
	EXTLOG_SUBDOMAIN    = "sub_domain"
	EXTLOG_HOST         = "host"
	EXTLOG_HOSTNAME     = "hostname"
	EXTLOG_CHANNEL      = "channel"
	EXTLOG_ENABLESSL    = false
	EXTLOG_TOKEN        = "token"
	EXTLOG_PORT         = 443
)

func TestNewExternalLoggingCreateFullMappingMock(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/external-logging").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := RequestBodyToJson(t, req)
			assertExternalLoggingFullRequest(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareExternalLoggingResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewExternalLoggingCreate().
		GroupId(EXTLOG_GROUPID).
		Service(EXTLOG_SERVICE).
		Enabled(EXTLOG_ENABLED).
		Config(prepareExternalLoggingConfig()).
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

	assertExternalLoggingResponse(t, response)
}

func TestNewExternalLoggingCustomMappingMock(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/external-logging").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := RequestBodyToJson(t, req)
			assertExternalLoggingCustomRequest(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareExternalLoggingResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewExternalLoggingCreate().
		GroupId(EXTLOG_GROUPID).
		Service(EXTLOG_SERVICE).
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

	assertExternalLoggingCustomResponse(t, response)
}

func TestNewExternalLoggingCustomMergedMappingMock(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/external-logging").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := RequestBodyToJson(t, req)
			assertExternalLoggingCustomMergedRequest(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareExternalLoggingMergedResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewExternalLoggingCreate().
		GroupId(EXTLOG_GROUPID).
		Service(EXTLOG_SERVICE).
		Enabled(EXTLOG_ENABLED).
		Config(prepareExternalLoggingConfig()).
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
	assertExternalLoggingCustomMergedResponse(t, response)
}

func prepareExternalLoggingResponse() string {
	return fmt.Sprintf(
		`{
            "code":"Created",
            "message":"External logging service has been added",
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

func prepareExternalLoggingMergedResponse() string {
	return fmt.Sprintf(
		`{
            "code":"Created",
            "message":"External logging service has been added",
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

func prepareExternalLoggingConfig() *externallogging.ExternalLoggingConfig {
	config := fivetran.NewExternalLoggingConfig()
	config.WorkspaceId(EXTLOG_WORKSPACEID)
	config.PrimaryKey(EXTLOG_PRIMARYKEY)

	return config
}

func prepareExternalLoggingCustomMergedConfig() *map[string]interface{} {
	config := make(map[string]interface{})

	config["fake_field"] = "unmapped-value"

	return &config
}

func assertExternalLoggingFullRequest(t *testing.T, request map[string]interface{}) {
	assertKey(t, "service", request, EXTLOG_SERVICE)
	assertKey(t, "group_id", request, EXTLOG_GROUPID)
	assertKey(t, "enabled", request, EXTLOG_ENABLED)

	config, ok := request["config"].(map[string]interface{})
	assertEqual(t, ok, true)

	assertKey(t, "workspace_id", config, EXTLOG_WORKSPACEID)
	assertKey(t, "primary_key", config, EXTLOG_PRIMARYKEY)
}

func assertExternalLoggingCustomRequest(t *testing.T, request map[string]interface{}) {
	assertKey(t, "service", request, EXTLOG_SERVICE)
	assertKey(t, "group_id", request, EXTLOG_GROUPID)
	assertKey(t, "enabled", request, EXTLOG_ENABLED)

	config, ok := request["config"].(map[string]interface{})

	assertEqual(t, ok, true)

	assertKey(t, "fake_field", config, "unmapped-value")
}

func assertExternalLoggingCustomMergedRequest(t *testing.T, request map[string]interface{}) {
	assertKey(t, "service", request, EXTLOG_SERVICE)
	assertKey(t, "group_id", request, EXTLOG_GROUPID)
	assertKey(t, "enabled", request, EXTLOG_ENABLED)

	config, ok := request["config"].(map[string]interface{})

	assertEqual(t, ok, true)

	assertKey(t, "workspace_id", config, EXTLOG_WORKSPACEID)
	assertKey(t, "primary_key", config, EXTLOG_PRIMARYKEY)
	assertKey(t, "fake_field", config, "unmapped-value")
}

func assertExternalLoggingResponse(t *testing.T, response externallogging.ExternalLoggingResponse) {
	assertEqual(t, response.Code, "Created")
	assertNotEmpty(t, response.Message)

	assertEqual(t, response.Data.Id, EXTLOG_GROUPID)
	assertEqual(t, response.Data.Service, EXTLOG_SERVICE)
	assertEqual(t, response.Data.Enabled, EXTLOG_ENABLED)

	assertEqual(t, response.Data.Config.WorkspaceId, EXTLOG_WORKSPACEID)
	assertEqual(t, response.Data.Config.PrimaryKey, EXTLOG_PRIMARYKEY)
}

func assertExternalLoggingCustomResponse(t *testing.T, response externallogging.ExternalLoggingCustomResponse) {
	assertEqual(t, response.Code, "Created")
	assertNotEmpty(t, response.Message)

	assertEqual(t, response.Data.Id, EXTLOG_GROUPID)
	assertEqual(t, response.Data.Service, EXTLOG_SERVICE)
	assertEqual(t, response.Data.Enabled, EXTLOG_ENABLED)

	assertEqual(t, response.Data.Config["workspace_id"], EXTLOG_WORKSPACEID)
	assertEqual(t, response.Data.Config["primary_key"], EXTLOG_PRIMARYKEY)
}

func assertExternalLoggingCustomMergedResponse(t *testing.T, response externallogging.ExternalLoggingCustomMergedResponse) {
	assertEqual(t, response.Code, "Created")
	assertNotEmpty(t, response.Message)

	assertEqual(t, response.Data.Id, EXTLOG_GROUPID)
	assertEqual(t, response.Data.Service, EXTLOG_SERVICE)
	assertEqual(t, response.Data.Enabled, EXTLOG_ENABLED)

	assertEqual(t, response.Data.Config.WorkspaceId, EXTLOG_WORKSPACEID)
	assertEqual(t, response.Data.Config.PrimaryKey, EXTLOG_PRIMARYKEY)
	assertKey(t, "fake_field", response.Data.CustomConfig, "unmapped-value")
}
