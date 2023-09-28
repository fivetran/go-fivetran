package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/dbt"
)

type DbtTransformationModifyService struct {
	c                   *Client
	dbtTransformationId *string
	schedule            *dbt.DbtTransformationSchedule
	runTests            *bool
	paused              *bool
}

type dbtTransformationModifyRequest struct {
	Schedule any   `json:"schedule,omitempty"`
	RunTests *bool `json:"run_tests,omitempty"`
	Paused   *bool `json:"paused,omitempty"`
}

func (c *Client) NewDbtTransformationModifyService() *DbtTransformationModifyService {
	return &DbtTransformationModifyService{c: c}
}

func (s *DbtTransformationModifyService) request() *dbtTransformationModifyRequest {
	var schedule interface{}

	if s.schedule != nil {
		schedule = s.schedule.Request()
	}

	return &dbtTransformationModifyRequest{
		Schedule: schedule,
		RunTests: s.runTests,
		Paused:   s.paused,
	}
}

func (s *DbtTransformationModifyService) DbtTransformationId(value string) *DbtTransformationModifyService {
	s.dbtTransformationId = &value
	return s
}

func (s *DbtTransformationModifyService) Schedule(value *dbt.DbtTransformationSchedule) *DbtTransformationModifyService {
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

func (s *DbtTransformationModifyService) Do(ctx context.Context) (dbt.DbtTransformationResponse, error) {
	var response dbt.DbtTransformationResponse

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
