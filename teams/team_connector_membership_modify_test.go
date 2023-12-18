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

func TestNewTeamConnectorModify(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/teams/team_id/connectors/connector_id").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertTeamConnectorModifyRequest(t, body)
			response := mock.NewResponse(req, http.StatusOK, prepareTeamConnectorModifyResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewTeamConnectorMembershipModify().
		TeamId("team_id").
		ConnectorId("connector_id").
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

	assertTeamConnectorModifyResponse(t, response)
}

func prepareTeamConnectorModifyResponse() string {
	return fmt.Sprintf(
		`{
            "code": "Success",
            "message": "Connector membership has been updated"
        }`,
	)
}

func assertTeamConnectorModifyRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "role", request, "Changed role")
}

func assertTeamConnectorModifyResponse(t *testing.T, response common.CommonResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertNotEmpty(t, response.Message)
}
