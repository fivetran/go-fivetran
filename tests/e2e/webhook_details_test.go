package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestWebhookDetailsE2E(t *testing.T) {
	webhookId := testutils.CreateWebhookAccount(t)

	result, err := testutils.Client.NewWebhookDetails().WebhookId(webhookId).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	testutils.AssertEqual(t, result.Code, "Success")
	testutils.AssertEqual(t, result.Data.Id, webhookId)
	testutils.AssertNotEmpty(t, result.Data.CreatedAt)
	testutils.AssertNotEmpty(t, result.Data.CreatedBy)
	testutils.AssertEmpty(t, result.Data.GroupId)
	testutils.AssertEqual(t, result.Data.Type, "account")
	testutils.AssertEqual(t, result.Data.Secret, "******")
	testutils.AssertEqual(t, result.Data.Url, "https://localhost:12345")
	testutils.AssertEqual(t, result.Data.Events, []string{"sync_start", "sync_end"})
	testutils.AssertEqual(t, result.Data.Active, false)

	t.Cleanup(func() { testutils.DeleteWebhook(t, webhookId) })
}
