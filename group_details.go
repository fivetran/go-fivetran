package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type GroupDetailsService struct {
	c       *Client
	groupID *string
}

type GroupDetailsResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"data"`
}

func (c *Client) NewGroupDetails() *GroupDetailsService {
	return &GroupDetailsService{c: c}
}

func (s *GroupDetailsService) GroupID(value string) *GroupDetailsService {
	s.groupID = &value
	return s
}

func (s *GroupDetailsService) Do(ctx context.Context) (GroupDetailsResponse, error) {
	var response GroupDetailsResponse

	if s.groupID == nil {
		return response, fmt.Errorf("missing required GroupID")
	}

	url := fmt.Sprintf("%v/groups/%v", s.c.baseURL, *s.groupID)
	expectedStatus := 200

	headers := make(map[string]string)
	headers["Authorization"] = s.c.authorization

	r := request{
		method:  "GET",
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
