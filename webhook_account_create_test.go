package fivetran_test

import (
    "context"
    "testing"
)

func TestNewWebhookAccountCreateE2E(t *testing.T) {
    created, err := Client.NewWebhookAccountCreate().
        Url("https://webhook.site/abe96072-249c-40bc-a12d-8b92750175e2").
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
    AssertEqual(t, created.Data.Url, "https://webhook.site/abe96072-249c-40bc-a12d-8b92750175e2")
    
    t.Cleanup(func() { DeleteWebhook(t, created.Data.Id) })
}
