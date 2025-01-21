package transformations_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/transformations"
	testutils "github.com/fivetran/go-fivetran/test_utils"
	
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestQuickstartPackagesListServiceDo(t *testing.T) {
	// arrange
	limit := 10
	cursor := "some_cursor"

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/transformations/package-metadata").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareQuickstartPackagesListResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewQuickstartPackagesList().
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
	assertQuickstartPackagesListResponse(t, response)
}

func prepareQuickstartPackagesListResponse() string {
	return `{
  "code": "Success",
  "message": "Operation performed.",
  "data": {
    "items": [
      {
        "id": "package_definition_id",
        "name": "package_definition_name",
        "version": "version",
        "connector_types": [
          "string"
        ],
        "output_model_names": [
          "string"
        ]
      }
    ],
    "next_cursor": "cursor_value"
  }
}`
}

func assertQuickstartPackagesListResponse(t *testing.T, response transformations.QuickstartPackagesListResponse) {
    testutils.AssertEqual(t, response.Code, "Success")
    testutils.AssertNotEmpty(t, response.Message)
    testutils.AssertEqual(t, response.Data.Items[0].Id, "package_definition_id")
    testutils.AssertEqual(t, response.Data.Items[0].Name, "package_definition_name")
    testutils.AssertEqual(t, response.Data.Items[0].Version, "version")
    testutils.AssertEqual(t, response.Data.Items[0].ConnectorTypes[0], "string")
    testutils.AssertEqual(t, response.Data.Items[0].OutputModelNames[0], "string")

		testutils.AssertEqual(t, response.Data.NextCursor, "cursor_value")
}