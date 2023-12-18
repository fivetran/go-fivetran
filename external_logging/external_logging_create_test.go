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
	EXTLOG_PROJECTID    = "project_id"
)

func TestNewExternalLoggingCreateFullMappingMock(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/external-logging").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
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
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)

	assertExternalLoggingResponse(t, response)
}

func TestNewExternalLoggingCustomMappingMock(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/external-logging").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
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
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)

	assertExternalLoggingCustomResponse(t, response)
}

func TestNewExternalLoggingCustomMergedMappingMock(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/external-logging").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
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
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
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
	testutils.AssertKey(t, "service", request, EXTLOG_SERVICE)
	testutils.AssertKey(t, "group_id", request, EXTLOG_GROUPID)
	testutils.AssertKey(t, "enabled", request, EXTLOG_ENABLED)

	config, ok := request["config"].(map[string]interface{})
	testutils.AssertEqual(t, ok, true)

	testutils.AssertKey(t, "workspace_id", config, EXTLOG_WORKSPACEID)
	testutils.AssertKey(t, "primary_key", config, EXTLOG_PRIMARYKEY)
}

func assertExternalLoggingCustomRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "service", request, EXTLOG_SERVICE)
	testutils.AssertKey(t, "group_id", request, EXTLOG_GROUPID)
	testutils.AssertKey(t, "enabled", request, EXTLOG_ENABLED)

	config, ok := request["config"].(map[string]interface{})

	testutils.AssertEqual(t, ok, true)

	testutils.AssertKey(t, "fake_field", config, "unmapped-value")
}

func assertExternalLoggingCustomMergedRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "service", request, EXTLOG_SERVICE)
	testutils.AssertKey(t, "group_id", request, EXTLOG_GROUPID)
	testutils.AssertKey(t, "enabled", request, EXTLOG_ENABLED)

	config, ok := request["config"].(map[string]interface{})

	testutils.AssertEqual(t, ok, true)

	testutils.AssertKey(t, "workspace_id", config, EXTLOG_WORKSPACEID)
	testutils.AssertKey(t, "primary_key", config, EXTLOG_PRIMARYKEY)
	testutils.AssertKey(t, "fake_field", config, "unmapped-value")
}

func assertExternalLoggingResponse(t *testing.T, response externallogging.ExternalLoggingResponse) {
	testutils.AssertEqual(t, response.Code, "Created")
	testutils.AssertNotEmpty(t, response.Message)

	testutils.AssertEqual(t, response.Data.Id, EXTLOG_GROUPID)
	testutils.AssertEqual(t, response.Data.Service, EXTLOG_SERVICE)
	testutils.AssertEqual(t, response.Data.Enabled, EXTLOG_ENABLED)

	testutils.AssertEqual(t, response.Data.Config.WorkspaceId, EXTLOG_WORKSPACEID)
	testutils.AssertEqual(t, response.Data.Config.PrimaryKey, EXTLOG_PRIMARYKEY)
}

func assertExternalLoggingCustomResponse(t *testing.T, response externallogging.ExternalLoggingCustomResponse) {
	testutils.AssertEqual(t, response.Code, "Created")
	testutils.AssertNotEmpty(t, response.Message)

	testutils.AssertEqual(t, response.Data.Id, EXTLOG_GROUPID)
	testutils.AssertEqual(t, response.Data.Service, EXTLOG_SERVICE)
	testutils.AssertEqual(t, response.Data.Enabled, EXTLOG_ENABLED)

	testutils.AssertEqual(t, response.Data.Config["workspace_id"], EXTLOG_WORKSPACEID)
	testutils.AssertEqual(t, response.Data.Config["primary_key"], EXTLOG_PRIMARYKEY)
}

func assertExternalLoggingCustomMergedResponse(t *testing.T, response externallogging.ExternalLoggingCustomMergedResponse) {
	testutils.AssertEqual(t, response.Code, "Created")
	testutils.AssertNotEmpty(t, response.Message)

	testutils.AssertEqual(t, response.Data.Id, EXTLOG_GROUPID)
	testutils.AssertEqual(t, response.Data.Service, EXTLOG_SERVICE)
	testutils.AssertEqual(t, response.Data.Enabled, EXTLOG_ENABLED)

	testutils.AssertEqual(t, response.Data.Config.WorkspaceId, EXTLOG_WORKSPACEID)
	testutils.AssertEqual(t, response.Data.Config.PrimaryKey, EXTLOG_PRIMARYKEY)
	testutils.AssertKey(t, "fake_field", response.Data.CustomConfig, "unmapped-value")
}
