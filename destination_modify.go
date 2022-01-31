package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

// DestinationModifyService implements the Destination Management, Modify a Destination API.
// Ref. https://fivetran.com/docs/rest-api/destinations#modifyadestination
type DestinationModifyService struct {
	c                 *Client
	destinationID     *string
	region            *string
	timeZoneOffset    *string
	config            *DestinationConfig
	trustCertificates *bool
	trustFingerprints *bool
	runSetupTests     *bool
}

type destinationModifyRequest struct {
	Region            *string                   `json:"region,omitempty"`
	TimeZoneOffset    *string                   `json:"time_zone_offset,omitempty"`
	Config            *destinationConfigRequest `json:"config,omitempty"`
	TrustCertificates *bool                     `json:"trust_certificates,omitempty"`
	TrustFingerprints *bool                     `json:"trust_fingerprints,omitempty"`
	RunSetupTests     *bool                     `json:"run_setup_tests,omitempty"`
}

type DestinationModifyResponse struct {
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

func (c *Client) NewDestinationModify() *DestinationModifyService {
	return &DestinationModifyService{c: c}
}

func (s *DestinationModifyService) request() *destinationModifyRequest {
	var config *destinationConfigRequest

	if s.config != nil {
		config = s.config.request()
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

func (s *DestinationModifyService) Config(value *DestinationConfig) *DestinationModifyService {
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

func (s *DestinationModifyService) Do(ctx context.Context) (DestinationModifyResponse, error) {
	var response DestinationModifyResponse

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

	r := request{
		method:  "PATCH",
		url:     url,
		body:    reqBody,
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
