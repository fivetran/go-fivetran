package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

type GroupAddUserService struct {
	c      *Client
	id     string
	Femail string `json:"email,omitempty"`
	Frole  string `json:"role,omitempty"`
}

type GroupAddUser struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c *Client) NewGroupAddUserService() *GroupAddUserService {
	return &GroupAddUserService{c: c}
}

func (s *GroupAddUserService) ID(id string) *GroupAddUserService {
	s.id = id
	return s
}

func (s *GroupAddUserService) Email(email string) *GroupAddUserService {
	s.Femail = email
	return s
}

func (s *GroupAddUserService) Role(role string) *GroupAddUserService {
	s.Frole = role
	return s
}

func (s *GroupAddUserService) Do(ctx context.Context) (GroupAddUser, error) {
	if s.id == "" {
		err := fmt.Errorf("missing required ID")
		return GroupAddUser{}, err
	}

	url := fmt.Sprintf("%v/groups/%v/users", s.c.baseURL, s.id)
	expectedStatus := 201
	headers := make(map[string]string)
	queries := make(map[string]string)

	headers["Authorization"] = s.c.authorization
	headers["Content-Type"] = "application/json"

	reqBody, err := json.Marshal(s)
	if err != nil {
		return GroupAddUser{}, err
	}

	r := Request{
		method:  "POST",
		url:     url,
		body:    reqBody,
		queries: queries,
		headers: headers,
	}

	respBody, respStatus, err := httpRequest(r, ctx)
	if err != nil {
		return GroupAddUser{}, err
	}

	var groupAddUser GroupAddUser
	if err := json.Unmarshal(respBody, &groupAddUser); err != nil {
		return GroupAddUser{}, err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected %v", respStatus, expectedStatus)
		return groupAddUser, err
	}

	return groupAddUser, nil
}
