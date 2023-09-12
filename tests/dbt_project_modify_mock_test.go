package tests

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestNewDbtProjectUpdateFullMappingMock(t *testing.T) {
	// arrange
	dbtProjectID := "dbt_project_id"
	dbtVersion := "dbt_version"
	groupID := "group_id"
	createdAt := "2023-01-01T00:00:00.743708Z"
	createdByID := "created_by_id"
	publicKey := "public_key"
	gitRemoteURL := "git_remote_url"
	gitBranch := "git_branch"
	defaultSchema := "default_schema"
	folderPath := "folder_path"
	targetName := "target_name"
	environmentVar := "DBT_VARIABLE_1=VALUE"
	threads := 1
	projectType := "GIT"
	status := "READY"
	errors := "error"

	projectResponse := fmt.Sprintf(
		`{
			"code": "Updated",
			"message": "DBT Project updated succesfully",
			"data": {
				"id": "%v",
				"group_id": "%v",
				"dbt_version": "%v",
				"created_at": "%v",
				"created_by_id": "%v",
				"public_key": "%v",
				"environment_vars": ["%v"],
				"default_schema": "%v",
				"target_name": "%v",
				"threads": %v,
				"type": "%v",
				"project_config": {
					"git_remote_url": "%v",
					"git_branch": "%v",
					"folder_path": "%v"
				},
				"status":"%v",
				"errors":["%v"]
			}
		}`,
		dbtProjectID,
		groupID,
		dbtVersion,
		createdAt,
		createdByID,
		publicKey,
		environmentVar,
		defaultSchema,
		targetName,
		strconv.Itoa(threads),
		projectType,
		gitRemoteURL,
		gitBranch,
		folderPath,
		status,
		errors,
	)

	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/dbt/projects/"+dbtProjectID).
		ThenCall(

			func(req *http.Request) (*http.Response, error) {
				body := requestBodyToJson(t, req)
				assertKey(t, "dbt_version", body, dbtVersion)
				assertKey(t, "target_name", body, targetName)
				assertKey(t, "threads", body, float64(threads))
				assertHasKey(t, body, "project_config")

				configRequest := body["project_config"].(map[string]interface{})

				assertHasNoKey(t, configRequest, "git_remote_url")
				assertKey(t, "git_branch", configRequest, gitBranch)
				assertKey(t, "folder_path", configRequest, folderPath)

				response := mock.NewResponse(req, http.StatusOK, projectResponse)
				return response, nil
			})

	// act

	response, err := ftClient.NewDbtProjectModify().
		DbtProjectID(dbtProjectID).
		DbtVersion(dbtVersion).
		TargetName(targetName).
		Threads(threads).
		EnvironmentVars([]string{environmentVar}).
		ProjectConfig(fivetran.NewDbtProjectConfig().
			GitRemoteUrl(gitRemoteURL). // This value should not be passed in request
			FolderPath(folderPath).
			GitBranch(gitBranch)).
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

	assertEqual(t, response.Data.ID, dbtProjectID)
	assertEqual(t, response.Data.TargetName, targetName)
	assertEqual(t, response.Data.DefaultSchema, defaultSchema)
	assertEqual(t, response.Data.GroupID, groupID)
	assertEqual(t, response.Data.CreatedAt, createdAt)
	assertEqual(t, response.Data.CreatedById, createdByID)
	assertEqual(t, response.Data.PublicKey, publicKey)
	assertEqual(t, response.Data.EnvironmentVars[0], environmentVar)
	assertEqual(t, response.Data.Type, projectType)

	assertEqual(t, response.Data.ProjectConfig.FolderPath, folderPath)
	assertEqual(t, response.Data.ProjectConfig.GitRemoteUrl, gitRemoteURL)
	assertEqual(t, response.Data.ProjectConfig.GitBranch, gitBranch)

	assertEqual(t, response.Data.Status, status)
	assertEqual(t, response.Data.Errors[0], errors)
}
