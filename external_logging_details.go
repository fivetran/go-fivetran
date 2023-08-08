package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

// ExternalLoggingDetailsService implements the Log Management, Retrieve Log Service details API.
// Ref. https://fivetran.com/docs/rest-api/log-service-management#retrievelogservicedetails
type ExternalLoggingDetailsService struct {
	c             *Client
	externalLoggingID *string
}

type ExternalLoggingDetailsResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID             string                    `json:"id"`
		GroupID        string                    `json:"group_id"`
		Service        string                    `json:"service"`
		Region         string                    `json:"region"`
		TimeZoneOffset string                    `json:"time_zone_offset"`
		SetupStatus    string                    `json:"setup_status"`
		Config         ExternalLoggingConfigResponse `json:"config"`
	} `json:"data"`
}

func (c *Client) NewExternalLoggingDetails() *ExternalLoggingDetailsService {
	return &ExternalLoggingDetailsService{c: c}
}

func (s *ExternalLoggingDetailsService) ExternalLoggingID(value string) *ExternalLoggingDetailsService {
	s.externalLoggingID = &value
	return s
}

func (s *ExternalLoggingDetailsService) Do(ctx context.Context) (ExternalLoggingDetailsResponse, error) {
	var response ExternalLoggingDetailsResponse

	if s.externalLoggingID == nil {
		return response, fmt.Errorf("missing required ExternalLoggingID")
	}

	url := fmt.Sprintf("%v/external-logging/%v", s.c.baseURL, *s.externalLoggingID)
	expectedStatus := 200

	headers := s.c.commonHeaders()
	headers["Accept"] = restAPIv2

	r := request{
		method:  "GET",
		url:     url,
		body:    nil,
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
