package privatelink_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/common"
	"github.com/fivetran/go-fivetran/tests/mock"
	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestPrivateLinkDeleteServiceDo(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodDelete, "/v1/private-links/123456").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, preparePrivateLinkDeleteResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewPrivateLinkDelete().
		PrivateLinkId("123456").
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	assertPrivateLinkDeleteResponse(t, response)
}

func preparePrivateLinkDeleteResponse() string {
	return `{
		"code": "Success",
		"message": "Private Link has been deleted"
	}`
}

func assertPrivateLinkDeleteResponse(t *testing.T, response common.CommonResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Message, "Private Link has been deleted")
}
