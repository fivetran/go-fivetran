package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/common"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestDbtTransformationDeleteService(t *testing.T) {
	//arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodDelete, "/v1/dbt/transformations/transformation_id").
		ThenCall(
			func(req *http.Request) (*http.Response, error) {
				response := mock.NewResponse(req, http.StatusOK, prepareDbtTransformationDeleteResponse("Success", "Dbt transformation has been deleted"))
				return response, nil
			})

	service := ftClient.NewDbtTransformationDeleteService().TransformationId("transformation_id")

	// act
	response, err := service.Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	// assert
	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)

	assertDbtTransformationDeleteResponse(t, response, "Success", "Dbt transformation has been deleted")
}

func TestRespStatusDbtTransformationDeleteService(t *testing.T) {
	//arrange
	ftClient, mockClient := CreateTestClient()
	mockClient.When(http.MethodDelete, "/v1/dbt/transformations/transformation_id").
		ThenCall(
			func(req *http.Request) (*http.Response, error) {
				response := mock.NewResponse(req, http.StatusNotFound, prepareDbtTransformationDeleteResponse("NotFound", "Cannot find dbt transformation with id 'transformation_id'"))
				return response, nil
			})

	service := ftClient.NewDbtTransformationDeleteService().TransformationId("transformation_id")

	// act
	response, err := service.Do(context.Background())

	// assert
	if err != nil {
		assertDbtTransformationDeleteResponse(t, response, "NotFound", "Cannot find dbt transformation with id 'transformation_id'")
	} else {
		t.Logf("%+v\n", response)
		t.Error(err)
	}
}

func prepareDbtTransformationDeleteResponse(code string, message string) string {
	var s = "{\"code\": \"" + code + "\" , \"message\": \"" + message + "\"}"
	return s
}

func assertDbtTransformationDeleteResponse(t *testing.T, response common.CommonResponse, expectCode string, expectMessage string) {
	assertEqual(t, response.Code, expectCode)
	if response.Message != expectMessage {
		t.Errorf("expected message '%s', got '%s'", expectMessage, response.Message)
	}
}
