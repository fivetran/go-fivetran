package externallogging

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ExternalLoggingSetupTestsService struct {
	httputils.HttpService
	externalLoggingId *string
}

func (s *ExternalLoggingSetupTestsService) ExternalLoggingId(value string) *ExternalLoggingSetupTestsService {
	s.externalLoggingId = &value
	return s
}

func (s *ExternalLoggingSetupTestsService) Do(ctx context.Context) (ExternalLoggingSetupTestsResponse, error) {
	return s.do(ctx, nil)
}

func (s *ExternalLoggingSetupTestsService) DoWithUserAgentSuffix(ctx context.Context, userAgentSuffix string) (ExternalLoggingSetupTestsResponse, error) {
	return s.do(ctx, &userAgentSuffix)
}

func (s *ExternalLoggingSetupTestsService) do(ctx context.Context, userAgentSuffix *string) (ExternalLoggingSetupTestsResponse, error) {
	var response ExternalLoggingSetupTestsResponse

	if s.externalLoggingId == nil {
		return response, fmt.Errorf("missing required externalLoggingId")
	}

	url := fmt.Sprintf("/external-logging/%v/test", *s.externalLoggingId)
	err := send(ctx, s.HttpService, userAgentSuffix, "POST", url, nil, nil, 200, &response)
	return response, err
}
