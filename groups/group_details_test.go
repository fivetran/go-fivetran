package groups_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/groups"
	testutils "github.com/fivetran/go-fivetran/test_utils"
	
	"github.com/fivetran/go-fivetran/tests/mock"
)

const (
	EXPECTED_GROUP_ID   = "projected_sickle"
	EXPECTED_GROUP_NAME = "Staging"
	EXPECTED_CREATED_AT = "2018-12-20T11:59:35.089589Z"
)

func TestGroupDetailsServiceDo(t *testing.T) {
	// arrange

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/groups/"+EXPECTED_GROUP_ID).
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareGroupDetailsResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewGroupDetails().
		GroupID(EXPECTED_GROUP_ID).
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	assertGroupDetailsResponse(t, response, EXPECTED_GROUP_ID, EXPECTED_GROUP_NAME, EXPECTED_CREATED_AT)
}

func prepareGroupDetailsResponse() string {
	return fmt.Sprintf(`{
		"code": "Success",
		"data": {
			"id": "%v",
			"name": "%v",
			"created_at": "%v"
		}
	}`,
		EXPECTED_GROUP_ID,
		EXPECTED_GROUP_NAME,
		EXPECTED_CREATED_AT)
}

func assertGroupDetailsResponse(t *testing.T, response groups.GroupDetailsResponse, expectedID, expectedName string, expectedCreatedAt string) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.ID, expectedID)
	testutils.AssertEqual(t, response.Data.Name, expectedName)
	testutils.AssertTimeEqual(t, response.Data.CreatedAt, expectedCreatedAt)
}
