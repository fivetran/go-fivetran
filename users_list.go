package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type UsersListService struct {
	c      *Client
	limit  *int
	cursor *string
}

type UsersListResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Items []struct {
			ID         string    `json:"id"`
			Email      string    `json:"email"`
			GivenName  string    `json:"given_name"`
			FamilyName string    `json:"family_name"`
			Verified   bool      `json:"verified"`
			Invited    bool      `json:"invited"`
			Picture    string    `json:"picture"`
			Phone      string    `json:"phone"`
			LoggedInAt time.Time `json:"logged_in_at"`
			CreatedAt  time.Time `json:"created_at"`
		} `json:"items"`
		NextCursor string `json:"next_cursor"`
	} `json:"data"`
}

func (c *Client) NewUsersList() *UsersListService {
	return &UsersListService{c: c}
}

func (s *UsersListService) Limit(value int) *UsersListService {
	s.limit = &value
	return s
}

func (s *UsersListService) Cursor(value string) *UsersListService {
	s.cursor = &value
	return s
}

func (s *UsersListService) Do(ctx context.Context) (UsersListResponse, error) {
	var response UsersListResponse
	url := fmt.Sprintf("%v/users", s.c.baseURL)
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

	r := request{
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
