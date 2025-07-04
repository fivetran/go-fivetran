package metadata_test

import (
  "context"
  "net/http"
  "testing"
  //"fmt"

  "github.com/fivetran/go-fivetran/metadata"
  testutils "github.com/fivetran/go-fivetran/test_utils"
  
  "github.com/fivetran/go-fivetran/tests/mock"
)

func TestMetadataDetailsServiceDo(t *testing.T) {
  // arrange

  ftClient, mockClient := testutils.CreateTestClient()
  handler := mockClient.When(http.MethodGet, "/v1/metadata/connector-types/google_ads").
    ThenCall(func(req *http.Request) (*http.Response, error) {
      response := mock.NewResponse(req, http.StatusOK, prepareMetadataDetailsResponse())
      return response, nil
    })

  // act
  response, err := ftClient.NewMetadataDetails().
    Service("google_ads").
    Do(context.Background())

  // assert
  if err != nil {
    t.Error(err)
  }

  interactions := mockClient.Interactions()
  testutils.AssertEqual(t, len(interactions), 1)
  testutils.AssertEqual(t, interactions[0].Handler, handler)
  testutils.AssertEqual(t, handler.Interactions, 1)
  assertMetadataDetailsResponse(t, response)
}

func prepareMetadataDetailsResponse() string {
  return `{
    "code": "Success",
    "data": {
        "id": "google_ads",
        "name": "Google Ads",
        "type": "Marketing",
        "description": "Google Ads is an online advertising platform",
        "icon_url": "https://fivetran.com/integrations/google_ads/resources/google-ads.png",
        "icons": [
            "https://fivetran.com/integrations/google_ads/resources/google-ads_512.png",
            "https://fivetran.com/integrations/google_ads/resources/google-ads_40.svg"
        ],
        "link_to_docs": "https://fivetran.com/docs/connectors/applications/google-ads",
        "connector_class": "standard",
        "supported_features": [
            {
                "id": "API_CONFIGURABLE",
                "notes": ""
            },
            {
                "id": "COLUMN_HASHING",
                "notes": ""
            },
            {
                "id": "DATA_BLOCKING",
                "notes": "Column level and table level"
            },
            {
                "id": "FIVETRAN_DATA_MODELS",
                "notes": "FIVETRAN_DATA_MODELS_Notes"
            },
            {
                "id": "RE_SYNC",
                "notes": "Connection and table level"
            },
            {
                "id": "AUTHORIZATION_VIA_API",
                "notes": ""
            }
        ],
        "config": {
            "type": "object",
            "description": "",
            "title": "Google Ads config object",
            "example": "config_example",
            "readonly": false,
            "properties": {
                "sync_mode": {
                    "type": "string",
                    "description": "Whether to sync all accounts or specific accounts.",
                    "example": "ManagerAccounts | AllAccounts | SpecificAccounts",
                    "readonly": false,
                    "enum": [
                        "ManagerAccounts",
                        "AllAccounts",
                        "SpecificAccounts"
                    ]
                },
                "schema": {
                    "type": "string",
                    "description": "Destination schema. Schema is permanent and cannot be changed after connection creation",
                    "title": "Destination schema",
                    "example": "schema_name",
                    "readonly": false
                },
                "reports": {
                    "type": "array",
                    "description": "The list of reports. Each report corresponds to a table within the schema to which connector will sync the data.",
                    "example": [
                        "{\n\"report_type\": campaign,\n\"fields\": [string],\n\"table\": table_2\n}"
                    ],
                    "readonly": false,
                    "items": {
                        "type": "object",
                        "description": "The list of reports. Each report corresponds to a table within the schema to which connector will sync the data.",
                        "example": "{\n\"report_type\": campaign,\n\"fields\": [string],\n\"table\": table_2\n}",
                        "readonly": false,
                        "properties": {
                            "report_type": {
                                "type": "string",
                                "description": "The name of the Google Ads report from which the connector will sync the data. [Possible report_type values](https://developers.google.com/adwords/api/docs/appendix/reports#report-types).",
                                "example": "campaign",
                                "readonly": false
                            },
                            "fields": {
                                "type": "array",
                                "description": "fields",
                                "example": [
                                    "string"
                                ],
                                "readonly": false,
                                "items": {
                                    "type": "string",
                                    "description": "",
                                    "example": "string",
                                    "readonly": false
                                }
                            },
                            "table": {
                                "type": "string",
                                "description": "The table name within the schema to which connector will sync the data of the specific report.",
                                "example": "table_2",
                                "readonly": false
                            }
                        }
                    }
                },
                "skip_empty_reports": {
                    "type": "boolean",
                    "description": "Toggles the [\"Skip empty reports\"](/docs/connectors/applications/google-ads#skipemptyreports) feature. Enabled by default",
                    "example": true,
                    "readonly": false
                },
                "manager_accounts": {
                    "type": "array",
                    "description": "manager_accounts",
                    "example": [
                        "string"
                    ],
                    "readonly": false,
                    "items": {
                        "type": "string",
                        "description": "",
                        "example": "string",
                        "readonly": false
                    }
                },
                "timeframe_months": {
                    "type": "string",
                    "description": "timeframe_months",
                    "example": "TWENTY_FOUR | SIX | ALL_TIME | TWELVE | THREE",
                    "readonly": false,
                    "enum": [
                        "TWENTY_FOUR",
                        "SIX",
                        "ALL_TIME",
                        "TWELVE",
                        "THREE"
                    ]
                },
                "accounts": {
                    "type": "array",
                    "description": "accounts",
                    "example": [
                        "string"
                    ],
                    "readonly": false,
                    "items": {
                        "type": "string",
                        "description": "",
                        "example": "string",
                        "readonly": false
                    }
                },
                "customer_id": {
                    "type": "string",
                    "description": "ID of the customer, can be retrieved from your AdWords dashboard.",
                    "example": "xxx-xxx-xxxx",
                    "readonly": false
                },
                "conversion_window_size": {
                    "type": "integer",
                    "description": "A period of time in days during which a conversion is recorded.",
                    "example": "30",
                    "readonly": false
                }
            },
            "required": [
                "schema"
            ]
        },
        "auth": {
            "type": "object",
            "description": "",
            "readonly": false,
            "properties": {
                "refresh_token": {
                    "type": "string",
                    "description": "refresh_token",
                    "example": "my_refresh_token",
                    "readonly": false
                },
                "client_access": {
                    "type": "object",
                    "description": "",
                    "readonly": false,
                    "properties": {
                        "developer_token": {
                            "type": "string",
                            "description": "developer_token",
                            "example": "string",
                            "readonly": false
                        }
                    }
                }
            }
        },
        "link_to_erd": "https://fivetran.com/docs/connectors/applications/google-ads#schemainformation",
        "service_status": "general_availability",
        "service_status_updated_at": "2022-06-10"
    }
}`
}

func assertMetadataDetailsResponse(t *testing.T, response metadata.ConnectorMetadataResponse) {
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