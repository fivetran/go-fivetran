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

func TestPrivateLinkUpdateServiceDo(t *testing.T) {
	// arrange

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/private-links/private_link_id").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, preparePrivateLinkUpdateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewPrivateLinkUpdate().
		PrivateLinkId("private_link_id").
		Config(preparePrivateLinkUpdateConfig()).
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	assertPrivateLinkUpdateResponse(t, response)
}

func TestPrivateLinkCustomUpdateService(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/private-links/private_link_id").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertPrivateLinkUpdateCustomRequest(t, body)
			response := mock.NewResponse(req, http.StatusOK, preparePrivateLinkUpdateCustomResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewPrivateLinkUpdate().
		PrivateLinkId("private_link_id").
		ConfigCustom(preparePrivateLinkCustomConfig()).
		DoCustom(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	// assert
	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)

	assertPrivateLinkUpdateCustomResponse(t, response)
}

func TestPrivateLinkCustomMergedUpdateService(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/private-links/private_link_id").
	ThenCall(func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertPrivateLinkUpdateCustomMergedRequest(t, body)
			response := mock.NewResponse(req, http.StatusOK, preparePrivateLinkUpdateMergedResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewPrivateLinkUpdate().
		PrivateLinkId("private_link_id").
		Config(preparePrivateLinkUpdateConfig()).
		ConfigCustom(preparePrivateLinkCustomConfig()).
		DoCustomMerged(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	// assert
	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)

	assertPrivateLinkUpdateCustomMergedResponse(t, response)
}

func preparePrivateLinkUpdateResponse() string {
	return `{
  "code": "Success",
  "data": {
    "id": "private_link_id",
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

func preparePrivateLinkUpdateCustomResponse() string {
	return `{
  "code": "Success",
  "data": {
    "id": "private_link_id",
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
      "connection_service_name_fake": "connection_service_name_fake"
    }
  }
}`
}

func preparePrivateLinkUpdateMergedResponse() string {
	return `{
  "code": "Success",
  "data": {
    "id": "private_link_id",
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
      "connection_service_name": "connection_service_name",
      "connection_service_name_fake": "connection_service_name_fake"
    }
  }
}`
}

func preparePrivateLinkUpdateConfig() *privatelink.PrivateLinkConfig {
	config := fivetran.NewPrivateLinkConfig()
	config.ConnectionServiceName("connection_service_name")

	return config
}

func preparePrivateLinkCustomConfig() *map[string]interface{} {
	config := make(map[string]interface{})

	config["connection_service_name_fake"] = "connection_service_name_fake"

	return &config
}

// assert Requests
func assertPrivateLinkUpdateRequest(t *testing.T, request map[string]interface{}) {
	config, ok := request["config"].(map[string]interface{})
	testutils.AssertEqual(t, ok, true)

	testutils.AssertKey(t, "connection_service_name", config, "connection_service_name")
}

func assertPrivateLinkUpdateCustomRequest(t *testing.T, request map[string]interface{}) {
	config, ok := request["config"].(map[string]interface{})

	testutils.AssertEqual(t, ok, true)

	testutils.AssertKey(t, "connection_service_name_fake", config, "connection_service_name_fake")
}

func assertPrivateLinkUpdateCustomMergedRequest(t *testing.T, request map[string]interface{}) {
	config, ok := request["config"].(map[string]interface{})

	testutils.AssertEqual(t, ok, true)

	testutils.AssertKey(t, "connection_service_name", config, "connection_service_name")
	testutils.AssertKey(t, "connection_service_name_fake", config, "connection_service_name_fake")
}

// assert Response
func assertPrivateLinkUpdateResponse(t *testing.T, response privatelink.PrivateLinkResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.Id, "private_link_id")
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

func assertPrivateLinkUpdateCustomResponse(t *testing.T, response privatelink.PrivateLinkCustomResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.Id, "private_link_id")
	testutils.AssertEqual(t, response.Data.Name, "name")
	testutils.AssertEqual(t, response.Data.Region, "region")
	testutils.AssertEqual(t, response.Data.Service, "service")
	testutils.AssertEqual(t, response.Data.CloudProvider, "cloud_provider")
	testutils.AssertEqual(t, response.Data.State, "state")
	testutils.AssertEqual(t, response.Data.StateSummary, "state_summary")
	testutils.AssertEqual(t, response.Data.CreatedAt, "2022-04-29T09:41:08.583Z")
	testutils.AssertEqual(t, response.Data.CreatedBy, "created_by")

	testutils.AssertKey(t, "connection_service_name_fake", response.Data.Config, "connection_service_name_fake")
}

func assertPrivateLinkUpdateCustomMergedResponse(t *testing.T, response privatelink.PrivateLinkCustomMergedResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.Id, "private_link_id")
	testutils.AssertEqual(t, response.Data.Name, "name")
	testutils.AssertEqual(t, response.Data.Region, "region")
	testutils.AssertEqual(t, response.Data.Service, "service")
	testutils.AssertEqual(t, response.Data.CloudProvider, "cloud_provider")
	testutils.AssertEqual(t, response.Data.State, "state")
	testutils.AssertEqual(t, response.Data.StateSummary, "state_summary")
	testutils.AssertEqual(t, response.Data.CreatedAt, "2022-04-29T09:41:08.583Z")
	testutils.AssertEqual(t, response.Data.CreatedBy, "created_by")

	testutils.AssertEqual(t, response.Data.Config.ConnectionServiceName, "connection_service_name")
	testutils.AssertKey(t, "connection_service_name_fake", response.Data.CustomConfig, "connection_service_name_fake")
}



