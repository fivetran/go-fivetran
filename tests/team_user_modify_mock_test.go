package tests

import (
    "context"
    "fmt"
    "net/http"
    "testing"

    "github.com/fivetran/go-fivetran"
    "github.com/fivetran/go-fivetran/tests/mock"
)

func TestNewTeamUserModify(t *testing.T) {
    // arrange
    ftClient, mockClient := CreateTestClient()
    handler := mockClient.When(http.MethodPatch, "/v1/teams/team_id/users/user_id").ThenCall(

        func(req *http.Request) (*http.Response, error) {
            body := requestBodyToJson(t, req)
            assertTeamUserModifyRequest(t, body)
            response := mock.NewResponse(req, http.StatusOK, prepareTeamUserModifyResponse())
            return response, nil
        })

    // act
    response, err := ftClient.NewTeamUserMembershipModify().
        TeamId("team_id").
        UserId("user_id").
        Role("Changed role").
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
    
    assertTeamUserModifyResponse(t, response)
}

func prepareTeamUserModifyResponse() string {
    return fmt.Sprintf(
        `{
            "code": "Success",
            "message": "User role has been updated"
        }`,
    )
}

func assertTeamUserModifyRequest(t *testing.T, request map[string]interface{}) {
    assertKey(t, "role", request, "Changed role")
}

func assertTeamUserModifyResponse(t *testing.T, response fivetran.TeamUserMembershipModifyResponse) {
    assertEqual(t, response.Code, "Success")
    assertNotEmpty(t, response.Message)
}