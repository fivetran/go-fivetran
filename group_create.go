package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type groupCreateService struct {
	c    *Client
	name *string
}

type groupCreateRequest struct {
	Name *string `json:"name,omitempty"`
}

type GroupCreateResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"data"`
}

func (c *Client) NewGroupCreate() *groupCreateService {
	return &groupCreateService{c: c}
}

func (s *groupCreateService) request() groupCreateRequest {
	return groupCreateRequest{
		Name: s.name,
	}
}

func (s *groupCreateService) Name(value string) *groupCreateService {
	s.name = &value
	return s
}

func (s *groupCreateService) Do(ctx context.Context) (GroupCreateResponse, error) {
	var response GroupCreateResponse
	url := fmt.Sprintf("%v/groups", s.c.baseURL)
	expectedStatus := 201

	headers := make(map[string]string)
	headers["Authorization"] = s.c.authorization
	headers["Content-Type"] = "application/json"

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return response, err
	}

	r := Request{
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
		err := fmt.Errorf("status code: %v; expected %v", respStatus, expectedStatus)
		return response, err
	}

	return response, nil
}
