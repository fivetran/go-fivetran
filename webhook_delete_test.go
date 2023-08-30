package fivetran_test

import (
    "context"
    "strings"
    "testing"
)

func TestNewWebhookDeleteE2E(t *testing.T) {
    webhookId := CreateWebhookAccount(t)

    deleted, err := Client.NewWebhookDelete().WebhookId(webhookId).Do(context.Background())

    if err != nil {
        t.Logf("%+v\n", deleted)
        t.Error(err)
    }

    AssertEqual(t, deleted.Code, "Success")
    AssertNotEmpty(t, deleted.Message)
    AssertEqual(t, strings.Contains(deleted.Message, webhookId), true)
}
