package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewDestinationSetupTestsE2E(t *testing.T) {
	destinationId := testutils.CreateTempDestination(t)
	response, err := testutils.Client.NewDestinationSetupTests().DestinationID(destinationId).
		TrustCertificates(true).
		TrustFingerprints(true).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertNotEmpty(t, response.Message)
	testutils.AssertEqual(t, response.Data.ID, destinationId)
	testutils.AssertEqual(t, response.Data.GroupID, destinationId)
	testutils.AssertEqual(t, response.Data.Service, "snowflake")
	testutils.AssertEqual(t, response.Data.Region, "GCP_US_EAST4")
	testutils.AssertEqual(t, response.Data.TimeZoneOffset, "+10")
	testutils.AssertEqual(t, response.Data.SetupStatus, "incomplete")
}
