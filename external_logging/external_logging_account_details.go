package externallogging

import (
	"context"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ExternalLoggingAccountDetailsService retrieves the account-level log service:
// GET /v1/external-logging/account. The config is returned as a map.
type ExternalLoggingAccountDetailsService struct {
	httputils.HttpService
}

func (s *ExternalLoggingAccountDetailsService) do(ctx context.Context, userAgentSuffix *string) (ExternalLoggingCustomResponse, error) {
	var response ExternalLoggingCustomResponse

	err := send(ctx, s.HttpService, userAgentSuffix, "GET", "/external-logging/account", nil, nil, 200, &response)

	return response, err
}

func (s *ExternalLoggingAccountDetailsService) DoCustom(ctx context.Context) (ExternalLoggingCustomResponse, error) {
	return s.do(ctx, nil)
}

func (s *ExternalLoggingAccountDetailsService) DoCustomWithUserAgentSuffix(ctx context.Context, userAgentSuffix string) (ExternalLoggingCustomResponse, error) {
	return s.do(ctx, &userAgentSuffix)
}
