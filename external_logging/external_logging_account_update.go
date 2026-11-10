package externallogging

import (
	"context"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ExternalLoggingAccountUpdateService updates the account-level log service:
// PATCH /v1/external-logging/account. The config is sent and returned as a map.
type ExternalLoggingAccountUpdateService struct {
	httputils.HttpService
	enabled       *bool
	runSetupTests *bool
	configCustom  *map[string]interface{}
}

func (s *ExternalLoggingAccountUpdateService) request() *externalLoggingCustomUpdateRequest {
	return &externalLoggingCustomUpdateRequest{
		externalLoggingUpdateRequestBase: externalLoggingUpdateRequestBase{
			Enabled:       s.enabled,
			RunSetupTests: s.runSetupTests,
		},
		Config: s.configCustom,
	}
}

func (s *ExternalLoggingAccountUpdateService) Enabled(value bool) *ExternalLoggingAccountUpdateService {
	s.enabled = &value
	return s
}

// RunSetupTests makes Fivetran run the setup tests right after updating the log
// service. The results are returned in the SetupTests field of the response.
func (s *ExternalLoggingAccountUpdateService) RunSetupTests(value bool) *ExternalLoggingAccountUpdateService {
	s.runSetupTests = &value
	return s
}

func (s *ExternalLoggingAccountUpdateService) ConfigCustom(value *map[string]interface{}) *ExternalLoggingAccountUpdateService {
	s.configCustom = value
	return s
}

func (s *ExternalLoggingAccountUpdateService) do(ctx context.Context, userAgentSuffix *string) (ExternalLoggingCustomResponse, error) {
	var response ExternalLoggingCustomResponse

	err := send(ctx, s.HttpService, userAgentSuffix, "PATCH", "/external-logging/account", s.request(), nil, 200, &response)

	return response, err
}

func (s *ExternalLoggingAccountUpdateService) DoCustom(ctx context.Context) (ExternalLoggingCustomResponse, error) {
	return s.do(ctx, nil)
}

func (s *ExternalLoggingAccountUpdateService) DoCustomWithUserAgentSuffix(ctx context.Context, userAgentSuffix string) (ExternalLoggingCustomResponse, error) {
	return s.do(ctx, &userAgentSuffix)
}
