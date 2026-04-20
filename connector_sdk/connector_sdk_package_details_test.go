package connectorsdk_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	connectorsdk "github.com/fivetran/go-fivetran/connector_sdk"
	testutils "github.com/fivetran/go-fivetran/test_utils"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestConnectorSdkPackageDetailsServiceDo(t *testing.T) {
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/connector-sdk/packages/happy_harmony").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, preparePackageDetailsResponse())
			return response, nil
		})

	response, err := ftClient.NewConnectorSdkPackageDetails().
		PackageID("happy_harmony").
		Do(context.Background())

	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	assertPackageResponse(t, response)
}

func preparePackageDetailsResponse() string {
	return fmt.Sprintf(`{
		"code": "Success",
		"data": {
			"id": "happy_harmony",
			"connection_id": "conn_123",
			"created_by": "user_1",
			"last_updated_by": "user_2",
			"created_at": "2024-01-01T00:00:00.000000Z",
			"updated_at": "2024-01-02T00:00:00.000000Z",
			"file_sha256_hash": "abc123def456"
		}
	}`)
}

func assertPackageResponse(t *testing.T, response connectorsdk.ConnectorSdkPackageResponse) {
	t.Helper()
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.ID, "happy_harmony")
	testutils.AssertEqual(t, response.Data.ConnectionID, "conn_123")
	testutils.AssertEqual(t, response.Data.CreatedBy, "user_1")
	testutils.AssertEqual(t, response.Data.LastUpdatedBy, "user_2")
	testutils.AssertEqual(t, response.Data.FileSha256Hash, "abc123def456")
}
