package fivetran_test

import (
	"context"
	"testing"
)

func TestNewExternalLoggingSetupTestsE2E(t *testing.T) {
	externalLoggingId := CreateTempExternalLogging(t)
	response, err := Client.NewExternalLoggingSetupTests().ExternalLoggingId(externalLoggingId).
		TrustCertificates(true).
		TrustFingerprints(true).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	AssertEqual(t, response.Code, "Success")
	AssertNotEmpty(t, response.Message)
	AssertEqual(t, response.Data.ID, externalLoggingId)
	AssertEqual(t, response.Data.GroupID, externalLoggingId)
	AssertEqual(t, response.Data.Service, "snowflake")
	AssertEqual(t, response.Data.Region, "GCP_US_EAST4")
	AssertEqual(t, response.Data.TimeZoneOffset, "+10")
	AssertEqual(t, response.Data.SetupStatus, "incomplete")
}
