package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

const (
	dbtProjectID   = "dbt_project_id"
	groupID        = "group_id"
	createdAt      = "2023-01-01T00:00:00.743708Z"
	createdByID    = "created_by_id"
	publicKey      = "public_key"
	gitRemoteURL   = "git_remote_url"
	gitBranch      = "git_branch"
	defaultSchema  = "default_schema"
	folderPath     = "folder_path"
	targetName     = "target_name"
	responseStatus = 200
	responseJSON   = `{
		"code": "Success",
		"data": {
			"id": "` + DBT_GROUP_ID + `",
			"group_id": "` + DBT_GROUP_ID + `",
			"created_at": "2023-08-24T14:15:22Z",
			"created_by_id": "123",
			"public_key": "public_key",
			"git_remote_url": "` + DBT_GIT_REMOTE_URL + `",
			"git_branch": "` + DBT_GIT_BRANCH + `",
			"default_schema": "` + DBT_DEFAULT_SCHEMA + `",
			"folder_path": "` + DBT_FOLDER_PATH + `",
			"target_name": "` + DBT_TARGET_NAME + `"
		}
	}`
)

func TestDbtProjectDetailsService(t *testing.T) {
	// arrange
	var expectedResponse fivetran.DbtProjectDetailsResponse
	err := json.Unmarshal([]byte(responseJSON), &expectedResponse)
	if err != nil {
		t.Fatalf("failed to unmarshal example response: %v", err)
	}

	client, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodGet, fmt.Sprintf("/v1/dbt/projects/%v", dbtProjectID)).
		ThenCall(

			func(req *http.Request) (*http.Response, error) {
				response := mock.NewResponse(req, http.StatusOK, prepareDbtDetailsProjectResponse())
				return response, nil
			})

	service := client.NewDbtDetails()

	service.DbtID(dbtProjectID)

	ctx := context.Background()

	// act
	response, err := service.Do(ctx)
	if err != nil {
		t.Fatalf("failed to get Dbt project details: %v", err)
	}

	// assert
	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)

	assertDbtProjectDetailsResponse(t, response, expectedResponse)
}

func prepareDbtDetailsProjectResponse() string {
	return fmt.Sprintf(
		`{
			"code": "Success",
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

func assertDbtProjectDetailsResponse(t *testing.T, actual, expected fivetran.DbtProjectDetailsResponse) {
	// Check individual fields of the response
	assertEqual(t, actual.Code, expected.Code)
	assertEqual(t, actual.Data.ID, expected.Data.ID)
	assertEqual(t, actual.Data.GroupID, expected.Data.GroupID)
	assertEqual(t, expected.Data.CreatedAt, actual.Data.CreatedAt)
	assertEqual(t, expected.Data.DefaultSchema, actual.Data.DefaultSchema)
	assertEqual(t, expected.Data.TargetName, actual.Data.TargetName)
	assertEqual(t, expected.Data.FolderPath, actual.Data.FolderPath)
	assertEqual(t, expected.Data.GitRemoteUrl, actual.Data.GitRemoteUrl)
}
