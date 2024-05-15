package privatelink_test

import (
	"context"
	"net/http"
	"testing"
	"github.com/fivetran/go-fivetran/private_link"
	"github.com/fivetran/go-fivetran/tests/mock"
	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestPrivateLinkDetailsServiceDo(t *testing.T) {
	// arrange

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/private-links/123456").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, preparePrivateLinkDetailsResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewPrivateLinkDetails().
		PrivateLinkId("123456").
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	assertPrivateLinkDetailsResponse(t, response)
}

func preparePrivateLinkDetailsResponse() string {
	return `{
  "code": "Success",
  "data": {
    "id": "123456",
    "name": "name",
    "region": "region",
    "service": "service",
    "account_id": "account_id",
    "cloud_provider": "cloud_provider",
    "state": "state",
    "state_summary": "state_summary",
    "created_at": "2022-04-29T09:41:08.583Z",
    "created_by": "created_by",
    "config": {
      "connection_service_name": "connection_service_name"
    }
  }
}`
}

func assertPrivateLinkDetailsResponse(t *testing.T, response privatelink.PrivateLinkResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.Id, "123456")
	testutils.AssertEqual(t, response.Data.Name, "name")
	testutils.AssertEqual(t, response.Data.Region, "region")
	testutils.AssertEqual(t, response.Data.Service, "service")
	testutils.AssertEqual(t, response.Data.CloudProvider, "cloud_provider")
	testutils.AssertEqual(t, response.Data.State, "state")
	testutils.AssertEqual(t, response.Data.StateSummary, "state_summary")
	testutils.AssertEqual(t, response.Data.CreatedAt, "2022-04-29T09:41:08.583Z")
	testutils.AssertEqual(t, response.Data.CreatedBy, "created_by")
	testutils.AssertEqual(t, response.Data.Config.ConnectionServiceName, "connection_service_name")
}
