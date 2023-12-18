package groups_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
	
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestGroupModifyServiceDo(t *testing.T) {
	// arrange
	const GROUP_MODIFY_GROUP_ID = "decent_dropsy"
	const GROUP_MODIFY_EXPECTED_GROUP_NAME = "New_Group_Name"
	const GROUP_MODIFY_CREATED_TIME = "2020-05-25T15:26:47.306509Z"

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/groups/"+GROUP_MODIFY_GROUP_ID).
		ThenCall(func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			testutils.AssertKey(t, "name", body, GROUP_MODIFY_EXPECTED_GROUP_NAME)

			response := mock.NewResponse(req, http.StatusOK, fmt.Sprintf(`
				{
					"code": "Success",
					"message": "Group has been updated",
					"data": {
						"id": "%s",
						"name": "%s",
						"created_at": "%s"
					}
				}`,
				GROUP_MODIFY_GROUP_ID,
				GROUP_MODIFY_EXPECTED_GROUP_NAME,
				GROUP_MODIFY_CREATED_TIME))
			return response, nil
		})

	// act
	response, err := ftClient.NewGroupModify().
		GroupID(GROUP_MODIFY_GROUP_ID).
		Name(GROUP_MODIFY_EXPECTED_GROUP_NAME).
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Message, "Group has been updated")
	testutils.AssertEqual(t, response.Data.ID, GROUP_MODIFY_GROUP_ID)
	testutils.AssertEqual(t, response.Data.Name, GROUP_MODIFY_EXPECTED_GROUP_NAME)
	testutils.AssertTimeEqual(t, response.Data.CreatedAt, GROUP_MODIFY_CREATED_TIME)
}
