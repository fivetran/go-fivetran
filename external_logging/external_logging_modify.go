package externallogging

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
	"github.com/fivetran/go-fivetran/utils"
)

// ExternalLoggingModifyService implements the Log Management, Modify a Log Service API.
// Ref. https://fivetran.com/docs/rest-api/log-service-management#updatealogservice
type ExternalLoggingModifyService struct {
	httputils.HttpService
	externalLoggingId *string
	enabled           *bool
	runSetupTests     *bool
	config            *ExternalLoggingConfig
	configCustom      *map[string]interface{}
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

func (s *ExternalLoggingModifyService) Config(value *ExternalLoggingConfig) *ExternalLoggingModifyService {
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

	url := fmt.Sprintf("/external-logging/%v", *s.externalLoggingId)
	err := s.HttpService.Do(ctx, "PATCH", url, req, nil, 200, &response)
	return err
}

func (s *ExternalLoggingModifyService) Do(ctx context.Context) (ExternalLoggingResponse, error) {
	var response ExternalLoggingResponse

	err := s.do(ctx, s.request(), &response)

	return response, err
}

func (s *ExternalLoggingModifyService) DoCustom(ctx context.Context) (ExternalLoggingCustomResponse, error) {
	var response ExternalLoggingCustomResponse

	err := s.do(ctx, s.requestCustom(), &response)

	return response, err
}

func (s *ExternalLoggingModifyService) DoCustomMerged(ctx context.Context) (ExternalLoggingCustomMergedResponse, error) {
	var response ExternalLoggingCustomMergedResponse

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
