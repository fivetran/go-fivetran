package privatelink

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
	"github.com/fivetran/go-fivetran/utils"
)

// PrivateLinkModifyService implements the Private Link Management, Modify a Private Link Service API.
// Ref. https://fivetran.com/docs/rest-api/private-link-management#updateaprivatelink
type PrivateLinkModifyService struct {
	httputils.HttpService
	privateLinkId 	  *string
	config            *PrivateLinkConfig
	configCustom      *map[string]interface{}
}

func (s *PrivateLinkModifyService) request() *privateLinkModifyRequest {
	var config interface{}

	if s.config != nil {
		config = s.config.Request()
	}

	return &privateLinkModifyRequest{
		Config:                       config,
	}
}

func (s *PrivateLinkModifyService) requestCustom() *privateLinkCustomModifyRequest {
	return &privateLinkCustomModifyRequest{
		Config:                       s.configCustom,
	}
}

func (s *PrivateLinkModifyService) requestCustomMerged() (*privateLinkCustomModifyRequest, error) {
	currentConfig := s.configCustom

	if s.config != nil {
		var err error
		currentConfig, err = s.config.Merge(currentConfig)
		if err != nil {
			return nil, err
		}
	}

	return &privateLinkCustomModifyRequest{
		Config:                           currentConfig,
	}, nil
}

func (s *PrivateLinkModifyService) PrivateLinkId(value string) *PrivateLinkModifyService {
	s.privateLinkId = &value
	return s
}

func (s *PrivateLinkModifyService) Config(value *PrivateLinkConfig) *PrivateLinkModifyService {
	s.config = value
	return s
}

func (s *PrivateLinkModifyService) ConfigCustom(value *map[string]interface{}) *PrivateLinkModifyService {
	s.configCustom = value
	return s
}

func (s *PrivateLinkModifyService) do(ctx context.Context, req, response any) error {
	if s.privateLinkId == nil {
		return fmt.Errorf("missing required privateLinkId")
	}

	url := fmt.Sprintf("/private-links/%v", *s.privateLinkId)
	err := s.HttpService.Do(ctx, "PATCH", url, req, nil, 200, &response)
	return err
}

func (s *PrivateLinkModifyService) Do(ctx context.Context) (PrivateLinkResponse, error) {
	var response PrivateLinkResponse

	err := s.do(ctx, s.request(), &response)

	return response, err
}

func (s *PrivateLinkModifyService) DoCustom(ctx context.Context) (PrivateLinkCustomResponse, error) {
	var response PrivateLinkCustomResponse

	err := s.do(ctx, s.requestCustom(), &response)

	return response, err
}

func (s *PrivateLinkModifyService) DoCustomMerged(ctx context.Context) (PrivateLinkCustomMergedResponse, error) {
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