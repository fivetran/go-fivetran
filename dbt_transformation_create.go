package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/dbt"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type DbtTransformationCreateService struct {
	c          *Client
	dbtModelId *string
	schedule   *dbt.DbtTransformationSchedule
	runTests   *bool
	paused     *bool
}

type dbtTransformationCreateRequest struct {
	DbtModelId *string `json:"dbt_model_id,omitempty"`
	Schedule   any     `json:"schedule,omitempty"`
	RunTests   *bool   `json:"run_tests,omitempty"`
	Paused     *bool   `json:"paused,omitempty"`
}

func (c *Client) NewDbtTransformationCreateService() *DbtTransformationCreateService {
	return &DbtTransformationCreateService{c: c}
}

func (s *DbtTransformationCreateService) request() *dbtTransformationCreateRequest {
	var schedule interface{}

	if s.schedule != nil {
		schedule = s.schedule.Request()
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

func (s *DbtTransformationCreateService) Schedule(value *dbt.DbtTransformationSchedule) *DbtTransformationCreateService {
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

func (s *DbtTransformationCreateService) Do(ctx context.Context) (dbt.DbtTransformationResponse, error) {
	var response dbt.DbtTransformationResponse
	url := fmt.Sprintf("%v/dbt/transformations", s.c.baseURL)
	expectedStatus := 201

	headers := s.c.commonHeaders()
	headers["Content-Type"] = "application/json"

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return response, err
	}

	r := httputils.Request{
		Method:           "POST",
		Url:              url,
		Body:             reqBody,
		Queries:          nil,
		Headers:          headers,
		Client:           s.c.httpClient,
		HandleRateLimits: s.c.handleRateLimits,
		MaxRetryAttempts: s.c.maxRetryAttempts,
	}

	respBody, respStatus, err := r.Do(ctx)
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
