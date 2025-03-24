package privatelink

import (
	"context"

	httputils "github.com/fivetran/go-fivetran/http_utils"
	"github.com/fivetran/go-fivetran/utils"
)

type PrivateLinkCreateService struct {
	httputils.HttpService
	name      	 *string
	region       *string
	service      *string
	config       *PrivateLinkConfig
	configCustom *map[string]interface{}
}


func (s *PrivateLinkCreateService) requestBase() privateLinkCreateRequestBase {
	return privateLinkCreateRequestBase{
		Name: 		s.name,
		Region: 	s.region,
		Service: 	s.service,
	}
}

func (s *PrivateLinkCreateService) request() *privateLinkCreateRequest {
	var config interface{}
	if s.config != nil {
		config = s.config.Request()
	}

	r := &privateLinkCreateRequest{
		privateLinkCreateRequestBase: s.requestBase(),
		Config:                       config,
	}

	return r
}

func (s *PrivateLinkCreateService) requestCustom() *privateLinkCustomCreateRequest {
	return &privateLinkCustomCreateRequest{
		privateLinkCreateRequestBase: s.requestBase(),
		Config:                       s.configCustom,
	}
}

func (s *PrivateLinkCreateService) requestCustomMerged() (*privateLinkCustomCreateRequest, error) {
	currentConfig := s.configCustom

	if s.config != nil {
		var err error
		currentConfig, err = s.config.Merge(currentConfig)
		if err != nil {
			return nil, err
		}
	}

	return &privateLinkCustomCreateRequest{
		privateLinkCreateRequestBase: s.requestBase(),
		Config:                       currentConfig,
	}, nil
}

func (s *PrivateLinkCreateService) Region(value string) *PrivateLinkCreateService {
	s.region = &value
	return s
}

func (s *PrivateLinkCreateService) Service(value string) *PrivateLinkCreateService {
	s.service = &value
	return s
}

func (s *PrivateLinkCreateService) Name(value string) *PrivateLinkCreateService {
	s.name = &value
	return s
}

func (s *PrivateLinkCreateService) Config(value *PrivateLinkConfig) *PrivateLinkCreateService {
	s.config = value
	return s
}

func (s *PrivateLinkCreateService) ConfigCustom(value *map[string]interface{}) *PrivateLinkCreateService {
	s.configCustom = value
	return s
}

func (s *PrivateLinkCreateService) do(ctx context.Context, req, response any) error {
	err := s.HttpService.Do(ctx, "POST", "/private-links", req, nil, 201, &response)
	return err
}

func (s *PrivateLinkCreateService) Do(ctx context.Context) (PrivateLinkResponse, error) {
	var response PrivateLinkResponse

	err := s.do(ctx, s.request(), &response)

	return response, err
}

func (s *PrivateLinkCreateService) DoCustom(ctx context.Context) (PrivateLinkCustomResponse, error) {
	var response PrivateLinkCustomResponse

	err := s.do(ctx, s.requestCustom(), &response)

	return response, err
}

func (s *PrivateLinkCreateService) DoCustomMerged(ctx context.Context) (PrivateLinkCustomMergedResponse, error) {
	var response PrivateLinkCustomMergedResponse

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