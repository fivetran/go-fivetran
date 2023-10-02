package fivetran_test

import (
	"context"
	"strings"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewExternalLoggingDeleteE2E(t *testing.T) {
	externalLoggingId := testutils.CreateExternalLogging(t)
	deleted, err := testutils.Client.NewExternalLoggingDelete().ExternalLoggingId(externalLoggingId).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}

	testutils.AssertEqual(t, deleted.Code, "Success")
	testutils.AssertNotEmpty(t, deleted.Message)
	testutils.AssertEqual(t, strings.Contains(deleted.Message, externalLoggingId), true)
}
