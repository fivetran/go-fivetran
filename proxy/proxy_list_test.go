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

func TestProxyListServiceDo(t *testing.T) {
	// arrange

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/proxy").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareProxyListResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewProxyList().
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	assertProxyListResponse(t, response)
}

func prepareProxyListResponse() string {
	return fmt.Sprintf(`{
    "code": "Success",
    "data": {
        "items": [
          {
            "id": "id",
            "account_id": "account_id",
            "registered_at": "registered_at",
            "region": "region",
            "created_by": "created_by",
            "display_name": "display_name"
          }
        ],
        "next_cursor": null
    }
}`)
}

func assertProxyListResponse(t *testing.T, response proxy.ProxyListResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.Items[0].Id, "id")
	testutils.AssertEqual(t, response.Data.Items[0].AccountId, "account_id")
	testutils.AssertEqual(t, response.Data.Items[0].RegisteredAt, "registered_at")
	testutils.AssertEqual(t, response.Data.Items[0].Region, "region")
	testutils.AssertEqual(t, response.Data.Items[0].CreatedBy, "created_by")
	testutils.AssertEqual(t, response.Data.Items[0].DisplayName, "display_name")
}