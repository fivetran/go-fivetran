package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestGroupDetailsServiceDo(t *testing.T) {
	// arrange
	expectedGroupID := "projected_sickle"
	expectedGroupName := "Staging"
	expectedCreatedAt, _ := time.Parse(time.RFC3339, "2018-12-20T11:59:35.089589Z")

	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/groups/"+expectedGroupID).
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareGroupDetailsResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewGroupDetails().
		GroupID(expectedGroupID).
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)
	assertGroupDetailsResponse(t, response, expectedGroupID, expectedGroupName, expectedCreatedAt)
}

func prepareGroupDetailsResponse() string {
	return fmt.Sprintf(`{
		"code": "Success",
		"data": {
			"id": "projected_sickle",
			"name": "Staging",
			"created_at": "2018-12-20T11:59:35.089589Z"
		}
	}`)
}

func assertGroupDetailsResponse(t *testing.T, response fivetran.GroupDetailsResponse, expectedID, expectedName string, expectedCreatedAt time.Time) {
	assertEqual(t, response.Code, "Success")
	assertEqual(t, response.Data.ID, expectedID)
	assertEqual(t, response.Data.Name, expectedName)
	assertEqual(t, response.Data.CreatedAt, expectedCreatedAt)
}
