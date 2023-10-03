package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestWebhookListE2E(t *testing.T) {
	webhookId := testutils.CreateWebhookAccount(t)

	result, err := testutils.Client.NewWebhookList().Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	testutils.AssertEqual(t, result.Code, "Success")
	testutils.AssertNotEmpty(t, result.Data.Items[0].Id)
	testutils.AssertNotEmpty(t, result.Data.Items[0].Events)
	testutils.AssertNotEmpty(t, result.Data.Items[0].CreatedAt)
	testutils.AssertNotEmpty(t, result.Data.Items[0].CreatedBy)
	testutils.AssertNotEmpty(t, result.Data.Items[0].Type)
	testutils.AssertNotEmpty(t, result.Data.Items[0].Secret)
	testutils.AssertNotEmpty(t, result.Data.Items[0].Url)

	t.Cleanup(func() { testutils.DeleteWebhook(t, webhookId) })
}
