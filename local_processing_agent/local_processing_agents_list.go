package localprocessingagent

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// LocalProcessingAgentListService Returns a list of local processing agents with usage within your Fivetran account.
// Ref. https://fivetran.com/docs/rest-api/local-processing-agent-management#listlocalprocessingagents
type LocalProcessingAgentListService struct {
	httputils.HttpService
	limit  *int
	cursor *string
}

func (s *LocalProcessingAgentListService) Limit(value int) *LocalProcessingAgentListService {
	s.limit = &value
	return s
}

func (s *LocalProcessingAgentListService) Cursor(value string) *LocalProcessingAgentListService {
	s.cursor = &value
	return s
}

func (s *LocalProcessingAgentListService) Do(ctx context.Context) (LocalProcessingAgentListResponse, error) {
	var response LocalProcessingAgentListResponse

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
	err := s.HttpService.Do(ctx, "GET", "/local-processing-agents", nil, queries, 200, &response)
	return response, err
}