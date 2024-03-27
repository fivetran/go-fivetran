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

func TestNewProxyConnectionCreate(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/proxy/proxy_id/connections").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertProxyConnectionCreateRequest(t, body)
			response := mock.NewResponse(req, http.StatusOK, prepareProxyConnectionCreateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewProxyConnectionMembershipCreate().
		ProxyId("proxy_id").
		ConnectionId("connection_id").
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

	assertProxyConnectionCreateResponse(t, response)
}

func prepareProxyConnectionCreateResponse() string {
	return fmt.Sprintf(
		`{
            "code": "Success",
            "message": "string"
        }`)
}

func assertProxyConnectionCreateRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "connection_id", request, "connection_id")
}

func assertProxyConnectionCreateResponse(t *testing.T, response common.CommonResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertNotEmpty(t, response.Message)
}
