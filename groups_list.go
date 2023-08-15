package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// GroupsListService implements the Group Management, List All Groups API.
// Ref. https://fivetran.com/docs/rest-api/groups#listallgroups
type GroupsListService struct {
	c      *Client
	limit  *int
	cursor *string
}

type GroupsListResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Items []struct {
			ID        string    `json:"id"`
			Name      string    `json:"name"`
			CreatedAt time.Time `json:"created_at"`
		} `json:"items"`
		NextCursor string `json:"next_cursor"`
	} `json:"data"`
}

func (c *Client) NewGroupsList() *GroupsListService {
	return &GroupsListService{c: c}
}

func (s *GroupsListService) Limit(value int) *GroupsListService {
	s.limit = &value
	return s
}

func (s *GroupsListService) Cursor(value string) *GroupsListService {
	s.cursor = &value
	return s
}

func (s *GroupsListService) Do(ctx context.Context) (GroupsListResponse, error) {
	var response GroupsListResponse
	url := fmt.Sprintf("%v/groups", s.c.baseURL)
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
