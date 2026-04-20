package connectorsdk_test

import (
	"context"
	"net/http"
	"strings"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestConnectorSdkPackageCreateServiceDo(t *testing.T) {
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/connector-sdk/packages").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			contentType := req.Header.Get("Content-Type")
			if !strings.Contains(contentType, "multipart/form-data") {
				t.Errorf("expected multipart/form-data content type, got %s", contentType)
			}
			response := mock.NewResponse(req, http.StatusCreated, preparePackageDetailsResponse())
			return response, nil
		})

	response, err := ftClient.NewConnectorSdkPackageCreate().
		FileContent(strings.NewReader("fake-zip-content")).
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
