package fivetran

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// F stands for Field
// needs to be exported because of json.Marshal()
type GroupCreateService struct {
	c     *Client
	Fname string `json:"name,omitempty"`
}

type GroupCreate struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"data"`
}

func (c *Client) NewGroupCreateService() *GroupCreateService {
	return &GroupCreateService{c: c}
}

func (s *GroupCreateService) Name(name string) *GroupCreateService {
	s.Fname = name
	return s
}

func (s *GroupCreateService) Do(ctx context.Context) (GroupCreate, error) {
	url := fmt.Sprintf("%v/groups", s.c.baseURL)
	expectedStatus := 201 // https://fivetran.height.app/T-96892
	headers := make(map[string]string)

	headers["Authorization"] = s.c.authorization
	headers["Content-Type"] = "application/json"

	reqBody, err := json.Marshal(s)
	if err != nil {
		return GroupCreate{}, err
	}

	r := Request{
		method:  "POST",
		url:     url,
		body:    bytes.NewReader(reqBody),
		queries: nil,
		headers: headers,
	}

	respBody, respStatus, err := httpRequest(r, ctx)
	if err != nil {
		return GroupCreate{}, err
	}

	var groupCreate GroupCreate
	if err := json.Unmarshal(respBody, &groupCreate); err != nil {
		return GroupCreate{}, err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected %v", respStatus, expectedStatus)
		return groupCreate, err
	}

	return groupCreate, nil
}
