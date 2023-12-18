package users_test

import (
    "context"
    "fmt"
    "net/http"
    "testing"

	"github.com/fivetran/go-fivetran/users"
    
    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

const (
	TEAM_CONNECTOR_ROLE = "Connector Collaborator"
)

func TestNewUserConnectorCreate(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/users/user_id/connectors").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertUserConnectorCreateRequest(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareUserConnectorCreateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewUserConnectorMembershipCreate().
		UserId("user_id").
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

	assertUserConnectorCreateResponse(t, response)
}

func prepareUserConnectorCreateResponse() string {
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

func assertUserConnectorCreateRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "id", request, "connector_id")
	testutils.AssertKey(t, "role", request, TEAM_CONNECTOR_ROLE)
}

func assertUserConnectorCreateResponse(t *testing.T, response users.UserConnectorMembershipCreateResponse) {
	testutils.AssertEqual(t, response.Code, "Created")
	testutils.AssertNotEmpty(t, response.Message)
	testutils.AssertEqual(t, response.Data.ConnectorId, "connector_id")
	testutils.AssertEqual(t, response.Data.Role, TEAM_CONNECTOR_ROLE)
	testutils.AssertNotEmpty(t, response.Data.CreatedAt)
}
