package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/destinations"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// DestinationModifyService implements the Destination Management, Modify a Destination API.
// Ref. https://fivetran.com/docs/rest-api/destinations#modifyadestination
type DestinationModifyService struct {
	c                 *Client
	destinationID     *string
	region            *string
	timeZoneOffset    *string
	config            *destinations.DestinationConfig
	trustCertificates *bool
	trustFingerprints *bool
	runSetupTests     *bool
}

type destinationModifyRequest struct {
	Region            *string `json:"region,omitempty"`
	TimeZoneOffset    *string `json:"time_zone_offset,omitempty"`
	Config            any     `json:"config,omitempty"`
	TrustCertificates *bool   `json:"trust_certificates,omitempty"`
	TrustFingerprints *bool   `json:"trust_fingerprints,omitempty"`
	RunSetupTests     *bool   `json:"run_setup_tests,omitempty"`
}

func (c *Client) NewDestinationModify() *DestinationModifyService {
	return &DestinationModifyService{c: c}
}

func (s *DestinationModifyService) request() *destinationModifyRequest {
	var config interface{}

	if s.config != nil {
		config = s.config.Request()
	}

	return &destinationModifyRequest{
		Region:            s.region,
		TimeZoneOffset:    s.timeZoneOffset,
		Config:            config,
		TrustCertificates: s.trustCertificates,
		TrustFingerprints: s.trustFingerprints,
		RunSetupTests:     s.runSetupTests,
	}
}

func (s *DestinationModifyService) DestinationID(value string) *DestinationModifyService {
	s.destinationID = &value
	return s
}

func (s *DestinationModifyService) Region(value string) *DestinationModifyService {
	s.region = &value
	return s
}

func (s *DestinationModifyService) TimeZoneOffset(value string) *DestinationModifyService {
	s.timeZoneOffset = &value
	return s
}

func (s *DestinationModifyService) Config(value *destinations.DestinationConfig) *DestinationModifyService {
	s.config = value
	return s
}

func (s *DestinationModifyService) TrustCertificates(value bool) *DestinationModifyService {
	s.trustCertificates = &value
	return s
}

func (s *DestinationModifyService) TrustFingerprints(value bool) *DestinationModifyService {
	s.trustFingerprints = &value
	return s
}

func (s *DestinationModifyService) RunSetupTests(value bool) *DestinationModifyService {
	s.runSetupTests = &value
	return s
}

func (s *DestinationModifyService) Do(ctx context.Context) (destinations.DestinationDetailsWithSetupTestsResponse, error) {
	var response destinations.DestinationDetailsWithSetupTestsResponse

	if s.destinationID == nil {
		return response, fmt.Errorf("missing required DestinationID")
	}

	url := fmt.Sprintf("%v/destinations/%v", s.c.baseURL, *s.destinationID)
	expectedStatus := 200

	headers := s.c.commonHeaders()
	headers["Content-Type"] = "application/json"
	headers["Accept"] = restAPIv2

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return response, err
	}

	r := httputils.Request{
		Method:           "PATCH",
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
