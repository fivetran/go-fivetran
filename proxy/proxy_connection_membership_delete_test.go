package proxy_test

import (
    "context"
    "fmt"
    "net/http"
    "testing"

	"github.com/fivetran/go-fivetran/common"
    
    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestProxyConnectionDeleteServiceDo(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodDelete, "/v1/proxy/proxy_id/connections/connection_id").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, `{"code": "Success"}`)
			return response, nil
		},
	)

	service := ftClient.NewProxyConnectionMembershipDelete().ProxyId("proxy_id").ConnectionId("connection_id")

	// act
	response, err := service.Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	assertProxyConnectionDeleteResponse(t, response, "Success")

	// Check that the expected interactions with the mock client occurred
	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
}

func TestProxyConnectionDeleteServiceDoMissingId(t *testing.T) {
	ftClient, _ := testutils.CreateTestClient()
	service := ftClient.NewProxyConnectionMembershipDelete()

	// Call the Do method to execute the request
	_, err := service.Do(context.Background())

	// Check for expected error
	expectedError := fmt.Errorf("missing required proxyId")
	testutils.AssertEqual(t, err, expectedError)
}

func TestProxyConnectionDeleteServiceDoMissingConnectionId(t *testing.T) {
	ftClient, _ := testutils.CreateTestClient()
	service := ftClient.NewProxyConnectionMembershipDelete().ProxyId("proxy_id")

	// Call the Do method to execute the request
	_, err := service.Do(context.Background())

	// Check for expected error
	expectedError := fmt.Errorf("missing required connectionId")
	testutils.AssertEqual(t, err, expectedError)
}

func assertProxyConnectionDeleteResponse(t *testing.T, response common.CommonResponse, code string) {
	testutils.AssertEqual(t, response.Code, code)
}
