package metadata_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	fivetran "github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/metadata"
	testutils "github.com/fivetran/go-fivetran/test_utils"

	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestLogServiceMetadataDetailsServiceDo(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/metadata/log-service-types/datadog_log").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareLogServiceMetadataDetailsResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewLogServiceMetadataDetails().
		Service("datadog_log").
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	assertLogServiceMetadataDetailsResponse(t, response)
}

func TestLogServiceMetadataDetailsServiceDoWithUserAgentSuffix(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	ftClient.CustomUserAgent("terraform-provider-fivetran/1.9.37")
	handler := mockClient.When(http.MethodGet, "/v1/metadata/log-service-types/datadog_log").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			testutils.AssertEqual(t, req.Header.Get("User-Agent"), fmt.Sprintf("%s terraform-provider-fivetran/1.9.37 fivetran_account_external_logging_v2", fivetran.DefaultUserAgent))
			response := mock.NewResponse(req, http.StatusOK, prepareLogServiceMetadataDetailsResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewLogServiceMetadataDetails().
		Service("datadog_log").
		DoWithUserAgentSuffix(context.Background(), "fivetran_account_external_logging_v2")

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	assertLogServiceMetadataDetailsResponse(t, response)
}

func TestLogServiceMetadataDetailsServiceMissingService(t *testing.T) {
	// arrange
	ftClient, _ := testutils.CreateTestClient()

	// act
	_, err := ftClient.NewLogServiceMetadataDetails().Do(context.Background())

	// assert
	if err == nil {
		t.Error("expected an error for a missing service")
	} else {
		testutils.AssertEqual(t, err.Error(), "missing required service")
	}
}

func prepareLogServiceMetadataDetailsResponse() string {
	return `{
    "code": "Success",
    "data": {
        "id": "datadog_log",
        "name": "Datadog",
        "type": "Log",
        "description": "Datadog is a monitoring and analytics platform",
        "icon_url": "https://fivetran.com/integrations/datadog_log/resources/datadog.png",
        "icons": [
            "https://fivetran.com/integrations/datadog_log/resources/datadog_512.png",
            "https://fivetran.com/integrations/datadog_log/resources/datadog_40.svg"
        ],
        "link_to_docs": "https://fivetran.com/docs/logs/datadog",
        "service_status": "general_availability",
        "service_status_updated_at": "2022-01-13",
        "config": {
            "type": "object",
            "title": "Datadog config object",
            "readonly": false,
            "required": [
                "api_key"
            ],
            "properties": {
                "api_key": {
                    "type": "string",
                    "description": "The Datadog API key",
                    "format": "password",
                    "readonly": false
                },
                "region": {
                    "type": "string",
                    "description": "The Datadog site",
                    "readonly": false,
                    "enum": [
                        "US1",
                        "EU"
                    ]
                }
            }
        }
    }
}`
}

func assertLogServiceMetadataDetailsResponse(t *testing.T, response metadata.LogServiceMetadataResponse) {
	testutils.AssertEqual(t, response.Code, "Success")

	data := response.Data.LogServiceMetadata
	testutils.AssertEqual(t, data.ID, "datadog_log")
	testutils.AssertEqual(t, data.Name, "Datadog")
	testutils.AssertEqual(t, data.Type, "Log")
	testutils.AssertEqual(t, data.Description, "Datadog is a monitoring and analytics platform")
	testutils.AssertEqual(t, data.IconURL, "https://fivetran.com/integrations/datadog_log/resources/datadog.png")
	testutils.AssertEqual(t, len(data.Icons), 2)
	testutils.AssertEqual(t, data.LinkToDocs, "https://fivetran.com/docs/logs/datadog")
	testutils.AssertEqual(t, data.ServiceStatus, "general_availability")
	testutils.AssertEqual(t, data.ServiceStatusUpdatedAt, "2022-01-13")

	testutils.AssertEqual(t, data.Config.Type, "object")
	testutils.AssertEqual(t, data.Config.Title, "Datadog config object")
	testutils.AssertEqual(t, len(data.Config.Required), 1)
	testutils.AssertEqual(t, data.Config.Required[0], "api_key")

	apiKey := data.Config.Properties["api_key"]
	if apiKey == nil {
		t.Fatal("expected the api_key property to be present")
	}
	testutils.AssertEqual(t, apiKey.Type, "string")
	testutils.AssertEqual(t, apiKey.Format, "password")

	region := data.Config.Properties["region"]
	if region == nil {
		t.Fatal("expected the region property to be present")
	}
	testutils.AssertEqual(t, len(region.Enum), 2)
	testutils.AssertEqual(t, region.Enum[0], "US1")
}
