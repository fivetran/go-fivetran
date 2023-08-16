package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

type DbtTransformationCreateService struct {
	c          *Client
	dbtModelId *string
	schedule   *DbtTransformationSchedule
	runTests   *bool
	projectId  *string
}

type dbtTransformationCreateRequest struct {
	DbtModelId *string                           `json:"dbt_model_id,omitempty"`
	Schedule   *dbtTransformationScheduleRequest `json:"schedule,omitempty"`
	RunTests   *bool                             `json:"run_tests,omitempty"`
	ProjectId  *string                           `json:"project_id,omitempty"`
}

type DbtTransformationCreateResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID              string                            `json:"id"`
		Status          string                            `json:"status"`
		Schedule        dbtTransformationScheduleResponse `json:"schedule"`
		LastRun         string                            `json:"last_run"`
		OutputModelName string                            `json:"output_model_name"`
		DbtProjectId    string                            `json:"dbt_project_id"`
		DbtModelId      string                            `json:"dbt_model_id"`
		NextRun         string                            `json:"next_run"`
		CreatedAt       string                            `json:"created_at"`
		ModelIds        []string                          `json:"model_ids"`
		ConnectorIds    []string                          `json:"connector_ids"`
	} `json:"data"`
}

func (c *Client) NewDbtTransformationCreateService() *DbtTransformationCreateService {
	return &DbtTransformationCreateService{c: c}
}

func (s *DbtTransformationCreateService) request() *dbtTransformationCreateRequest {
	var schedule *dbtTransformationScheduleRequest

	if s.schedule != nil {
		schedule = s.schedule.request()
	}

	return &dbtTransformationCreateRequest{
		DbtModelId: s.dbtModelId,
		Schedule:   schedule,
		RunTests:   s.runTests,
		ProjectId:  s.projectId,
	}
}

func (s *DbtTransformationCreateService) DbtModelId(value string) *DbtTransformationCreateService {
	s.dbtModelId = &value
	return s
}

func (s *DbtTransformationCreateService) Schedule(value *DbtTransformationSchedule) *DbtTransformationCreateService {
	s.schedule = value
	return s
}

func (s *DbtTransformationCreateService) RunTests(value bool) *DbtTransformationCreateService {
	s.runTests = &value
	return s
}

func (s *DbtTransformationCreateService) ProjectId(value string) *DbtTransformationCreateService {
	s.projectId = &value
	return s
}

func (s *DbtTransformationCreateService) Do(ctx context.Context) (DbtTransformationCreateResponse, error) {
	var response DbtTransformationCreateResponse
	url := fmt.Sprintf("%v/dbt/transformations", s.c.baseURL)
	expectedStatus := 201

	headers := s.c.commonHeaders()
	headers["Content-Type"] = "application/json"

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return response, err
	}

	r := request{
		method:           "POST",
		url:              url,
		body:             reqBody,
		queries:          nil,
		headers:          headers,
		client:           s.c.httpClient,
		handleRateLimits: s.c.handleRateLimits,
		maxRetryAttempts: s.c.maxRetryAttempts,
	}

	respBody, respStatus, err := r.httpRequest(ctx)
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return response, err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected: %v", respStatus, expectedStatus)
		return response, err
	}

	return response, nil
}
