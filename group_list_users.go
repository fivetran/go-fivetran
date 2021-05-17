package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type GroupListUsersService struct {
	c      *Client
	id     string
	limit  int
	cursor string
}

type GroupListUsers struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Items []struct {
			ID         string      `json:"id"`
			Email      string      `json:"email"`
			GivenName  string      `json:"given_name"`
			FamilyName string      `json:"family_name"`
			Verified   bool        `json:"verified"`
			Invited    bool        `json:"invited"`
			Picture    interface{} `json:"picture"`
			Phone      interface{} `json:"phone"`
			LoggedInAt time.Time   `json:"logged_in_at"`
			CreatedAt  time.Time   `json:"created_at"`
		} `json:"items"`
		NextCursor string `json:"next_cursor"`
	} `json:"data"`
}

func (c *Client) NewGroupListUsersService() *GroupListUsersService {
	return &GroupListUsersService{c: c}
}

func (s *GroupListUsersService) ID(id string) *GroupListUsersService {
	s.id = id
	return s
}

func (s *GroupListUsersService) Limit(limit int) *GroupListUsersService {
	s.limit = limit
	return s
}

func (s *GroupListUsersService) Cursor(cursor string) *GroupListUsersService {
	s.cursor = cursor
	return s
}

func (s *GroupListUsersService) Do(ctx context.Context) (GroupListUsers, error) {
	if s.id == "" {
		err := fmt.Errorf("missing required ID")
		return GroupListUsers{}, err
	}

	url := fmt.Sprintf("%v/groups/%v/users", s.c.baseURL, s.id)
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
		return GroupListUsers{}, err
	}

	var groupListUsers GroupListUsers
	if err := json.Unmarshal(respBody, &groupListUsers); err != nil {
		return GroupListUsers{}, err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected %v", respStatus, expectedStatus)
		return groupListUsers, err
	}

	return groupListUsers, nil
}
