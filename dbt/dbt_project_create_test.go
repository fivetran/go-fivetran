package dbt_test

import (
    "context"
    "fmt"
    "net/http"
    "testing"
    "strconv"

	"github.com/fivetran/go-fivetran"
    
    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewDbtProjectCreateFullMappingMock(t *testing.T) {
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
			"code": "Created",
			"message": "DBT Project created succesfully",
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
				"status": "%v",
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

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/dbt/projects").
		ThenCall(

			func(req *http.Request) (*http.Response, error) {
				body := testutils.RequestBodyToJson(t, req)
				testutils.AssertKey(t, "group_id", body, groupID)
				testutils.AssertKey(t, "dbt_version", body, dbtVersion)
				testutils.AssertKey(t, "default_schema", body, defaultSchema)
				testutils.AssertKey(t, "target_name", body, targetName)
				testutils.AssertKey(t, "threads", body, float64(threads))
				testutils.AssertKey(t, "type", body, projectType)
				testutils.AssertHasKey(t, body, "project_config")

				configRequest := body["project_config"].(map[string]interface{})

				testutils.AssertKey(t, "git_remote_url", configRequest, gitRemoteURL)
				testutils.AssertKey(t, "git_branch", configRequest, gitBranch)
				testutils.AssertKey(t, "folder_path", configRequest, folderPath)

				response := mock.NewResponse(req, http.StatusCreated, projectResponse)
				return response, nil
			})

	// act

	response, err := ftClient.NewDbtProjectCreate().
		GroupID(groupID).
		DbtVersion(dbtVersion).
		DefaultSchema(defaultSchema).
		TargetName(targetName).
		Threads(threads).
		EnvironmentVars([]string{environmentVar}).
		Type("GIT").
		ProjectConfig(fivetran.NewDbtProjectConfig().
			FolderPath(folderPath).
			GitRemoteUrl(gitRemoteURL).
			GitBranch(gitBranch)).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	// assert
	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)

	testutils.AssertEqual(t, response.Data.ID, dbtProjectID)
	testutils.AssertEqual(t, response.Data.TargetName, targetName)
	testutils.AssertEqual(t, response.Data.DefaultSchema, defaultSchema)
	testutils.AssertEqual(t, response.Data.GroupId, groupID)
	testutils.AssertEqual(t, response.Data.CreatedAt, createdAt)
	testutils.AssertEqual(t, response.Data.CreatedById, createdByID)
	testutils.AssertEqual(t, response.Data.PublicKey, publicKey)
	testutils.AssertEqual(t, response.Data.EnvironmentVars[0], environmentVar)
	testutils.AssertEqual(t, response.Data.Type, projectType)

	testutils.AssertEqual(t, response.Data.ProjectConfig.FolderPath, folderPath)
	testutils.AssertEqual(t, response.Data.ProjectConfig.GitRemoteUrl, gitRemoteURL)
	testutils.AssertEqual(t, response.Data.ProjectConfig.GitBranch, gitBranch)

	testutils.AssertEqual(t, response.Data.Status, status)
	testutils.AssertEqual(t, response.Data.Errors[0], errors)
}
