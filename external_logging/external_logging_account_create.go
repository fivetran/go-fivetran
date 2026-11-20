package externallogging

import (
	"context"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ExternalLoggingAccountCreateService creates the account-level log service:
// POST /v1/external-logging/account. The config is sent and returned as a map.
type ExternalLoggingAccountCreateService struct {
	httputils.HttpService
	service       *string
	enabled       *bool
	runSetupTests *bool
	configCustom  *map[string]interface{}
}

// externalLoggingAccountCreateRequest is the account-level create body. Unlike
// the group-level one it has no group_id, and it accepts run_setup_tests.
type externalLoggingAccountCreateRequest struct {
	Service       *string                 `json:"service,omitempty"`
	Enabled       *bool                   `json:"enabled,omitempty"`
	RunSetupTests *bool                   `json:"run_setup_tests,omitempty"`
	Config        *map[string]interface{} `json:"config,omitempty"`
}

func (s *ExternalLoggingAccountCreateService) request() *externalLoggingAccountCreateRequest {
	return &externalLoggingAccountCreateRequest{
		Service:       s.service,
		Enabled:       s.enabled,
		RunSetupTests: s.runSetupTests,
		Config:        s.configCustom,
	}
}

func (s *ExternalLoggingAccountCreateService) Service(value string) *ExternalLoggingAccountCreateService {
	s.service = &value
	return s
}

func (s *ExternalLoggingAccountCreateService) Enabled(value bool) *ExternalLoggingAccountCreateService {
	s.enabled = &value
	return s
}

// RunSetupTests makes Fivetran run the setup tests right after creating the log
// service. The results are returned in the SetupTests field of the response.
func (s *ExternalLoggingAccountCreateService) RunSetupTests(value bool) *ExternalLoggingAccountCreateService {
	s.runSetupTests = &value
	return s
}

func (s *ExternalLoggingAccountCreateService) ConfigCustom(value *map[string]interface{}) *ExternalLoggingAccountCreateService {
	s.configCustom = value
	return s
}

func (s *ExternalLoggingAccountCreateService) do(ctx context.Context, userAgentSuffix *string) (ExternalLoggingCustomResponse, error) {
	var response ExternalLoggingCustomResponse

	err := send(ctx, s.HttpService, userAgentSuffix, "POST", "/external-logging/account", s.request(), nil, 201, &response)

	return response, err
}

func (s *ExternalLoggingAccountCreateService) DoCustom(ctx context.Context) (ExternalLoggingCustomResponse, error) {
	return s.do(ctx, nil)
}

func (s *ExternalLoggingAccountCreateService) DoCustomWithUserAgentSuffix(ctx context.Context, userAgentSuffix string) (ExternalLoggingCustomResponse, error) {
	return s.do(ctx, &userAgentSuffix)
}
