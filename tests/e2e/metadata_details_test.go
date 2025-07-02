package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewMetadataDetailsE2E(t *testing.T) {
  response, err := testutils.Client.NewMetadataDetails().Service("google_ads").Do(context.Background())
  if err != nil {
    t.Logf("%+v\n", response)
	t.Error(err)
  }

  testutils.AssertEqual(t, response.Code, "Success")
  testutils.AssertEqual(t, response.Data.ID, "google_ads")
  testutils.AssertEqual(t, response.Data.Type,"Marketing")
  testutils.AssertEqual(t, response.Data.Name, "Google Ads")
  testutils.AssertNotEmpty(t, response.Data.Description)
  testutils.AssertNotEmpty(t, response.Data.IconURL)
  testutils.AssertNotEmpty(t, response.Data.LinkToDocs)
  testutils.AssertNotEmpty(t, response.Data.LinkToErd)
  testutils.AssertNotEmpty(t, response.Data.Icons[0])
  testutils.AssertNotEmpty(t, response.Data.Icons[1])
  testutils.AssertEqual(t, response.Data.ConnectorClass, "standard")
  testutils.AssertEqual(t, response.Data.ServiceStatus, "general_availability")
  testutils.AssertNotEmpty(t, response.Data.ServiceStatusUpdatedAt)
  testutils.AssertNotEmpty(t, response.Data.SupportedFeatures[0].Id)
  testutils.AssertNotEmpty(t, response.Data.SupportedFeatures[0].Notes)
  testutils.AssertNotEmpty(t, response.Data.SupportedFeatures[1].Id)
  testutils.AssertNotEmpty(t, response.Data.SupportedFeatures[1].Notes)
  testutils.AssertNotEmpty(t, response.Data.SupportedFeatures[2].Id)
  testutils.AssertNotEmpty(t, response.Data.SupportedFeatures[2].Notes)

  testutils.AssertEqual(t, response.Data.Config.Type, "object")
  testutils.AssertEqual(t, response.Data.Config.Description, "")
  testutils.AssertEqual(t, response.Data.Config.Title, "Google Ads config object")
  testutils.AssertEqual(t, response.Data.Config.Readonly, false)
  testutils.AssertEqual(t, response.Data.Config.Required[0], "schema")

  testutils.AssertEqual(t, response.Data.Config.Properties["sync_mode"].Type, "string")
  testutils.AssertNotEmpty(t, response.Data.Config.Properties["sync_mode"].Description)
  testutils.AssertEqual(t, response.Data.Config.Properties["sync_mode"].Readonly, false)
  testutils.AssertEqual(t, response.Data.Config.Properties["sync_mode"].Enum[0], "ManagerAccounts")
  testutils.AssertEqual(t, response.Data.Config.Properties["sync_mode"].Enum[1], "AllAccounts")
  testutils.AssertEqual(t, response.Data.Config.Properties["sync_mode"].Enum[2], "SpecificAccounts")

  testutils.AssertEqual(t, response.Data.Config.Properties["manager_accounts"].Type, "array")
  testutils.AssertNotEmpty(t, response.Data.Config.Properties["manager_accounts"].Description)
  testutils.AssertEqual(t, response.Data.Config.Properties["manager_accounts"].Readonly, false)
  testutils.AssertEqual(t, response.Data.Config.Properties["manager_accounts"].Items.Type, "string")
  testutils.AssertEqual(t, response.Data.Config.Properties["manager_accounts"].Items.Description, "")
  testutils.AssertEqual(t, response.Data.Config.Properties["manager_accounts"].Items.Readonly, false)


  testutils.AssertEqual(t, response.Data.Config.Properties["reports"].Type, "array")
  testutils.AssertNotEmpty(t, response.Data.Config.Properties["reports"].Description)
  testutils.AssertEqual(t, response.Data.Config.Properties["reports"].Readonly, false)
  testutils.AssertEqual(t, response.Data.Config.Properties["reports"].Items.Type, "object")
  testutils.AssertNotEmpty(t, response.Data.Config.Properties["reports"].Items.Description)
  testutils.AssertEqual(t, response.Data.Config.Properties["reports"].Items.Readonly, false)
  testutils.AssertEqual(t, response.Data.Config.Properties["reports"].Items.Properties["report_type"].Type, "string")
  testutils.AssertNotEmpty(t, response.Data.Config.Properties["reports"].Items.Properties["report_type"].Description)
  testutils.AssertEqual(t, response.Data.Config.Properties["reports"].Items.Properties["report_type"].Readonly, false)
  testutils.AssertEqual(t, response.Data.Config.Properties["reports"].Items.Properties["fields"].Type, "array")
  testutils.AssertNotEmpty(t, response.Data.Config.Properties["reports"].Items.Properties["fields"].Description)
  testutils.AssertEqual(t, response.Data.Config.Properties["reports"].Items.Properties["fields"].Readonly, false)
  testutils.AssertEqual(t, response.Data.Config.Properties["reports"].Items.Properties["fields"].Items.Type, "string")
  testutils.AssertNotEmpty(t, response.Data.Config.Properties["reports"].Items.Properties["fields"].Items.Description)
  testutils.AssertEqual(t, response.Data.Config.Properties["reports"].Items.Properties["fields"].Items.Readonly, false)

  testutils.AssertEqual(t, response.Data.Auth.Type, "object")
  testutils.AssertNotEmpty(t, response.Data.Auth.Description)
  testutils.AssertEqual(t, response.Data.Auth.Readonly, false)
  testutils.AssertEqual(t, response.Data.Auth.Properties["refresh_token"].Type, "string")
  testutils.AssertNotEmpty(t, response.Data.Auth.Properties["refresh_token"].Description)
  testutils.AssertEqual(t, response.Data.Auth.Properties["refresh_token"].Readonly, false)

  testutils.AssertEqual(t, response.Data.Auth.Properties["client_access"].Type, "object")
  testutils.AssertNotEmpty(t, response.Data.Auth.Properties["client_access"].Description)
  testutils.AssertEqual(t, response.Data.Auth.Properties["client_access"].Readonly, false)
  testutils.AssertEqual(t, response.Data.Auth.Properties["client_access"].Properties["developer_token"].Type, "string")
  testutils.AssertNotEmpty(t, response.Data.Auth.Properties["client_access"].Properties["developer_token"].Description)
  testutils.AssertEqual(t, response.Data.Auth.Properties["client_access"].Properties["developer_token"].Readonly, false)
}
