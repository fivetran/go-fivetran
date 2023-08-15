package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

type DbtTransformationModifyService struct {
	c                   *Client
	dbtTransformationId *string
	schedule            *DbtTransformationSchedule
	runTests            *bool
}

type dbtTransformationModifyRequest struct {
	Schedule *dbtTransformationScheduleRequest `json:"schedule,omitempty"`
	RunTests *bool                             `json:"run_tests,omitempty"`
}

type DbtTransformationModifyResponse struct {
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
		RunTests        string                            `json:"run_tests"`
		ModelIds        []string                          `json:"model_ids"`
		ConnectorIds    []string                          `json:"connector_ids"`
	} `json:"data"`
}

func (c *Client) NewDbtTransformationModifyService() *DbtTransformationModifyService {
	return &DbtTransformationModifyService{c: c}
}

func (s *DbtTransformationModifyService) request() *dbtTransformationModifyRequest {
	var schedule *dbtTransformationScheduleRequest

	if s.schedule != nil {
		schedule = s.schedule.request()
	}

	return &dbtTransformationModifyRequest{
		Schedule: schedule,
		RunTests: s.runTests,
	}
}

func (s *DbtTransformationModifyService) DbtTransformationId(value string) *DbtTransformationModifyService {
	s.dbtTransformationId = &value
	return s
}

func (s *DbtTransformationModifyService) Schedule(value *DbtTransformationSchedule) *DbtTransformationModifyService {
	s.schedule = value
	return s
}

func (s *DbtTransformationModifyService) RunTests(value bool) *DbtTransformationModifyService {
	s.runTests = &value
	return s
}

func (s *DbtTransformationModifyService) Do(ctx context.Context) (DbtTransformationModifyResponse, error) {
	var response DbtTransformationModifyResponse

	if s.dbtTransformationId == nil {
		return response, fmt.Errorf("missing required dbt transformation ID")
	}

	url := fmt.Sprintf("%v/dbt/transformations/%v", s.c.baseURL, *s.dbtTransformationId)
	expectedStatus := 200

	headers := s.c.commonHeaders()
	headers["Content-Type"] = "application/json"

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return response, err
	}

	r := request{
		method:           "PATCH",
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
