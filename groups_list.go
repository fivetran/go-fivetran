package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type GroupsListService struct {
	c      *Client
	limit  int
	cursor string
}

type GroupsList struct {
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

func (c *Client) NewGroupsListService() *GroupsListService {
	return &GroupsListService{c: c}
}

func (s *GroupsListService) Limit(limit int) *GroupsListService {
	s.limit = limit
	return s
}

func (s *GroupsListService) Cursor(cursor string) *GroupsListService {
	s.cursor = cursor
	return s
}

func (s *GroupsListService) Do(ctx context.Context) (GroupsList, error) {
	url := fmt.Sprintf("%v/groups", s.c.baseURL)
	expectedStatus := 200
	headers := make(map[string]string)
	queries := make(map[string]string)

	headers["Authorization"] = s.c.authorization

	if s.cursor != "" {
		queries["cursor"] = s.cursor
	}

	if s.limit != 0 {
		queries["limit"] = fmt.Sprint(s.limit)
	}

	r := Request{
		method:  "GET",
		url:     url,
		body:    nil,
		queries: queries,
		headers: headers,
	}

	respBody, respStatus, err := httpRequest(r, ctx)
	if err != nil {
		return GroupsList{}, err
	}

	var groupsList GroupsList
	if err := json.Unmarshal(respBody, &groupsList); err != nil {
		return GroupsList{}, err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected %v", respStatus, expectedStatus)
		return groupsList, err
	}

	return groupsList, nil
}
