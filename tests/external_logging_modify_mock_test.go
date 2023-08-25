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
			response := mock.NewResponse(req, http.StatusOK, prepareExternalLoggingModifyResponse())
			return response, nil
		})

	externalLoggingConfig := fivetran.NewExternalLoggingConfig()
	externalLoggingConfig = externalLoggingConfig.
			WorkspaceId(EXTLOG_WORKSPACEID).
			PrimaryKey(EXTLOG_PRIMARYKEY)

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

func prepareExternalLoggingModifyResponse() string {
	return fmt.Sprintf(`{
				"code": "Success",
				"message": "External logging service has been updated",
				"data": {
					"id": "%v",
					"service": "%v",
					"enabled": %v
				}
			}`,
    	EXTLOG_GROUPID,
    	EXTLOG_SERVICE,
    	EXTLOG_ENABLED,
    )
}

func assertExternalLoggingModifyResponse(t *testing.T, response fivetran.ExternalLoggingModifyResponse) {
	assertEqual(t, response.Code, "Success")
	assertEqual(t, response.Message, "External logging service has been updated")
    assertEqual(t, response.Data.Id, EXTLOG_GROUPID)
    assertEqual(t, response.Data.Service, EXTLOG_SERVICE)
    assertEqual(t, response.Data.Enabled, EXTLOG_ENABLED)
}
