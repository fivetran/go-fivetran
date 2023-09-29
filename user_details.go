package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
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

	r := httputils.Request{
		Method:           "GET",
		Url:              url,
		Body:             nil,
		Queries:          nil,
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
