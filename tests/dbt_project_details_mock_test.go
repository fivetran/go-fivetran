package tests

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestDbtProjectDetailsService(t *testing.T) {
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

	responseJSON := fmt.Sprintf(
		`{
			"code": "Success",
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

	client, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodGet, fmt.Sprintf("/v1/dbt/projects/%v", dbtProjectID)).
		ThenCall(
			func(req *http.Request) (*http.Response, error) {
				response := mock.NewResponse(req, http.StatusOK, responseJSON)
				return response, nil
			})

	service := client.NewDbtProjectDetails()

	service.DbtProjectID(dbtProjectID)

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

	// Check individual fields of the response
	assertEqual(t, response.Data.ID, dbtProjectID)
	assertEqual(t, response.Data.GroupId, groupID)
	assertEqual(t, response.Data.CreatedAt, createdAt)
	assertEqual(t, response.Data.DefaultSchema, defaultSchema)
	assertEqual(t, response.Data.TargetName, targetName)
	assertEqual(t, response.Data.Threads, threads)
	assertEqual(t, response.Data.EnvironmentVars[0], environmentVar)
	assertEqual(t, response.Data.CreatedById, createdByID)
	assertEqual(t, response.Data.DbtVersion, dbtVersion)
	assertEqual(t, response.Data.Type, projectType)
	assertEqual(t, response.Data.PublicKey, publicKey)
	assertEqual(t, response.Data.ProjectConfig.GitBranch, gitBranch)
	assertEqual(t, response.Data.ProjectConfig.FolderPath, folderPath)
	assertEqual(t, response.Data.ProjectConfig.GitRemoteUrl, gitRemoteURL)
	assertEqual(t, response.Data.Status, status)
	assertEqual(t, response.Data.Errors[0], errors)
}
