package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

// DestinationCreateService implements the Destination Management, Create a Destination API.
// Ref. https://fivetran.com/docs/rest-api/destinations#createadestination
type DestinationCreateService struct {
	c                 *Client
	groupID           *string
	service           *string
	region            *string
	timeZoneOffset    *string
	config            *DestinationConfig
	trustCertificates *bool
	trustFingerprints *bool
	runSetupTests     *bool
}

type destinationCreateRequest struct {
	GroupID           *string                   `json:"group_id,omitempty"`
	Service           *string                   `json:"service,omitempty"`
	Region            *string                   `json:"region,omitempty"`
	TimeZoneOffset    *string                   `json:"time_zone_offset,omitempty"`
	Config            *destinationConfigRequest `json:"config,omitempty"`
	TrustCertificates *bool                     `json:"trust_certificates,omitempty"`
	TrustFingerprints *bool                     `json:"trust_fingerprints,omitempty"`
	RunSetupTests     *bool                     `json:"run_setup_tests,omitempty"`
}

type DestinationCreateResponse struct {
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

func (c *Client) NewDestinationCreate() *DestinationCreateService {
	return &DestinationCreateService{c: c}
}

func (s *DestinationCreateService) request() *destinationCreateRequest {
	var config *destinationConfigRequest

	if s.config != nil {
		config = s.config.request()
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

func (s *DestinationCreateService) Config(value *DestinationConfig) *DestinationCreateService {
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

func (s *DestinationCreateService) Do(ctx context.Context) (DestinationCreateResponse, error) {
	var response DestinationCreateResponse
	url := fmt.Sprintf("%v/destinations", s.c.baseURL)
	expectedStatus := 201

	headers := s.c.commonHeaders()
	headers["Content-Type"] = "application/json"
	headers["Accept"] = restAPIv2

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return response, err
	}

	r := request{
		method:  "POST",
		url:     url,
		body:    reqBody,
		queries: nil,
		headers: headers,
		client:  s.c.httpClient,
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
