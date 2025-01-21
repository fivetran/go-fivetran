package transformations_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/transformations"
	testutils "github.com/fivetran/go-fivetran/test_utils"
	
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestTransformationProjectsListServiceDo(t *testing.T) {
	// arrange
	limit := 10
	cursor := "some_cursor"

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/transformation-projects").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareTransformationProjectsListResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewTransformationProjectsList().
		Limit(limit).
		Cursor(cursor).
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	assertTransformationProjectsListResponse(t, response)
}

func prepareTransformationProjectsListResponse() string {
	return `{
  "code": "Success",
  "message": "Operation performed.",
  "data": {
    "items": [
      {
        "id": "string",
        "type": "DBT_GIT",
        "created_at": "2019-08-24T14:15:22Z",
        "created_by_id": "string",
        "group_id": "string"
      }
    ],
    "next_cursor": "cursor_value"
  }
}`
}

func assertTransformationProjectsListResponse(t *testing.T, response transformations.TransformationProjectsListResponse) {
  testutils.AssertEqual(t, response.Code, "Success")
  testutils.AssertEqual(t, response.Data.Items[0].Id, "string")
  testutils.AssertEqual(t, response.Data.Items[0].ProjectType, "DBT_GIT")
  testutils.AssertEqual(t, response.Data.Items[0].CreatedAt, "2019-08-24T14:15:22Z")
  testutils.AssertEqual(t, response.Data.Items[0].GroupId, "string")
  testutils.AssertEqual(t, response.Data.Items[0].CreatedById, "string")

	testutils.AssertEqual(t, response.Data.NextCursor, "cursor_value")
}