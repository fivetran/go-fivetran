package proxy_test

import (
    "context"
    "fmt"
    "net/http"
    "testing"

	"github.com/fivetran/go-fivetran/proxy"
    
    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestProxyRegeneratesSecretsServiceDo(t *testing.T) {
	// arrange

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/proxy/proxy_id").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareProxyRegenerateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewProxyRegenerateSecrets().
		ProxyId("proxy_id").
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	assertProxyRegenerateResponse(t, response)
}

func prepareProxyRegenerateResponse() string {
	return fmt.Sprintf(`{
  				"code": "Created",
  				"message": "Operation performed.",
  				"data": {
    				"client_cert": "client_cert",
    				"agent_id": "proxy_id",
    				"proxy_server_uri": "proxy_server_uri",
    				"auth_token": "auth_token",
    				"client_private_key": "client_private_key"
  				}
			}`)
}

func assertProxyRegenerateResponse(t *testing.T, response proxy.ProxyCreateResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.AgentId, "agent_id")
	testutils.AssertEqual(t, response.Data.AuthToken, "auth_token")
	testutils.AssertEqual(t, response.Data.ClientCert, "client_cert")
	testutils.AssertEqual(t, response.Data.ClientPrivateKey, "client_private_key")
}