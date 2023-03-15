package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestNewDbtProjectCreate(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/dbt/projects").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := requestBodyToJson(t, req)
			assertDbtProjectRequest(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareDbtProjectResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewDbtProjectCreateService().
		GroupID("_moonbeam_bright").
		DbtVersion("1.0.1").
		GitRemoteUrl("git@some-host.com/project.git").
		GitBranch("main").
		DefaultSchema("some_schema").
		FolderPath("some-folder-in-git-repo").
		TargetName("some-name").
		Threads(1).
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

	assertDbtProjectResponse(t, response)
}

func assertDbtProjectRequest(t *testing.T, request map[string]interface{}) {
	assertKey(t, "group_id", request, "_moonbeam_bright")
	assertKey(t, "dbt_version", request, "1.0.1")
	assertKey(t, "git_remote_url", request, "git@some-host.com/project.git")
	assertKey(t, "git_branch", request, "main")
	assertKey(t, "default_schema", request, "some_schema")
	assertKey(t, "folder_path", request, "some-folder-in-git-repo")
	assertKey(t, "target_name", request, "some-name")
	assertKey(t, "threads", request, float64(1))
}

func assertDbtProjectResponse(t *testing.T, response fivetran.DbtProjectCreateResponse) {

	assertEqual(t, response.Code, "Success")
	assertNotEmpty(t, response.Message)

	assertEqual(t, response.Data.ID, "_moonbeam_project")
	assertEqual(t, response.Data.GroupID, "_moonbeam_bright")
}

func prepareDbtProjectResponse() string {
	return `{
		"code": "Success",
		"message": "DBT Project has been created",
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
