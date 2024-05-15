package privatelink_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/private_link"
	"github.com/fivetran/go-fivetran/tests/mock"
	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestPrivateLinkCreateServiceDo(t *testing.T) {
	// arrange

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/private-links").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusCreated, preparePrivateLinkCreateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewPrivateLinkCreate().
		Name("name").
		Region("region").
		Service("service").
		Config(preparePrivateLinkConfig()).
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	assertPrivateLinkCreateResponse(t, response)
}

func preparePrivateLinkConfig() *privatelink.PrivateLinkConfig {
	config := fivetran.NewPrivateLinkConfig()
	config.ConnectionServiceName("connection_service_name")

	return config
}

func preparePrivateLinkCreateResponse() string {
	return `{
    "code": "Success",
    "data": {
        "id": "id",
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

func assertPrivateLinkCreateResponse(t *testing.T, response privatelink.PrivateLinkResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.Id, "id")
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
