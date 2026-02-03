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

func (s *ExternalLoggingDetailsService) do(ctx context.Context, response any) error {
	if s.externalLoggingId == nil {
		return fmt.Errorf("missing required externalLoggingId")
	}

	url := fmt.Sprintf("/external-logging/%v", *s.externalLoggingId)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return err
}

func (s *ExternalLoggingDetailsService) Do(ctx context.Context) (ExternalLoggingResponse, error) {
	var response ExternalLoggingResponse

	err := s.do(ctx, &response)

	return response, err
}

func (s *ExternalLoggingDetailsService) DoCustom(ctx context.Context) (ExternalLoggingCustomResponse, error) {
	var response ExternalLoggingCustomResponse

	err := s.do(ctx, &response)

	return response, err
}

func (s *ExternalLoggingDetailsService) DoCustomMerged(ctx context.Context) (ExternalLoggingCustomMergedResponse, error) {
	var response ExternalLoggingCustomMergedResponse

	err := s.do(ctx, &response)

	if err == nil {
		err = utils.FetchFromMap(&response.Data.CustomConfig, &response.Data.Config)
	}

	return response, err
}
