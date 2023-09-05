package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

type DbtModelDetailsService struct {
	c       *Client
	modelId *string
}

type DbtModelDetailsResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID        string `json:"id"`
		ModelName string `json:"model_name"`
		Scheduled bool   `json:"scheduled"`
	} `json:"data"`
}

func (c *Client) NewDbtModelDetails() *DbtModelDetailsService {
	return &DbtModelDetailsService{c: c}
}

func (s *DbtModelDetailsService) ModelId(value string) *DbtModelDetailsService {
	s.modelId = &value
	return s
}

func (s *DbtModelDetailsService) Do(ctx context.Context) (DbtModelDetailsResponse, error) {
	var response DbtModelDetailsResponse

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
