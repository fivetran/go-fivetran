package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
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

type DestinationSetupTestsResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID             string `json:"id"`
		GroupID        string `json:"group_id"`
		Service        string `json:"service"`
		Region         string `json:"region"`
		TimeZoneOffset string `json:"time_zone_offset"`
		SetupStatus    string `json:"setup_status"`
		SetupTests     []struct {
			Title   string `json:"title"`
			Status  string `json:"status"`
			Message string `json:"message"`
		} `json:"setup_tests"`
		Config DestinationConfigResponse `json:"config"`
	} `json:"data"`
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

func (s *DestinationSetupTestsService) Do(ctx context.Context) (DestinationSetupTestsResponse, error) {
	var response DestinationSetupTestsResponse

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
