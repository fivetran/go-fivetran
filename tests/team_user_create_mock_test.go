package tests

import (
    "context"
    "fmt"
    "net/http"
    "testing"

    "github.com/fivetran/go-fivetran"
    "github.com/fivetran/go-fivetran/tests/mock"
)

const (
    TEAM_USER_ROLE = "Destination Administrator"
)

func TestNewTeamUserCreate(t *testing.T) {
    // arrange
    ftClient, mockClient := CreateTestClient()
    handler := mockClient.When(http.MethodPost, "/v1/teams/team_id/users").ThenCall(

        func(req *http.Request) (*http.Response, error) {
            body := requestBodyToJson(t, req)
            assertTeamUserCreateRequest(t, body)
            response := mock.NewResponse(req, http.StatusCreated, prepareTeamUserCreateResponse())
            return response, nil
        })

    // act
    response, err := ftClient.NewTeamUserMembershipCreate().
        TeamId("team_id").
        UserId("user_id").
        Role(TEAM_USER_ROLE).
        Do(context.Background())

    if err != nil {
        t.Logf("%+v\n", response)
        t.Error(err)
    }

    // assert
    interactions := mockClient.Interactions()
    assertEqual(t, len(interactions), 1)
    assertEqual(t, interactions[0].Handler, handler)
    assertEqual(t, handler.Interactions, 1)
    
    assertTeamUserCreateResponse(t, response)
}

func prepareTeamUserCreateResponse() string {
    return fmt.Sprintf(
        `{
            "code": "Created",
            "message": "User has been added to the team",
            "data": {
                "user_id": "user_id",
                "role": "%v"
            }
        }`,
        TEAM_USER_ROLE,
    )
}

func assertTeamUserCreateRequest(t *testing.T, request map[string]interface{}) {
    assertKey(t, "user_id", request, "user_id")
    assertKey(t, "role", request, TEAM_USER_ROLE)
}

func assertTeamUserCreateResponse(t *testing.T, response fivetran.TeamUserMembershipCreateResponse) {
    assertEqual(t, response.Code, "Created")
    assertNotEmpty(t, response.Message)
    assertEqual(t, response.Data.UserId, "user_id")
    assertEqual(t, response.Data.Role, TEAM_USER_ROLE)
}