package roles_test

import (
    "context"
    "net/http"
    "testing"

    "github.com/fivetran/go-fivetran/roles"
    testutils "github.com/fivetran/go-fivetran/test_utils"
    
    "github.com/fivetran/go-fivetran/tests/mock"
)

func TestRolesListServiceDo(t *testing.T) {
    ftClient, mockClient := testutils.CreateTestClient()

    responseData := `{
      "code": "Success",
      "data": {
         "items": [
        {
            "name": "Account Administrator",
            "description": "Can view and change account information, including billing, users, roles, API access, and security settings. Can create, manage, and delete destinations and connectors. Can manage transformations and logs.",
            "is_custom": false,
            "scope": ["ACCOUNT"],
            "is_deprecated": null,
            "replacement_role_name": null
        },
        {
            "name": "Destination Reviewer",
            "description": "Can view the destinations that you are invited to and their associated connectors. Cannot create, delete, or manage destinations or connectors. Cannot access account information.",
            "is_custom": false,
            "scope": ["DESTINATION"],
            "is_deprecated": false,
            "replacement_role_name": "test"
        }
        ],
        "next_cursor": "eyJza2lwIjoxfQ"
       } 
    }`

    handler := mockClient.When(http.MethodGet, "/v1/roles").ThenCall(
        func(req *http.Request) (*http.Response, error) {
            response := mock.NewResponse(req, http.StatusOK, responseData)
            return response, nil
        },
    )

    service := ftClient.NewRolesList()

    response, err := service.Do(context.Background())
    if err != nil {
        t.Error(err)
    }

    flag := false
    expectedResponse := roles.RolesListResponse{
        Code: "Success",
        Data: struct {
            Items []struct {
                Name                string    `json:"name"`
                Description         string    `json:"description"`
                IsCustom            *bool     `json:"is_custom"`
                Scope               []string  `json:"scope"`
                IsDeprecated        *bool     `json:"is_deprecated"`
                ReplacementRoleName string    `json:"replacement_role_name"`
            } `json:"items"`
            NextCursor string `json:"next_cursor"`
        }{
            Items: []struct {
                Name                string    `json:"name"`
                Description         string    `json:"description"`
                IsCustom            *bool     `json:"is_custom"`
                Scope               []string  `json:"scope"`
                IsDeprecated        *bool     `json:"is_deprecated"`
                ReplacementRoleName string    `json:"replacement_role_name"`
            }{
                {
                    Name:                   "Account Administrator",
                    Description:            "Can view and change account information, including billing, users, roles, API access, and security settings. Can create, manage, and delete destinations and connectors. Can manage transformations and logs.",
                    IsCustom:               &flag,
                    Scope:                  []string{"ACCOUNT"},
                    IsDeprecated:           nil,
                    ReplacementRoleName:    "",
                },
                {
                    Name:                   "Destination Reviewer",
                    Description:            "Can view the destinations that you are invited to and their associated connectors. Cannot create, delete, or manage destinations or connectors. Cannot access account information.",
                    IsCustom:               &flag,
                    Scope:                  []string{"DESTINATION"},
                    IsDeprecated:           &flag,
                    ReplacementRoleName:    "test",
                },
            },
            NextCursor: "eyJza2lwIjoxfQ",
        },
    }

    assertRolesListResponse(t, response, expectedResponse)

    interactions := mockClient.Interactions()
    testutils.AssertEqual(t, len(interactions), 1)
    testutils.AssertEqual(t, interactions[0].Handler, handler)
    testutils.AssertEqual(t, handler.Interactions, 1)

}

func assertRolesListResponse(t *testing.T, response roles.RolesListResponse, expected roles.RolesListResponse) {
    testutils.AssertEqual(t, response.Code, expected.Code)
    testutils.AssertEqual(t, response.Data.NextCursor, expected.Data.NextCursor)

    // Assert items
    testutils.AssertEqual(t, len(response.Data.Items), len(expected.Data.Items))
    for i, item := range response.Data.Items {
        testutils.AssertEqual(t, item.Name, expected.Data.Items[i].Name)
        testutils.AssertEqual(t, item.Description, expected.Data.Items[i].Description)
        testutils.AssertEqual(t, item.IsCustom, expected.Data.Items[i].IsCustom)
        testutils.AssertEqual(t, item.Scope, expected.Data.Items[i].Scope)
        testutils.AssertEqual(t, item.IsDeprecated, expected.Data.Items[i].IsDeprecated)
        testutils.AssertEqual(t, item.ReplacementRoleName, expected.Data.Items[i].ReplacementRoleName)
    }
}
