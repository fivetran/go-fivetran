package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

const (
	DBT_GROUP_ID       = "_moonbeam"
	DBT_VERSION        = "1.3.1"
	DBT_GIT_REMOTE_URL = "https://github.com/fivetran/dbt_demo"
	DBT_GIT_BRANCH     = "main"
	DBT_DEFAULT_SCHEMA = "schema"
	DBT_FOLDER_PATH    = "/"
	DBT_TARGET_NAME    = "target"
	DBT_THREADS        = 4
)

func TestNewDbtProjectCreateFullMappingMock(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/dbt/projects").
		ThenCall(

			func(req *http.Request) (*http.Response, error) {
				body := requestBodyToJson(t, req)
				assertDbtProjectRequest(t, body)
				response := mock.NewResponse(req, http.StatusCreated, prepareDbtProjectResponse())
				return response, nil
			})

	// act
	response, err := ftClient.NewDbtProjectCreate().
		GroupID(DBT_GROUP_ID).
		DbtVersion(DBT_VERSION).
		GitRemoteUrl(DBT_GIT_REMOTE_URL).
		GitBranch(DBT_GIT_BRANCH).
		DefaultSchema(DBT_DEFAULT_SCHEMA).
		FolderPath(DBT_FOLDER_PATH).
		TargetName(DBT_TARGET_NAME).
		Threads(DBT_THREADS).
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

func prepareDbtProjectResponse() string {
	return fmt.Sprintf(
		`{
			"code": "Created",
			"message": "Dbt project has been created",
			"data": {
			  "id": "%v",
			  "folder_path": "%v",
			  "created_at": "2023-08-24T14:15:22Z",
			  "target_name": "%v",
			  "git_remote_url": "%v",
			  "default_schema": "%v",
			  "group_id": "%v",
			  "public_key": "public_key",
			  "created_by_id": "123",
			  "git_branch": "%v"
			}
		  }`,
		DBT_GROUP_ID,
		DBT_FOLDER_PATH,
		DBT_TARGET_NAME,
		DBT_GIT_REMOTE_URL,
		DBT_DEFAULT_SCHEMA,
		DBT_GROUP_ID,
		DBT_GIT_BRANCH,
	)
}

func assertDbtProjectRequest(t *testing.T, request map[string]interface{}) {
	assertKey(t, "group_id", request, DBT_GROUP_ID)
	assertKey(t, "dbt_version", request, DBT_VERSION)
	assertKey(t, "git_remote_url", request, DBT_GIT_REMOTE_URL)
	assertKey(t, "git_branch", request, DBT_GIT_BRANCH)
	assertKey(t, "default_schema", request, DBT_DEFAULT_SCHEMA)
	assertKey(t, "folder_path", request, DBT_FOLDER_PATH)
	assertKey(t, "target_name", request, DBT_TARGET_NAME)
	assertKey(t, "threads", request, float64(DBT_THREADS))
}

func assertDbtProjectResponse(t *testing.T, response fivetran.DbtProjectCreateResponse) {
	assertEqual(t, response.Code, "Created")
	assertNotEmpty(t, response.Message)

	assertEqual(t, response.Data.ID, DBT_GROUP_ID)
	assertEqual(t, response.Data.FolderPath, DBT_FOLDER_PATH)
	assertEqual(t, response.Data.TargetName, DBT_TARGET_NAME)
	assertEqual(t, response.Data.GitRemoteUrl, DBT_GIT_REMOTE_URL)
	assertEqual(t, response.Data.DefaultSchema, DBT_DEFAULT_SCHEMA)
	assertEqual(t, response.Data.GroupID, DBT_GROUP_ID)
	assertEqual(t, response.Data.GitBranch, DBT_GIT_BRANCH)
}
