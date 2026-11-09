package externallogging

import (
	"context"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ExternalLoggingAccountDeleteService deletes the account-level log service:
// DELETE /v1/external-logging/account.
type ExternalLoggingAccountDeleteService struct {
	httputils.HttpService
}

func (s *ExternalLoggingAccountDeleteService) do(ctx context.Context, userAgentSuffix *string) (common.CommonResponse, error) {
	var response common.CommonResponse

	err := send(ctx, s.HttpService, userAgentSuffix, "DELETE", "/external-logging/account", nil, nil, 200, &response)

	return response, err
}

func (s *ExternalLoggingAccountDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	return s.do(ctx, nil)
}

func (s *ExternalLoggingAccountDeleteService) DoWithUserAgentSuffix(ctx context.Context, userAgentSuffix string) (common.CommonResponse, error) {
	return s.do(ctx, &userAgentSuffix)
}
