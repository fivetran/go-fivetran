package fivetran_test

import (
    "context"
    "testing"
)

func TestNewWebhookGroupCreateE2E(t *testing.T) {
    created, err := Client.NewWebhookGroupCreate().
        Url("https://localhost:12345").
        Secret("my_secret").
        GroupId(PredefinedGroupId).
        Active(false).
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
    AssertEqual(t, created.Data.Type, "group")
    AssertEqual(t, created.Data.Active, false)
    AssertEqual(t, created.Data.GroupId, PredefinedGroupId)
    AssertEqual(t, created.Data.Secret, "******")
    AssertEqual(t, created.Data.Url, "https://localhost:12345")
    AssertEqual(t, created.Data.Events, []string{"sync_start","sync_end"})
    
    t.Cleanup(func() { DeleteWebhook(t, created.Data.Id) })
}
