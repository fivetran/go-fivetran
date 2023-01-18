package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestUserClearPhoneSetPhoneMock(t *testing.T) {
	ftClient, _ := CreateTestClient()

	// act
	_, err := ftClient.NewUserModify().
		UserID("user_id").
		ClearPhone().
		Phone("+12345678").
		Do(context.Background())

	// assert
	assertIsNotNil(t, err)
	assertEqual(t, err.Error(), "can't 'set phone' and 'clear phone' in one request")
}

func TestUserClearPictureSetPictureMock(t *testing.T) {
	ftClient, _ := CreateTestClient()

	// act
	_, err := ftClient.NewUserModify().
		UserID("user_id").
		ClearPicture().
		Picture("http://some.picture.url/file.png").
		Do(context.Background())

	// assert
	assertIsNotNil(t, err)
	assertEqual(t, err.Error(), "can't 'set picture' and 'clear picture' in one request")
}

func TestUserClearPictureAndPhoneMock(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/users/user_id").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := requestBodyToJson(t, req)
			assertUserUpdateRequest(t, body)
			response := mock.NewResponse(req, http.StatusOK, prepareUserUpdateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewUserModify().
		UserID("user_id").
		ClearPhone().
		ClearPicture().
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
}

func prepareUserUpdateResponse() string {
	return `{
		"code": "Success",
		"data": {
			"id": "user_id",
			"email": "john.white@mycompany.com",
			"given_name": "John",
			"family_name": "White",
			"verified": false,
			"invited": true,
			"picture": null,
			"phone": null,
			"role": "Account Administrator",
			"logged_in_at": null,
			"created_at": "2019-01-20T16:03:36.786936Z",
			"active": true
		}
	}`
}

func assertUserUpdateRequest(t *testing.T, request map[string]interface{}) {
	assertHasKey(t, request, "phone")
	assertKeyValue(t, request, "phone", nil)
	assertHasKey(t, request, "picture")
	assertKeyValue(t, request, "picture", nil)
}
