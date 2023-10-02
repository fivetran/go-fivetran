package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestWebhookTestE2E(t *testing.T) {
	webhookId := testutils.CreateWebhookAccount(t)

	response, err := testutils.Client.NewWebhookTest().
		WebhookId(webhookId).
		Event("sync_start").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.Succeed, false)

	t.Cleanup(func() { testutils.DeleteWebhook(t, webhookId) })
}
