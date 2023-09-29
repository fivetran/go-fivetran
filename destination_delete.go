package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
)

// DestinationDeleteService implements the Destination Management, Delete a destination API.
// Ref. https://fivetran.com/docs/rest-api/destinations#deleteadestination
type DestinationDeleteService struct {
	c             *Client
	destinationID *string
}

func (c *Client) NewDestinationDelete() *DestinationDeleteService {
	return &DestinationDeleteService{c: c}
}

func (s *DestinationDeleteService) DestinationID(value string) *DestinationDeleteService {
	s.destinationID = &value
	return s
}

func (s *DestinationDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse

	if s.destinationID == nil {
		return response, fmt.Errorf("missing required DestinationID")
	}

	url := fmt.Sprintf("%v/destinations/%v", s.c.baseURL, *s.destinationID)
	expectedStatus := 200

	headers := s.c.commonHeaders()

	r := request{
		method:           "DELETE",
		url:              url,
		body:             nil,
		queries:          nil,
		headers:          headers,
		client:           s.c.httpClient,
		handleRateLimits: s.c.handleRateLimits,
		maxRetryAttempts: s.c.maxRetryAttempts,
	}

	respBody, respStatus, err := r.httpRequest(ctx)
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
