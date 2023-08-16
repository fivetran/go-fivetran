package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// ConnectorModifyService implements the Connector Management, Modify a Connector API.
// Ref. https://fivetran.com/docs/rest-api/connectors#modifyaconnector
type ConnectorModifyService struct {
	c                 *Client
	connectorID       *string
	paused            *bool
	syncFrequency     *int
	dailySyncTime     *string
	trustCertificates *bool
	trustFingerprints *bool
	isHistoricalSync  *bool
	scheduleType      *string
	runSetupTests     *bool
	pauseAfterTrial   *bool
	config            *ConnectorConfig
	auth              *ConnectorAuth
	configCustom      *map[string]interface{}
	authCustom        *map[string]interface{}
}

type connectorModifyRequestBase struct {
	Paused            *bool   `json:"paused,omitempty"`
	SyncFrequency     *int    `json:"sync_frequency,omitempty"`
	DailySyncTime     *string `json:"daily_sync_time,omitempty"`
	TrustCertificates *bool   `json:"trust_certificates,omitempty"`
	TrustFingerprints *bool   `json:"trust_fingerprints,omitempty"`
	IsHistoricalSync  *bool   `json:"is_historical_sync,omitempty"`
	ScheduleType      *string `json:"schedule_type,omitempty"`
	RunSetupTests     *bool   `json:"run_setup_tests,omitempty"`
	PauseAfterTrial   *bool   `json:"pause_after_trial,omitempty"`
}

type connectorModifyRequest struct {
	connectorModifyRequestBase
	Config *connectorConfigRequest `json:"config,omitempty"`
	Auth   *connectorAuthRequest   `json:"auth,omitempty"`
}

type connectorCustomModifyRequest struct {
	connectorModifyRequestBase
	Config *map[string]interface{} `json:"config,omitempty"`
	Auth   *map[string]interface{} `json:"auth,omitempty"`
}

type ConnectorModifyResponseDataBase struct {
	ID              string    `json:"id"`
	GroupID         string    `json:"group_id"`
	Service         string    `json:"service"`
	ServiceVersion  *int      `json:"service_version"`
	Schema          string    `json:"schema"`
	ConnectedBy     string    `json:"connected_by"`
	CreatedAt       time.Time `json:"created_at"`
	SucceededAt     time.Time `json:"succeeded_at"`
	FailedAt        time.Time `json:"failed_at"`
	SyncFrequency   *int      `json:"sync_frequency"`
	Paused          *bool     `json:"paused"`
	PauseAfterTrial *bool     `json:"pause_after_trial"`
	ScheduleType    string    `json:"schedule_type"`
	Status          struct {
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
}

type ConnectorModifyResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ConnectorModifyResponseDataBase
		Config ConnectorConfigResponse `json:"config"`
	} `json:"data"`
}

type ConnectorCustomModifyResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ConnectorModifyResponseDataBase
		Config map[string]interface{} `json:"config"`
	} `json:"data"`
}

type ConnectorCustomMergedModifyResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ConnectorModifyResponseDataBase
		CustomConfig map[string]interface{}  `json:"config"`
		Config       ConnectorConfigResponse // no mapping here
	} `json:"data"`
}

func (c *Client) NewConnectorModify() *ConnectorModifyService {
	return &ConnectorModifyService{c: c}
}

func (s *ConnectorModifyService) requestBase() connectorModifyRequestBase {
	return connectorModifyRequestBase{
		Paused:            s.paused,
		SyncFrequency:     s.syncFrequency,
		DailySyncTime:     s.dailySyncTime,
		TrustCertificates: s.trustCertificates,
		TrustFingerprints: s.trustFingerprints,
		IsHistoricalSync:  s.isHistoricalSync,
		ScheduleType:      s.scheduleType,
		RunSetupTests:     s.runSetupTests,
		PauseAfterTrial:   s.pauseAfterTrial,
	}
}

func (s *ConnectorModifyService) request() *connectorModifyRequest {
	var config *connectorConfigRequest
	if s.config != nil {
		config = s.config.request()
	}

	var auth *connectorAuthRequest
	if s.auth != nil {
		auth = s.auth.request()
	}

	return &connectorModifyRequest{
		connectorModifyRequestBase: s.requestBase(),
		Config:                     config,
		Auth:                       auth,
	}
}

func (s *ConnectorModifyService) requestCustom() *connectorCustomModifyRequest {
	return &connectorCustomModifyRequest{
		connectorModifyRequestBase: s.requestBase(),
		Config:                     s.configCustom,
		Auth:                       s.authCustom,
	}
}

