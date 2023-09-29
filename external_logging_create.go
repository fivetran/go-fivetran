package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	externallogging "github.com/fivetran/go-fivetran/external_logging"
	"github.com/fivetran/go-fivetran/utils"
)

// ExternalLoggingCreateService implements the Log Management, Create a Log Service API.
// Ref. https://fivetran.com/docs/rest-api/log-service-management#createalogservice
type ExternalLoggingCreateService struct {
	c            *Client
	groupId      *string
	service      *string
	enabled      *bool
	config       *externallogging.ExternalLoggingConfig
	configCustom *map[string]interface{}
}

/* requests */
type externalLoggingCreateRequestBase struct {
	Id      *string `json:"id,omitempty"`
	GroupId *string `json:"group_id,omitempty"`
	Service *string `json:"service,omitempty"`
	Enabled *bool   `json:"enabled,omitempty"`
}

type externalLoggingCreateRequest struct {
	externalLoggingCreateRequestBase
	Config any `json:"config,omitempty"`
}

type externalLoggingCustomCreateRequest struct {
	externalLoggingCreateRequestBase
	Config *map[string]interface{} `json:"config,omitempty"`
}

/* responses */

func (c *Client) NewExternalLoggingCreate() *ExternalLoggingCreateService {
	return &ExternalLoggingCreateService{c: c}
}

func (s *ExternalLoggingCreateService) requestBase() externalLoggingCreateRequestBase {
	return externalLoggingCreateRequestBase{
		GroupId: s.groupId,
		Service: s.service,
		Enabled: s.enabled,
	}
}

func (s *ExternalLoggingCreateService) request() *externalLoggingCreateRequest {
	var config interface{}
	if s.config != nil {
		config = s.config.Request()
	}

	r := &externalLoggingCreateRequest{
		externalLoggingCreateRequestBase: s.requestBase(),
		Config:                           config,
	}

	return r
}

func (s *ExternalLoggingCreateService) requestCustom() *externalLoggingCustomCreateRequest {
	return &externalLoggingCustomCreateRequest{
		externalLoggingCreateRequestBase: s.requestBase(),
		Config:                           s.configCustom,
	}
}

func (s *ExternalLoggingCreateService) requestCustomMerged() (*externalLoggingCustomCreateRequest, error) {
	currentConfig := s.configCustom

	if s.config != nil {
		var err error
		currentConfig, err = s.config.Merge(currentConfig)
		if err != nil {
			return nil, err
		}
	}

	return &externalLoggingCustomCreateRequest{
		externalLoggingCreateRequestBase: s.requestBase(),
		Config:                           currentConfig,
	}, nil
}

func (s *ExternalLoggingCreateService) GroupId(value string) *ExternalLoggingCreateService {
	s.groupId = &value
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

func (s *ExternalLoggingCreateService) Config(value *externallogging.ExternalLoggingConfig) *ExternalLoggingCreateService {
	s.config = value
	return s
}

func (s *ExternalLoggingCreateService) ConfigCustom(value *map[string]interface{}) *ExternalLoggingCreateService {
	s.configCustom = value
	return s
}

func (s *ExternalLoggingCreateService) do(ctx context.Context, req, response any) error {
	url := fmt.Sprintf("%v/external-logging", s.c.baseURL)
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

	if err := json.Unmarshal(respBody, &response); err != nil {
		return err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected: %v", respStatus, expectedStatus)
		return err
	}

	return nil
}

func (s *ExternalLoggingCreateService) Do(ctx context.Context) (externallogging.ExternalLoggingResponse, error) {
	var response externallogging.ExternalLoggingResponse

	err := s.do(ctx, s.request(), &response)

	return response, err
}

func (s *ExternalLoggingCreateService) DoCustom(ctx context.Context) (externallogging.ExternalLoggingCustomResponse, error) {
	var response externallogging.ExternalLoggingCustomResponse

	err := s.do(ctx, s.requestCustom(), &response)

	return response, err
}

func (s *ExternalLoggingCreateService) DoCustomMerged(ctx context.Context) (externallogging.ExternalLoggingCustomMergedResponse, error) {
	var response externallogging.ExternalLoggingCustomMergedResponse

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
