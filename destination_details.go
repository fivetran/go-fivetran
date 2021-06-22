package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

type destinationDetailsService struct {
	c             *Client
	destinationID *string
}

type DestinationDetailsResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID             string                    `json:"id"`
		GroupID        string                    `json:"group_id"`
		Service        string                    `json:"service"`
		Region         string                    `json:"region"`
		TimeZoneOffset string                    `json:"time_zone_offset"`
		SetupStatus    string                    `json:"setup_status"`
		Config         DestinationConfigResponse `json:"config"`
	} `json:"data"`
}

func (c *Client) NewDestinationDetails() *destinationDetailsService {
	return &destinationDetailsService{c: c}
}

func (s *destinationDetailsService) DestinationID(value string) *destinationDetailsService {
	s.destinationID = &value
	return s
}

func (s *destinationDetailsService) Do(ctx context.Context) (DestinationDetailsResponse, error) {
	var response DestinationDetailsResponse

	if s.destinationID == nil {
		return response, fmt.Errorf("missing required DestinationID")
	}

	url := fmt.Sprintf("%v/destinations/%v", s.c.baseURL, *s.destinationID)
	expectedStatus := 200

	headers := make(map[string]string)
	headers["Authorization"] = s.c.authorization

	r := Request{
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
