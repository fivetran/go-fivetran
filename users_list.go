package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
	"github.com/fivetran/go-fivetran/users"
)

// UsersListService implements the User Management, List All Users API.
// Ref. https://fivetran.com/docs/rest-api/users#listallusers
type UsersListService struct {
	c      *Client
	limit  *int
	cursor *string
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

func (s *UsersListService) Do(ctx context.Context) (users.UsersListResponse, error) {
	var response users.UsersListResponse
	url := fmt.Sprintf("%v/users", s.c.baseURL)
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
