package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

// ExternalLoggingSetupTestsService implements the Log Management, Run Log service setup tests API.
// Ref. https://fivetran.com/docs/rest-api/log-service-management#runlogservicesetuptests
type ExternalLoggingSetupTestsService struct {
	c                 *Client
	externalLoggingID     *string
	trustCertificates *bool
	trustFingerprints *bool
}

type externalLoggingSetupTestsRequest struct {
	TrustCertificates *bool `json:"trust_certificates,omitempty"`
	TrustFingerprints *bool `json:"trust_fingerprints,omitempty"`
}

type ExternalLoggingSetupTestsResponse struct {
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
		Config ExternalLoggingConfigResponse `json:"config"`
	} `json:"data"`
}

func (c *Client) NewExternalLoggingSetupTests() *ExternalLoggingSetupTestsService {
	return &ExternalLoggingSetupTestsService{c: c}
}

func (s *ExternalLoggingSetupTestsService) request() *externalLoggingSetupTestsRequest {
	return &externalLoggingSetupTestsRequest{
		TrustCertificates: s.trustCertificates,
		TrustFingerprints: s.trustFingerprints,
	}
}

func (s *ExternalLoggingSetupTestsService) ExternalLoggingID(value string) *ExternalLoggingSetupTestsService {
	s.externalLoggingID = &value
	return s
}

func (s *ExternalLoggingSetupTestsService) TrustCertificates(value bool) *ExternalLoggingSetupTestsService {
	s.trustCertificates = &value
	return s
}

func (s *ExternalLoggingSetupTestsService) TrustFingerprints(value bool) *ExternalLoggingSetupTestsService {
	s.trustFingerprints = &value
	return s
}

func (s *ExternalLoggingSetupTestsService) Do(ctx context.Context) (ExternalLoggingSetupTestsResponse, error) {
	var response ExternalLoggingSetupTestsResponse

	if s.externalLoggingID == nil {
		return response, fmt.Errorf("missing required ExternalLoggingID")
	}

	url := fmt.Sprintf("%v/external-logging/%v/test", s.c.baseURL, *s.externalLoggingID)
	expectedStatus := 200

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
