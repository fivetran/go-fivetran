package metadata

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type MetadataDetailsService struct {
	httputils.HttpService
	service 	*string
}

func (s *MetadataDetailsService) Service(value string) *MetadataDetailsService {
	s.service = &value
	return s
}

func (s *MetadataDetailsService) Do(ctx context.Context) (ConnectorMetadataResponse, error) {
	var response ConnectorMetadataResponse
    if s.service == nil {
        return response, fmt.Errorf("missing required service")
    }

	url := fmt.Sprintf("/metadata/connector-types/%v", *s.service)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}