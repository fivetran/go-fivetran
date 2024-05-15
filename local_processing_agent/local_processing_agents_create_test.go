package localprocessingagent_test

import (
    "context"
    "net/http"
    "testing"
	"github.com/fivetran/go-fivetran/local_processing_agent"
    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewLocalProcessingAgentCreateMappingMock(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/local-processing-agents").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertLocalProcessingAgentCreateRequest(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareLocalProcessingAgentCreateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewLocalProcessingAgentCreate().
		GroupId("group_id").
		DisplayName("display_name").
		EnvType("DOCKER").
		AcceptTerms(true).
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

	assertLocalProcessingAgentCreateResponse(t, response)
}

func prepareLocalProcessingAgentCreateResponse() string {
	return `{
		    "code": "Success",
		    "message": "Success",
		    "data": {
       			"id": "id",
       			"display_name": "display_name",
       			"group_id": "group_id",
       			"registered_at": "1970-01-01T00:00:00.000000Z",
       			"files": {
          			"config_json": "config_json",
          			"auth_json": "auth_json",
          			"docker_compose_yaml": "docker_compose_yaml"
       			}
    		}
		}`
}

func assertLocalProcessingAgentCreateRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "group_id", request, "group_id")
	testutils.AssertKey(t, "display_name", request, "display_name")
	testutils.AssertKey(t, "env_type", request, "DOCKER")
	testutils.AssertKey(t, "accept_terms", request, true)
}

func assertLocalProcessingAgentCreateResponse(t *testing.T, response localprocessingagent.LocalProcessingAgentCreateResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertNotEmpty(t, response.Message)

	testutils.AssertEqual(t, response.Data.Id, "id")
	testutils.AssertEqual(t, response.Data.DisplayName, "display_name")
	testutils.AssertEqual(t, response.Data.GroupId, "group_id")
	testutils.AssertEqual(t, response.Data.RegisteredAt, "1970-01-01T00:00:00.000000Z")
	testutils.AssertEqual(t, response.Data.Files.ConfigJson, "config_json")
	testutils.AssertEqual(t, response.Data.Files.AuthJson, "auth_json")
	testutils.AssertEqual(t, response.Data.Files.DockerComposeYaml, "docker_compose_yaml")
}
