package fivetran_test

import (
	"context"
	"strings"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewWebhookDeleteE2E(t *testing.T) {
	webhookId := testutils.CreateWebhookAccount(t)

	deleted, err := testutils.Client.NewWebhookDelete().WebhookId(webhookId).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}

	testutils.AssertEqual(t, deleted.Code, "Success")
	testutils.AssertNotEmpty(t, deleted.Message)
	testutils.AssertEqual(t, strings.Contains(deleted.Message, webhookId), true)
}
