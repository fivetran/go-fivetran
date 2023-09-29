package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/users"
)

// UserDetailsService implements the User Management, Retrieve user details API.
// Ref. https://fivetran.com/docs/rest-api/users#retrieveuserdetails
type UserDetailsService struct {
	c      *Client
	userID *string
}

func (c *Client) NewUserDetails() *UserDetailsService {
	return &UserDetailsService{c: c}
}

func (s *UserDetailsService) UserID(value string) *UserDetailsService {
	s.userID = &value
	return s
}

func (s *UserDetailsService) Do(ctx context.Context) (users.UserDetailsResponse, error) {
	var response users.UserDetailsResponse

	if s.userID == nil {
		return response, fmt.Errorf("missing required UserId")
	}

	url := fmt.Sprintf("%v/users/%v", s.c.baseURL, *s.userID)
	expectedStatus := 200

	headers := s.c.commonHeaders()

	r := request{
		method:           "GET",
		url:              url,
		body:             nil,
		queries:          nil,
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
