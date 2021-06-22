package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type groupListUsersService struct {
	c       *Client
	groupID *string
	limit   *int
	cursor  *string
}

type GroupListUsersResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Items []struct {
			ID         string `json:"id"`
			Email      string `json:"email"`
			GivenName  string `json:"given_name"`
			FamilyName string `json:"family_name"`
			Verified   bool   `json:"verified"`
			Invited    bool   `json:"invited"`
			// Picture    interface{} `json:"picture"`
			// Phone      interface{} `json:"phone"`
			Picture    string    `json:"picture"`
			Phone      string    `json:"phone"`
			LoggedInAt time.Time `json:"logged_in_at"`
			CreatedAt  time.Time `json:"created_at"`
		} `json:"items"`
		NextCursor string `json:"next_cursor"`
	} `json:"data"`
}

func (c *Client) NewGroupListUsers() *groupListUsersService {
	return &groupListUsersService{c: c}
}

func (s *groupListUsersService) GroupID(value string) *groupListUsersService {
	s.groupID = &value
	return s
}

func (s *groupListUsersService) Limit(value int) *groupListUsersService {
	s.limit = &value
	return s
}

func (s *groupListUsersService) Cursor(value string) *groupListUsersService {
	s.cursor = &value
	return s
}

func (s *groupListUsersService) Do(ctx context.Context) (GroupListUsersResponse, error) {
	var response GroupListUsersResponse

	if s.groupID == nil {
		return response, fmt.Errorf("missing required GroupID")
	}

	url := fmt.Sprintf("%v/groups/%v/users", s.c.baseURL, *s.groupID)
	expectedStatus := 200

	headers := make(map[string]string)
	headers["Authorization"] = s.c.authorization

	queries := make(map[string]string)
	if s.cursor != nil {
		queries["cursor"] = *s.cursor
	}
	if s.limit != nil {
		queries["limit"] = fmt.Sprint(*s.limit)
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
		return response, err
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return response, err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected %v", respStatus, expectedStatus)
		return response, err
	}

	return response, nil
}
