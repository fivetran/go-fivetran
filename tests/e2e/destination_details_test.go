package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewDestinationDetailsE2E(t *testing.T) {
	destinationId := testutils.CreateTempDestination(t)

	details, err := testutils.Client.NewDestinationDetails().DestinationID(destinationId).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", details)
		t.Error(err)
	}

	testutils.AssertEqual(t, details.Code, "Success")
	testutils.AssertEqual(t, details.Data.ID, testutils.PredefinedGroupId)
	testutils.AssertEqual(t, details.Data.GroupID, testutils.PredefinedGroupId)
	testutils.AssertEqual(t, details.Data.Service, "snowflake")
	testutils.AssertEqual(t, details.Data.Region, "GCP_US_EAST4")
	testutils.AssertEqual(t, details.Data.TimeZoneOffset, "+10")
	testutils.AssertEqual(t, details.Data.SetupStatus, "incomplete")
	testutils.AssertEqual(t, details.Data.Config.Database, "fivetran")
	testutils.AssertEqual(t, details.Data.Config.Password, "******")
	testutils.AssertEqual(t, details.Data.Config.Port, "443")
	testutils.AssertEqual(t, details.Data.Config.Host, "your-account.snowflakecomputing.com")
	testutils.AssertEqual(t, details.Data.Config.User, "fivetran_user")
}
