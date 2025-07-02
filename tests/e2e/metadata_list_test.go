package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewMetadataListE2E(t *testing.T) {
	response, err := testutils.Client.NewMetadataList().Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

  	testutils.AssertEqual(t, response.Code, "Success")
  	testutils.AssertEqual(t, response.Data.Items[0].ID, "15five")
  	testutils.AssertEqual(t, response.Data.Items[0].Type,"HumanResources")
  	testutils.AssertEqual(t, response.Data.Items[0].Name, "15Five")
  	testutils.AssertEqual(t, response.Data.Items[0].Description, "15Five is an employee performance management platform.")
  	testutils.AssertNotEmpty(t, response.Data.Items[0].IconURL)
  	testutils.AssertNotEmpty(t, response.Data.Items[0].LinkToDocs)
  	testutils.AssertNotEmpty(t, response.Data.Items[0].LinkToErd)
  	testutils.AssertNotEmpty(t, response.Data.Items[0].Icons[0])
  	testutils.AssertNotEmpty(t, response.Data.Items[0].Icons[1])
  	testutils.AssertEqual(t, response.Data.Items[0].ConnectorClass, "lite")
  	testutils.AssertEqual(t, response.Data.Items[0].ServiceStatus, "general_availability")
  	testutils.AssertNotEmpty(t, response.Data.Items[0].ServiceStatusUpdatedAt)

  	testutils.AssertNotEmpty(t, response.Data.Items[0].SupportedFeatures[0].Id)
  	testutils.AssertNotEmpty(t, response.Data.Items[0].SupportedFeatures[0].Notes)
  	testutils.AssertNotEmpty(t, response.Data.Items[0].SupportedFeatures[1].Id)
  	testutils.AssertNotEmpty(t, response.Data.Items[0].SupportedFeatures[1].Notes)
  	testutils.AssertNotEmpty(t, response.Data.Items[0].SupportedFeatures[2].Id)
  	testutils.AssertNotEmpty(t, response.Data.Items[0].SupportedFeatures[2].Notes)
}
