package groups_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/groups"
	"github.com/fivetran/go-fivetran/tests/mock"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

const (
	GROUP_LIST_PRIVATE_LINK_ID           	= "123456789123456789"
	GROUP_LIST_PRIVATE_LINK_NAME         	= "name"
	GROUP_LIST_PRIVATE_LINK_GROUP_ID     	= "group"
	GROUP_LIST_PRIVATE_LINK_CLOUD  			= "cloud"
	GROUP_LIST_PRIVATE_LINK_REGION  	 	= "region"
	GROUP_LIST_PRIVATE_LINK_STATE  			= "state"
	GROUP_LIST_PRIVATE_LINK_STATE_SUMMARY  	= "state_summary"
	GROUP_LIST_PRIVATE_LINK_CREATED_AT   	= "2018-01-15T11:00:27.329220Z"
	GROUP_LIST_PRIVATE_LINK_CREATED_BY   	= "created_by"
)

func TestGroupListPrivateLinksServiceDo(t *testing.T) {
	// arrange
	groupID := "projected_sickle"
	limit := 10
	cursor := "eyJza2lwIjoxfQ"

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, fmt.Sprintf("/v1/groups/%s/private-links", groupID)).
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareGroupListPrivateLinksResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewGroupListPrivateLinks().
		GroupID(groupID).
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
	assertGroupListPrivateLinksResponse(t, response)
}

func prepareGroupListPrivateLinksResponse() string {
	value := fmt.Sprintf(`{
		"code": "Success",
		"data": {
			"items": [
				{
            		"id": "%v",
            		"name": "%v",
            		"group_id": "%v",
            		"cloud_provider": "%v",
            		"region": "%v",
            		"state": "%v",
            		"state_summary": "%v",
            		"created_at": "%v",
            		"created_by": "%v"
				}
			],
			"next_cursor": "eyJza2lwIjoyfQ"
		}
	}`,
		GROUP_LIST_PRIVATE_LINK_ID,
		GROUP_LIST_PRIVATE_LINK_NAME,
		GROUP_LIST_PRIVATE_LINK_GROUP_ID,
		GROUP_LIST_PRIVATE_LINK_CLOUD,
		GROUP_LIST_PRIVATE_LINK_REGION,
		GROUP_LIST_PRIVATE_LINK_STATE,
		GROUP_LIST_PRIVATE_LINK_STATE_SUMMARY,
		GROUP_LIST_PRIVATE_LINK_CREATED_AT,
		GROUP_LIST_PRIVATE_LINK_CREATED_BY)
	return value
}

func assertGroupListPrivateLinksResponse(t *testing.T, response groups.GroupListPrivateLinksResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, len(response.Data.Items), 1)
	item := response.Data.Items[0]
	testutils.AssertEqual(t, item.Id, GROUP_LIST_PRIVATE_LINK_ID)
	testutils.AssertEqual(t, item.Name, GROUP_LIST_PRIVATE_LINK_NAME)
	testutils.AssertEqual(t, item.GroupId, GROUP_LIST_PRIVATE_LINK_GROUP_ID)
	testutils.AssertEqual(t, item.CloudProvider, GROUP_LIST_PRIVATE_LINK_CLOUD)
	testutils.AssertEqual(t, item.Region, GROUP_LIST_PRIVATE_LINK_REGION)
	testutils.AssertEqual(t, item.State, GROUP_LIST_PRIVATE_LINK_STATE)
	testutils.AssertEqual(t, item.StateSummary, GROUP_LIST_PRIVATE_LINK_STATE_SUMMARY)
	testutils.AssertEqual(t, item.CreatedAt, GROUP_LIST_PRIVATE_LINK_CREATED_AT)
	testutils.AssertEqual(t, item.CreatedBy, GROUP_LIST_PRIVATE_LINK_CREATED_BY)
}

