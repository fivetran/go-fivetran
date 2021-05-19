package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

type DestinationDetailsService struct {
	c             *Client
	destinationID string
}

type DestinationDetails struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID             string            `json:"id"`
		GroupID        string            `json:"group_id"`
		Service        string            `json:"service"`
		Region         string            `json:"region"`
		TimeZoneOffset string            `json:"time_zone_offset"`
		SetupStatus    string            `json:"setup_status"`
		Config         DestinationConfig `json:"config"`
	} `json:"data"`
}

func (c *Client) NewDestinationDetailsService() *DestinationDetailsService {
	return &DestinationDetailsService{c: c}
}

func (s *DestinationDetailsService) DestinationID(destinationID string) *DestinationDetailsService {
	s.destinationID = destinationID
	return s
}

func (s *DestinationDetailsService) Do(ctx context.Context) (DestinationDetails, error) {
	if s.destinationID == "" { // we don't validate business rules (unless it is strictly necessary)
		err := fmt.Errorf("missing required DestinationID")
		return DestinationDetails{}, err
	}

	url := fmt.Sprintf("%v/destinations/%v", s.c.baseURL, s.destinationID)
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
		return DestinationDetails{}, err
	}

	var destinationDetails DestinationDetails
	if err := json.Unmarshal(respBody, &destinationDetails); err != nil {
		return DestinationDetails{}, err
	}

	// converts destinationDetails.Data.Config.Fport to int. Should be removed
	// when https://fivetran.height.app/T-97508 fixed.
	switch destinationDetails.Data.Config.Fport.(type) {
	case string:
		destinationDetails.Data.Config.Fport, err = strconv.Atoi(destinationDetails.Data.Config.Fport.(string))
		if err != nil {
			return DestinationDetails{}, err
		}

	default:
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected %v", respStatus, expectedStatus)
		return destinationDetails, err
	}

	return destinationDetails, nil
}
