package fivetran_test

import (
	"context"
	"testing"
)

func TestWebhookTestsE2E(t *testing.T) {
	webhookId := CreateWebhookAccount(t)
	
	response, err := Client.NewWebhookTest().
			WebhookId(webhookId).
			Event("sync_start").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	AssertEqual(t, response.Code, "Success")
	AssertEqual(t, response.Data.Succeed, false)

	t.Cleanup(func() { DeleteWebhook(t, webhookId) })
}
