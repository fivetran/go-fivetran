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
    TEAM_NAME         = "Finance Team"
    TEAM_DESCRIPTION  = "Finance Team description"
    TEAM_ROLE         = "Account Analyst"
)

func TestNewTeamCreate(t *testing.T) {
    // arrange
    ftClient, mockClient := CreateTestClient()
    handler := mockClient.When(http.MethodPost, "/v1/teams").ThenCall(

        func(req *http.Request) (*http.Response, error) {
            body := requestBodyToJson(t, req)
            assertTeamCreateRequest(t, body)
            response := mock.NewResponse(req, http.StatusCreated, prepareTeamCreateResponse())
            return response, nil
        })

    // act
    response, err := ftClient.NewTeamsCreate().
        Name(TEAM_NAME).
        Description(TEAM_DESCRIPTION).
        Role(TEAM_ROLE).
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
    
    assertTeamCreateResponse(t, response)
}

func prepareTeamCreateResponse() string {
    return fmt.Sprintf(
        `{
            "code": "Success",
            "message": "Team has been created",
            "data": {
                "id": "clarification_expand",
                "name": "%v",
                "description": "%v",
                "role": "%v"
            }
        }`,
        TEAM_NAME,
        TEAM_DESCRIPTION,
        TEAM_ROLE,
    )
}

func assertTeamCreateRequest(t *testing.T, request map[string]interface{}) {
    assertKey(t, "name", request, TEAM_NAME)
    assertKey(t, "description", request, TEAM_DESCRIPTION)
    assertKey(t, "role", request, TEAM_ROLE)
}

func assertTeamCreateResponse(t *testing.T, response fivetran.TeamsCreateResponse) {
    assertEqual(t, response.Code, "Success")
    assertNotEmpty(t, response.Message)

    assertNotEmpty(t, response.Data.Id)
    assertEqual(t, response.Data.Name, TEAM_NAME)
    assertEqual(t, response.Data.Description, TEAM_DESCRIPTION)
    assertEqual(t, response.Data.Role, TEAM_ROLE)
}