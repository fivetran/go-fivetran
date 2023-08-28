package fivetran_test

import (
	"context"
	"testing"
)

func TestWebhookTestE2E(t *testing.T) {
	webhookId := CreateTempWebhook(t)
	
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
	AssertEqual(t, response.Data.Status, 200)
	AssertEqual(t, response.Data.Message, "SUCCESS")

	t.Cleanup(func() { DeleteWebhook(t, webhookId) })
}
