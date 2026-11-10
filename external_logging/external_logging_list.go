package externallogging

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ExternalLoggingListService struct {
	httputils.HttpService
	limit  *int
	cursor *string
}

func (s *ExternalLoggingListService) Limit(value int) *ExternalLoggingListService {
	s.limit = &value
	return s
}

func (s *ExternalLoggingListService) Cursor(value string) *ExternalLoggingListService {
	s.cursor = &value
	return s
}

func (s *ExternalLoggingListService) Do(ctx context.Context) (ExternalLoggingListResponse, error) {
	return s.do(ctx, nil)
}

func (s *ExternalLoggingListService) DoWithUserAgentSuffix(ctx context.Context, userAgentSuffix string) (ExternalLoggingListResponse, error) {
	return s.do(ctx, &userAgentSuffix)
}

func (s *ExternalLoggingListService) do(ctx context.Context, userAgentSuffix *string) (ExternalLoggingListResponse, error) {
	var response ExternalLoggingListResponse
	var queries map[string]string = nil
	if s.cursor != nil || s.limit != nil {
		queries = make(map[string]string)
		if s.cursor != nil {
			queries["cursor"] = *s.cursor
		}
		if s.limit != nil {
			queries["limit"] = fmt.Sprintf("%v", *s.limit)
		}
	}
	err := send(ctx, s.HttpService, userAgentSuffix, "GET", "/external-logging", nil, queries, 200, &response)
	return response, err
}
