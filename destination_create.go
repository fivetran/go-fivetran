package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

type destinationCreateService struct {
	c                 *Client
	groupID           *string
	service           *string
	region            *string
	timeZoneOffset    *string
	config            *destinationConfig
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

func (c *Client) NewDestinationCreate() *destinationCreateService {
	return &destinationCreateService{c: c}
}

func (s *destinationCreateService) request() *destinationCreateRequest {
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

func (s *destinationCreateService) GroupID(value string) *destinationCreateService {
	s.groupID = &value
	return s
}

func (s *destinationCreateService) Service(value string) *destinationCreateService {
	s.service = &value
	return s
}

func (s *destinationCreateService) Region(value string) *destinationCreateService {
	s.region = &value
	return s
}

func (s *destinationCreateService) TimeZoneOffset(value string) *destinationCreateService {
	s.timeZoneOffset = &value
	return s
}

func (s *destinationCreateService) Config(value *destinationConfig) *destinationCreateService {
	s.config = value
	return s
}

func (s *destinationCreateService) TrustCertificates(value bool) *destinationCreateService {
	s.trustCertificates = &value
	return s
}

func (s *destinationCreateService) TrustFingerprints(value bool) *destinationCreateService {
	s.trustFingerprints = &value
	return s
}

func (s *destinationCreateService) RunSetupTests(value bool) *destinationCreateService {
	s.runSetupTests = &value
	return s
}

func (s *destinationCreateService) Do(ctx context.Context) (DestinationCreateResponse, error) {
	var response DestinationCreateResponse
	url := fmt.Sprintf("%v/destinations", s.c.baseURL)
	expectedStatus := 201

	headers := make(map[string]string)
	headers["Authorization"] = s.c.authorization
	headers["Content-Type"] = "application/json"

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return response, err
	}

	r := Request{
		method:  "POST",
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
