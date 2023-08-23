package fivetran_test

import (
	"context"
	"testing"
)

func TestNewExternalLoggingSetupTestsE2E(t *testing.T) {
	externalLoggingId := CreateTempExternalLogging(t)
	response, err := Client.NewExternalLoggingSetupTests().ExternalLoggingId(externalLoggingId).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	AssertEqual(t, response.Code, "Success")
	AssertNotEmpty(t, response.Message)
}
