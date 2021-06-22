package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

type destinationDeleteService struct {
	c             *Client
	destinationID *string
}

type DestinationDeleteResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c *Client) NewDestinationDelete() *destinationDeleteService {
	return &destinationDeleteService{c: c}
}

func (s *destinationDeleteService) DestinationID(value string) *destinationDeleteService {
	s.destinationID = &value
	return s
}

func (s *destinationDeleteService) Do(ctx context.Context) (DestinationDeleteResponse, error) {
	var response DestinationDeleteResponse

	if s.destinationID == nil {
		return response, fmt.Errorf("missing required DestinationID")
	}

	url := fmt.Sprintf("%v/destinations/%v", s.c.baseURL, *s.destinationID)
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
