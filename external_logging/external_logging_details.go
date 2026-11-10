package externallogging

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
	"github.com/fivetran/go-fivetran/utils"
)

type ExternalLoggingDetailsService struct {
	httputils.HttpService
	externalLoggingId *string
}

func (s *ExternalLoggingDetailsService) ExternalLoggingId(value string) *ExternalLoggingDetailsService {
	s.externalLoggingId = &value
	return s
}

func (s *ExternalLoggingDetailsService) do(ctx context.Context, userAgentSuffix *string, response any) error {
	if s.externalLoggingId == nil {
		return fmt.Errorf("missing required externalLoggingId")
	}

	url := fmt.Sprintf("/external-logging/%v", *s.externalLoggingId)
	return send(ctx, s.HttpService, userAgentSuffix, "GET", url, nil, nil, 200, &response)
}

func (s *ExternalLoggingDetailsService) Do(ctx context.Context) (ExternalLoggingResponse, error) {
	var response ExternalLoggingResponse

	err := s.do(ctx, nil, &response)

	return response, err
}

func (s *ExternalLoggingDetailsService) DoCustom(ctx context.Context) (ExternalLoggingCustomResponse, error) {
	var response ExternalLoggingCustomResponse

	err := s.do(ctx, nil, &response)

	return response, err
}

func (s *ExternalLoggingDetailsService) DoCustomMerged(ctx context.Context) (ExternalLoggingCustomMergedResponse, error) {
	return s.doCustomMerged(ctx, nil)
}

func (s *ExternalLoggingDetailsService) doCustomMerged(ctx context.Context, userAgentSuffix *string) (ExternalLoggingCustomMergedResponse, error) {
	var response ExternalLoggingCustomMergedResponse

	err := s.do(ctx, userAgentSuffix, &response)

	if err == nil {
		err = utils.FetchFromMap(&response.Data.CustomConfig, &response.Data.Config)
	}

	return response, err
}

func (s *ExternalLoggingDetailsService) DoWithUserAgentSuffix(ctx context.Context, userAgentSuffix string) (ExternalLoggingResponse, error) {
	var response ExternalLoggingResponse

	err := s.do(ctx, &userAgentSuffix, &response)

	return response, err
}

func (s *ExternalLoggingDetailsService) DoCustomWithUserAgentSuffix(ctx context.Context, userAgentSuffix string) (ExternalLoggingCustomResponse, error) {
	var response ExternalLoggingCustomResponse

	err := s.do(ctx, &userAgentSuffix, &response)

	return response, err
}

func (s *ExternalLoggingDetailsService) DoCustomMergedWithUserAgentSuffix(ctx context.Context, userAgentSuffix string) (ExternalLoggingCustomMergedResponse, error) {
	return s.doCustomMerged(ctx, &userAgentSuffix)
}
