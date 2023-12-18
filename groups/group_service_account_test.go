package groups_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
	
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestGroupServiceAccountServiceDo(t *testing.T) {
	// arrange
	serviceAccount := "g-group_id@fivetran-production.iam.gserviceaccount.com"

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/groups/"+EXPECTED_GROUP_ID+"/service-account").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, fmt.Sprintf(
				`{
					"code": "Success",
					"data": {
						"service_account": "%v"
					}
				}`,
				serviceAccount))
			return response, nil
		})

	// act
	response, err := ftClient.NewGroupServiceAccount().
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

	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.ServiceAccount, serviceAccount)
}
