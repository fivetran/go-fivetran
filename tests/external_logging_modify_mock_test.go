package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

const (
    GROUPID         = "group_id"
    SERVICE         = "service"
    ENABLED         = true

    WORKSPACEID     = "workspace_id"
    PRIMARYKEY      = "primary_key"
    LOGGROUPNAME    = "log_group_name"
    ROLEARN         = "role_arn"
    EXTERNALID      = "external_id"
    REGION          = "region"
    APIKEY          = "api_key"
    SUBDOMAIN       = "sub_domain"
    HOST            = "host"
    HOSTNAME        = "hostname"
    CHANNEL         = "channel"
    ENABLESSL       = "enable_ssl"
    TOKEN           = "token"
    PORT            = 443
)

func TestExternalLoggingModifyService(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/external-logging/log_id").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareExternalLoggingModifyResponse())
			return response, nil
		})

	externalLoggingConfig := fivetran.NewExternalLoggingConfig()
	externalLoggingConfig = externalLoggingConfig.
			WorkspaceId(WORKSPACEID).
			PrimaryKey(PRIMARYKEY)

	service := ftClient.NewExternalLoggingModify().
		Enabled(true).
		Config(externalLoggingConfig)

	// act
	response, err := service.Do(context.Background())

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

func assertExternalLoggingModifyResponse() string {
	return fmt.Sprintf(`{
		"code": "Success",
		"message": "External logging service has been updated",
		"data": {
			"id": "%v",
			"service": "%v",
			"enabled": "%v"
		}
	}`,
    GROUPID,
    SERVICE,
    ENABLED
    )
}

func assertExternalLoggingModifyResponse(t *testing.T, response fivetran.ExternalLoggingModifyResponse) {
	assertEqual(t, response.Code, "Success")
	assertEqual(t, response.Message, "External logging service has been updated")
    assertEqual(t, response.Data.Id, GROUPID)
    assertEqual(t, response.Data.Service, SERVICE)
    assertEqual(t, response.Data.Enabled, ENABLED)
}
