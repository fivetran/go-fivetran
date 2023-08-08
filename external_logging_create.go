package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

// ExternalLoggingCreateService implements the Log Management, Create a Log Service API.
// Ref. https://fivetran.com/docs/rest-api/log-service-management#createalogservice
type ExternalLoggingCreateService struct {
	c                 *Client
	groupID           *string
	service           *string
	region            *string
	timeZoneOffset    *string
	enable 			  *bool
	config            *ExternalLoggingConfig
}

type externalLoggingCreateRequest struct {
	GroupID           *string                  		`json:"group_id,omitempty"`
	Service           *string                  		`json:"service,omitempty"`
	Config            *externalLoggingConfigRequest `json:"config,omitempty"`
}

type ExternalLoggingCreateResponse struct {
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

func (c *Client) NewExternalLoggingCreate() *ExternalLoggingCreateService {
	return &ExternalLoggingCreateService{c: c}
}

func (s *ExternalLoggingCreateService) request() *externalLoggingCreateRequest {
	var config *externalLoggingConfigRequest

	if s.config != nil {
		config = s.config.request()
	}

	return &externalLoggingCreateRequest{
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

func (s *ExternalLoggingCreateService) GroupID(value string) *ExternalLoggingCreateService {
	s.groupID = &value
	return s
}

func (s *ExternalLoggingCreateService) Service(value string) *ExternalLoggingCreateService {
	s.service = &value
	return s
}

func (s *ExternalLoggingCreateService) Region(value string) *ExternalLoggingCreateService {
	s.region = &value
	return s
}

func (s *ExternalLoggingCreateService) TimeZoneOffset(value string) *ExternalLoggingCreateService {
	s.timeZoneOffset = &value
	return s
}

func (s *ExternalLoggingCreateService) Config(value *ExternalLoggingConfig) *ExternalLoggingCreateService {
	s.config = value
	return s
}

func (s *ExternalLoggingCreateService) TrustCertificates(value bool) *ExternalLoggingCreateService {
	s.trustCertificates = &value
	return s
}

func (s *ExternalLoggingCreateService) TrustFingerprints(value bool) *ExternalLoggingCreateService {
	s.trustFingerprints = &value
	return s
}

func (s *ExternalLoggingCreateService) RunSetupTests(value bool) *ExternalLoggingCreateService {
	s.runSetupTests = &value
	return s
}

func (s *ExternalLoggingCreateService) Do(ctx context.Context) (ExternalLoggingCreateResponse, error) {
	var response ExternalLoggingCreateResponse
	url := fmt.Sprintf("%v/external-logging", s.c.baseURL)
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
