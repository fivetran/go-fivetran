package fivetran_test

import (
	"context"
	"testing"
)

func TestWebhookModifyE2E(t *testing.T) {
	webhookId := CreateWebhookAccount(t)

	updated, err := Client.NewWebhookModify().
            WebhookId(webhookId).
            Url("https://localhost:12345").
            Secret("my_secret").
            Active(false).
            Events([]string{"sync_start","sync_end"}).
		Do(context.Background())
	

	if err != nil {
		t.Logf("%+v\n", updated)
		t.Error(err)
	}

    AssertEqual(t, updated.Code, "Success")
    AssertNotEmpty(t, updated.Message)
    AssertEqual(t, updated.Data.Id, webhookId)
    AssertNotEmpty(t, updated.Data.CreatedAt)
    AssertNotEmpty(t, updated.Data.CreatedBy)
    AssertNotEmpty(t, updated.Data.Type)
    AssertEqual(t, updated.Data.Active, false)
    AssertEqual(t, updated.Data.Secret, "******")
    AssertEqual(t, updated.Data.Url, "https://localhost:12345")
    AssertEqual(t, updated.Data.Events, []string{"sync_start","sync_end"})

    t.Cleanup(func() { DeleteWebhook(t, webhookId) })
}