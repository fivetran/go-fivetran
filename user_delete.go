package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
)

// UserDeleteService implements the User Management, Delete a user API.
// Ref. https://fivetran.com/docs/rest-api/users#deleteauser
type UserDeleteService struct {
	c      *Client
	userID *string
}

func (c *Client) NewUserDelete() *UserDeleteService {
	return &UserDeleteService{c: c}
}

func (s *UserDeleteService) UserID(value string) *UserDeleteService {
	s.userID = &value
	return s
}

func (s *UserDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse

	if s.userID == nil {
		return response, fmt.Errorf("missing required UserId")
	}

	url := fmt.Sprintf("%v/users/%v", s.c.baseURL, *s.userID)
	expectedStatus := 200

	headers := s.c.commonHeaders()

	r := request{
		method:           "DELETE",
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
