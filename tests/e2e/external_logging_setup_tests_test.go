package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewExternalLoggingSetupTestsE2E(t *testing.T) {
	externalLoggingId := testutils.CreateTempExternalLogging(t)
	response, err := testutils.Client.NewExternalLoggingSetupTests().ExternalLoggingId(externalLoggingId).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertNotEmpty(t, response.Message)
}
