package hybriddeploymentagent

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type HybridDeploymentAgentReAuthService struct {
	httputils.HttpService
	agentId 	*string
	authType 	*string
}

func (s *HybridDeploymentAgentReAuthService) request() hybridDeploymentAgentReAuthRequest {
	return hybridDeploymentAgentReAuthRequest{
		AuthType: 		s.authType,
	}
}

func (s *HybridDeploymentAgentReAuthService) AgentId(value string) *HybridDeploymentAgentReAuthService {
	s.agentId = &value
	return s
}

func (s *HybridDeploymentAgentReAuthService) AuthType(value string) *HybridDeploymentAgentReAuthService {
	s.authType = &value
	return s
}

func (s *HybridDeploymentAgentReAuthService) Do(ctx context.Context) (HybridDeploymentAgentCreateResponse, error) {
	var response HybridDeploymentAgentCreateResponse

	if s.agentId == nil {
		return response, fmt.Errorf("missing required agentId")
	}

	url := fmt.Sprintf("/hybrid-deployment-agents/%v/re-auth", *s.agentId)
	err := s.HttpService.Do(ctx, "POST", url, s.request(), nil, 200, &response)
	return response, err
}