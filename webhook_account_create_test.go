package fivetran_test

import (
    "context"
    "testing"
)

func TestNewWebhookAccountCreateE2E(t *testing.T) {
    created, err := Client.NewWebhookAccountCreate().
        Url("https://your-host.your-domain/webhook").
        Secret("my_secret").
        Active(true).
        Events([]string{"sync_start","sync_end"}).
        Do(context.Background())

    if err != nil {
        t.Logf("%+v\n", created)
        t.Error(err)
    }

    AssertEqual(t, created.Code, "Success")
    AssertNotEmpty(t, created.Message)
    AssertNotEmpty(t, created.Data.Id)
    AssertNotEmpty(t, created.Data.Events)
    AssertNotEmpty(t, created.Data.CreatedAt)
    AssertNotEmpty(t, created.Data.CreatedBy)
    AssertEqual(t, created.Data.Type, "account")
    AssertEqual(t, created.Data.Active, true)
    AssertEqual(t, created.Data.Secret, "******")
    AssertEqual(t, created.Data.Url, "https://your-host.your-domain/webhook")
    
    t.Cleanup(func() { DeleteWebhook(t, created.Data.Id) })
}
