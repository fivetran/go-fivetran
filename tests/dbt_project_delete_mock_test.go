package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestDbtProjectDeleteServiceDo(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	projectID := "project123"

	handler := mockClient.When(http.MethodDelete, fmt.Sprintf("/v1/projects/%s", projectID)).
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareProjectDeleteResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewDbtProjectDelete().
		ProjectID(projectID).
		Do(context.Background())

	// assert
	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)

	assertProjectDeleteResponse(t, response)
}

func prepareProjectDeleteResponse() string {
	return fmt.Sprintf(
		`{
			"code": "Success",
			"message": "Project has been deleted"
		}`)
}

func assertProjectDeleteResponse(t *testing.T, response fivetran.DbtProjectDeleteResponse) {
	assertEqual(t, response.Code, "Success")
	assertEqual(t, response.Message, "Project has been deleted")
}
