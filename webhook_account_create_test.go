package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewWebhookAccountCreateE2E(t *testing.T) {
	created, err := testutils.Client.NewWebhookAccountCreate().
		Url("https://localhost:12345").
		Secret("my_secret").
		Active(false).
		Events([]string{"sync_start", "sync_end"}).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}

	testutils.AssertEqual(t, created.Code, "Success")
	testutils.AssertNotEmpty(t, created.Message)
	testutils.AssertNotEmpty(t, created.Data.Id)
	testutils.AssertNotEmpty(t, created.Data.Events)
	testutils.AssertNotEmpty(t, created.Data.CreatedAt)
	testutils.AssertNotEmpty(t, created.Data.CreatedBy)
	testutils.AssertEqual(t, created.Data.Type, "account")
	testutils.AssertEqual(t, created.Data.Active, false)
	testutils.AssertEqual(t, created.Data.Secret, "******")
	testutils.AssertEqual(t, created.Data.Url, "https://localhost:12345")
	testutils.AssertEqual(t, created.Data.Events, []string{"sync_start", "sync_end"})

	t.Cleanup(func() { testutils.DeleteWebhook(t, created.Data.Id) })
}
