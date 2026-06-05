package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestWebhookUpdateE2E(t *testing.T) {
	webhookId := testutils.CreateWebhookAccount(t)

	updated, err := testutils.Client.NewWebhookUpdate().
		WebhookId(webhookId).
		Url("https://example.com").
		Secret("my_secret").
		Active(false).
		Events([]string{"sync_start", "sync_end"}).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", updated)
		t.Error(err)
	}

	testutils.AssertEqual(t, updated.Code, "Success")
	testutils.AssertNotEmpty(t, updated.Message)
	testutils.AssertEqual(t, updated.Data.Id, webhookId)
	testutils.AssertNotEmpty(t, updated.Data.CreatedAt)
	testutils.AssertNotEmpty(t, updated.Data.CreatedBy)
	testutils.AssertNotEmpty(t, updated.Data.Type)
	testutils.AssertEqual(t, updated.Data.Active, false)
	testutils.AssertEqual(t, updated.Data.Secret, "******")
	testutils.AssertEqual(t, updated.Data.Url, "https://example.com")
	testutils.AssertEqual(t, updated.Data.Events, []string{"sync_start", "sync_end"})

	t.Cleanup(func() { testutils.DeleteWebhook(t, webhookId) })
}
