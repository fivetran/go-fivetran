package hybriddeploymentagent

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
	"github.com/fivetran/go-fivetran/common"
)

// HybridDeploymentAgentResetCredentialsService Regenerates authentication keys for the specified hybrid deployment agent.
// Ref. https://fivetran.com/docs/rest-api/api-reference/hybrid-deployment-agent-management/reset-local-processing-agent-credentials
type HybridDeploymentAgentResetCredentialsService struct {
	httputils.HttpService
	agentId *string
}

func (s *HybridDeploymentAgentResetCredentialsService) AgentId(value string) *HybridDeploymentAgentResetCredentialsService {
	s.agentId = &value
	return s
}

func (s *HybridDeploymentAgentResetCredentialsService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse

	if s.agentId == nil {
		return response, fmt.Errorf("missing required agentId")
	}

	url := fmt.Sprintf("/hybrid-deployment-agents/%v/reset-credentials", *s.agentId)
	err := s.HttpService.Do(ctx, "POST", url, nil, nil, 200, &response)
	return response, err
}