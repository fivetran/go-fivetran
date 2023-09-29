package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	externallogging "github.com/fivetran/go-fivetran/external_logging"
	httputils "github.com/fivetran/go-fivetran/http_utils"
	"github.com/fivetran/go-fivetran/utils"
)

// ExternalLoggingModifyService implements the Log Management, Modify a Log Service API.
// Ref. https://fivetran.com/docs/rest-api/log-service-management#updatealogservice
type ExternalLoggingModifyService struct {
	c                 *Client
	externalLoggingId *string
	enabled           *bool
	runSetupTests     *bool
	config            *externallogging.ExternalLoggingConfig
	configCustom      *map[string]interface{}
}

type externalLoggingModifyRequestBase struct {
	Enabled       *bool `json:"enabled,omitempty"`
	RunSetupTests *bool `json:"run_setup_tests,omitempty"`
}

type externalLoggingModifyRequest struct {
	externalLoggingModifyRequestBase
	Config any `json:"config,omitempty"`
}

type externalLoggingCustomModifyRequest struct {
	externalLoggingModifyRequestBase
	Config *map[string]interface{} `json:"config,omitempty"`
}

type ExternalLoggingModifyResponseDataBased struct {
	Id      string `json:"id"`
	Service string `json:"service"`
	Enabled bool   `json:"enabled"`
}

type ExternalLoggingModifyResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ExternalLoggingModifyResponseDataBased
		Config externallogging.ExternalLoggingConfigResponse `json:"config"`
	} `json:"data"`
}

type ExternalLoggingModifyCustomResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ExternalLoggingModifyResponseDataBased
		Config map[string]interface{} `json:"config"`
	} `json:"data"`
}

type ExternalLoggingModifyCustomMergedResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ExternalLoggingModifyResponseDataBased
		CustomConfig map[string]interface{}                        `json:"config"`
		Config       externallogging.ExternalLoggingConfigResponse // no mapping here
	} `json:"data"`
}

func (c *Client) NewExternalLoggingModify() *ExternalLoggingModifyService {
	return &ExternalLoggingModifyService{c: c}
}

func (s *ExternalLoggingModifyService) requestBase() externalLoggingModifyRequestBase {
	return externalLoggingModifyRequestBase{
		Enabled: s.enabled,
	}
}

func (s *ExternalLoggingModifyService) request() *externalLoggingModifyRequest {
	var config interface{}

	if s.config != nil {
		config = s.config.Request()
	}

	return &externalLoggingModifyRequest{
		externalLoggingModifyRequestBase: s.requestBase(),
		Config:                           config,
	}
}

func (s *ExternalLoggingModifyService) requestCustom() *externalLoggingCustomModifyRequest {
	return &externalLoggingCustomModifyRequest{
		externalLoggingModifyRequestBase: s.requestBase(),
		Config:                           s.configCustom,
	}
}

func (s *ExternalLoggingModifyService) requestCustomMerged() (*externalLoggingCustomModifyRequest, error) {
	currentConfig := s.configCustom

	if s.config != nil {
		var err error
		currentConfig, err = s.config.Merge(currentConfig)
		if err != nil {
			return nil, err
		}
	}

	return &externalLoggingCustomModifyRequest{
		externalLoggingModifyRequestBase: s.requestBase(),
		Config:                           currentConfig,
	}, nil
}

func (s *ExternalLoggingModifyService) ExternalLoggingId(value string) *ExternalLoggingModifyService {
	s.externalLoggingId = &value
	return s
}

func (s *ExternalLoggingModifyService) Enabled(value bool) *ExternalLoggingModifyService {
	s.enabled = &value
	return s
}

func (s *ExternalLoggingModifyService) Config(value *externallogging.ExternalLoggingConfig) *ExternalLoggingModifyService {
	s.config = value
	return s
}

func (s *ExternalLoggingModifyService) ConfigCustom(value *map[string]interface{}) *ExternalLoggingModifyService {
	s.configCustom = value
	return s
}

func (s *ExternalLoggingModifyService) RunSetupTests(value bool) *ExternalLoggingModifyService {
	s.runSetupTests = &value
	return s
}

func (s *ExternalLoggingModifyService) do(ctx context.Context, req, response any) error {
	if s.externalLoggingId == nil {
		return fmt.Errorf("missing required externalLoggingId")
	}

	url := fmt.Sprintf("%v/external-logging/%v", s.c.baseURL, *s.externalLoggingId)
	expectedStatus := 200

	headers := s.c.commonHeaders()
	headers["Content-Type"] = "application/json"
	headers["Accept"] = restAPIv2

	reqBody, err := json.Marshal(req)
	if err != nil {
		return err
	}

	r := httputils.Request{
		Method:           "PATCH",
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

func (s *ExternalLoggingModifyService) Do(ctx context.Context) (externallogging.ExternalLoggingResponse, error) {
	var response externallogging.ExternalLoggingResponse

	err := s.do(ctx, s.request(), &response)

	return response, err
}

func (s *ExternalLoggingModifyService) DoCustom(ctx context.Context) (externallogging.ExternalLoggingCustomResponse, error) {
	var response externallogging.ExternalLoggingCustomResponse

	err := s.do(ctx, s.requestCustom(), &response)

	return response, err
}

func (s *ExternalLoggingModifyService) DoCustomMerged(ctx context.Context) (externallogging.ExternalLoggingCustomMergedResponse, error) {
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
