package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

type DbtProjectsListService struct {
	c      *Client
	limit  *int
	cursor *string
}

type DbtProjectsListResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Items []struct {
			ID          string `json:"id"`
			GroupId     string `json:"group_id"`
			CreatedAt   string `json:"created_at"`
			CreatedById string `json:"created_by_id"`
		} `json:"items"`
		NextCursor string `json:"next_cursor"`
	} `json:"data"`
}

func (c *Client) NewDbtProjectsList() *DbtProjectsListService {
	return &DbtProjectsListService{c: c}
}

func (s *DbtProjectsListService) Limit(value int) *DbtProjectsListService {
	s.limit = &value
	return s
}

func (s *DbtProjectsListService) Cursor(value string) *DbtProjectsListService {
	s.cursor = &value
	return s
}

func (s *DbtProjectsListService) Do(ctx context.Context) (DbtProjectsListResponse, error) {
	var response DbtProjectsListResponse
	url := fmt.Sprintf("%v/dbt/projects", s.c.baseURL)
	expectedStatus := 200

	headers := s.c.commonHeaders()

	queries := make(map[string]string)
	if s.cursor != nil {
		queries["cursor"] = *s.cursor
	}
	if s.limit != nil {
		queries["limit"] = fmt.Sprint(*s.limit)
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