func (s *ConnectorModifyService) requestCustomMerged() (*connectorCustomModifyRequest, error) {
	currentConfig := s.configCustom

	if s.config != nil {
		var err error
		currentConfig, err = s.config.merge(currentConfig)
		if err != nil {
			return nil, err
		}
	}

	currentAuth := s.authCustom
	if s.auth != nil {
		var err error
		currentAuth, err = s.auth.merge(currentAuth)
		if err != nil {
			return nil, err
		}
	}

	return &connectorCustomModifyRequest{
		connectorModifyRequestBase: s.requestBase(),
		Config:                     currentConfig,
		Auth:                       currentAuth,
	}, nil
}

func (s *ConnectorModifyService) ConnectorID(value string) *ConnectorModifyService {
	s.connectorID = &value
	return s
}

func (s *ConnectorModifyService) Paused(value bool) *ConnectorModifyService {
	s.paused = &value
	return s
}

func (s *ConnectorModifyService) SyncFrequency(value int) *ConnectorModifyService {
	s.syncFrequency = &value
	return s
}

func (s *ConnectorModifyService) DailySyncTime(value string) *ConnectorModifyService {
	s.dailySyncTime = &value
	return s
}

func (s *ConnectorModifyService) Config(value *ConnectorConfig) *ConnectorModifyService {
	s.config = value
	return s
}

func (s *ConnectorModifyService) Auth(value *ConnectorAuth) *ConnectorModifyService {
	s.auth = value
	return s
}

func (s *ConnectorModifyService) ConfigCustom(value *map[string]interface{}) *ConnectorModifyService {
	s.configCustom = value
	return s
}

func (s *ConnectorModifyService) AuthCustom(value *map[string]interface{}) *ConnectorModifyService {
	s.authCustom = value
	return s
}

func (s *ConnectorModifyService) TrustCertificates(value bool) *ConnectorModifyService {
	s.trustCertificates = &value
	return s
}

func (s *ConnectorModifyService) TrustFingerprints(value bool) *ConnectorModifyService {
	s.trustFingerprints = &value
	return s
}

func (s *ConnectorModifyService) IsHistoricalSync(value bool) *ConnectorModifyService {
	s.isHistoricalSync = &value
	return s
}

func (s *ConnectorModifyService) ScheduleType(value string) *ConnectorModifyService {
	s.scheduleType = &value
	return s
}

func (s *ConnectorModifyService) RunSetupTests(value bool) *ConnectorModifyService {
	s.runSetupTests = &value
	return s
}

func (s *ConnectorModifyService) PauseAfterTrial(value bool) *ConnectorModifyService {
	s.pauseAfterTrial = &value
	return s
}

func (s *ConnectorModifyService) do(ctx context.Context, req, response any) error {

	if s.connectorID == nil {
		return fmt.Errorf("missing required ConnectorID")
	}

	url := fmt.Sprintf("%v/connectors/%v", s.c.baseURL, *s.connectorID)
	expectedStatus := 200

	headers := s.c.commonHeaders()
	headers["Content-Type"] = "application/json"
	headers["Accept"] = restAPIv2

	reqBody, err := json.Marshal(req)
	if err != nil {
		return err
	}

	r := request{
		method:           "PATCH",
		url:              url,
		body:             reqBody,
		queries:          nil,
		headers:          headers,
		client:           s.c.httpClient,
		handleRateLimits: s.c.handleRateLimits,
		maxRetryAttempts: s.c.maxRetryAttempts,
	}

	respBody, respStatus, err := r.httpRequest(ctx)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected: %v", respStatus, expectedStatus)
		return err
	}

	return nil
}

func (s *ConnectorModifyService) Do(ctx context.Context) (ConnectorModifyResponse, error) {
	var response ConnectorModifyResponse

	err := s.do(ctx, s.request(), &response)

	return response, err
}

func (s *ConnectorModifyService) DoCustom(ctx context.Context) (ConnectorCustomModifyResponse, error) {
	var response ConnectorCustomModifyResponse

	err := s.do(ctx, s.requestCustom(), &response)

	return response, err
}

func (s *ConnectorModifyService) DoCustomMerged(ctx context.Context) (ConnectorCustomMergedModifyResponse, error) {
	var response ConnectorCustomMergedModifyResponse

	req, err := s.requestCustomMerged()

	if err != nil {
		return response, err
	}

	err = s.do(ctx, req, &response)

	if err == nil {
		err = FetchFromMap(&response.Data.CustomConfig, &response.Data.Config)
	}

	return response, err
}
