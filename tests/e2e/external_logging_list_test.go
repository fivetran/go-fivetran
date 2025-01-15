package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewExternalLoggingListE2E(t *testing.T) {
	externalLoggingId := testutils.CreateTempExternalLogging(t)
	result, err := testutils.Client.NewExternalLoggingList().Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	testutils.AssertEqual(t, result.Code, "Success")
	testutils.AssertHasLength(t, result.Data.Items, 1)
	testutils.AssertEmpty(t, result.Message)
	testutils.AssertEqual(t, result.Data.Items[0].Id, externalLoggingId)
	testutils.AssertEqual(t, result.Data.Items[0].Service, "azure_monitor_log")
	testutils.AssertEqual(t, result.Data.Items[0].Enabled, true)

	testutils.AssertEmpty(t, result.Data.NextCursor)
}
