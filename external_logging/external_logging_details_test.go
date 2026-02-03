package externallogging_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	externallogging "github.com/fivetran/go-fivetran/external_logging"
	testutils "github.com/fivetran/go-fivetran/test_utils"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestExternalLoggingDetailsService(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/external-logging/log_id").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req,
				http.StatusOK,
				prepareExternalLoggingDetailsMergedResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewExternalLoggingDetails().
		ExternalLoggingId("log_id").
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

	assertExternalLoggingDetailsResponse(t, response)
}

func TestExternalLoggingCustomDetailsService(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/external-logging/log_id").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req,
				http.StatusOK,
				prepareExternalLoggingDetailsMergedResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewExternalLoggingDetails().
		ExternalLoggingId("log_id").
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

	assertExternalLoggingDetailsCustomResponse(t, response)
}

func TestExternalLoggingCustomMergedDetailsService(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/external-logging/log_id").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req,
				http.StatusOK,
				prepareExternalLoggingDetailsMergedResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewExternalLoggingDetails().
		ExternalLoggingId("log_id").
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

	assertExternalLoggingDetailsCustomMergedResponse(t, response)
}

func assertExternalLoggingDetailsResponse(t *testing.T, response externallogging.ExternalLoggingResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.Id, EXTLOG_GROUPID)
	testutils.AssertEqual(t, response.Data.Service, EXTLOG_SERVICE)
	testutils.AssertEqual(t, response.Data.Enabled, EXTLOG_ENABLED)

	testutils.AssertEqual(t, response.Data.Config.WorkspaceId, EXTLOG_WORKSPACEID)
	testutils.AssertEqual(t, response.Data.Config.PrimaryKey, EXTLOG_PRIMARYKEY)
}

func assertExternalLoggingDetailsCustomResponse(t *testing.T, response externallogging.ExternalLoggingCustomResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.Id, EXTLOG_GROUPID)
	testutils.AssertEqual(t, response.Data.Service, EXTLOG_SERVICE)
	testutils.AssertEqual(t, response.Data.Enabled, EXTLOG_ENABLED)

	testutils.AssertKey(t, "workspace_id", response.Data.Config, EXTLOG_WORKSPACEID)
	testutils.AssertKey(t, "primary_key", response.Data.Config, EXTLOG_PRIMARYKEY)
	testutils.AssertKey(t, "fake_field", response.Data.Config, "unmapped-value")
}

func assertExternalLoggingDetailsCustomMergedResponse(t *testing.T, response externallogging.ExternalLoggingCustomMergedResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.Id, EXTLOG_GROUPID)
	testutils.AssertEqual(t, response.Data.Service, EXTLOG_SERVICE)
	testutils.AssertEqual(t, response.Data.Enabled, EXTLOG_ENABLED)

	testutils.AssertEqual(t, response.Data.Config.WorkspaceId, EXTLOG_WORKSPACEID)
	testutils.AssertEqual(t, response.Data.Config.PrimaryKey, EXTLOG_PRIMARYKEY)
	testutils.AssertHasNoKey(t, response.Data.CustomConfig, "workspace_id")
	testutils.AssertHasNoKey(t, response.Data.CustomConfig, "primary_key")
	testutils.AssertKey(t, "fake_field", response.Data.CustomConfig, "unmapped-value")
}

func prepareExternalLoggingDetailsMergedResponse() string {
	return fmt.Sprintf(
		`{
			"code": "Success",
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
