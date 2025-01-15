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

func TestNewTransformationProjectCreateFullMappingMock(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/transformation-projects").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertTransformationProjectFullRequest(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareTransformationProjectResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewTransformationProjectCreate().
		GroupId("group_id").
		ProjectType("DBT_GIT").
		RunTests(true).
		ProjectConfig(prepareTransformationProjectConfig()).
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

	assertTransformationProjectResponse(t, response)
}

func TestNewTransformationProjectCustomMappingMock(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/transformation-projects").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertTransformationProjectCustomRequest(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareTransformationProjectResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewTransformationProjectCreate().
		GroupId("group_id").
		ProjectType("DBT_GIT").
		RunTests(true).
		ProjectConfigCustom(prepareTransformationProjectCustomMergedConfig()).
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

	assertTransformationProjectCustomResponse(t, response)
}

func TestNewTransformationProjectCustomMergedMappingMock(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/transformation-projects").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertTransformationProjectCustomMergedRequest(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareTransformationProjectMergedResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewTransformationProjectCreate().
		GroupId("group_id").
		ProjectType("DBT_GIT").
		RunTests(true).
		ProjectConfig(prepareTransformationProjectConfig()).
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
	assertTransformationProjectCustomMergedResponse(t, response)
}

func prepareTransformationProjectResponse() string {
	return `{
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
      "dbt_version": "1.0.0",
      "default_schema": "schema"
    }
  }
}`
}

func prepareTransformationProjectMergedResponse() string {
	return `{
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
      "dbt_version": "1.0.0",
      "default_schema": "schema",
      "fake_field": "unmapped-value"
    }
  }
}`
}

func prepareTransformationProjectConfig() *transformations.TransformationProjectConfig {
	config := fivetran.NewTransformationProjectConfig()
	config.DbtVersion("1.0.0")
	config.DefaultSchema("schema")

	return config
}

func prepareTransformationProjectCustomMergedConfig() *map[string]interface{} {
	config := make(map[string]interface{})

	config["fake_field"] = "unmapped-value"

	return &config
}

func assertTransformationProjectFullRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "type", request, "DBT_GIT")
	testutils.AssertKey(t, "group_id", request, "group_id")
	testutils.AssertKey(t, "run_tests", request, true)

	config, ok := request["project_config"].(map[string]interface{})
	testutils.AssertEqual(t, ok, true)

	testutils.AssertKey(t, "dbt_version", config, "1.0.0")
	testutils.AssertKey(t, "default_schema", config, "schema")
}

func assertTransformationProjectCustomRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "type", request, "DBT_GIT")
	testutils.AssertKey(t, "group_id", request, "group_id")
	testutils.AssertKey(t, "run_tests", request, true)

	config, ok := request["project_config"].(map[string]interface{})

	testutils.AssertEqual(t, ok, true)

	testutils.AssertKey(t, "fake_field", config, "unmapped-value")
}

func assertTransformationProjectCustomMergedRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "type", request, "DBT_GIT")
	testutils.AssertKey(t, "group_id", request, "group_id")
	testutils.AssertKey(t, "run_tests", request, true)

	config, ok := request["project_config"].(map[string]interface{})

	testutils.AssertEqual(t, ok, true)

	testutils.AssertKey(t, "dbt_version", config, "1.0.0")
	testutils.AssertKey(t, "default_schema", config, "schema")
	testutils.AssertKey(t, "fake_field", config, "unmapped-value")
}

func assertTransformationProjectResponse(t *testing.T, response transformations.TransformationProjectResponse) {
    testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertNotEmpty(t, response.Message)
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
    testutils.AssertEqual(t, response.Data.ProjectConfig.DbtVersion, "1.0.0")
    testutils.AssertEqual(t, response.Data.ProjectConfig.DefaultSchema, "schema")
}

func assertTransformationProjectCustomResponse(t *testing.T, response transformations.TransformationProjectCustomResponse) {
    testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertNotEmpty(t, response.Message)
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
	testutils.AssertEqual(t, response.Data.ProjectConfig["dbt_version"], "1.0.0")
	testutils.AssertEqual(t, response.Data.ProjectConfig["default_schema"], "schema")
}

func assertTransformationProjectCustomMergedResponse(t *testing.T, response transformations.TransformationProjectCustomMergedResponse) {
    testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertNotEmpty(t, response.Message)
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
    testutils.AssertEqual(t, response.Data.ProjectConfig.DbtVersion, "1.0.0")
    testutils.AssertEqual(t, response.Data.ProjectConfig.DefaultSchema, "schema")
	testutils.AssertKey(t, "fake_field", response.Data.ProjectConfigCustom, "unmapped-value")
}
