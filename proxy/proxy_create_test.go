package proxy_test

import (
    "context"
    "net/http"
    "testing"

	"github.com/fivetran/go-fivetran/proxy"
    "github.com/fivetran/go-fivetran/tests/mock"
    testutils "github.com/fivetran/go-fivetran/test_utils"
)

const (
	PROXY_NAME       = "Test Proxy"
	PROXY_GROUP      = "group_region"
)

func TestNewProxyCreate(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/proxy").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertProxyCreateRequest(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareProxyCreateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewProxyCreate().
		DisplayName(PROXY_NAME).
		GroupRegion(PROXY_GROUP).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	// assert
	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)

	assertProxyCreateResponse(t, response)
}

func prepareProxyCreateResponse() string {
	return `{
    		"code": "Success",
    		"data": {
        		"agent_id": "id",
        		"auth_token": "auth_token",
        		"proxy_server_uri": "proxy_server_uri"
    		}
		}`
}

func assertProxyCreateRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "display_name", request, PROXY_NAME)
	testutils.AssertKey(t, "group_region", request, PROXY_GROUP)
}

func assertProxyCreateResponse(t *testing.T, response proxy.ProxyCreateResponse) {
	testutils.AssertEqual(t, response.Code, "Success")

	testutils.AssertEqual(t, response.Data.AgentId, "id")
	testutils.AssertEqual(t, response.Data.AuthToken, "auth_token")
	testutils.AssertEqual(t, response.Data.ProxyServerUri, "proxy_server_uri")
}
