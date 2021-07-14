package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// ConnectorCreateService implements the Connector Management, Create a Connector API.
// Ref. https://fivetran.com/docs/rest-api/connectors#createaconnector
type ConnectorCreateService struct {
	c                 *Client
	service           *string
	groupID           *string
	trustCertificates *bool
	trustFingerprints *bool
	runSetupTests     *bool
	paused            *bool
	config            *ConnectorConfig
	auth              *ConnectorAuth
}

type connectorCreateRequest struct {
	Service           *string                 `json:"service,omitempty"`
	GroupID           *string                 `json:"group_id,omitempty"`
	TrustCertificates *bool                   `json:"trust_certificates,omitempty"`
	TrustFingerprints *bool                   `json:"trust_fingerprints,omitempty"`
	RunSetupTests     *bool                   `json:"run_setup_tests,omitempty"`
	Paused            *bool                   `json:"paused,omitempty"`
	Config            *connectorConfigRequest `json:"config,omitempty"`
	Auth              *connectorAuthRequest   `json:"auth,omitempty"`
}

type ConnectorCreateResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID             string    `json:"id"`
		GroupID        string    `json:"group_id"`
		Service        string    `json:"service"`
		ServiceVersion int       `json:"service_version"`
		Schema         string    `json:"schema"`
		ConnectedBy    string    `json:"connected_by"`
		CreatedAt      time.Time `json:"created_at"`
		SucceededAt    time.Time `json:"succeeded_at"`
		FailedAt       time.Time `json:"failed_at"`
		SyncFrequency  int       `json:"sync_frequency"`
		ScheduleType   string    `json:"schedule_type"`
		Status         struct {
			SetupState       string `json:"setup_state"`
			SyncState        string `json:"sync_state"`
			UpdateState      string `json:"update_state"`
			IsHistoricalSync *bool  `json:"is_historical_sync"`
			Tasks            []struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			} `json:"tasks"`
			Warnings []struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			} `json:"warnings"`
		} `json:"status"`
		SetupTests []struct {
			Title   string `json:"title"`
			Status  string `json:"status"`
			Message string `json:"message"`
		} `json:"setup_tests"`
		Config ConnectorConfigResponse `json:"config"`
	} `json:"data"`
}

func (c *Client) NewConnectorCreate() *ConnectorCreateService {
	return &ConnectorCreateService{c: c}
}

func (s *ConnectorCreateService) request() *connectorCreateRequest {
	var config *connectorConfigRequest
	if s.config != nil {
		config = s.config.request()
	}

	var auth *connectorAuthRequest
	if s.auth != nil {
		auth = s.auth.request()
	}

	return &connectorCreateRequest{
		Service:           s.service,
		GroupID:           s.groupID,
		TrustCertificates: s.trustCertificates,
		TrustFingerprints: s.trustFingerprints,
		RunSetupTests:     s.runSetupTests,
		Paused:            s.paused,
		Config:            config,
		Auth:              auth,
	}
}

func (s *ConnectorCreateService) Service(value string) *ConnectorCreateService {
	s.service = &value
	return s
}

func (s *ConnectorCreateService) GroupID(value string) *ConnectorCreateService {
	s.groupID = &value
	return s
}

func (s *ConnectorCreateService) TrustCertificates(value bool) *ConnectorCreateService {
	s.trustCertificates = &value
	return s
}

func (s *ConnectorCreateService) TrustFingerprints(value bool) *ConnectorCreateService {
	s.trustFingerprints = &value
	return s
}

func (s *ConnectorCreateService) RunSetupTests(value bool) *ConnectorCreateService {
	s.runSetupTests = &value
	return s
}

func (s *ConnectorCreateService) Paused(value bool) *ConnectorCreateService {
	s.paused = &value
	return s
}

func (s *ConnectorCreateService) Config(value *ConnectorConfig) *ConnectorCreateService {
	s.config = value
	return s
}

func (s *ConnectorCreateService) Auth(value *ConnectorAuth) *ConnectorCreateService {
	s.auth = value
	return s
}

func (s *ConnectorCreateService) Do(ctx context.Context) (ConnectorCreateResponse, error) {
	var response ConnectorCreateResponse
	url := fmt.Sprintf("%v/connectors", s.c.baseURL)
	expectedStatus := 201

	headers := make(map[string]string)
	headers["Authorization"] = s.c.authorization
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json;version=2"

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
