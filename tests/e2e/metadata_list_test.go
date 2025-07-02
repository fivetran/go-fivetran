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
  	testutils.AssertEqual(t, response.Data.Items[0].IconURL, "https://fivetran.com/integrations/coil_connectors/resources/15five/resources/15five.svg")
  	testutils.AssertEqual(t, response.Data.Items[0].LinkToDocs, "https://fivetran.com/docs/connectors/applications/15five")
  	testutils.AssertEqual(t, response.Data.Items[0].LinkToErd, "https://fivetran.com/docs/connectors/applications/15five#schemainformation")
  	testutils.AssertEqual(t, response.Data.Items[0].Icons[0], "https://fivetran.com/integrations/coil_connectors/resources/15five/resources/15five_512.png")
  	testutils.AssertEqual(t, response.Data.Items[0].Icons[1],"https://fivetran.com/integrations/coil_connectors/resources/15five/resources/15five_40.svg")
  	testutils.AssertEqual(t, response.Data.Items[0].ConnectorClass, "lite")
  	testutils.AssertEqual(t, response.Data.Items[0].ServiceStatus, "general_availability")
  	testutils.AssertEqual(t, response.Data.Items[0].ServiceStatusUpdatedAt,"2023-10-11")

  	testutils.AssertEqual(t, response.Data.Items[0].SupportedFeatures[0].Id, "API_CONFIGURABLE")
  	testutils.AssertEqual(t, response.Data.Items[0].SupportedFeatures[0].Notes, "")
  	testutils.AssertEqual(t, response.Data.Items[0].SupportedFeatures[1].Id, "CAPTURE_DELETES")
  	testutils.AssertEqual(t, response.Data.Items[0].SupportedFeatures[1].Notes, "All tables")
  	testutils.AssertEqual(t, response.Data.Items[0].SupportedFeatures[2].Id, "COLUMN_HASHING")
  	testutils.AssertEqual(t, response.Data.Items[0].SupportedFeatures[2].Notes, "")
  	testutils.AssertEqual(t, response.Data.Items[0].SupportedFeatures[3].Id, "DATA_BLOCKING")
  	testutils.AssertEqual(t, response.Data.Items[0].SupportedFeatures[3].Notes, "Column level")
  	testutils.AssertEqual(t, response.Data.Items[0].SupportedFeatures[4].Id, "RE_SYNC")
  	testutils.AssertEqual(t, response.Data.Items[0].SupportedFeatures[4].Notes, "Connection level and table level. A parent table re-sync includes all associated child tables.")
  	testutils.AssertEqual(t, response.Data.Items[0].SupportedFeatures[5].Id, "AUTHORIZATION_VIA_API")
  	testutils.AssertEqual(t, response.Data.Items[0].SupportedFeatures[5].Notes, "")
}
