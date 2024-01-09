package privatelinks_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/private_links"
	"github.com/fivetran/go-fivetran/tests/mock"
	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestPrivateLinksModifyServiceDo(t *testing.T) {
	// arrange

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/private-links/" + GROUP_LIST_PRIVATE_LINK_ID).
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, fmt.Sprintf(`{
				"code": "Success",
				"data": {
			       		"id": "%v",
			       		"name": "%v",
			       		"group_id": "%v",
			       		"cloud_provider": "%v",
       					"service": "%v",
       					"region": "%v",
       					"state": "%v",
       					"state_summary": "%v",
       					"created_at": "%v",
       					"created_by": "%v",
       					"config" : {
       						"workspace_url": "%v"
       					}
					}
				}`,
				GROUP_LIST_PRIVATE_LINK_ID,
				GROUP_LIST_PRIVATE_LINK_NAME,
				GROUP_LIST_PRIVATE_LINK_GROUP_ID,
				GROUP_LIST_PRIVATE_LINK_CLOUD,
				GROUP_LIST_PRIVATE_LINK_SERVICE,
				GROUP_LIST_PRIVATE_LINK_REGION,
				GROUP_LIST_PRIVATE_LINK_STATE,
				GROUP_LIST_PRIVATE_LINK_STATE_SUMMARY,
				GROUP_LIST_PRIVATE_LINK_CREATED_AT,
				GROUP_LIST_PRIVATE_LINK_CREATED_BY,
				GROUP_LIST_PRIVATE_LINK_WORKSPACE_ID))
			return response, nil
		})

	// act
	response, err := ftClient.NewPrivateLinksModify().
		PrivateLinkId(GROUP_LIST_PRIVATE_LINK_ID).
		Config(preparePrivateLinksModifyConfig()).
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
	testutils.AssertEqual(t, response.Data.Id, GROUP_LIST_PRIVATE_LINK_ID)
	testutils.AssertEqual(t, response.Data.Name, GROUP_LIST_PRIVATE_LINK_NAME)
	testutils.AssertEqual(t, response.Data.GroupId, GROUP_LIST_PRIVATE_LINK_GROUP_ID)
	testutils.AssertEqual(t, response.Data.CloudProvider, GROUP_LIST_PRIVATE_LINK_CLOUD)
	testutils.AssertEqual(t, response.Data.Service, GROUP_LIST_PRIVATE_LINK_SERVICE)
	testutils.AssertEqual(t, response.Data.Region, GROUP_LIST_PRIVATE_LINK_REGION)
	testutils.AssertEqual(t, response.Data.State, GROUP_LIST_PRIVATE_LINK_STATE)
	testutils.AssertEqual(t, response.Data.StateSummary, GROUP_LIST_PRIVATE_LINK_STATE_SUMMARY)
	testutils.AssertEqual(t, response.Data.CreatedAt, GROUP_LIST_PRIVATE_LINK_CREATED_AT)
	testutils.AssertEqual(t, response.Data.CreatedBy, GROUP_LIST_PRIVATE_LINK_CREATED_BY)
	testutils.AssertEqual(t, response.Data.Config.WorkspaceUrl, GROUP_LIST_PRIVATE_LINK_WORKSPACE_ID)
}

func preparePrivateLinksModifyConfig() *privatelinks.PrivateLinksConfig {
	config := fivetran.NewPrivateLinksConfig()
	config.WorkspaceUrl(GROUP_LIST_PRIVATE_LINK_WORKSPACE_ID)

	return config
}