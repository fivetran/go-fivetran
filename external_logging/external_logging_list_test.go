package externallogging_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/external_logging"
	testutils "github.com/fivetran/go-fivetran/test_utils"
	
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestExternalLoggingListServiceDo(t *testing.T) {
	// arrange
	limit := 10
	cursor := "some_cursor"

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/external-logging").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareExternalLoggingListResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewExternalLoggingList().
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
	assertExternalLoggingListResponse(t, response)
}

func prepareExternalLoggingListResponse() string {
	return `{
  "code": "Success",
  "message": "Operation performed.",
  "data": {
    "items": [
      {
        "id": "log_id",
        "service": "string",
        "enabled": true
      }
    ],
    "next_cursor": "cursor_value"
  }
}`
}

func assertExternalLoggingListResponse(t *testing.T, response externallogging.ExternalLoggingListResponse) {
	testutils.AssertEqual(t, response.Code, "Success")

  testutils.AssertEqual(t, response.Code, "Success")
  testutils.AssertEqual(t, response.Data.Items[0].Id, "log_id")
  testutils.AssertEqual(t, response.Data.Items[0].Service, "string")
  testutils.AssertEqual(t, response.Data.Items[0].Enabled, true)

	testutils.AssertEqual(t, response.Data.NextCursor, "cursor_value")
}