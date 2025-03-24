package externallogging

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ExternalLoggingDeleteService struct {
	httputils.HttpService
	externalLoggingId *string
}

func (s *ExternalLoggingDeleteService) ExternalLoggingId(value string) *ExternalLoggingDeleteService {
	s.externalLoggingId = &value
	return s
}

func (s *ExternalLoggingDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	
	if s.externalLoggingId == nil {
		return response, fmt.Errorf("missing required externalLoggingId")
	}

	url := fmt.Sprintf("/external-logging/%v", *s.externalLoggingId)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}
