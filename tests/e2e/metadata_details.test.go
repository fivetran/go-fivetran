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
  testutils.AssertEqual(t, response.Data.Description, "Google Ads is an online advertising platform")
  testutils.AssertEqual(t, response.Data.IconURL, "https://fivetran.com/integrations/google_ads/resources/google-ads.png")
  testutils.AssertEqual(t, response.Data.LinkToDocs, "https://fivetran.com/docs/connectors/applications/google-ads")
  testutils.AssertEqual(t, response.Data.LinkToErd, "https://fivetran.com/docs/connectors/applications/google-ads#schemainformation")
  testutils.AssertEqual(t, response.Data.Icons[0], "https://fivetran.com/integrations/google_ads/resources/google-ads_512.png")
  testutils.AssertEqual(t, response.Data.Icons[1],"https://fivetran.com/integrations/google_ads/resources/google-ads_40.svg")
  testutils.AssertEqual(t, response.Data.ConnectorClass, "standard")
  testutils.AssertEqual(t, response.Data.ServiceStatus, "general_availability")
  testutils.AssertEqual(t, response.Data.ServiceStatusUpdatedAt,"2022-06-10")
  testutils.AssertEqual(t, response.Data.SupportedFeatures[0].Id, "API_CONFIGURABLE")
  testutils.AssertEqual(t, response.Data.SupportedFeatures[0].Notes, "")
  testutils.AssertEqual(t, response.Data.SupportedFeatures[1].Id, "COLUMN_HASHING")
  testutils.AssertEqual(t, response.Data.SupportedFeatures[1].Notes, "")
  testutils.AssertEqual(t, response.Data.SupportedFeatures[2].Id, "DATA_BLOCKING")
  testutils.AssertEqual(t, response.Data.SupportedFeatures[2].Notes, "Column level and table level")
  testutils.AssertEqual(t, response.Data.SupportedFeatures[3].Id, "FIVETRAN_DATA_MODELS")
  testutils.AssertEqual(t, response.Data.SupportedFeatures[3].Notes, "FIVETRAN_DATA_MODELS_Notes")
  testutils.AssertEqual(t, response.Data.SupportedFeatures[4].Id, "RE_SYNC")
  testutils.AssertEqual(t, response.Data.SupportedFeatures[4].Notes, "Connection and table level")
  testutils.AssertEqual(t, response.Data.SupportedFeatures[5].Id, "AUTHORIZATION_VIA_API")
  testutils.AssertEqual(t, response.Data.SupportedFeatures[5].Notes, "")

  testutils.AssertEqual(t, response.Data.Config.Type, "object")
  testutils.AssertEqual(t, response.Data.Config.Description, "")
  testutils.AssertEqual(t, response.Data.Config.Title, "Google Ads config object")
  testutils.AssertEqual(t, response.Data.Config.Readonly, false)
  testutils.AssertEqual(t, response.Data.Config.Required[0], "schema")

  testutils.AssertEqual(t, response.Data.Config.Properties["sync_mode"].Type, "string")
  testutils.AssertEqual(t, response.Data.Config.Properties["sync_mode"].Description, "Whether to sync all accounts or specific accounts.")
  testutils.AssertEqual(t, response.Data.Config.Properties["sync_mode"].Readonly, false)
  testutils.AssertEqual(t, response.Data.Config.Properties["sync_mode"].Enum[0], "ManagerAccounts")
  testutils.AssertEqual(t, response.Data.Config.Properties["sync_mode"].Enum[1], "AllAccounts")
  testutils.AssertEqual(t, response.Data.Config.Properties["sync_mode"].Enum[2], "SpecificAccounts")

  testutils.AssertEqual(t, response.Data.Config.Properties["manager_accounts"].Type, "array")
  testutils.AssertEqual(t, response.Data.Config.Properties["manager_accounts"].Description, "manager_accounts")
  testutils.AssertEqual(t, response.Data.Config.Properties["manager_accounts"].Readonly, false)
  testutils.AssertEqual(t, response.Data.Config.Properties["manager_accounts"].Items.Type, "string")
  testutils.AssertEqual(t, response.Data.Config.Properties["manager_accounts"].Items.Description, "")
  testutils.AssertEqual(t, response.Data.Config.Properties["manager_accounts"].Items.Readonly, false)


  testutils.AssertEqual(t, response.Data.Config.Properties["reports"].Type, "array")
  testutils.AssertEqual(t, response.Data.Config.Properties["reports"].Description, "The list of reports. Each report corresponds to a table within the schema to which connector will sync the data.")
  testutils.AssertEqual(t, response.Data.Config.Properties["reports"].Readonly, false)
  testutils.AssertEqual(t, response.Data.Config.Properties["reports"].Items.Type, "object")
  testutils.AssertEqual(t, response.Data.Config.Properties["reports"].Items.Description, "The list of reports. Each report corresponds to a table within the schema to which connector will sync the data.")
  testutils.AssertEqual(t, response.Data.Config.Properties["reports"].Items.Readonly, false)
  testutils.AssertEqual(t, response.Data.Config.Properties["reports"].Items.Properties["report_type"].Type, "string")
  testutils.AssertEqual(t, response.Data.Config.Properties["reports"].Items.Properties["report_type"].Description, "The name of the Google Ads report from which the connector will sync the data. [Possible report_type values](https://developers.google.com/adwords/api/docs/appendix/reports#report-types).")
  testutils.AssertEqual(t, response.Data.Config.Properties["reports"].Items.Properties["report_type"].Readonly, false)
  testutils.AssertEqual(t, response.Data.Config.Properties["reports"].Items.Properties["fields"].Type, "array")
  testutils.AssertEqual(t, response.Data.Config.Properties["reports"].Items.Properties["fields"].Description, "fields")
  testutils.AssertEqual(t, response.Data.Config.Properties["reports"].Items.Properties["fields"].Readonly, false)
  testutils.AssertEqual(t, response.Data.Config.Properties["reports"].Items.Properties["fields"].Items.Type, "string")
  testutils.AssertEqual(t, response.Data.Config.Properties["reports"].Items.Properties["fields"].Items.Description, "")
  testutils.AssertEqual(t, response.Data.Config.Properties["reports"].Items.Properties["fields"].Items.Readonly, false)

  testutils.AssertEqual(t, response.Data.Auth.Type, "object")
  testutils.AssertEqual(t, response.Data.Auth.Description, "")
  testutils.AssertEqual(t, response.Data.Auth.Readonly, false)
  testutils.AssertEqual(t, response.Data.Auth.Properties["refresh_token"].Type, "string")
  testutils.AssertEqual(t, response.Data.Auth.Properties["refresh_token"].Description, "refresh_token")
  testutils.AssertEqual(t, response.Data.Auth.Properties["refresh_token"].Readonly, false)

  testutils.AssertEqual(t, response.Data.Auth.Properties["client_access"].Type, "object")
  testutils.AssertEqual(t, response.Data.Auth.Properties["client_access"].Description, "")
  testutils.AssertEqual(t, response.Data.Auth.Properties["client_access"].Readonly, false)
  testutils.AssertEqual(t, response.Data.Auth.Properties["client_access"].Properties["developer_token"].Type, "string")
  testutils.AssertEqual(t, response.Data.Auth.Properties["client_access"].Properties["developer_token"].Description, "developer_token")
  testutils.AssertEqual(t, response.Data.Auth.Properties["client_access"].Properties["developer_token"].Readonly, false)
}
