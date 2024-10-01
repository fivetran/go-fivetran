package hybriddeploymentagent

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// HybridDeploymentAgentDeleteService Deletes the specified local processing agent from your Fivetran account.
// Ref. https://fivetran.com/docs/rest-api/hybrid-deployment-agent-management#deleteahybriddeploymentagent
type HybridDeploymentAgentDeleteService struct {
	httputils.HttpService
	agentId   *string
}

func (s *HybridDeploymentAgentDeleteService) AgentId(value string) *HybridDeploymentAgentDeleteService {
	s.agentId = &value
	return s
}

func (s *HybridDeploymentAgentDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	
	if s.agentId == nil {
		return response, fmt.Errorf("missing required agentId")
	}

	url := fmt.Sprintf("/hybrid-deployment-agents/%v", *s.agentId)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}
