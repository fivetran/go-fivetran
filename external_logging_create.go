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
	enabled 		  *bool
	config            *ExternalLoggingConfig
}

type externalLoggingCreateRequest struct {
	GroupID           *string                  		`json:"group_id,omitempty"`
	Service           *string                  		`json:"service,omitempty"`
	Enabled 		  *bool							`json:"enabled,omitempty"` 				
	Config            *externalLoggingConfigRequest `json:"config,omitempty"`
}

type ExternalLoggingCreateResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID             string `json:"id"`
		Service        string `json:"service"`
		Enabled        bool   `json:"enabled"`
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
		Enabled:           s.enabled,
		Config:            config
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

func (s *ExternalLoggingCreateService) Enabled(value bool) *ExternalLoggingCreateService {
	s.enabled = &value
	return s
}

func (s *ExternalLoggingCreateService) Config(value *ExternalLoggingConfig) *ExternalLoggingCreateService {
	s.config = value
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
