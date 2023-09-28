package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/dbt"
)

type DbtTransformationDetailsService struct {
	c                *Client
	transformationId *string
}

func (c *Client) NewDbtTransformationDetailsService() *DbtTransformationDetailsService {
	return &DbtTransformationDetailsService{c: c}
}

func (s *DbtTransformationDetailsService) TransformationId(value string) *DbtTransformationDetailsService {
	s.transformationId = &value
	return s
}

func (s *DbtTransformationDetailsService) Do(ctx context.Context) (dbt.DbtTransformationResponse, error) {
	var response dbt.DbtTransformationResponse

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
