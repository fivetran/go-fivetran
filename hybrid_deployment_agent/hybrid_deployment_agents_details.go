package hybriddeploymentagent

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type HybridDeploymentAgentDetailsService struct {
	httputils.HttpService
	agentId *string
}

func (s *HybridDeploymentAgentDetailsService) AgentId(value string) *HybridDeploymentAgentDetailsService {
	s.agentId = &value
	return s
}

func (s *HybridDeploymentAgentDetailsService) Do(ctx context.Context) (HybridDeploymentAgentDetailsResponse, error) {
	var response HybridDeploymentAgentDetailsResponse

	if s.agentId == nil {
		return response, fmt.Errorf("missing required agentId")
	}

	url := fmt.Sprintf("/hybrid-deployment-agents/%v", *s.agentId)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}
