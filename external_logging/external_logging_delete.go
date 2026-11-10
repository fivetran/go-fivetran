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
	return s.do(ctx, nil)
}

func (s *ExternalLoggingDeleteService) DoWithUserAgentSuffix(ctx context.Context, userAgentSuffix string) (common.CommonResponse, error) {
	return s.do(ctx, &userAgentSuffix)
}

func (s *ExternalLoggingDeleteService) do(ctx context.Context, userAgentSuffix *string) (common.CommonResponse, error) {
	var response common.CommonResponse

	if s.externalLoggingId == nil {
		return response, fmt.Errorf("missing required externalLoggingId")
	}

	url := fmt.Sprintf("/external-logging/%v", *s.externalLoggingId)
	err := send(ctx, s.HttpService, userAgentSuffix, "DELETE", url, nil, nil, 200, &response)
	return response, err
}
