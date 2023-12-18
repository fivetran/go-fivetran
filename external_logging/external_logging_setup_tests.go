package externallogging

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ExternalLoggingSetupTestsService implements the Log Management, Run Log service setup tests API.
// Ref. https://fivetran.com/docs/rest-api/log-service-management#runlogservicesetuptests
type ExternalLoggingSetupTestsService struct {
	httputils.HttpService
	externalLoggingId *string
}

func (s *ExternalLoggingSetupTestsService) ExternalLoggingId(value string) *ExternalLoggingSetupTestsService {
	s.externalLoggingId = &value
	return s
}

func (s *ExternalLoggingSetupTestsService) Do(ctx context.Context) (ExternalLoggingSetupTestsResponse, error) {
	var response ExternalLoggingSetupTestsResponse

	if s.externalLoggingId == nil {
		return response, fmt.Errorf("missing required externalLoggingId")
	}

	url := fmt.Sprintf("/external-logging/%v/test", *s.externalLoggingId)
	err := s.HttpService.Do(ctx, "POST", url, nil, nil, 200, &response)
	return response, err
}