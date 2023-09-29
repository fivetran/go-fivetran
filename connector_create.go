package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/connectors"
	httputils "github.com/fivetran/go-fivetran/http_utils"
	"github.com/fivetran/go-fivetran/utils"
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
	config            *connectors.ConnectorConfig
	auth              *connectors.ConnectorAuth
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
	Config any `json:"config,omitempty"`
	Auth   any `json:"auth,omitempty"`
}

type connectorCustomCreateRequest struct {
	connectorCreateRequestBase
	Config *map[string]interface{} `json:"config,omitempty"`
	Auth   *map[string]interface{} `json:"auth,omitempty"`
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
	var config interface{}
	if s.config != nil {
		config = s.config.Request()
	}

	var auth interface{}
	if s.auth != nil {
		auth = s.auth.Request()
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
		currentConfig, err = s.config.Merge(currentConfig)
		if err != nil {
			return nil, err
		}
	}

	currentAuth := s.authCustom
	if s.auth != nil {
		var err error
		currentAuth, err = s.auth.Merge(currentAuth)
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

func (s *ConnectorCreateService) Config(value *connectors.ConnectorConfig) *ConnectorCreateService {
	s.config = value
	return s
}

func (s *ConnectorCreateService) ConfigCustom(value *map[string]interface{}) *ConnectorCreateService {
	s.configCustom = value
	return s
}

func (s *ConnectorCreateService) Auth(value *connectors.ConnectorAuth) *ConnectorCreateService {
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

	r := httputils.Request{
		Method:           "POST",
		Url:              url,
		Body:             reqBody,
		Queries:          nil,
		Headers:          headers,
		Client:           s.c.httpClient,
		HandleRateLimits: s.c.handleRateLimits,
		MaxRetryAttempts: s.c.maxRetryAttempts,
	}

	respBody, respStatus, err := r.Do(ctx)
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

func (s *ConnectorCreateService) Do(ctx context.Context) (connectors.DetailsWithConfigResponse, error) {
	var response connectors.DetailsWithConfigResponse

	err := s.do(ctx, s.request(), &response)

	return response, err
}

func (s *ConnectorCreateService) DoCustom(ctx context.Context) (connectors.DetailsWithCustomConfigResponse, error) {
	var response connectors.DetailsWithCustomConfigResponse

	err := s.do(ctx, s.requestCustom(), &response)

	return response, err
}

func (s *ConnectorCreateService) DoCustomMerged(ctx context.Context) (connectors.DetailsWithCustomMergedConfigResponse, error) {
	var response connectors.DetailsWithCustomMergedConfigResponse

	req, err := s.requestCustomMerged()

	if err != nil {
		return response, err
	}

	err = s.do(ctx, req, &response)

	if err == nil {
		err = utils.FetchFromMap(&response.Data.CustomConfig, &response.Data.Config)
	}

	return response, err
}
