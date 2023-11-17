package groups_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/common"
	testutils "github.com/fivetran/go-fivetran/test_utils"
	"github.com/fivetran/go-fivetran/tests"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestGroupDeleteServiceDo(t *testing.T) {
	// arrange
	var expectedGroupID = "group_id"
	ftClient, mockClient := tests.CreateTestClient()
	handler := mockClient.When(http.MethodDelete, "/v1/groups/"+expectedGroupID).
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareGroupDeleteResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewGroupDelete().
		GroupID(expectedGroupID).
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	assertGroupDeleteResponse(t, response)
}

func prepareGroupDeleteResponse() string {
	return `{
		"code": "Success",
		"message": "Group has been deleted"
	}`
}

func assertGroupDeleteResponse(t *testing.T, response common.CommonResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Message, "Group has been deleted")
}
