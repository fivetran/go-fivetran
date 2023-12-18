package externallogging

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ExternalLoggingDetailsService implements the Log Management, Retrieve Log Service details API.
// Ref. https://fivetran.com/docs/rest-api/log-service-management#retrievelogservicedetails
type ExternalLoggingDetailsService struct {
	httputils.HttpService
	externalLoggingId *string
}

func (s *ExternalLoggingDetailsService) ExternalLoggingId(value string) *ExternalLoggingDetailsService {
	s.externalLoggingId = &value
	return s
}

func (s *ExternalLoggingDetailsService) Do(ctx context.Context) (ExternalLoggingResponse, error) {
	var response ExternalLoggingResponse

	if s.externalLoggingId == nil {
		return response, fmt.Errorf("missing required externalLoggingId")
	}

	url := fmt.Sprintf("/external-logging/%v", *s.externalLoggingId)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}
