package fivetran_test

import (
	"context"
	"testing"
)

func TestWebhookDetailsE2E(t *testing.T) {
	webhookId := CreateWebhookAccount(t)

	result, err := Client.NewWebhookDetails().WebhookId(webhookId).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	AssertEqual(t, result.Code, "Success")
	AssertNotEmpty(t, result.Data.Id)
    AssertNotEmpty(t, result.Data.Events)
    AssertNotEmpty(t, result.Data.CreatedAt)
    AssertNotEmpty(t, result.Data.CreatedBy)
    AssertNotEmpty(t, result.Data.Type)
    AssertNotEmpty(t, result.Data.Secret)
    AssertNotEmpty(t, result.Data.Url)

    t.Cleanup(func() { DeleteWebhook(t, webhookId) })
}
