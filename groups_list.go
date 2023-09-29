package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	"github.com/fivetran/go-fivetran/groups"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// GroupsListService implements the Group Management, List All Groups API.
// Ref. https://fivetran.com/docs/rest-api/groups#listallgroups
type GroupsListService struct {
	c      *Client
	limit  *int
	cursor *string
}

type GroupsListResponse struct {
	common.CommonResponse
	Data struct {
		Items      []groups.GroupItem `json:"items"`
		NextCursor string             `json:"next_cursor"`
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
