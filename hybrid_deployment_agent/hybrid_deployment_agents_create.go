package hybriddeploymentagent

import (
	"context"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type HybridDeploymentAgentCreateService struct {
	httputils.HttpService
	groupId 	     *string
	displayName      *string
	envType      	 *string
	authType      	 *string
	acceptTerms      *bool
}

func (s *HybridDeploymentAgentCreateService) request() hybridDeploymentAgentCreateRequest {
	return hybridDeploymentAgentCreateRequest{
		GroupId: 		s.groupId,
		DisplayName: 	s.displayName,
		EnvType: 		s.envType,
		AuthType: 		s.authType,
		AcceptTerms: 	s.acceptTerms,
	}
}

func (s *HybridDeploymentAgentCreateService) GroupId(value string) *HybridDeploymentAgentCreateService {
	s.groupId = &value
	return s
}

func (s *HybridDeploymentAgentCreateService) DisplayName(value string) *HybridDeploymentAgentCreateService {
	s.displayName = &value
	return s
}

func (s *HybridDeploymentAgentCreateService) EnvType(value string) *HybridDeploymentAgentCreateService {
	s.envType = &value
	return s
}

func (s *HybridDeploymentAgentCreateService) AuthType(value string) *HybridDeploymentAgentCreateService {
	s.authType = &value
	return s
}

func (s *HybridDeploymentAgentCreateService) AcceptTerms(value bool) *HybridDeploymentAgentCreateService {
	s.acceptTerms = &value
	return s
}

func (s *HybridDeploymentAgentCreateService) Do(ctx context.Context) (HybridDeploymentAgentCreateResponse, error) {
	var response HybridDeploymentAgentCreateResponse
	err := s.HttpService.Do(ctx, "POST", "/hybrid-deployment-agents", s.request(), nil, 201, &response)
	return response, err
}