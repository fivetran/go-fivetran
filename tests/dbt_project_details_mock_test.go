package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestDbtProjectDetailsMock(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/dbt/projects/_moonbeam_project").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(
				req,
				http.StatusOK,
				prepareDbtProjectDetailsResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewDbtProjectDetailsService().
		DbtProjectID("_moonbeam_project").
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

	assertDbtTransformationDetailsResponse(t, response)
}

func assertDbtTransformationDetailsResponse(t *testing.T, response fivetran.DbtProjectDetailsResponse) {

	assertEqual(t, response.Code, "Success")

	assertEqual(t, response.Data.ID, "_moonbeam_project")
	assertEqual(t, response.Data.GroupID, "_moonbeam_bright")
	assertEqual(t, response.Data.CreatedById, "_accountworthy")
	assertEqual(t, response.Data.PublicKey, "ssh-public-key")
	assertEqual(t, response.Data.GitRemoteUrl, "git@some-host.com/project.git")
	assertEqual(t, response.Data.GitBranch, "main")
	assertEqual(t, response.Data.DefaultSchema, "some_schema")
	assertEqual(t, response.Data.FolderPath, "some-folder-in-git-repo")
	assertEqual(t, response.Data.TargetName, "some-name")
}

func prepareDbtProjectDetailsResponse() string {
	return `{
		"code": "Success",
		"data": {
			"id": "_moonbeam_project",
			"group_id": "_moonbeam_bright",
			"created_at": "2022-04-29T11:24:41.312868Z",
			"created_by_id": "_accountworthy",
			"public_key": "ssh-public-key",
			"git_remote_url": "git@some-host.com/project.git",
			"git_branch": "main",
			"default_schema": "some_schema",
			"folder_path": "some-folder-in-git-repo",
			"target_name": "some-name"
		}
	 }`
}
