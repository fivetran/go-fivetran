package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/destinations"
)

// DestinationSetupTestsService implements the Destination Management, Run destination setup tests API.
// Ref. https://fivetran.com/docs/rest-api/destinations#rundestinationsetuptests
type DestinationSetupTestsService struct {
	c                 *Client
	destinationID     *string
	trustCertificates *bool
	trustFingerprints *bool
}

type destinationSetupTestsRequest struct {
	TrustCertificates *bool `json:"trust_certificates,omitempty"`
	TrustFingerprints *bool `json:"trust_fingerprints,omitempty"`
}

func (c *Client) NewDestinationSetupTests() *DestinationSetupTestsService {
	return &DestinationSetupTestsService{c: c}
}

func (s *DestinationSetupTestsService) request() *destinationSetupTestsRequest {
	return &destinationSetupTestsRequest{
		TrustCertificates: s.trustCertificates,
		TrustFingerprints: s.trustFingerprints,
	}
}

func (s *DestinationSetupTestsService) DestinationID(value string) *DestinationSetupTestsService {
	s.destinationID = &value
	return s
}

func (s *DestinationSetupTestsService) TrustCertificates(value bool) *DestinationSetupTestsService {
	s.trustCertificates = &value
	return s
}

func (s *DestinationSetupTestsService) TrustFingerprints(value bool) *DestinationSetupTestsService {
	s.trustFingerprints = &value
	return s
}

func (s *DestinationSetupTestsService) Do(ctx context.Context) (destinations.DestinationDetailsWithSetupTestsResponse, error) {
	var response destinations.DestinationDetailsWithSetupTestsResponse

	if s.destinationID == nil {
		return response, fmt.Errorf("missing required DestinationID")
	}

	url := fmt.Sprintf("%v/destinations/%v/test", s.c.baseURL, *s.destinationID)
	expectedStatus := 200

	headers := s.c.commonHeaders()
	headers["Content-Type"] = "application/json"
	headers["Accept"] = restAPIv2

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return response, err
	}

	r := request{
		method:           "POST",
		url:              url,
		body:             reqBody,
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
