package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/dbt"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type DbtProjectsListService struct {
	c      *Client
	limit  *int
	cursor *string
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

func (s *DbtProjectsListService) Do(ctx context.Context) (dbt.DbtProjectsListResponse, error) {
	var response dbt.DbtProjectsListResponse
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

	r := httputils.Request{
		Method:           "GET",
		Url:              url,
		Body:             nil,
		Queries:          queries,
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
