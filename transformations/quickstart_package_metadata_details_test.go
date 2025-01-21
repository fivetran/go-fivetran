package transformations_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/transformations"
	testutils "github.com/fivetran/go-fivetran/test_utils"
	
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestQuickstartPackageDetailsServiceDo(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/transformations/package-metadata/package_definition_id").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareQuickstartPackageDetailsResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewQuickstartPackageDetails().
		PackageDefinitionId("package_definition_id").
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	assertQuickstartPackageDetailsResponse(t, response)
}

func prepareQuickstartPackageDetailsResponse() string {
	return `{
  "code": "Success",
  "message": "Operation performed.",
  "data": {
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
}`
}

func assertQuickstartPackageDetailsResponse(t *testing.T, response transformations.QuickstartPackageResponse) {
    testutils.AssertEqual(t, response.Code, "Success")
    testutils.AssertNotEmpty(t, response.Message)
    testutils.AssertEqual(t, response.Data.Id, "package_definition_id")
    testutils.AssertEqual(t, response.Data.Name, "package_definition_name")
    testutils.AssertEqual(t, response.Data.Version, "version")
    testutils.AssertEqual(t, response.Data.ConnectorTypes[0], "string")
    testutils.AssertEqual(t, response.Data.OutputModelNames[0], "string")
}