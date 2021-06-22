package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type groupModifyService struct {
	c       *Client
	groupID *string
	name    *string
}

type groupModifyRequest struct {
	Name *string `json:"name,omitempty"`
}

type GroupModifyResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"data"`
}

func (c *Client) NewGroupModify() *groupModifyService {
	return &groupModifyService{c: c}
}

func (s *groupModifyService) request() groupModifyRequest {
	return groupModifyRequest{
		Name: s.name,
	}
}

func (s *groupModifyService) GroupID(value string) *groupModifyService {
	s.groupID = &value
	return s
}

func (s *groupModifyService) Name(value string) *groupModifyService {
	s.name = &value
	return s
}

func (s *groupModifyService) Do(ctx context.Context) (GroupModifyResponse, error) {
	var response GroupModifyResponse

	if s.groupID == nil {
		return response, fmt.Errorf("missing required GroupID")
	}

	url := fmt.Sprintf("%v/groups/%v", s.c.baseURL, *s.groupID)
	expectedStatus := 200

	headers := make(map[string]string)
	headers["Authorization"] = s.c.authorization
	headers["Content-Type"] = "application/json"

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return response, err
	}

	r := Request{
		method:  "PATCH",
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
