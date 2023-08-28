package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

const (
	ExpectedGroupID = "decent_dropsy"
)

func TestGroupDeleteServiceDo(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodDelete, "/v1/groups/"+ExpectedGroupID).
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareGroupDeleteResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewGroupDelete().
		GroupID(ExpectedGroupID).
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)
	assertGroupDeleteResponse(t, response)
}

func prepareGroupDeleteResponse() string {
	return fmt.Sprintf(`{
		"code": "Success",
		"message": "Group has been deleted"
	}`)
}

func assertGroupDeleteResponse(t *testing.T, response fivetran.GroupDeleteResponse) {
	assertEqual(t, response.Code, "Success")
	assertEqual(t, response.Message, "Group has been deleted")
}
