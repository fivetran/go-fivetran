package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

type groupRemoveUserService struct {
	c       *Client
	groupID *string
	userID  *string
}

type GroupRemoveUserResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c *Client) NewGroupRemoveUser() *groupRemoveUserService {
	return &groupRemoveUserService{c: c}
}

func (s *groupRemoveUserService) GroupID(value string) *groupRemoveUserService {
	s.groupID = &value
	return s
}

func (s *groupRemoveUserService) UserID(value string) *groupRemoveUserService {
	s.userID = &value
	return s
}

func (s *groupRemoveUserService) Do(ctx context.Context) (GroupRemoveUserResponse, error) {
	var response GroupRemoveUserResponse

	if s.groupID == nil {
		return response, fmt.Errorf("missing required GroupID")
	}
	if s.userID == nil {
		return response, fmt.Errorf("missing required UserID")
	}

	url := fmt.Sprintf("%v/groups/%v/users/%v", s.c.baseURL, *s.groupID, *s.userID)
	expectedStatus := 200

	headers := make(map[string]string)
	headers["Authorization"] = s.c.authorization

	r := Request{
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
