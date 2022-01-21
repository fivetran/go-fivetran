package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// UserDetailsService implements the User Management, Retrieve user details API.
// Ref. https://fivetran.com/docs/rest-api/users#retrieveuserdetails
type UserDetailsService struct {
	c      *Client
	userID *string
}

type UserDetailsResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID         string    `json:"id"`
		Email      string    `json:"email"`
		GivenName  string    `json:"given_name"`
		FamilyName string    `json:"family_name"`
		Verified   *bool     `json:"verified"`
		Invited    *bool     `json:"invited"`
		Picture    string    `json:"picture"`
		Phone      string    `json:"phone"`
		LoggedInAt time.Time `json:"logged_in_at"`
		CreatedAt  time.Time `json:"created_at"`
		Role	   string    `json:"role"`
	} `json:"data"`
}

func (c *Client) NewUserDetails() *UserDetailsService {
	return &UserDetailsService{c: c}
}

func (s *UserDetailsService) UserID(value string) *UserDetailsService {
	s.userID = &value
	return s
}

func (s *UserDetailsService) Do(ctx context.Context) (UserDetailsResponse, error) {
	var response UserDetailsResponse

	if s.userID == nil {
		return response, fmt.Errorf("missing required UserId")
	}

	url := fmt.Sprintf("%v/users/%v", s.c.baseURL, *s.userID)
	expectedStatus := 200

	headers := s.c.commonHeaders()

	r := request{
		method:  "GET",
		url:     url,
		body:    nil,
		queries: nil,
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
		err := fmt.Errorf("status code: %v; expected: %v", respStatus, expectedStatus)
		return response, err
	}

	return response, nil
}
