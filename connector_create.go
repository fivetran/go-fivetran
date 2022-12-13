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
	syncFrequency     *int
	dailySyncTime     *string
	pauseAfterTrial   *bool
	config            *ConnectorConfig
	auth              *ConnectorAuth
	configCustom      *map[string]interface{}
	authCustom        *map[string]interface{}
}

type connectorCreateRequestBase struct {
	Service           *string `json:"service,omitempty"`
	GroupID           *string `json:"group_id,omitempty"`
	TrustCertificates *bool   `json:"trust_certificates,omitempty"`
	TrustFingerprints *bool   `json:"trust_fingerprints,omitempty"`
	RunSetupTests     *bool   `json:"run_setup_tests,omitempty"`
	Paused            *bool   `json:"paused,omitempty"`
	SyncFrequency     *int    `json:"sync_frequency,omitempty"`
	DailySyncTime     *string `json:"daily_sync_time,omitempty"`
	PauseAfterTrial   *bool   `json:"pause_after_trial,omitempty"`
}

type connectorCreateRequest struct {
	connectorCreateRequestBase
	Config *connectorConfigRequest `json:"config,omitempty"`
	Auth   *connectorAuthRequest   `json:"auth,omitempty"`
}

type connectorCustomCreateRequest struct {
	connectorCreateRequestBase
	Config *map[string]interface{} `json:"config,omitempty"`
	Auth   *map[string]interface{} `json:"auth,omitempty"`
}

type ConnectorCreateResponseDataBase struct {
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
	ScheduleType    string    `json:"schedule_type"`
	Paused          *bool     `json:"paused"`
	PauseAfterTrial *bool     `json:"pause_after_trial"`
	DailySyncTime   string    `json:"daily_sync_time"`
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

type ConnectorCreateResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ConnectorCreateResponseDataBase
		Config ConnectorConfigResponse `json:"config"`
	} `json:"data"`
}

type ConnectorCustomCreateResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ConnectorCreateResponseDataBase
		Config map[string]interface{} `json:"config"`
	} `json:"data"`
}

type ConnectorCustomMergedCreateResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ConnectorCreateResponseDataBase
		CustomConfig map[string]interface{}  `json:"config"`
		Config       ConnectorConfigResponse // no mapping here
	} `json:"data"`
}

func (c *Client) NewConnectorCreate() *ConnectorCreateService {
	return &ConnectorCreateService{c: c}
}

func (s *ConnectorCreateService) requestBase() connectorCreateRequestBase {
	return connectorCreateRequestBase{
		Service:           s.service,
		GroupID:           s.groupID,
		TrustCertificates: s.trustCertificates,
		TrustFingerprints: s.trustFingerprints,
		RunSetupTests:     s.runSetupTests,
		Paused:            s.paused,
		SyncFrequency:     s.syncFrequency,
		DailySyncTime:     s.dailySyncTime,
		PauseAfterTrial:   s.pauseAfterTrial,
	}
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

	r := &connectorCreateRequest{
		connectorCreateRequestBase: s.requestBase(),
		Config:                     config,
		Auth:                       auth,
	}

	return r
}

func (s *ConnectorCreateService) requestCustom() *connectorCustomCreateRequest {
	return &connectorCustomCreateRequest{
		connectorCreateRequestBase: s.requestBase(),
		Config:                     s.configCustom,
		Auth:                       s.authCustom,
	}
}

func (s *ConnectorCreateService) requestCustomMerged() (*connectorCustomCreateRequest, error) {
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

	return &connectorCustomCreateRequest{
		connectorCreateRequestBase: s.requestBase(),
		Config:                     currentConfig,
		Auth:                       currentAuth,
	}, nil
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

func (s *ConnectorCreateService) SyncFrequency(value int) *ConnectorCreateService {
	s.syncFrequency = &value
	return s
}

func (s *ConnectorCreateService) DailySyncTime(value string) *ConnectorCreateService {
	s.dailySyncTime = &value
	return s
}

func (s *ConnectorCreateService) PauseAfterTrial(value bool) *ConnectorCreateService {
	s.pauseAfterTrial = &value
	return s
}

func (s *ConnectorCreateService) Config(value *ConnectorConfig) *ConnectorCreateService {
	s.config = value
	return s
}

func (s *ConnectorCreateService) ConfigCustom(value *map[string]interface{}) *ConnectorCreateService {
	s.configCustom = value
	return s
}

func (s *ConnectorCreateService) Auth(value *ConnectorAuth) *ConnectorCreateService {
	s.auth = value
	return s
}

func (s *ConnectorCreateService) AuthCustom(value *map[string]interface{}) *ConnectorCreateService {
	s.authCustom = value
	return s
}

func (s *ConnectorCreateService) do(ctx context.Context, req, response any) error {
	url := fmt.Sprintf("%v/connectors", s.c.baseURL)
	expectedStatus := 201

	headers := s.c.commonHeaders()
	headers["Content-Type"] = "application/json"
	headers["Accept"] = restAPIv2

	reqBody, err := json.Marshal(req)
	if err != nil {
		return err
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
		return err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected: %v", respStatus, expectedStatus)
		return err
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return err
	}

	return nil
}

func (s *ConnectorCreateService) Do(ctx context.Context) (ConnectorCreateResponse, error) {
	var response ConnectorCreateResponse

	err := s.do(ctx, s.request(), &response)

	return response, err
}

func (s *ConnectorCreateService) DoCustom(ctx context.Context) (ConnectorCustomCreateResponse, error) {
	var response ConnectorCustomCreateResponse

	err := s.do(ctx, s.requestCustom(), &response)

	return response, err
}

func (s *ConnectorCreateService) DoCustomMerged(ctx context.Context) (ConnectorCustomMergedCreateResponse, error) {
	var response ConnectorCustomMergedCreateResponse

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
