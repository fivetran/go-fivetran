package externallogging

import (
	"context"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ExternalLoggingAccountSetupTestsService runs the setup tests of the
// account-level log service: POST /v1/external-logging/account/test.
type ExternalLoggingAccountSetupTestsService struct {
	httputils.HttpService
}

func (s *ExternalLoggingAccountSetupTestsService) do(ctx context.Context, userAgentSuffix *string) (ExternalLoggingSetupTestsResponse, error) {
	var response ExternalLoggingSetupTestsResponse

	err := send(ctx, s.HttpService, userAgentSuffix, "POST", "/external-logging/account/test", nil, nil, 200, &response)

	return response, err
}

func (s *ExternalLoggingAccountSetupTestsService) Do(ctx context.Context) (ExternalLoggingSetupTestsResponse, error) {
	return s.do(ctx, nil)
}

func (s *ExternalLoggingAccountSetupTestsService) DoWithUserAgentSuffix(ctx context.Context, userAgentSuffix string) (ExternalLoggingSetupTestsResponse, error) {
	return s.do(ctx, &userAgentSuffix)
}
