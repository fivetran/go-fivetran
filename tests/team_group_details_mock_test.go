package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestTeamGroupDetailsServiceDo(t *testing.T) {
	// arrange

	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/teams/team_id/groups/group_id").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareTeamGroupDetailsResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewTeamGroupMembershipDetails().
		TeamId("team_id").
		GroupId("group_id").
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)
	assertTeamGroupDetailsResponse(t, response)
}

func prepareTeamGroupDetailsResponse() string {
	return fmt.Sprintf(`{
    		"code": "Success",
    		"data": {
          		"id": "group_id",
          		"role": "Destination Administrator",
          		"created_at": "2020-05-25T15:26:47.306509Z"
    		}
		}`)
}

func assertTeamGroupDetailsResponse(t *testing.T, response fivetran.TeamGroupMembershipDetailsResponse) {
	assertEqual(t, response.Code, "Success")
	assertEqual(t, response.Data.GroupId, "group_id")
	assertEqual(t, response.Data.Role, "Destination Administrator")
	assertEqual(t, response.Data.CreatedAt, "2020-05-25T15:26:47.306509Z")
}