package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestTeamListServiceDo(t *testing.T) {
	// arrange

	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/teams").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareTeamListResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewTeamsList().
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)
	assertTeamListResponse(t, response)
}

func prepareTeamListResponse() string {
	return fmt.Sprintf(`{
  	"code": "Success",
  		"data": {
    		"items":[
        	{
          		"id": "abode_abolish",
          		"name": "Head Team",
          		"description": "Head Team description",
          		"role": "Account Administrator"
        	},
        	{
          		"id": "numerator_handiness",
          		"name": "Finance Team",
          		"description": "Finance Team description",
          		"role": "Account Reviewer"
        	}],
    		"next_cursor": null
  		}
	}`)
}

func assertTeamListResponse(t *testing.T, response fivetran.TeamsListResponse) {
	assertEqual(t, response.Code, "Success")
	assertEqual(t, response.Data.Items[0].Id, "abode_abolish")
	assertEqual(t, response.Data.Items[0].Name, "Head Team")
	assertEqual(t, response.Data.Items[0].Description, "Head Team description")
	assertEqual(t, response.Data.Items[0].Role, "Account Administrator")
	assertEqual(t, response.Data.Items[1].Id, "numerator_handiness")
	assertEqual(t, response.Data.Items[1].Name, "Finance Team")
	assertEqual(t, response.Data.Items[1].Description, "Finance Team description")
	assertEqual(t, response.Data.Items[1].Role, "Account Reviewer")
}
