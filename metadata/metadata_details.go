package metadata

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type MetadataDetailsService struct {
	httputils.HttpService
	service *string
}

func (s *MetadataDetailsService) Service(value string) *MetadataDetailsService {
	s.service = &value
	return s
}

func (s *MetadataDetailsService) Do(ctx context.Context) (ConnectorMetadataResponse, error) {
	return s.do(ctx)
}

func (s *MetadataDetailsService) DoWithUserAgentSuffix(ctx context.Context, userAgentSuffix string) (ConnectorMetadataResponse, error) {
	return s.doWithUserAgentSuffix(ctx, userAgentSuffix)
}

func (s *MetadataDetailsService) do(ctx context.Context) (ConnectorMetadataResponse, error) {
	var response ConnectorMetadataResponse
	if s.service == nil {
		return response, fmt.Errorf("missing required service")
	}

	url := fmt.Sprintf("/metadata/connector-types/%v", *s.service)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}

func (s *MetadataDetailsService) doWithUserAgentSuffix(ctx context.Context, userAgentSuffix string) (ConnectorMetadataResponse, error) {
	var response ConnectorMetadataResponse
	if s.service == nil {
		return response, fmt.Errorf("missing required service")
	}

	url := fmt.Sprintf("/metadata/connector-types/%v", *s.service)
	err := s.HttpService.DoWithUserAgentSuffix(ctx, userAgentSuffix, "GET", url, nil, nil, 200, &response)
	return response, err
}
