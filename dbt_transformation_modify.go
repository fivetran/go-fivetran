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
	paused              *bool
}

type dbtTransformationModifyRequest struct {
	Schedule *dbtTransformationScheduleRequest `json:"schedule,omitempty"`
	RunTests *bool                             `json:"run_tests,omitempty"`
	Paused   *bool                             `json:"paused,omitempty"`
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

func (s *DbtTransformationModifyService) Paused(value bool) *DbtTransformationModifyService {
	s.paused = &value
	return s
}

func (s *DbtTransformationModifyService) Do(ctx context.Context) (DbtTransformationResponse, error) {
	var response DbtTransformationResponse

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
