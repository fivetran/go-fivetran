package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// GroupModifyService implements the Group Management, Modify a Group API.
// Ref. https://fivetran.com/docs/rest-api/groups#modifyagroup
type GroupModifyService struct {
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

func (c *Client) NewGroupModify() *GroupModifyService {
	return &GroupModifyService{c: c}
}

func (s *GroupModifyService) request() *groupModifyRequest {
	return &groupModifyRequest{
		Name: s.name,
	}
}

func (s *GroupModifyService) GroupID(value string) *GroupModifyService {
	s.groupID = &value
	return s
}

func (s *GroupModifyService) Name(value string) *GroupModifyService {
	s.name = &value
	return s
}

func (s *GroupModifyService) Do(ctx context.Context) (GroupModifyResponse, error) {
	var response GroupModifyResponse

	if s.groupID == nil {
		return response, fmt.Errorf("missing required GroupID")
	}

	url := fmt.Sprintf("%v/groups/%v", s.c.baseURL, *s.groupID)
	expectedStatus := 200

	headers := s.c.commonHeaders()
	headers["Content-Type"] = "application/json"

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return response, err
	}

	r := request{
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
		err := fmt.Errorf("status code: %v; expected: %v", respStatus, expectedStatus)
		return response, err
	}

	return response, nil
}
