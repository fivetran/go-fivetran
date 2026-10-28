package metadata

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type DestinationMetadataDetailsService struct {
	httputils.HttpService
	service *string
}

func (s *DestinationMetadataDetailsService) Service(value string) *DestinationMetadataDetailsService {
	s.service = &value
	return s
}

func (s *DestinationMetadataDetailsService) Do(ctx context.Context) (DestinationMetadataResponse, error) {
	return s.do(ctx)
}

func (s *DestinationMetadataDetailsService) DoWithUserAgentSuffix(ctx context.Context, userAgentSuffix string) (DestinationMetadataResponse, error) {
	return s.doWithUserAgentSuffix(ctx, userAgentSuffix)
}

func (s *DestinationMetadataDetailsService) do(ctx context.Context) (DestinationMetadataResponse, error) {
	var response DestinationMetadataResponse
	if s.service == nil {
		return response, fmt.Errorf("missing required service")
	}

	url := fmt.Sprintf("/metadata/destination-types/%v", *s.service)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}

func (s *DestinationMetadataDetailsService) doWithUserAgentSuffix(ctx context.Context, userAgentSuffix string) (DestinationMetadataResponse, error) {
	var response DestinationMetadataResponse
	if s.service == nil {
		return response, fmt.Errorf("missing required service")
	}

	url := fmt.Sprintf("/metadata/destination-types/%v", *s.service)
	err := s.HttpService.DoWithUserAgentSuffix(ctx, userAgentSuffix, "GET", url, nil, nil, 200, &response)
	return response, err
}
