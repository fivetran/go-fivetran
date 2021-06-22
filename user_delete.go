package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

// UserDeleteService implements the User Management, Delete a user API.
// Ref. https://fivetran.com/docs/rest-api/users#deleteauser
type UserDeleteService struct {
	c      *Client
	userID *string
}

type userDeleteResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c *Client) NewUserDelete() *UserDeleteService {
	return &UserDeleteService{c: c}
}

func (s *UserDeleteService) UserID(value string) *UserDeleteService {
	s.userID = &value
	return s
}

func (s *UserDeleteService) Do(ctx context.Context) (userDeleteResponse, error) {
	var response userDeleteResponse

	if s.userID == nil {
		return response, fmt.Errorf("missing required UserId")
	}

	url := fmt.Sprintf("%v/users/%v", s.c.baseURL, *s.userID)
	expectedStatus := 200

	headers := make(map[string]string)
	headers["Authorization"] = s.c.authorization

	r := request{
		method:  "DELETE",
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
		err := fmt.Errorf("status code: %v; expected %v", respStatus, expectedStatus)
		return response, err
	}

	return response, nil
}
