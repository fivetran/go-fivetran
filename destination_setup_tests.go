package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

type destinationSetupTestsService struct {
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

func (c *Client) NewDestinationSetupTests() *destinationSetupTestsService {
	return &destinationSetupTestsService{c: c}
}

func (s *destinationSetupTestsService) request() destinationSetupTestsRequest {
	return destinationSetupTestsRequest{
		TrustCertificates: s.trustCertificates,
		TrustFingerprints: s.trustFingerprints,
	}
}

func (s *destinationSetupTestsService) DestinationID(value string) *destinationSetupTestsService {
	s.destinationID = &value
	return s
}

func (s *destinationSetupTestsService) TrustCertificates(value bool) *destinationSetupTestsService {
	s.trustCertificates = &value
	return s
}

func (s *destinationSetupTestsService) TrustFingerprints(value bool) *destinationSetupTestsService {
	s.trustFingerprints = &value
	return s
}

func (s *destinationSetupTestsService) Do(ctx context.Context) (DestinationSetupTestsResponse, error) {
	var response DestinationSetupTestsResponse

	if s.destinationID == nil {
		return response, fmt.Errorf("missing required DestinationID")
	}

	url := fmt.Sprintf("%v/destinations/%v/test", s.c.baseURL, *s.destinationID)
	expectedStatus := 200

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
