package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

type GroupDeleteService struct {
	c       *Client
	groupID *string
}

type groupDeleteResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c *Client) NewGroupDelete() *GroupDeleteService {
	return &GroupDeleteService{c: c}
}

func (s *GroupDeleteService) GroupID(value string) *GroupDeleteService {
	s.groupID = &value
	return s
}

func (s *GroupDeleteService) Do(ctx context.Context) (groupDeleteResponse, error) {
	var response groupDeleteResponse

	if s.groupID == nil {
		return response, fmt.Errorf("missing required GroupID")
	}

	url := fmt.Sprintf("%v/groups/%v", s.c.baseURL, *s.groupID)
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
