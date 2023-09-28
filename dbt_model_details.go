package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/dbt"
)

type DbtModelDetailsService struct {
	c       *Client
	modelId *string
}

func (c *Client) NewDbtModelDetails() *DbtModelDetailsService {
	return &DbtModelDetailsService{c: c}
}

func (s *DbtModelDetailsService) ModelId(value string) *DbtModelDetailsService {
	s.modelId = &value
	return s
}

func (s *DbtModelDetailsService) Do(ctx context.Context) (dbt.DbtModelDetailsResponse, error) {
	var response dbt.DbtModelDetailsResponse

	if s.modelId == nil {
		return response, fmt.Errorf("missing required ModelId")
	}

	url := fmt.Sprintf("%v/dbt/models/%v", s.c.baseURL, *s.modelId)
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
