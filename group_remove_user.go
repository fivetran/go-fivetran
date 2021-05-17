package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

type GroupRemoveUserService struct {
	c      *Client
	id     string
	userID string
}

type GroupRemoveUser struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c *Client) NewGroupRemoveUserService() *GroupRemoveUserService {
	return &GroupRemoveUserService{c: c}
}

func (s *GroupRemoveUserService) ID(id string) *GroupRemoveUserService {
	s.id = id
	return s
}

func (s *GroupRemoveUserService) UserID(userID string) *GroupRemoveUserService {
	s.userID = userID
	return s
}

func (s *GroupRemoveUserService) Do(ctx context.Context) (GroupRemoveUser, error) {
	if s.id == "" {
		err := fmt.Errorf("missing required ID")
		return GroupRemoveUser{}, err
	}

	if s.userID == "" {
		err := fmt.Errorf("missing required UserID")
		return GroupRemoveUser{}, err
	}

	url := fmt.Sprintf("%v/groups/%v/users/%v", s.c.baseURL, s.id, s.userID)
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
		return GroupRemoveUser{}, err
	}

	var groupRemoveUser GroupRemoveUser
	if err := json.Unmarshal(respBody, &groupRemoveUser); err != nil {
		return GroupRemoveUser{}, err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected %v", respStatus, expectedStatus)
		return groupRemoveUser, err
	}

	return groupRemoveUser, nil
}
