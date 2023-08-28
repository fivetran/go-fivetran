package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

type DbtTransformationDetailsService struct {
	c                *Client
	transformationId *string
}

type DbtTransformationResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID              string                            `json:"id"`
		Status          string                            `json:"status"`
		Schedule        DbtTransformationScheduleResponse `json:"schedule"`
		LastRun         string                            `json:"last_run"`
		OutputModelName string                            `json:"output_model_name"`
		DbtProjectId    string                            `json:"dbt_project_id"`
		DbtModelId      string                            `json:"dbt_model_id"`
		NextRun         string                            `json:"next_run"`
		CreatedAt       string                            `json:"created_at"`
		ModelIds        []string                          `json:"model_ids"`
		ConnectorIds    []string                          `json:"connector_ids"`
		RunTests        bool                              `json:"run_tests"`
		Paused          bool                              `json:"paused"`
	} `json:"data"`
}

func (c *Client) NewDbtTransformationDetailsService() *DbtTransformationDetailsService {
	return &DbtTransformationDetailsService{c: c}
}

func (s *DbtTransformationDetailsService) TransformationId(value string) *DbtTransformationDetailsService {
	s.transformationId = &value
	return s
}

func (s *DbtTransformationDetailsService) Do(ctx context.Context) (DbtTransformationResponse, error) {
	var response DbtTransformationResponse

	if s.transformationId == nil {
		return response, fmt.Errorf("missing required transformation id")
	}

	url := fmt.Sprintf("%v/dbt/transformations/%v", s.c.baseURL, *s.transformationId)
	expectedStatus := 200

	headers := s.c.commonHeaders()

	r := request{
		method:           "GET",
		url:              url,
		body:             nil,
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
