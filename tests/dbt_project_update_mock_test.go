package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestDbtProjectUpdateMock(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(
		http.MethodPatch,
		"/v1/dbt/projects/_moonbeam_project").
		ThenCall(

			func(req *http.Request) (*http.Response, error) {
				body := requestBodyToJson(t, req)
				assertDbtProjectUpdateRequest(t, body)
				response := mock.NewResponse(
					req,
					http.StatusOK,
					prepareDbtProjectUpdateResponse())
				return response, nil
			})

	// act
	response, err := ftClient.NewDbtProjectModifyService().
		DbtProjectID("_moonbeam_project").
		GroupID("_moonbeam_updated").
		DbtVersion("1.0.2").
		GitRemoteUrl("git@some-host.com/project.git.updated").
		GitBranch("develop").
		DefaultSchema("updated_schema").
		FolderPath("updated-folder-path").
		TargetName("updated-name").
		Threads(2).
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

	assertDbtProjectUpdateResponse(t, response)
}

func assertDbtProjectUpdateRequest(t *testing.T, request map[string]interface{}) {
	assertKeyValue(t, request, "group_id", "_moonbeam_updated")
	assertKeyValue(t, request, "dbt_version", "1.0.2")
	assertKeyValue(t, request, "git_remote_url", "git@some-host.com/project.git.updated")
	assertKeyValue(t, request, "default_schema", "updated_schema")
	assertKeyValue(t, request, "git_branch", "develop")
	assertKeyValue(t, request, "folder_path", "updated-folder-path")
	assertKeyValue(t, request, "target_name", "updated-name")
	assertKeyValue(t, request, "threads", float64(2))
}

func assertDbtProjectUpdateResponse(t *testing.T, response fivetran.DbtProjectModifyResponse) {
	assertEqual(t, response.Code, "Success")

	assertEqual(t, response.Data.ID, "_moonbeam_project")
	assertEqual(t, response.Data.GroupID, "_moonbeam_updated")
	assertEqual(t, response.Data.CreatedById, "_accountworthy")
	assertEqual(t, response.Data.PublicKey, "ssh-public-key")
	assertEqual(t, response.Data.GitRemoteUrl, "git@some-host.com/project.git.updated")
	assertEqual(t, response.Data.GitBranch, "develop")
	assertEqual(t, response.Data.DefaultSchema, "updated_schema")
	assertEqual(t, response.Data.FolderPath, "updated-folder-path")
	assertEqual(t, response.Data.TargetName, "updated-name")
}

func prepareDbtProjectUpdateResponse() string {
	return `{
		"code": "Success",
		"message": "DBT Project has been updated",
		"data": {
			"id": "_moonbeam_project",
			"group_id": "_moonbeam_updated",
			"created_at": "2022-04-29T11:24:41.312868Z",
			"created_by_id": "_accountworthy",
			"public_key": "ssh-public-key",
			"git_remote_url": "git@some-host.com/project.git.updated",
			"git_branch": "develop",
			"default_schema": "updated_schema",
			"folder_path": "updated-folder-path",
			"target_name": "updated-name"
		}
	 }`
}
