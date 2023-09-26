package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestTeamGroupListServiceDo(t *testing.T) {
	// arrange

	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/teams/team_id/groups").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareTeamGroupListResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewTeamGroupMembershipsList().
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
	assertTeamGroupListResponse(t, response)
}

func prepareTeamGroupListResponse() string {
	return fmt.Sprintf(`{
    		"code": "Success",
    		"data": {
      		"items": [
        		{
          			"id": "group_id_1",
          			"role": "Destination Administrator",
          			"created_at": "2020-05-25T15:26:47.306509Z"
        		},
        		{
          			"id": "group_id_2",
          			"role": "Destination Reviewer",
          			"created_at": "2020-05-25T15:26:47.306509Z"
        		}
      		],
      		"next_cursor": null
    		}
		}`)
}

func assertTeamGroupListResponse(t *testing.T, response fivetran.TeamGroupMembershipsListResponse) {
	assertEqual(t, response.Code, "Success")
	assertEqual(t, response.Data.Items[0].GroupId, "group_id_1")
	assertEqual(t, response.Data.Items[0].Role, "Destination Administrator")
	assertEqual(t, response.Data.Items[0].CreatedAt, "2020-05-25T15:26:47.306509Z")

	assertEqual(t, response.Data.Items[1].GroupId, "group_id_2")
	assertEqual(t, response.Data.Items[1].Role, "Destination Reviewer")
	assertEqual(t, response.Data.Items[1].CreatedAt, "2020-05-25T15:26:47.306509Z")
}