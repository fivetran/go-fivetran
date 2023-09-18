package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestTeamUserListServiceDo(t *testing.T) {
	// arrange

	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/teams/team_id/users").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareTeamUserListResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewTeamUserMembershipsList().
		TeamId("team_id").
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)
	assertTeamUserListResponse(t, response)
}

func prepareTeamUserListResponse() string {
	return fmt.Sprintf(`{
    		"code": "Success",
    		"data": {
      		"items": [
        		{
          			"user_id": "user_id_1",
          			"role": "Team Member"
        		},
        		{
          			"user_id": "user_id_2",
          			"role": "Team Manager"
        		}
      		],
      		"next_cursor": null
    		}
		}`)
}

func assertTeamUserListResponse(t *testing.T, response fivetran.TeamUserMembershipsListResponse) {
	assertEqual(t, response.Code, "Success")
	assertEqual(t, response.Data.Items[0].UserId, "user_id_1")
	assertEqual(t, response.Data.Items[0].Role, "Team Member")

	assertEqual(t, response.Data.Items[1].UserId, "user_id_2")
	assertEqual(t, response.Data.Items[1].Role, "Team Manager")
}