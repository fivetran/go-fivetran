package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

// DestinationDetailsService implements the Destination Management, Retrieve destination details API.
// Ref. https://fivetran.com/docs/rest-api/destinations#retrievedestinationdetails
type DestinationDetailsService struct {
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

func (c *Client) NewDestinationDetails() *DestinationDetailsService {
	return &DestinationDetailsService{c: c}
}

func (s *DestinationDetailsService) DestinationID(value string) *DestinationDetailsService {
	s.destinationID = &value
	return s
}

func (s *DestinationDetailsService) Do(ctx context.Context) (DestinationDetailsResponse, error) {
	var response DestinationDetailsResponse

	if s.destinationID == nil {
		return response, fmt.Errorf("missing required DestinationID")
	}

	url := fmt.Sprintf("%v/destinations/%v", s.c.baseURL, *s.destinationID)
	expectedStatus := 200

	headers := s.c.commonHeaders()
	headers["Accept"] = restAPIv2

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
		err := fmt.Errorf("status code: %v; expected: %v", respStatus, expectedStatus)
		return response, err
	}

	return response, nil
}
