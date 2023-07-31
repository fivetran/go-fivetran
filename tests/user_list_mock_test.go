package tests

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestUsersListServiceDo(t *testing.T) {
	ftClient, mockClient := CreateTestClient()

	responseData := `{
		"code": "Success",
		"data": {
			"items": [
				{
					"id": "nozzle_eat",
					"email": "john@mycompany.com",
					"given_name": "John",
					"family_name": "White",
					"verified": true,
					"invited": true,
					"picture": null,
					"phone": null,
					"role": "Account Administrator",
					"logged_in_at": "2019-01-03T08:44:45.369Z",
					"created_at": "2018-01-15T11:00:27.329220Z",
					"active": true
				},
				{
					"id": "prophecies_falsely",
					"email": "robert@mycompany.com",
					"given_name": "Robert",
					"family_name": "Brown",
					"verified": true,
					"invited": true,
					"picture": null,
					"phone": null,
					"role": null,
					"logged_in_at": "2018-12-12T12:06:15.337Z",
					"created_at": "2018-01-24T20:43:32.963843Z",
					"active": true
				}
			],
			"next_cursor": "eyJza2lwIjoyfQ"
		}
	}`

	handler := mockClient.When(http.MethodGet, "/v1/users").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, responseData)
			return response, nil
		},
	)

	service := ftClient.NewUsersList()

	response, err := service.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	flag := true
	expectedResponse := fivetran.UsersListResponse{
		Code: "Success",
		Data: struct {
			Items []struct {
				ID         string    `json:"id"`
				Email      string    `json:"email"`
				GivenName  string    `json:"given_name"`
				FamilyName string    `json:"family_name"`
				Verified   *bool     `json:"verified"`
				Invited    *bool     `json:"invited"`
				Picture    string    `json:"picture"`
				Phone      string    `json:"phone"`
				Role       string    `json:"role"`
				LoggedInAt time.Time `json:"logged_in_at"`
				CreatedAt  time.Time `json:"created_at"`
			} `json:"items"`
			NextCursor string `json:"next_cursor"`
		}{
			Items: []struct {
				ID         string    `json:"id"`
				Email      string    `json:"email"`
				GivenName  string    `json:"given_name"`
				FamilyName string    `json:"family_name"`
				Verified   *bool     `json:"verified"`
				Invited    *bool     `json:"invited"`
				Picture    string    `json:"picture"`
				Phone      string    `json:"phone"`
				Role       string    `json:"role"`
				LoggedInAt time.Time `json:"logged_in_at"`
				CreatedAt  time.Time `json:"created_at"`
			}{
				{
					ID:         "nozzle_eat",
					Email:      "john@mycompany.com",
					GivenName:  "John",
					FamilyName: "White",
					Verified:   &flag,
					Invited:    &flag,
					Picture:    "",
					Phone:      "",
					Role:       "Account Administrator",
					LoggedInAt: parseTime("2019-01-03T08:44:45.369Z"),
					CreatedAt:  parseTime("2018-01-15T11:00:27.329220Z"),
				},
				{
					ID:         "prophecies_falsely",
					Email:      "robert@mycompany.com",
					GivenName:  "Robert",
					FamilyName: "Brown",
					Verified:   &flag,
					Invited:    &flag,
					Picture:    "",
					Phone:      "",
					Role:       "",
					LoggedInAt: parseTime("2018-12-12T12:06:15.337Z"),
					CreatedAt:  parseTime("2018-01-24T20:43:32.963843Z"),
				},
			},
			NextCursor: "eyJza2lwIjoyfQ",
		},
	}

	assertUsersListResponse(t, response, expectedResponse)

	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)

}

func assertUsersListResponse(t *testing.T, response fivetran.UsersListResponse, expected fivetran.UsersListResponse) {
	assertEqual(t, response.Code, expected.Code)
	assertEqual(t, response.Data.NextCursor, expected.Data.NextCursor)

	// Assert items
	assertEqual(t, len(response.Data.Items), len(expected.Data.Items))
	for i, item := range response.Data.Items {
		assertEqual(t, item.ID, expected.Data.Items[i].ID)
		assertEqual(t, item.Email, expected.Data.Items[i].Email)
		assertEqual(t, item.GivenName, expected.Data.Items[i].GivenName)
		assertEqual(t, item.FamilyName, expected.Data.Items[i].FamilyName)

		if item.Verified != nil {
			assertEqual(t, *item.Verified, *expected.Data.Items[i].Verified)
		} else {
			assertNil(t, expected.Data.Items[i].Verified)
		}

		if item.Invited != nil {
			assertEqual(t, *item.Invited, *expected.Data.Items[i].Invited)
		} else {
			assertNil(t, expected.Data.Items[i].Invited)
		}

		assertEqual(t, item.Picture, expected.Data.Items[i].Picture)
		assertEqual(t, item.Phone, expected.Data.Items[i].Phone)
		assertEqual(t, item.Role, expected.Data.Items[i].Role)
		assertTimeEqual(t, item.LoggedInAt, expected.Data.Items[i].LoggedInAt)
		assertTimeEqual(t, item.CreatedAt, expected.Data.Items[i].CreatedAt)
	}
}

func assertNil(t *testing.T, b *bool) {
	if b != nil {
		t.Errorf("Expected nil, got: %v", *b)
	}
}

func assertTimeEqual(t *testing.T, actual, expected time.Time) {
	if !actual.Equal(expected) {
		t.Errorf("Expected time: %v, got: %v", expected, actual)
	}
}
