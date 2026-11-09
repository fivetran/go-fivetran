package metadata

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type LogServiceMetadataDetailsService struct {
	httputils.HttpService
	service *string
}

func (s *LogServiceMetadataDetailsService) Service(value string) *LogServiceMetadataDetailsService {
	s.service = &value
	return s
}

func (s *LogServiceMetadataDetailsService) Do(ctx context.Context) (LogServiceMetadataResponse, error) {
	return s.do(ctx)
}

func (s *LogServiceMetadataDetailsService) DoWithUserAgentSuffix(ctx context.Context, userAgentSuffix string) (LogServiceMetadataResponse, error) {
	return s.doWithUserAgentSuffix(ctx, userAgentSuffix)
}

func (s *LogServiceMetadataDetailsService) do(ctx context.Context) (LogServiceMetadataResponse, error) {
	var response LogServiceMetadataResponse
	if s.service == nil {
		return response, fmt.Errorf("missing required service")
	}

	url := fmt.Sprintf("/metadata/log-service-types/%v", *s.service)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}

func (s *LogServiceMetadataDetailsService) doWithUserAgentSuffix(ctx context.Context, userAgentSuffix string) (LogServiceMetadataResponse, error) {
	var response LogServiceMetadataResponse
	if s.service == nil {
		return response, fmt.Errorf("missing required service")
	}

	url := fmt.Sprintf("/metadata/log-service-types/%v", *s.service)
	err := s.HttpService.DoWithUserAgentSuffix(ctx, userAgentSuffix, "GET", url, nil, nil, 200, &response)
	return response, err
}
