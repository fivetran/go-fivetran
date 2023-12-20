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

const (
	WEBHOOK_URL    = "https://localhost:12345"
	WEBHOOK_EVENT  = "sync_start"
	WEBHOOK_ACTIVE = false
	WEBHOOK_SECRET = "my_secret"
	WEBHOOK_GROUP  = "test_group"
)

func TestNewWebhookAccountCreate(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/webhooks/account").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertWebhookAccountRequest(t, body)
			response := mock.NewResponse(req, http.StatusOK, prepareWebhookAccountResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewWebhookAccountCreate().
		Url(WEBHOOK_URL).
		Secret(WEBHOOK_SECRET).
		Active(WEBHOOK_ACTIVE).
		Events([]string{WEBHOOK_EVENT}).
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

	assertWebhookAccountResponse(t, response)
}

func TestNewWebhookGroupCreate(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/webhooks/group/"+WEBHOOK_GROUP).ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertWebhookGroupRequest(t, body)
			response := mock.NewResponse(req, http.StatusOK, prepareWebhookGroupResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewWebhookGroupCreate().
		Url(WEBHOOK_URL).
		Secret(WEBHOOK_SECRET).
		Active(WEBHOOK_ACTIVE).
		GroupId(WEBHOOK_GROUP).
		Events([]string{WEBHOOK_EVENT}).
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
	assertWebhookGroupResponse(t, response)
}

func prepareWebhookGroupResponse() string {
	return fmt.Sprintf(
		`{
            "code": "Success",
            "message": "Group webhook has been created",
            "data": {
                "id": "program_quoth",
                "type": "group",
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

func prepareWebhookAccountResponse() string {
	return fmt.Sprintf(
		`{
            "code": "Success",
            "message": "Account webhook has been created",
            "data": {
                "id": "program_quoth",
                "type": "account",
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
		WEBHOOK_URL, // id
		WEBHOOK_EVENT,
		WEBHOOK_ACTIVE,
	)
}

func assertWebhookAccountRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "url", request, WEBHOOK_URL)
	testutils.AssertKey(t, "active", request, WEBHOOK_ACTIVE)
}

func assertWebhookGroupRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "url", request, WEBHOOK_URL)
	testutils.AssertKey(t, "active", request, WEBHOOK_ACTIVE)
}

func assertWebhookAccountResponse(t *testing.T, response webhooks.WebhookResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertNotEmpty(t, response.Message)

	testutils.AssertNotEmpty(t, response.Data.Id)
	testutils.AssertNotEmpty(t, response.Data.CreatedAt)
	testutils.AssertNotEmpty(t, response.Data.CreatedBy)

	testutils.AssertEqual(t, response.Data.Events, []string{WEBHOOK_EVENT})
	testutils.AssertEqual(t, response.Data.Url, WEBHOOK_URL)
	testutils.AssertEqual(t, response.Data.Secret, "******")
	testutils.AssertEqual(t, response.Data.Type, "account")
}

func assertWebhookGroupResponse(t *testing.T, response webhooks.WebhookResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertNotEmpty(t, response.Message)

	testutils.AssertNotEmpty(t, response.Data.Id)
	testutils.AssertNotEmpty(t, response.Data.CreatedAt)
	testutils.AssertNotEmpty(t, response.Data.CreatedBy)

	testutils.AssertEqual(t, response.Data.Url, WEBHOOK_URL)
	testutils.AssertEqual(t, response.Data.Active, WEBHOOK_ACTIVE)
	testutils.AssertEqual(t, response.Data.GroupId, WEBHOOK_GROUP)
	testutils.AssertEqual(t, response.Data.Secret, "******")
	testutils.AssertEqual(t, response.Data.Type, "group")
	testutils.AssertEqual(t, response.Data.Events, []string{WEBHOOK_EVENT})
}
