package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
)

type DbtTransformationDeleteService struct {
	c                *Client
	transformationId *string
}

func (c *Client) NewDbtTransformationDeleteService() *DbtTransformationDeleteService {
	return &DbtTransformationDeleteService{c: c}
}

func (s *DbtTransformationDeleteService) TransformationId(value string) *DbtTransformationDeleteService {
	s.transformationId = &value
	return s
}

func (s *DbtTransformationDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse

	if s.transformationId == nil {
		return response, fmt.Errorf("missing required transformation id")
	}

	url := fmt.Sprintf("%v/dbt/transformations/%v", s.c.baseURL, *s.transformationId)
	expectedStatus := 200

	headers := s.c.commonHeaders()

	r := request{
		method:           "DELETE",
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
