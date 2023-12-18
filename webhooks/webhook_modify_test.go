package webhooks_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	
	"github.com/fivetran/go-fivetran/tests/mock"
	"github.com/fivetran/go-fivetran/webhooks"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestWebhookModifyService(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/webhooks/webhook_id").ThenCall(
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
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)

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
            }
        }`,
		WEBHOOK_GROUP,
		WEBHOOK_URL,
		WEBHOOK_EVENT,
		WEBHOOK_ACTIVE,
	)
}

func assertWebhookModifyResponse(t *testing.T, response webhooks.WebhookResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertNotEmpty(t, response.Message)

	testutils.AssertNotEmpty(t, response.Data.Id)
	testutils.AssertNotEmpty(t, response.Data.CreatedAt)
	testutils.AssertNotEmpty(t, response.Data.CreatedBy)

	testutils.AssertEqual(t, response.Data.Url, WEBHOOK_URL)
	testutils.AssertEqual(t, response.Data.Active, WEBHOOK_ACTIVE)
	testutils.AssertEqual(t, response.Data.GroupId, WEBHOOK_GROUP)
	testutils.AssertEqual(t, response.Data.Secret, "******")
	testutils.AssertEqual(t, response.Data.Type, "account")
	testutils.AssertEqual(t, response.Data.Events, []string{WEBHOOK_EVENT})
}
