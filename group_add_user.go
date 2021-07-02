package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
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

type GroupAddUserResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
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

func (s *GroupAddUserService) Do(ctx context.Context) (GroupAddUserResponse, error) {
	var response GroupAddUserResponse

	if s.groupID == nil {
		return response, fmt.Errorf("missing required GroupID")
	}

	url := fmt.Sprintf("%v/groups/%v/users", s.c.baseURL, *s.groupID)
	expectedStatus := 200

	headers := make(map[string]string)
	headers["Authorization"] = s.c.authorization
	headers["Content-Type"] = "application/json"

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return response, err
	}

	r := request{
		method:  "POST",
		url:     url,
		body:    reqBody,
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
