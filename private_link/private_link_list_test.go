package privatelink_test

import (
    "context"
    "net/http"
    "testing"
		"github.com/fivetran/go-fivetran/private_link" 
    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestPrivateLinkListServiceDo(t *testing.T) {
	// arrange

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/private-links").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, preparePrivateLinkListResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewPrivateLinkList().
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	assertPrivateLinkListResponse(t, response)
}

func preparePrivateLinkListResponse() string {
	return `{
	    "code": "Success",
	    "data": {
	        "items": [
	          {
	            "id": "123456",
	            "name": "name",
	            "region": "region",
	            "service": "service",
	            "account_id": "account_id",
	            "cloud_provider": "cloud_provider",
	            "host": "host",
	            "state": "state",
	            "state_summary": "state_summary",
	            "created_at": "2022-04-29T09:41:08.583Z",
	            "created_by": "created_by"
	          }
	        ],
	        "next_cursor": null
	    }
	}`
}

func assertPrivateLinkListResponse(t *testing.T, response privatelink.PrivateLinkListResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.Items[0].Id, "123456")
	testutils.AssertEqual(t, response.Data.Items[0].Name, "name")
	testutils.AssertEqual(t, response.Data.Items[0].Region, "region")
	testutils.AssertEqual(t, response.Data.Items[0].Service, "service")
	testutils.AssertEqual(t, response.Data.Items[0].CloudProvider, "cloud_provider")
	testutils.AssertEqual(t, response.Data.Items[0].Host, "host")
	testutils.AssertEqual(t, response.Data.Items[0].State, "state")
	testutils.AssertEqual(t, response.Data.Items[0].StateSummary, "state_summary")
	testutils.AssertEqual(t, response.Data.Items[0].CreatedAt, "2022-04-29T09:41:08.583Z")
	testutils.AssertEqual(t, response.Data.Items[0].CreatedBy, "created_by")
}