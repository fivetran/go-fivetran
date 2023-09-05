package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

type DbtModelsListService struct {
	c         *Client
	projectId *string
	limit     *int
	cursor    *string
}

type DbtModelsListResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Items []struct {
			ID        string `json:"id"`
			ModelName string `json:"model_name"`
			Scheduled bool   `json:"scheduled"`
		} `json:"items"`
		NextCursor string `json:"next_cursor"`
	} `json:"data"`
}

func (c *Client) NewDbtModelsList() *DbtModelsListService {
	return &DbtModelsListService{c: c}
}

func (s *DbtModelsListService) ProjectId(value string) *DbtModelsListService {
	s.projectId = &value
	return s
}

func (s *DbtModelsListService) Limit(value int) *DbtModelsListService {
	s.limit = &value
	return s
}

func (s *DbtModelsListService) Cursor(value string) *DbtModelsListService {
	s.cursor = &value
	return s
}

func (s *DbtModelsListService) Do(ctx context.Context) (DbtModelsListResponse, error) {
	var response DbtModelsListResponse
	url := fmt.Sprintf("%v/dbt/models", s.c.baseURL)
	expectedStatus := 200

	headers := s.c.commonHeaders()

	queries := make(map[string]string)
	if s.cursor != nil {
		queries["cursor"] = *s.cursor
	}
	if s.limit != nil {
		queries["limit"] = fmt.Sprint(*s.limit)
	}
	if s.projectId != nil {
		queries["project_id"] = *s.projectId
	}

	r := request{
		method:           "GET",
		url:              url,
		body:             nil,
		queries:          queries,
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
