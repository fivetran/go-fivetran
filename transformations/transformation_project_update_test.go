package transformations_test

import (
    "context"
    "net/http"
    "testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/transformations"
    
    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestTransformationProjectUpdateService(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/transformation-projects/project_id").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertTransformationProjectUpdateRequest(t, body)
			response := mock.NewResponse(req, http.StatusOK, prepareTransformationProjectUpdateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewTransformationProjectUpdate().
		ProjectId("project_id").
		RunTests(true).
		ProjectConfig(prepareTransformationProjectUpdateConfig()).
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

	assertTransformationProjectUpdateResponse(t, response)
}

func TestTransformationProjectCustomUpdateService(t *testing.T) {
	// arrange
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/transformation-projects/project_id").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertTransformationProjectUpdateCustomRequest(t, body)
			response := mock.NewResponse(req, http.StatusOK, prepareTransformationProjectUpdateMergedResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewTransformationProjectUpdate().
		ProjectId("project_id").
		RunTests(true).
		ProjectConfigCustom(prepareTransformationProjectCustomConfig()).
		DoCustom(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	// assert
	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)

	assertTransformationProjectUpdateCustomResponse(t, response)
}

func TestTransformationProjectCustomMergedUpdateService(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/transformation-projects/project_id").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertTransformationProjectUpdateCustomMergedRequest(t, body)
			response := mock.NewResponse(req, http.StatusOK, prepareTransformationProjectUpdateMergedResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewTransformationProjectUpdate().
		ProjectId("project_id").
		RunTests(true).
		ProjectConfig(prepareTransformationProjectUpdateConfig()).
		ProjectConfigCustom(prepareTransformationProjectCustomMergedConfig()).
		DoCustomMerged(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	// assert
	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)

	assertTransformationProjectUpdateCustomMergedResponse(t, response)
}

func prepareTransformationProjectUpdateResponse() string {
	return  `{
  "code": "Success",
  "message": "Operation performed.",
  "data": {
    "id": "project_id",
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
      "folder_path": "folder_path"
    }
  }
}`
}

func prepareTransformationProjectUpdateMergedResponse() string {
	return `{
  "code": "Success",
  "message": "Operation performed.",
  "data": {
    "id": "project_id",
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
      "folder_path": "folder_path",
      "fake_field": "unmapped-value"
    }
  }
}`
}

func prepareTransformationProjectUpdateConfig() *transformations.TransformationProjectConfig {
	config := fivetran.NewTransformationProjectConfig()
	config.FolderPath("folder_path")

	return config
}

func prepareTransformationProjectCustomConfig() *map[string]interface{} {
	config := make(map[string]interface{})

	config["fake_field"] = "unmapped-value"

	return &config
}

// assert Requests
func assertTransformationProjectUpdateRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "run_tests", request, true)

	config, ok := request["project_config"].(map[string]interface{})
	testutils.AssertEqual(t, ok, true)

	testutils.AssertKey(t, "folder_path", config, "folder_path")
}

func assertTransformationProjectUpdateCustomRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "run_tests", request, true)

	config, ok := request["project_config"].(map[string]interface{})

	testutils.AssertEqual(t, ok, true)

	testutils.AssertKey(t, "fake_field", config, "unmapped-value")
}

func assertTransformationProjectUpdateCustomMergedRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "run_tests", request, true)

	config, ok := request["project_config"].(map[string]interface{})

	testutils.AssertEqual(t, ok, true)

	testutils.AssertKey(t, "folder_path", config, "folder_path")
	testutils.AssertKey(t, "fake_field", config, "unmapped-value")
}

// assert Response
func assertTransformationProjectUpdateResponse(t *testing.T, response transformations.TransformationProjectResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.Id, "project_id")

	testutils.AssertEqual(t, response.Data.ProjectConfig.FolderPath, "folder_path")
}

func assertTransformationProjectUpdateCustomResponse(t *testing.T, response transformations.TransformationProjectCustomResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.Id, "project_id")

	testutils.AssertKey(t, "fake_field", response.Data.ProjectConfig, "unmapped-value")
}

func assertTransformationProjectUpdateCustomMergedResponse(t *testing.T, response transformations.TransformationProjectCustomMergedResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.Id, "project_id")

	testutils.AssertEqual(t, response.Data.ProjectConfig.FolderPath, "folder_path")
	testutils.AssertKey(t, "fake_field", response.Data.ProjectConfigCustom, "unmapped-value")
}
