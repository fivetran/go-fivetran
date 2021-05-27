package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// F stands for Field
// needs to be exported because of json.Marshal()
type ConnectorModifyService struct {
	c                  *Client
	connectorID        *string
	Fpaused            *bool            `json:"paused,omitempty"`
	FsyncFrequency     *int             `json:"sync_frequency,omitempty"`
	FdailySyncTime     *string          `json:"daily_sync_time,omitempty"`
	Fconfig            *ConnectorConfig `json:"config,omitempty"`
	Fauth              *ConnectorAuth   `json:"auth,omitempty"`
	FtrustCertificates *bool            `json:"trust_certificates,omitempty"`
	FtrustFingerprints *bool            `json:"trust_fingerprints,omitempty"`
	FisHistoricalSync  *bool            `json:"is_historical_sync,omitempty"`
	FscheduleType      *string          `json:"schedule_type,omitempty"`
	FrunSetupTests     *bool            `json:"run_setup_tests,omitempty"`
}

type ConnectorModify struct {
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

func (c *Client) NewConnectorModifyService() *ConnectorModifyService {
	return &ConnectorModifyService{c: c}
}

func (s *ConnectorModifyService) ConnectorID(connectorID string) *ConnectorModifyService {
	s.connectorID = &connectorID
	return s
}

func (s *ConnectorModifyService) Paused(paused bool) *ConnectorModifyService {
	s.Fpaused = &paused
	return s
}

func (s *ConnectorModifyService) SyncFrequency(syncFrequency int) *ConnectorModifyService {
	s.FsyncFrequency = &syncFrequency
	return s
}

func (s *ConnectorModifyService) DailySyncTime(dailySyncTime string) *ConnectorModifyService {
	s.FdailySyncTime = &dailySyncTime
	return s
}

func (s *ConnectorModifyService) Config(config *ConnectorConfig) *ConnectorModifyService {
	s.Fconfig = config
	return s
}

func (s *ConnectorModifyService) Auth(auth *ConnectorAuth) *ConnectorModifyService {
	s.Fauth = auth
	return s
}

func (s *ConnectorModifyService) TrustCertificates(trustCertificates bool) *ConnectorModifyService {
	s.FtrustCertificates = &trustCertificates
	return s
}

func (s *ConnectorModifyService) TrustFingerprints(trustFingerprints bool) *ConnectorModifyService {
	s.FtrustFingerprints = &trustFingerprints
	return s
}

func (s *ConnectorModifyService) IsHistoricalSync(isHistoricalSync bool) *ConnectorModifyService {
	s.FisHistoricalSync = &isHistoricalSync
	return s
}

func (s *ConnectorModifyService) ScheduleType(scheduleType string) *ConnectorModifyService {
	s.FscheduleType = &scheduleType
	return s
}

func (s *ConnectorModifyService) RunSetupTests(runSetupTests bool) *ConnectorModifyService {
	s.FrunSetupTests = &runSetupTests
	return s
}

func (s *ConnectorModifyService) Do(ctx context.Context) (ConnectorModify, error) {
	if s.connectorID == nil { // we don't validate business rules (unless it is strictly necessary)
		err := fmt.Errorf("missing required ConnectorID")
		return ConnectorModify{}, err
	}

	url := fmt.Sprintf("%v/connectors/%v", s.c.baseURL, *s.connectorID)
	expectedStatus := 200
	headers := make(map[string]string)

	headers["Authorization"] = s.c.authorization
	headers["Content-Type"] = "application/json"

	reqBody, err := json.Marshal(s)
	if err != nil {
		return ConnectorModify{}, err
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
		return ConnectorModify{}, err
	}

	var connectorModify ConnectorModify
	if err := json.Unmarshal(respBody, &connectorModify); err != nil {
		return ConnectorModify{}, err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected %v", respStatus, expectedStatus)
		return connectorModify, err
	}

	return connectorModify, nil
}
