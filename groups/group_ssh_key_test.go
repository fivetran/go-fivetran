package groups_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
	
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestGroupSshKeyServiceDo(t *testing.T) {
	// arrange
	sshPublicKey := `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC2l5Tq4JWBHyTb46aGRQ== fivetran user key\\n` // we have to escape char here for serialization
	expectedKey := `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC2l5Tq4JWBHyTb46aGRQ== fivetran user key\n`   // we expect \n at the end of key value

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/groups/"+EXPECTED_GROUP_ID+"/public-key").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, fmt.Sprintf(`{
				"code": "Success",
				"data": {
					"public_key": "%v"
				}
			}`,
				sshPublicKey))
			return response, nil
		})

	// act
	response, err := ftClient.NewGroupSshPublicKey().
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
	testutils.AssertEqual(t, response.Data.PublicKey, expectedKey)
}
