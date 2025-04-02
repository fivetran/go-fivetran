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

func TestProxyConnectionListServiceDo(t *testing.T) {
	// arrange

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/proxy/proxy_id/connections").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareProxyConnectionListResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewProxyConnectionMembershipsList().
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
	assertProxyConnectionListResponse(t, response)
}

func prepareProxyConnectionListResponse() string {
	return fmt.Sprintf(`{
    			"code": "Success",
    			"data": {
        			"items": [
          				{
            				"connection_id": "string"
          				}
        			],
        		"next_cursor": null
        	}
		}`)
}

func assertProxyConnectionListResponse(t *testing.T, response proxy.ProxyConnectionMembershipsListResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.Items[0].ConnectionId, "string")
}