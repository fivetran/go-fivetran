package webhooks_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/common"
	
	"github.com/fivetran/go-fivetran/tests/mock"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestWebhookDeleteServiceDo(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodDelete, "/v1/webhooks/webhook_id").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, `{"code": "Success", "message": "Webhook with id 'webhook_id' has been deleted"}`)
			return response, nil
		},
	)

	service := ftClient.NewWebhookDelete().WebhookId("webhook_id")

	// act
	response, err := service.Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	assertWebhookDeleteResponse(t, response, "Success", "Webhook with id 'webhook_id' has been deleted")

	// Check that the expected interactions with the mock client occurred
	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
}

func TestWebhookDeleteServiceDoMissingId(t *testing.T) {
	// Create a test client
	ftClient, _ := testutils.CreateTestClient()

	// Create the WebhookDeleteService without setting the webhook ID
	service := ftClient.NewWebhookDelete()

	// Call the Do method to execute the request
	_, err := service.Do(context.Background())

	// Check for expected error
	expectedError := fmt.Errorf("missing required webhookId")
	testutils.AssertEqual(t, err, expectedError)
}

func assertWebhookDeleteResponse(t *testing.T, response common.CommonResponse, code string, massage string) {
	testutils.AssertEqual(t, response.Code, code)
	testutils.AssertEqual(t, response.Message, massage)
}
