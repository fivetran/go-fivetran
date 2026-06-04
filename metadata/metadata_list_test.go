package metadata_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/metadata"
	testutils "github.com/fivetran/go-fivetran/test_utils"
	
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestMetadataListServiceDo(t *testing.T) {
	// arrange
	limit := 10
	cursor := "some_cursor"

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/metadata/connector-types").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareMetadataListResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewMetadataList().
		Limit(limit).
		Cursor(cursor).
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	assertMetadataListResponse(t, response)
}

func prepareMetadataListResponse() string {
	return `{
    "code": "Success",
    "data": {
        "items": [
            {
                "id": "15five",
                "name": "15Five",
                "type": "HumanResources",
                "description": "15Five is an employee performance management platform.",
                "icon_url": "https://fivetran.com/integrations/coil_connectors/resources/15five/resources/15five.svg",
                "icons": [
                    "https://fivetran.com/integrations/coil_connectors/resources/15five/resources/15five_512.png",
                    "https://fivetran.com/integrations/coil_connectors/resources/15five/resources/15five_40.svg"
                ],
                "link_to_docs": "https://fivetran.com/docs/connectors/applications/15five",
                "connector_class": "lite",
                "supported_features": [
                    {
                        "id": "API_CONFIGURABLE",
                        "notes": ""
                    },
                    {
                        "id": "CAPTURE_DELETES",
                        "notes": "All tables"
                    },
                    {
                        "id": "COLUMN_HASHING",
                        "notes": ""
                    },
                    {
                        "id": "DATA_BLOCKING",
                        "notes": "Column level"
                    },
                    {
                        "id": "RE_SYNC",
                        "notes": "Connection level and table level. A parent table re-sync includes all associated child tables."
                    },
                    {
                        "id": "AUTHORIZATION_VIA_API",
                        "notes": ""
                    }
                ],
                "config": {
                    "type": "object",
                    "description": "",
                    "readonly": false,
                    "properties": {
                        "api_key": {
                            "type": "string",
                            "description": "api_key",
                            "readonly": false,
                            "nullable": true,
                            "format": "password",
                            "immutable": true
                        },
                        "region": {
                            "type": "string",
                            "description": "region",
                            "readonly": false
                        }
                    }
                },
                "auth": {
                    "type": "object",
                    "description": "",
                    "readonly": false,
                    "properties": {
                        "credentials": {
                            "type": "object",
                            "description": "",
                            "readonly": false,
                            "properties": {
                                "token": {
                                    "type": "string",
                                    "description": "token",
                                    "readonly": false,
                                    "nullable": true,
                                    "format": "password",
                                    "immutable": true
                                },
                                "label": {
                                    "type": "string",
                                    "description": "label",
                                    "readonly": false
                                }
                            }
                        }
                    }
                },
                "link_to_erd": "https://fivetran.com/docs/connectors/applications/15five#schemainformation",
                "service_status": "general_availability",
                "service_status_updated_at": "2023-10-11"
            }
        ],
        "next_cursor": "cursor_value"
    }
}`
}

func assertMetadataListResponse(t *testing.T, response metadata.ConnectorMetadataListResponse) {
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

  testutils.AssertEqual(t, response.Data.Items[0].Config.Type, "object")
  testutils.AssertEqual(t, response.Data.Items[0].Config.Readonly, false)
  testutils.AssertEqual(t, response.Data.Items[0].Config.Nullable, false)
  testutils.AssertEqual(t, response.Data.Items[0].Config.Format, "")
  testutils.AssertEqual(t, response.Data.Items[0].Config.Immutable, false)
  testutils.AssertEqual(t, response.Data.Items[0].Config.Properties["api_key"].Type, "string")
  testutils.AssertEqual(t, response.Data.Items[0].Config.Properties["api_key"].Readonly, false)
  testutils.AssertEqual(t, response.Data.Items[0].Config.Properties["api_key"].Nullable, true)
  testutils.AssertEqual(t, response.Data.Items[0].Config.Properties["api_key"].Format, "password")
  testutils.AssertEqual(t, response.Data.Items[0].Config.Properties["api_key"].Immutable, true)
  testutils.AssertEqual(t, response.Data.Items[0].Config.Properties["region"].Type, "string")
  testutils.AssertEqual(t, response.Data.Items[0].Config.Properties["region"].Readonly, false)
  testutils.AssertEqual(t, response.Data.Items[0].Config.Properties["region"].Nullable, false)
  testutils.AssertEqual(t, response.Data.Items[0].Config.Properties["region"].Format, "")
  testutils.AssertEqual(t, response.Data.Items[0].Config.Properties["region"].Immutable, false)

  testutils.AssertEqual(t, response.Data.Items[0].Auth.Type, "object")
  testutils.AssertEqual(t, response.Data.Items[0].Auth.Readonly, false)
  testutils.AssertEqual(t, response.Data.Items[0].Auth.Properties["credentials"].Type, "object")
  testutils.AssertEqual(t, response.Data.Items[0].Auth.Properties["credentials"].Readonly, false)
  testutils.AssertEqual(t, response.Data.Items[0].Auth.Properties["credentials"].Properties["token"].Type, "string")
  testutils.AssertEqual(t, response.Data.Items[0].Auth.Properties["credentials"].Properties["token"].Readonly, false)
  testutils.AssertEqual(t, response.Data.Items[0].Auth.Properties["credentials"].Properties["token"].Nullable, true)
  testutils.AssertEqual(t, response.Data.Items[0].Auth.Properties["credentials"].Properties["token"].Format, "password")
  testutils.AssertEqual(t, response.Data.Items[0].Auth.Properties["credentials"].Properties["token"].Immutable, true)
  testutils.AssertEqual(t, response.Data.Items[0].Auth.Properties["credentials"].Properties["label"].Type, "string")
  testutils.AssertEqual(t, response.Data.Items[0].Auth.Properties["credentials"].Properties["label"].Readonly, false)
  testutils.AssertEqual(t, response.Data.Items[0].Auth.Properties["credentials"].Properties["label"].Nullable, false)
  testutils.AssertEqual(t, response.Data.Items[0].Auth.Properties["credentials"].Properties["label"].Format, "")
  testutils.AssertEqual(t, response.Data.Items[0].Auth.Properties["credentials"].Properties["label"].Immutable, false)

	testutils.AssertEqual(t, response.Data.NextCursor, "cursor_value")
}
