package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// F stands for Field
// needs to be exported because of json.Marshal()
type ConnectorCreateService struct {
	c                  *Client
	Fservice           *string          `json:"service,omitempty"`
	FgroupID           *string          `json:"group_id,omitempty"`
	FtrustCertificates *bool            `json:"trust_certificates,omitempty"`
	FtrustFingerprints *bool            `json:"trust_fingerprints,omitempty"`
	FrunSetupTests     *bool            `json:"run_setup_tests,omitempty"`
	Fpaused            *bool            `json:"paused,omitempty"`
	Fconfig            *ConnectorConfig `json:"config,omitempty"`
	Fauth              *ConnectorAuth   `json:"auth,omitempty"`
}

type ConnectorCreate struct {
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
		Config ConnectorConfig `json:"config"`
	} `json:"data"`
}

func (c *Client) NewConnectorCreateService() *ConnectorCreateService {
	return &ConnectorCreateService{c: c}
}

func (s *ConnectorCreateService) Service(service string) *ConnectorCreateService {
	s.Fservice = &service
	return s
}

func (s *ConnectorCreateService) GroupID(groupID string) *ConnectorCreateService {
	s.FgroupID = &groupID
	return s
}

func (s *ConnectorCreateService) TrustCertificates(trustCertificates bool) *ConnectorCreateService {
	s.FtrustCertificates = &trustCertificates
	return s
}

func (s *ConnectorCreateService) TrustFingerprints(trustFingerprints bool) *ConnectorCreateService {
	s.FtrustFingerprints = &trustFingerprints
	return s
}

func (s *ConnectorCreateService) RunSetupTests(runSetupTests bool) *ConnectorCreateService {
	s.FrunSetupTests = &runSetupTests
	return s
}

func (s *ConnectorCreateService) Paused(paused bool) *ConnectorCreateService {
	s.Fpaused = &paused
	return s
}

func (s *ConnectorCreateService) Config(config *ConnectorConfig) *ConnectorCreateService {
	s.Fconfig = config
	return s
}

func (s *ConnectorCreateService) Auth(auth *ConnectorAuth) *ConnectorCreateService {
	s.Fauth = auth
	return s
}

func (s *ConnectorCreateService) Do(ctx context.Context) (ConnectorCreate, error) {
	url := fmt.Sprintf("%v/connectors", s.c.baseURL)
	expectedStatus := 201
	headers := make(map[string]string)

	headers["Authorization"] = s.c.authorization
	headers["Content-Type"] = "application/json"

	reqBody, err := json.Marshal(s)
	if err != nil {
		return ConnectorCreate{}, err
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
		return ConnectorCreate{}, err
	}

	var connectorCreate ConnectorCreate
	if err := json.Unmarshal(respBody, &connectorCreate); err != nil {
		return ConnectorCreate{}, err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected %v", respStatus, expectedStatus)
		return connectorCreate, err
	}

	return connectorCreate, nil
}
