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
    AssertEqual(t, result.Data.Id, webhookId)
    AssertNotEmpty(t, result.Data.CreatedAt)
    AssertNotEmpty(t, result.Data.CreatedBy)
    AssertEmpty(t, result.Data.GroupId)
    AssertEqual(t, result.Data.Type, "account")
    AssertEqual(t, result.Data.Secret, "******")
    AssertEqual(t, result.Data.Url, "https://localhost:12345")
    AssertEqual(t, result.Data.Events, []string{"sync_start","sync_end"})
    AssertEqual(t, result.Data.Active, false)    

    t.Cleanup(func() { DeleteWebhook(t, webhookId) })
}