package privatelink

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
	"github.com/fivetran/go-fivetran/utils"
)

type PrivateLinkUpdateService struct {
	httputils.HttpService
	privateLinkId 	  *string
	config            *PrivateLinkConfig
	configCustom      *map[string]interface{}
}

func (s *PrivateLinkUpdateService) request() *privateLinkUpdateRequest {
	var config interface{}

	if s.config != nil {
		config = s.config.Request()
	}

	return &privateLinkUpdateRequest{
		Config:                       config,
	}
}

func (s *PrivateLinkUpdateService) requestCustom() *privateLinkCustomUpdateRequest {
	return &privateLinkCustomUpdateRequest{
		Config:                       s.configCustom,
	}
}

func (s *PrivateLinkUpdateService) requestCustomMerged() (*privateLinkCustomUpdateRequest, error) {
	currentConfig := s.configCustom

	if s.config != nil {
		var err error
		currentConfig, err = s.config.Merge(currentConfig)
		if err != nil {
			return nil, err
		}
	}

	return &privateLinkCustomUpdateRequest{
		Config:                           currentConfig,
	}, nil
}

func (s *PrivateLinkUpdateService) PrivateLinkId(value string) *PrivateLinkUpdateService {
	s.privateLinkId = &value
	return s
}

func (s *PrivateLinkUpdateService) Config(value *PrivateLinkConfig) *PrivateLinkUpdateService {
	s.config = value
	return s
}

func (s *PrivateLinkUpdateService) ConfigCustom(value *map[string]interface{}) *PrivateLinkUpdateService {
	s.configCustom = value
	return s
}

func (s *PrivateLinkUpdateService) do(ctx context.Context, req, response any) error {
	if s.privateLinkId == nil {
		return fmt.Errorf("missing required privateLinkId")
	}

	url := fmt.Sprintf("/private-links/%v", *s.privateLinkId)
	err := s.HttpService.Do(ctx, "PATCH", url, req, nil, 200, &response)
	return err
}

func (s *PrivateLinkUpdateService) Do(ctx context.Context) (PrivateLinkResponse, error) {
	var response PrivateLinkResponse

	err := s.do(ctx, s.request(), &response)

	return response, err
}

func (s *PrivateLinkUpdateService) DoCustom(ctx context.Context) (PrivateLinkCustomResponse, error) {
	var response PrivateLinkCustomResponse

	err := s.do(ctx, s.requestCustom(), &response)

	return response, err
}

func (s *PrivateLinkUpdateService) DoCustomMerged(ctx context.Context) (PrivateLinkCustomMergedResponse, error) {
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