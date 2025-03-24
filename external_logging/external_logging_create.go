package externallogging

import (
	"context"

	httputils "github.com/fivetran/go-fivetran/http_utils"
	"github.com/fivetran/go-fivetran/utils"
)

type ExternalLoggingCreateService struct {
	httputils.HttpService
	groupId      *string
	service      *string
	enabled      *bool
	config       *ExternalLoggingConfig
	configCustom *map[string]interface{}
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

func (s *ExternalLoggingCreateService) Config(value *ExternalLoggingConfig) *ExternalLoggingCreateService {
	s.config = value
	return s
}

func (s *ExternalLoggingCreateService) ConfigCustom(value *map[string]interface{}) *ExternalLoggingCreateService {
	s.configCustom = value
	return s
}

func (s *ExternalLoggingCreateService) do(ctx context.Context, req, response any) error {
	err := s.HttpService.Do(ctx, "POST", "/external-logging", req, nil, 201, &response)
	return err
}

func (s *ExternalLoggingCreateService) Do(ctx context.Context) (ExternalLoggingResponse, error) {
	var response ExternalLoggingResponse

	err := s.do(ctx, s.request(), &response)

	return response, err
}

func (s *ExternalLoggingCreateService) DoCustom(ctx context.Context) (ExternalLoggingCustomResponse, error) {
	var response ExternalLoggingCustomResponse

	err := s.do(ctx, s.requestCustom(), &response)

	return response, err
}

func (s *ExternalLoggingCreateService) DoCustomMerged(ctx context.Context) (ExternalLoggingCustomMergedResponse, error) {
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
