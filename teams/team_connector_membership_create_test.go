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

const (
	TEAM_CONNECTOR_ROLE = "Connector Collaborator"
)

func TestNewTeamConnectorCreate(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/teams/team_id/connectors").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertTeamConnectorCreateRequest(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareTeamConnectorCreateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewTeamConnectorMembershipCreate().
		TeamId("team_id").
		ConnectorId("connector_id").
		Role(TEAM_CONNECTOR_ROLE).
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

	assertTeamConnectorCreateResponse(t, response)
}

func prepareTeamConnectorCreateResponse() string {
	return fmt.Sprintf(
		`{
            "code": "Created",
            "message": "Connector membership has been created",
            "data": {
                "id": "connector_id",
                "role": "%v",
                "created_at": "2021-09-29T10:50:51.397153Z"
            }
        }`,
		TEAM_CONNECTOR_ROLE,
	)
}

func assertTeamConnectorCreateRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "id", request, "connector_id")
	testutils.AssertKey(t, "role", request, TEAM_CONNECTOR_ROLE)
}

func assertTeamConnectorCreateResponse(t *testing.T, response teams.TeamConnectorMembershipCreateResponse) {
	testutils.AssertEqual(t, response.Code, "Created")
	testutils.AssertNotEmpty(t, response.Message)
	testutils.AssertEqual(t, response.Data.ConnectorId, "connector_id")
	testutils.AssertEqual(t, response.Data.Role, TEAM_CONNECTOR_ROLE)
	testutils.AssertNotEmpty(t, response.Data.CreatedAt)
}
