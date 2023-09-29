package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/destinations"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// DestinationDetailsService implements the Destination Management, Retrieve destination details API.
// Ref. https://fivetran.com/docs/rest-api/destinations#retrievedestinationdetails
type DestinationDetailsService struct {
	c             *Client
	destinationID *string
}

func (c *Client) NewDestinationDetails() *DestinationDetailsService {
	return &DestinationDetailsService{c: c}
}

func (s *DestinationDetailsService) DestinationID(value string) *DestinationDetailsService {
	s.destinationID = &value
	return s
}

func (s *DestinationDetailsService) Do(ctx context.Context) (destinations.DestinationDetailsResponse, error) {
	var response destinations.DestinationDetailsResponse

	if s.destinationID == nil {
		return response, fmt.Errorf("missing required DestinationID")
	}

	url := fmt.Sprintf("%v/destinations/%v", s.c.baseURL, *s.destinationID)
	expectedStatus := 200

	headers := s.c.commonHeaders()
	headers["Accept"] = restAPIv2

	r := httputils.Request{
		Method:           "GET",
		Url:              url,
		Body:             nil,
		Queries:          nil,
		Headers:          headers,
		Client:           s.c.httpClient,
		HandleRateLimits: s.c.handleRateLimits,
		MaxRetryAttempts: s.c.maxRetryAttempts,
	}

	respBody, respStatus, err := r.Do(ctx)
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
