package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewDestinationsListE2E(t *testing.T) {
	destinationId := testutils.CreateTempDestination(t)
	result, err := testutils.Client.NewDestinationsList().Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	testutils.AssertEqual(t, result.Code, "Success")
	testutils.AssertHasLength(t, result.Data.Items, 1)
	testutils.AssertEmpty(t, result.Message)
	testutils.AssertEqual(t, result.Code, "Success")
	testutils.AssertEqual(t, result.Data.Items[0].ID, testutils.PredefinedGroupId)
	testutils.AssertEqual(t, result.Data.Items[0].GroupID, testutils.PredefinedGroupId)
	testutils.AssertEqual(t, result.Data.Items[0].Service, "snowflake")
	testutils.AssertEqual(t, result.Data.Items[0].Region, "GCP_US_EAST4")
	testutils.AssertEqual(t, result.Data.Items[0].TimeZoneOffset, "+10")
	testutils.AssertEqual(t, result.Data.Items[0].DaylightSavingTimeEnabled, true)
	testutils.AssertEmpty(t, result.Data.Items[0].PrivateLinkId)
	testutils.AssertEmpty(t, result.Data.Items[0].HybridDeploymentAgentId)
	testutils.AssertEqual(t, result.Data.Items[0].NetworkingMethod, "Directly")
	testutils.AssertEqual(t, result.Data.Items[0].SetupStatus, "incomplete")

	testutils.AssertEmpty(t, result.Data.NextCursor)
}
