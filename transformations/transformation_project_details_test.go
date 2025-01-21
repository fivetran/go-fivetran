package transformations_test

import (
    "context"
    "fmt"
    "net/http"
    "testing"

    "github.com/fivetran/go-fivetran/transformations"
    
    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestTransformationProjectDetailsServiceDo(t *testing.T) {
    // arrange

    ftClient, mockClient := testutils.CreateTestClient()
    handler := mockClient.When(http.MethodGet, "/v1/transformation-projects/project_id").
        ThenCall(func(req *http.Request) (*http.Response, error) {
            response := mock.NewResponse(req, http.StatusOK, prepareTransformationProjectDetailsResponse())
            return response, nil
        })

    // act
    response, err := ftClient.NewTransformationProjectDetails().
        ProjectId("project_id").
        Do(context.Background())

    // assert
    if err != nil {
        t.Error(err)
    }

    interactions := mockClient.Interactions()
    testutils.AssertEqual(t, len(interactions), 1)
    testutils.AssertEqual(t, interactions[0].Handler, handler)
    testutils.AssertEqual(t, handler.Interactions, 1)
    assertTransformationProjectDetailsResponse(t, response)
}

func prepareTransformationProjectDetailsResponse() string {
    return fmt.Sprintf(`{
  "code": "Success",
  "message": "Operation performed.",
  "data": {
    "id": "string",
    "type": "DBT_GIT",
    "status": "NOT_READY",
    "errors": [
      "string"
    ],
    "created_at": "2019-08-24T14:15:22Z",
    "group_id": "string",
    "setup_tests": [
      {
        "title": "Test Title",
        "status": "FAILED",
        "message": "Error message",
        "details": "Error details"
      }
    ],
    "created_by_id": "string",
    "project_config": {
      "dbt_version": "string",
      "default_schema": "string",
      "git_remote_url": "string",
      "folder_path": "string",
      "git_branch": "string",
      "threads": 0,
      "target_name": "string",
      "environment_vars": [
        "string"
      ],
      "public_key": "string"
    }
  }
}`)
}

func assertTransformationProjectDetailsResponse(t *testing.T, response transformations.TransformationProjectResponse) {
    testutils.AssertEqual(t, response.Code, "Success")
    testutils.AssertEqual(t, response.Data.Id, "string")
    testutils.AssertEqual(t, response.Data.ProjectType, "DBT_GIT")
    testutils.AssertEqual(t, response.Data.CreatedAt, "2019-08-24T14:15:22Z")
    testutils.AssertEqual(t, response.Data.GroupId, "string")
    testutils.AssertEqual(t, response.Data.CreatedById, "string")
    testutils.AssertEqual(t, response.Data.Status, "NOT_READY")
    testutils.AssertEqual(t, response.Data.SetupTests[0].Title, "Test Title")
    testutils.AssertEqual(t, response.Data.SetupTests[0].Status, "FAILED")
    testutils.AssertEqual(t, response.Data.SetupTests[0].Message, "Error message")
    testutils.AssertEqual(t, response.Data.SetupTests[0].Details, "Error details")
    testutils.AssertEqual(t, response.Data.Errors[0], "string")
    testutils.AssertEqual(t, response.Data.ProjectConfig.DbtVersion, "string")
    testutils.AssertEqual(t, response.Data.ProjectConfig.DefaultSchema, "string")
    testutils.AssertEqual(t, response.Data.ProjectConfig.GitRemoteUrl, "string")
    testutils.AssertEqual(t, response.Data.ProjectConfig.FolderPath, "string")
    testutils.AssertEqual(t, response.Data.ProjectConfig.GitBranch, "string")
    testutils.AssertEqual(t, response.Data.ProjectConfig.TargetName, "string")
    testutils.AssertEqual(t, response.Data.ProjectConfig.EnvironmentVars[0], "string")
    testutils.AssertEqual(t, response.Data.ProjectConfig.PublicKey, "string")
    testutils.AssertEqual(t, response.Data.ProjectConfig.Threads, 0)
}