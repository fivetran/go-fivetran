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

func TestTeamDeleteRoleInAccountServiceDo(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodDelete, "/v1/teams/team_id/role").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, `{"code": "Success", "message": "Team role in account has been removed"}`)
			return response, nil
		},
	)

	service := ftClient.NewTeamsDeleteRoleInAccount().TeamId("team_id")

	// act
	response, err := service.Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	assertTeamDeleteRoleInAccountResponse(t, response, "Success", "Team role in account has been removed")

	// Check that the expected interactions with the mock client occurred
	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
}

func TestTeamDeleteRoleInAccountServiceDoMissingId(t *testing.T) {
	// Create a test client
	ftClient, _ := testutils.CreateTestClient()

	// Create the TeamDeleteService without setting the Team ID
	service := ftClient.NewTeamsDeleteRoleInAccount()

	// Call the Do method to execute the request
	_, err := service.Do(context.Background())

	// Check for expected error
	expectedError := fmt.Errorf("missing required teamId")
	testutils.AssertEqual(t, err, expectedError)
}

func assertTeamDeleteRoleInAccountResponse(t *testing.T, response common.CommonResponse, code string, massage string) {
	testutils.AssertEqual(t, response.Code, code)
	testutils.AssertEqual(t, response.Message, massage)
}
