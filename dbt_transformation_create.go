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
	paused     *bool
}

type dbtTransformationCreateRequest struct {
	DbtModelId *string                           `json:"dbt_model_id,omitempty"`
	Schedule   *dbtTransformationScheduleRequest `json:"schedule,omitempty"`
	RunTests   *bool                             `json:"run_tests,omitempty"`
	Paused     *bool                             `json:"paused,omitempty"`
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
		Paused:     s.paused,
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

func (s *DbtTransformationCreateService) Paused(value bool) *DbtTransformationCreateService {
	s.paused = &value
	return s
}

func (s *DbtTransformationCreateService) Do(ctx context.Context) (DbtTransformationResponse, error) {
	var response DbtTransformationResponse
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
