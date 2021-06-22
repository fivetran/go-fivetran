package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type connectorModifyService struct {
	c                 *Client
	connectorID       *string
	paused            *bool
	syncFrequency     *int
	dailySyncTime     *string
	config            *connectorConfig
	auth              *connectorAuth
	trustCertificates *bool
	trustFingerprints *bool
	isHistoricalSync  *bool
	scheduleType      *string
	runSetupTests     *bool
}

type connectorModifyRequest struct {
	Paused            *bool                   `json:"paused,omitempty"`
	SyncFrequency     *int                    `json:"sync_frequency,omitempty"`
	DailySyncTime     *string                 `json:"daily_sync_time,omitempty"`
	Config            *connectorConfigRequest `json:"config,omitempty"`
	Auth              *connectorAuthRequest   `json:"auth,omitempty"`
	TrustCertificates *bool                   `json:"trust_certificates,omitempty"`
	TrustFingerprints *bool                   `json:"trust_fingerprints,omitempty"`
	IsHistoricalSync  *bool                   `json:"is_historical_sync,omitempty"`
	ScheduleType      *string                 `json:"schedule_type,omitempty"`
	RunSetupTests     *bool                   `json:"run_setup_tests,omitempty"`
}

type ConnectorModifyResponse struct {
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
			IsHistoricalSync bool   `json:"is_historical_sync"`
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

func (c *Client) NewConnectorModify() *connectorModifyService {
	return &connectorModifyService{c: c}
}

func (s *connectorModifyService) request() *connectorModifyRequest {
	var config *connectorConfigRequest
	if s.config != nil {
		config = s.config.request()
	}

	var auth *connectorAuthRequest
	if s.auth != nil {
		auth = s.auth.request()
	}

	return &connectorModifyRequest{
		Paused:            s.paused,
		SyncFrequency:     s.syncFrequency,
		DailySyncTime:     s.dailySyncTime,
		Config:            config,
		Auth:              auth,
		TrustCertificates: s.trustCertificates,
		TrustFingerprints: s.trustFingerprints,
		IsHistoricalSync:  s.isHistoricalSync,
		ScheduleType:      s.scheduleType,
		RunSetupTests:     s.runSetupTests,
	}
}

func (s *connectorModifyService) ConnectorID(value string) *connectorModifyService {
	s.connectorID = &value
	return s
}

func (s *connectorModifyService) Paused(value bool) *connectorModifyService {
	s.paused = &value
	return s
}

func (s *connectorModifyService) SyncFrequency(value int) *connectorModifyService {
	s.syncFrequency = &value
	return s
}

func (s *connectorModifyService) DailySyncTime(value string) *connectorModifyService {
	s.dailySyncTime = &value
	return s
}

func (s *connectorModifyService) Config(value *connectorConfig) *connectorModifyService {
	s.config = value
	return s
}

func (s *connectorModifyService) Auth(value *connectorAuth) *connectorModifyService {
	s.auth = value
	return s
}

func (s *connectorModifyService) TrustCertificates(value bool) *connectorModifyService {
	s.trustCertificates = &value
	return s
}

func (s *connectorModifyService) TrustFingerprints(value bool) *connectorModifyService {
	s.trustFingerprints = &value
	return s
}

func (s *connectorModifyService) IsHistoricalSync(value bool) *connectorModifyService {
	s.isHistoricalSync = &value
	return s
}

func (s *connectorModifyService) ScheduleType(value string) *connectorModifyService {
	s.scheduleType = &value
	return s
}

func (s *connectorModifyService) RunSetupTests(value bool) *connectorModifyService {
	s.runSetupTests = &value
	return s
}

func (s *connectorModifyService) Do(ctx context.Context) (ConnectorModifyResponse, error) {
	var response ConnectorModifyResponse

	if s.connectorID == nil {
		return response, fmt.Errorf("missing required ConnectorID")
	}

	url := fmt.Sprintf("%v/connectors/%v", s.c.baseURL, *s.connectorID)
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
