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

func TestProxyDetailsServiceDo(t *testing.T) {
	// arrange

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/proxy/proxy_id").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareProxyDetailsResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewProxyDetails().
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
	assertProxyDetailsResponse(t, response)
}

func prepareProxyDetailsResponse() string {
	return fmt.Sprintf(`{
  		"code": "Success",
  		"data": {
            "id": "id",
            "account_id": "account_id",
            "registred_at": "registred_at",
            "region": "region",
            "token": "token",
            "salt": "salt",
            "created_by": "created_by",
            "display_name": "display_name"
  			}
		}`)
}

func assertProxyDetailsResponse(t *testing.T, response proxy.ProxyDetailsResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.Id, "id")
	testutils.AssertEqual(t, response.Data.AccountId, "account_id")
	testutils.AssertEqual(t, response.Data.RegistredAt, "registred_at")
	testutils.AssertEqual(t, response.Data.Region, "region")
	testutils.AssertEqual(t, response.Data.Token, "token")
	testutils.AssertEqual(t, response.Data.Salt, "salt")
	testutils.AssertEqual(t, response.Data.CreatedBy, "created_by")
	testutils.AssertEqual(t, response.Data.DisplayName, "display_name")
}