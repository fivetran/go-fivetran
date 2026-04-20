package connectorsdk_test

import (
	"context"
	"net/http"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestConnectorSdkPackageListServiceDo(t *testing.T) {
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/connector-sdk/packages").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			testutils.AssertEqual(t, req.URL.Query().Get("cursor"), "next_page")
			testutils.AssertEqual(t, req.URL.Query().Get("limit"), "10")
			response := mock.NewResponse(req, http.StatusOK, preparePackageListResponse())
			return response, nil
		})

	response, err := ftClient.NewConnectorSdkPackageList().
		Cursor("next_page").
		Limit(10).
		Do(context.Background())

	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, len(response.Data.Items), 1)
	testutils.AssertEqual(t, response.Data.Items[0].ID, "happy_harmony")
	testutils.AssertEqual(t, response.Data.Items[0].FileSha256Hash, "abc123def456")
	testutils.AssertEqual(t, response.Data.NextCursor, "cursor_value")
}

func preparePackageListResponse() string {
	return `{
		"code": "Success",
		"data": {
			"items": [{
				"id": "happy_harmony",
				"connection_id": "conn_123",
				"created_by": "user_1",
				"last_updated_by": "user_2",
				"created_at": "2024-01-01T00:00:00.000000Z",
				"updated_at": "2024-01-02T00:00:00.000000Z",
				"file_sha256_hash": "abc123def456"
			}],
			"next_cursor": "cursor_value"
		}
	}`
}
