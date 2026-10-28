package metadata_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/metadata"
	testutils "github.com/fivetran/go-fivetran/test_utils"

	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestDestinationMetadataListServiceDo(t *testing.T) {
	// arrange
	limit := 10
	cursor := "some_cursor"

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/metadata/destination-types").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			testutils.AssertEqual(t, req.URL.Query().Get("limit"), "10")
			testutils.AssertEqual(t, req.URL.Query().Get("cursor"), cursor)
			response := mock.NewResponse(req, http.StatusOK, prepareDestinationMetadataListResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewDestinationMetadataList().
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
	assertDestinationMetadataListResponse(t, response)
}

func prepareDestinationMetadataListResponse() string {
	return `{
    "code": "Success",
    "data": {
        "items": [
            {
                "id": "snowflake",
                "name": "Snowflake",
                "type": "Warehouse",
                "description": "Snowflake is a cloud data warehouse",
                "icon_url": "https://fivetran.com/integrations/snowflake/resources/snowflake.png",
                "icons": [
                    "https://fivetran.com/integrations/snowflake/resources/snowflake_512.png"
                ],
                "link_to_docs": "https://fivetran.com/docs/destinations/snowflake",
                "service_status": "general_availability",
                "service_status_updated_at": "2022-01-13",
                "supported_regions": [
                    "GCP_US_EAST4",
                    "AWS_US_EAST_1"
                ],
                "supports_hybrid_deployment": false,
                "supports_private_link": true,
                "private_link_types": [
                    "SELF_SERVICE",
                    "CUSTOMER_MANAGED"
                ],
                "supported_networking_methods": [
                    "Directly",
                    "PrivateLink"
                ]
            },
            {
                "id": "sql_server_warehouse",
                "name": "SQL Server",
                "type": "Warehouse",
                "service_status": "general_availability",
                "service_status_updated_at": "2023-05-02",
                "supported_regions": [
                    "AWS_US_EAST_1"
                ],
                "supports_hybrid_deployment": true,
                "supports_private_link": false,
                "private_link_types": [],
                "supported_networking_methods": [
                    "Directly",
                    "SshTunnel"
                ]
            }
        ],
        "next_cursor": "eyJza2lwIjoyfQ"
    }
}`
}

func assertDestinationMetadataListResponse(t *testing.T, response metadata.DestinationMetadataListResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.NextCursor, "eyJza2lwIjoyfQ")
	testutils.AssertEqual(t, len(response.Data.Items), 2)

	snowflake := response.Data.Items[0]
	testutils.AssertEqual(t, snowflake.ID, "snowflake")
	testutils.AssertEqual(t, snowflake.Name, "Snowflake")
	testutils.AssertEqual(t, snowflake.Type, "Warehouse")
	testutils.AssertEqual(t, snowflake.SupportsPrivateLink, true)
	testutils.AssertEqual(t, len(snowflake.PrivateLinkTypes), 2)
	testutils.AssertEqual(t, len(snowflake.SupportedRegions), 2)
	// The list endpoint omits config.
	testutils.AssertEqual(t, snowflake.Config.Type, "")
	if snowflake.ExternalSecretsKeysConfig != nil {
		t.Error("expected the list endpoint to omit external_secrets_keys_config")
	}

	sqlServer := response.Data.Items[1]
	testutils.AssertEqual(t, sqlServer.ID, "sql_server_warehouse")
	testutils.AssertEqual(t, sqlServer.SupportsHybridDeployment, true)
	testutils.AssertEqual(t, sqlServer.SupportsPrivateLink, false)
	testutils.AssertEqual(t, len(sqlServer.PrivateLinkTypes), 0)
	testutils.AssertEqual(t, len(sqlServer.SupportedNetworkingMethods), 2)
}
