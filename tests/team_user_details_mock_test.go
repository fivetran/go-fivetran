package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestTeamUserDetailsServiceDo(t *testing.T) {
	// arrange

	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/teams/team_id/users/user_id").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareTeamUserDetailsResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewTeamUserMembershipDetails().
		TeamId("team_id").
		UserId("user_id").
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)
	assertTeamUserDetailsResponse(t, response)
}

func prepareTeamUserDetailsResponse() string {
	return fmt.Sprintf(`{
    		"code": "Success",
    		"data": {
      			"user_id": "user_id",
      			"role": "Team Member"
    		}
		}`)
}

func assertTeamUserDetailsResponse(t *testing.T, response fivetran.TeamUserMembershipDetailsResponse) {
	assertEqual(t, response.Code, "Success")
	assertEqual(t, response.Data.UserId, "user_id")
	assertEqual(t, response.Data.Role, "Team Member")
}