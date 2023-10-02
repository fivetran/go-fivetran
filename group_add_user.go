package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// GroupAddUserService implements the Group Management, Add a User to a Group API.
// Ref. https://fivetran.com/docs/rest-api/groups#addausertoagroup
type GroupAddUserService struct {
	c       *Client
	groupID *string
	email   *string
	role    *string
}

type groupAddUserRequest struct {
	Email *string `json:"email,omitempty"`
	Role  *string `json:"role,omitempty"`
}

func (c *Client) NewGroupAddUser() *GroupAddUserService {
	return &GroupAddUserService{c: c}
}

func (s *GroupAddUserService) request() *groupAddUserRequest {
	return &groupAddUserRequest{
		Email: s.email,
		Role:  s.role,
	}
}

func (s *GroupAddUserService) GroupID(value string) *GroupAddUserService {
	s.groupID = &value
	return s
}

func (s *GroupAddUserService) Email(value string) *GroupAddUserService {
	s.email = &value
	return s
}

func (s *GroupAddUserService) Role(value string) *GroupAddUserService {
	s.role = &value
	return s
}

func (s *GroupAddUserService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse

	if s.groupID == nil {
		return response, fmt.Errorf("missing required GroupID")
	}

	url := fmt.Sprintf("%v/groups/%v/users", s.c.baseURL, *s.groupID)
	expectedStatus := 200

	headers := s.c.commonHeaders()
	headers["Content-Type"] = "application/json"

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return response, err
	}

	r := httputils.Request{
		Method:           "POST",
		Url:              url,
		Body:             reqBody,
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
