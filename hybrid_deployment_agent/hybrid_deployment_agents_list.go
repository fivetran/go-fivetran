package hybriddeploymentagent

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type HybridDeploymentAgentListService struct {
	httputils.HttpService
	limit  *int
	cursor *string
}

func (s *HybridDeploymentAgentListService) Limit(value int) *HybridDeploymentAgentListService {
	s.limit = &value
	return s
}

func (s *HybridDeploymentAgentListService) Cursor(value string) *HybridDeploymentAgentListService {
	s.cursor = &value
	return s
}

func (s *HybridDeploymentAgentListService) Do(ctx context.Context) (HybridDeploymentAgentListResponse, error) {
	var response HybridDeploymentAgentListResponse

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
	err := s.HttpService.Do(ctx, "GET", "/hybrid-deployment-agents", nil, queries, 200, &response)
	return response, err
}