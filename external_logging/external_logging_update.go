package externallogging

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
	"github.com/fivetran/go-fivetran/utils"
)

type ExternalLoggingUpdateService struct {
	httputils.HttpService
	externalLoggingId *string
	enabled           *bool
	runSetupTests     *bool
	config            *ExternalLoggingConfig
	configCustom      *map[string]interface{}
}

func (s *ExternalLoggingUpdateService) requestBase() externalLoggingUpdateRequestBase {
	return externalLoggingUpdateRequestBase{
		Enabled: s.enabled,
	}
}

func (s *ExternalLoggingUpdateService) request() *externalLoggingUpdateRequest {
	var config interface{}

	if s.config != nil {
		config = s.config.Request()
	}

	return &externalLoggingUpdateRequest{
		externalLoggingUpdateRequestBase: s.requestBase(),
		Config:                           config,
	}
}

func (s *ExternalLoggingUpdateService) requestCustom() *externalLoggingCustomUpdateRequest {
	return &externalLoggingCustomUpdateRequest{
		externalLoggingUpdateRequestBase: s.requestBase(),
		Config:                           s.configCustom,
	}
}

func (s *ExternalLoggingUpdateService) requestCustomMerged() (*externalLoggingCustomUpdateRequest, error) {
	currentConfig := s.configCustom

	if s.config != nil {
		var err error
		currentConfig, err = s.config.Merge(currentConfig)
		if err != nil {
			return nil, err
		}
	}

	return &externalLoggingCustomUpdateRequest{
		externalLoggingUpdateRequestBase: s.requestBase(),
		Config:                           currentConfig,
	}, nil
}

func (s *ExternalLoggingUpdateService) ExternalLoggingId(value string) *ExternalLoggingUpdateService {
	s.externalLoggingId = &value
	return s
}

func (s *ExternalLoggingUpdateService) Enabled(value bool) *ExternalLoggingUpdateService {
	s.enabled = &value
	return s
}

func (s *ExternalLoggingUpdateService) Config(value *ExternalLoggingConfig) *ExternalLoggingUpdateService {
	s.config = value
	return s
}

func (s *ExternalLoggingUpdateService) ConfigCustom(value *map[string]interface{}) *ExternalLoggingUpdateService {
	s.configCustom = value
	return s
}

func (s *ExternalLoggingUpdateService) RunSetupTests(value bool) *ExternalLoggingUpdateService {
	s.runSetupTests = &value
	return s
}

func (s *ExternalLoggingUpdateService) do(ctx context.Context, req, response any) error {
	if s.externalLoggingId == nil {
		return fmt.Errorf("missing required externalLoggingId")
	}

	url := fmt.Sprintf("/external-logging/%v", *s.externalLoggingId)
	err := s.HttpService.Do(ctx, "PATCH", url, req, nil, 200, &response)
	return err
}

func (s *ExternalLoggingUpdateService) Do(ctx context.Context) (ExternalLoggingResponse, error) {
	var response ExternalLoggingResponse

	err := s.do(ctx, s.request(), &response)

	return response, err
}

func (s *ExternalLoggingUpdateService) DoCustom(ctx context.Context) (ExternalLoggingCustomResponse, error) {
	var response ExternalLoggingCustomResponse

	err := s.do(ctx, s.requestCustom(), &response)

	return response, err
}

func (s *ExternalLoggingUpdateService) DoCustomMerged(ctx context.Context) (ExternalLoggingCustomMergedResponse, error) {
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
