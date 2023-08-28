package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestWebhookModifyService(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "v1/webhooks/webhook_id").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareWebhookModifyResponse())
			return response, nil
		})

	service := ftClient.NewWebhookModify().
        	WebhookId("webhook_id").
        	Url(WEBHOOK_URL).
        	Secret(WEBHOOK_SECRET).
        	Active(WEBHOOK_ACTIVE).
        	Events([]string{WEBHOOK_EVENT})

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

	assertWebhookModifyResponse(t, response)
}

func prepareWebhookModifyResponse() string {
    return fmt.Sprintf(
        `{
            "code": "Success",
            "message": "Webhook has been updated",
            "data": {
                "id": "program_quoth",
                "type": "account",
                "group_id": "%v",
                "url": "%v",
                "events": [
                    "%v"
                ],
                "active": %v,
                "secret": "******",
                "created_at": "2022-04-29T09:41:08.583Z",
                "created_by": "_airworthy"
        }`,
        WEBHOOK_GROUP,
        WEBHOOK_URL,
        WEBHOOK_EVENT,
        WEBHOOK_ACTIVE,
    )
}

func assertWebhookModifyResponse(t *testing.T, response fivetran.WebhookModifyResponse) {
    assertEqual(t, response.Code, "Created")
    assertNotEmpty(t, response.Message)

    assertNotEmpty(t, response.Data.Id)
    assertNotEmpty(t, response.Data.CreatedAt)
    assertNotEmpty(t, response.Data.CreatedBy)

    assertEqual(t, response.Data.Url, WEBHOOK_URL)
    assertEqual(t, response.Data.Active, WEBHOOK_ACTIVE)
    assertEqual(t, response.Data.GroupId, WEBHOOK_GROUP)
    assertEqual(t, response.Data.Secret, "******")
    assertEqual(t, response.Data.Type, "group")
    assertEqual(t, response.Data.Events[0], WEBHOOK_EVENT)
}
