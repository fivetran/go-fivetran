package fivetran_test

import (
    "context"
    "strings"
    "testing"
)

func TestNewExternalLoggingDeleteE2E(t *testing.T) {
    externalLoggingId := CreateExternalLogging(t)
    deleted, err := Client.NewExternalLoggingDelete().ExternalLoggingId(externalLoggingId).Do(context.Background())

    if err != nil {
        t.Logf("%+v\n", deleted)
        t.Error(err)
    }

    AssertEqual(t, deleted.Code, "Success")
    AssertNotEmpty(t, deleted.Message)
    AssertEqual(t, strings.Contains(deleted.Message, externalLoggingId), true)
}
