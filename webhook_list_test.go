package fivetran_test

import (
	"context"
	"testing"
)

func TestWebhookListE2E(t *testing.T) {
	webhookId := CreateWebhookAccount(t)

	result, err := Client.NewWebhookList().Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	AssertEqual(t, result.Code, "Success")
	AssertNotEmpty(t, result.Data.Items[0].Id)
    AssertNotEmpty(t, result.Data.Items[0].Events)
    AssertNotEmpty(t, result.Data.Items[0].CreatedAt)
    AssertNotEmpty(t, result.Data.Items[0].CreatedBy)
    AssertNotEmpty(t, result.Data.Items[0].Type)
    AssertNotEmpty(t, result.Data.Items[0].Secret)
    AssertNotEmpty(t, result.Data.Items[0].Url)

    t.Cleanup(func() { DeleteWebhook(t, webhookId) })
}
