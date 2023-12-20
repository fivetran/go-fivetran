package teams_test

import (
    "context"
    "fmt"
    "net/http"
    "testing"

	"github.com/fivetran/go-fivetran/teams"
    
    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestTeamListServiceDo(t *testing.T) {
	// arrange

	ftClient, mockClient := testutils.CreateTestClient()
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
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
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

func assertTeamListResponse(t *testing.T, response teams.TeamsListResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.Items[0].Id, "abode_abolish")
	testutils.AssertEqual(t, response.Data.Items[0].Name, "Head Team")
	testutils.AssertEqual(t, response.Data.Items[0].Description, "Head Team description")
	testutils.AssertEqual(t, response.Data.Items[0].Role, "Account Administrator")
	testutils.AssertEqual(t, response.Data.Items[1].Id, "numerator_handiness")
	testutils.AssertEqual(t, response.Data.Items[1].Name, "Finance Team")
	testutils.AssertEqual(t, response.Data.Items[1].Description, "Finance Team description")
	testutils.AssertEqual(t, response.Data.Items[1].Role, "Account Reviewer")
}
