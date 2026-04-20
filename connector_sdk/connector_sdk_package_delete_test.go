package connectorsdk_test

import (
	"context"
	"net/http"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestConnectorSdkPackageDeleteServiceDo(t *testing.T) {
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodDelete, "/v1/connector-sdk/packages/happy_harmony").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, `{
				"code": "Success",
				"message": "Package has been deleted"
			}`)
			return response, nil
		})

	response, err := ftClient.NewConnectorSdkPackageDelete().
		PackageID("happy_harmony").
		Do(context.Background())

	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Message, "Package has been deleted")
}
