package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

type DestinationDeleteService struct {
	c             *Client
	destinationID string
}

type DestinationDelete struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c *Client) NewDestinationDeleteService() *DestinationDeleteService {
	return &DestinationDeleteService{c: c}
}

func (s *DestinationDeleteService) DestinationID(destinationID string) *DestinationDeleteService {
	s.destinationID = destinationID
	return s
}

func (s *DestinationDeleteService) Do(ctx context.Context) (DestinationDelete, error) {
	if s.destinationID == "" { // we don't validate business rules (unless it is strictly necessary)
		err := fmt.Errorf("missing required DestinationID")
		return DestinationDelete{}, err
	}

	url := fmt.Sprintf("%v/destinations/%v", s.c.baseURL, s.destinationID)
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
		return DestinationDelete{}, err
	}

	var destinationDelete DestinationDelete
	if err := json.Unmarshal(respBody, &destinationDelete); err != nil {
		return DestinationDelete{}, err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected %v", respStatus, expectedStatus)
		return destinationDelete, err
	}

	return destinationDelete, nil
}
