package metadata_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/metadata"
	testutils "github.com/fivetran/go-fivetran/test_utils"

	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestLogServiceMetadataListServiceDo(t *testing.T) {
	// arrange
	limit := 10
	cursor := "some_cursor"

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/metadata/log-service-types").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			testutils.AssertEqual(t, req.URL.Query().Get("limit"), "10")
			testutils.AssertEqual(t, req.URL.Query().Get("cursor"), cursor)
			response := mock.NewResponse(req, http.StatusOK, prepareLogServiceMetadataListResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewLogServiceMetadataList().
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
	assertLogServiceMetadataListResponse(t, response)
}

func prepareLogServiceMetadataListResponse() string {
	return `{
    "code": "Success",
    "data": {
        "items": [
            {
                "id": "cloudwatch",
                "name": "Amazon CloudWatch",
                "type": "Log",
                "icon_url": "https://fivetran.com/integrations/cloudwatch/resources/cloudwatch.png",
                "icons": [
                    "https://fivetran.com/integrations/cloudwatch/resources/cloudwatch_512.png"
                ],
                "link_to_docs": "https://fivetran.com/docs/logs/cloudwatch",
                "service_status": "general_availability",
                "service_status_updated_at": "2022-01-13"
            },
            {
                "id": "datadog_log",
                "name": "Datadog",
                "type": "Log",
                "description": "Datadog is a monitoring and analytics platform",
                "service_status": "beta",
                "service_status_updated_at": "2023-05-02"
            }
        ],
        "next_cursor": "eyJza2lwIjoyfQ"
    }
}`
}

func assertLogServiceMetadataListResponse(t *testing.T, response metadata.LogServiceMetadataListResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.NextCursor, "eyJza2lwIjoyfQ")
	testutils.AssertEqual(t, len(response.Data.Items), 2)

	cloudwatch := response.Data.Items[0]
	testutils.AssertEqual(t, cloudwatch.ID, "cloudwatch")
	testutils.AssertEqual(t, cloudwatch.Name, "Amazon CloudWatch")
	testutils.AssertEqual(t, cloudwatch.Type, "Log")
	testutils.AssertEqual(t, len(cloudwatch.Icons), 1)
	testutils.AssertEqual(t, cloudwatch.LinkToDocs, "https://fivetran.com/docs/logs/cloudwatch")
	// The list endpoint omits config.
	testutils.AssertEqual(t, cloudwatch.Config.Type, "")

	datadog := response.Data.Items[1]
	testutils.AssertEqual(t, datadog.ID, "datadog_log")
	testutils.AssertEqual(t, datadog.Description, "Datadog is a monitoring and analytics platform")
	testutils.AssertEqual(t, datadog.ServiceStatus, "beta")
	testutils.AssertEqual(t, datadog.ServiceStatusUpdatedAt, "2023-05-02")
}
