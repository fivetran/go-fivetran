package fivetran_test

import (
	"context"
	"testing"
)

func TestWebhookModifyE2E(t *testing.T) {
	webhookId := CreateTempWebhook(t)

	updated, err := Client.NewWebhookModify().
            WebhookId(webhookId).
            Url("https://webhook.site/abe96072-249c-40bc-a12d-8b92750175e2").                   // Unstable test url
            Secret("my_secret").
            Active(true).
            Events([]string{"sync_start","sync_end"}).
		Do(context.Background())
	

	if err != nil {
		t.Logf("%+v\n", updated)
		t.Error(err)
	}

    AssertEqual(t, updated.Code, "Success")
    AssertNotEmpty(t, updated.Message)
    AssertEqual(t, updated.Data.Id, webhookId)
    AssertNotEmpty(t, updated.Data.Events)
    AssertNotEmpty(t, updated.Data.CreatedAt)
    AssertNotEmpty(t, updated.Data.CreatedBy)
    AssertNotEmpty(t, updated.Data.Type)
    AssertEqual(t, updated.Data.Active, true)
    AssertEqual(t, updated.Data.Secret, "******")
    AssertEqual(t, updated.Data.Url, "https://webhook.site/abe96072-249c-40bc-a12d-8b92750175e2")           // Unstable test url

    t.Cleanup(func() { DeleteWebhook(t, webhookId) })
}