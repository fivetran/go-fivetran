package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

type destinationModifyService struct {
	c                 *Client
	destinationID     *string
	region            *string
	timeZoneOffset    *string
	config            *destinationConfig
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

func (c *Client) NewDestinationModify() *destinationModifyService {
	return &destinationModifyService{c: c}
}

func (s *destinationModifyService) request() destinationModifyRequest {
	var config *destinationConfigRequest

	if s.config != nil {
		config = s.config.request()
	}

	return destinationModifyRequest{
		Region:            s.region,
		TimeZoneOffset:    s.timeZoneOffset,
		Config:            config,
		TrustCertificates: s.trustCertificates,
		TrustFingerprints: s.trustFingerprints,
		RunSetupTests:     s.runSetupTests,
	}
}

func (s *destinationModifyService) DestinationID(value string) *destinationModifyService {
	s.destinationID = &value
	return s
}

func (s *destinationModifyService) Region(value string) *destinationModifyService {
	s.region = &value
	return s
}

func (s *destinationModifyService) TimeZoneOffset(value string) *destinationModifyService {
	s.timeZoneOffset = &value
	return s
}

func (s *destinationModifyService) Config(value *destinationConfig) *destinationModifyService {
	s.config = value
	return s
}

func (s *destinationModifyService) TrustCertificates(value bool) *destinationModifyService {
	s.trustCertificates = &value
	return s
}

func (s *destinationModifyService) TrustFingerprints(value bool) *destinationModifyService {
	s.trustFingerprints = &value
	return s
}

func (s *destinationModifyService) RunSetupTests(value bool) *destinationModifyService {
	s.runSetupTests = &value
	return s
}

func (s *destinationModifyService) Do(ctx context.Context) (DestinationModifyResponse, error) {
	var response DestinationModifyResponse

	if s.destinationID == nil {
		return response, fmt.Errorf("missing required DestinationID")
	}

	url := fmt.Sprintf("%v/destinations/%v", s.c.baseURL, *s.destinationID)
	expectedStatus := 200

	headers := make(map[string]string)
	headers["Authorization"] = s.c.authorization
	headers["Content-Type"] = "application/json"

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return response, err
	}

	r := Request{
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
		err := fmt.Errorf("status code: %v; expected %v", respStatus, expectedStatus)
		return response, err
	}

	return response, nil
}
