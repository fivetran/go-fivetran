package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/destinations"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// DestinationCreateService implements the Destination Management, Create a Destination API.
// Ref. https://fivetran.com/docs/rest-api/destinations#createadestination
type DestinationCreateService struct {
	c                 *Client
	groupID           *string
	service           *string
	region            *string
	timeZoneOffset    *string
	config            *destinations.DestinationConfig
	trustCertificates *bool
	trustFingerprints *bool
	runSetupTests     *bool
}

type destinationCreateRequest struct {
	GroupID           *string `json:"group_id,omitempty"`
	Service           *string `json:"service,omitempty"`
	Region            *string `json:"region,omitempty"`
	TimeZoneOffset    *string `json:"time_zone_offset,omitempty"`
	Config            any     `json:"config,omitempty"`
	TrustCertificates *bool   `json:"trust_certificates,omitempty"`
	TrustFingerprints *bool   `json:"trust_fingerprints,omitempty"`
	RunSetupTests     *bool   `json:"run_setup_tests,omitempty"`
}

func (c *Client) NewDestinationCreate() *DestinationCreateService {
	return &DestinationCreateService{c: c}
}

func (s *DestinationCreateService) request() *destinationCreateRequest {
	var config interface{}
	if s.config != nil {
		config = s.config.Request()
	}

	return &destinationCreateRequest{
		GroupID:           s.groupID,
		Service:           s.service,
		Region:            s.region,
		TimeZoneOffset:    s.timeZoneOffset,
		Config:            config,
		TrustCertificates: s.trustCertificates,
		TrustFingerprints: s.trustFingerprints,
		RunSetupTests:     s.runSetupTests,
	}
}

func (s *DestinationCreateService) GroupID(value string) *DestinationCreateService {
	s.groupID = &value
	return s
}

func (s *DestinationCreateService) Service(value string) *DestinationCreateService {
	s.service = &value
	return s
}

func (s *DestinationCreateService) Region(value string) *DestinationCreateService {
	s.region = &value
	return s
}

func (s *DestinationCreateService) TimeZoneOffset(value string) *DestinationCreateService {
	s.timeZoneOffset = &value
	return s
}

func (s *DestinationCreateService) Config(value *destinations.DestinationConfig) *DestinationCreateService {
	s.config = value
	return s
}

func (s *DestinationCreateService) TrustCertificates(value bool) *DestinationCreateService {
	s.trustCertificates = &value
	return s
}

func (s *DestinationCreateService) TrustFingerprints(value bool) *DestinationCreateService {
	s.trustFingerprints = &value
	return s
}

func (s *DestinationCreateService) RunSetupTests(value bool) *DestinationCreateService {
	s.runSetupTests = &value
	return s
}

func (s *DestinationCreateService) Do(ctx context.Context) (destinations.DestinationDetailsWithSetupTestsResponse, error) {
	var response destinations.DestinationDetailsWithSetupTestsResponse
	url := fmt.Sprintf("%v/destinations", s.c.baseURL)
	expectedStatus := 201

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
