package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/destinations"
	httputils "github.com/fivetran/go-fivetran/http_utils"
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

	r := httputils.Request{
		Method:           "POST",
		Url:              url,
		Body:             reqBody,
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
