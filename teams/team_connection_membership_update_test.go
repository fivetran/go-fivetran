package teams_test

import (
    "context"
    "fmt"
    "net/http"
    "testing"

	"github.com/fivetran/go-fivetran/common"
    
    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewTeamConnectionUpdate(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/teams/team_id/connections/connection_id").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertTeamConnectionUpdateRequest(t, body)
			response := mock.NewResponse(req, http.StatusOK, prepareTeamConnectionUpdateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewTeamConnectionMembershipUpdate().
		TeamId("team_id").
		ConnectionId("connection_id").
		Role("Changed role").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	// assert
	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)

	assertTeamConnectionUpdateResponse(t, response)
}

func prepareTeamConnectionUpdateResponse() string {
	return fmt.Sprintf(
		`{
            "code": "Success",
            "message": "Connection membership has been updated"
        }`,
	)
}

func assertTeamConnectionUpdateRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "role", request, "Changed role")
}

func assertTeamConnectionUpdateResponse(t *testing.T, response common.CommonResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertNotEmpty(t, response.Message)
}
