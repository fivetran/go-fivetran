package localprocessingagent

import (
	"context"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// LocalProcessingAgentCreateService Creates a new local processing agent within your Fivetran account.
// Ref. https://fivetran.com/docs/rest-api/local-processing-agent-management#createalocalprocessingagent
type LocalProcessingAgentCreateService struct {
	httputils.HttpService
	groupId 	     *string
	displayName      *string
	envType      	 *string
	acceptTerms      *bool
}

func (s *LocalProcessingAgentCreateService) request() localProcessingAgentCreateRequest {
	return localProcessingAgentCreateRequest{
		GroupId: 		s.groupId,
		DisplayName: 	s.displayName,
		EnvType: 		s.envType,
		AcceptTerms: 	s.acceptTerms,
	}
}

func (s *LocalProcessingAgentCreateService) GroupId(value string) *LocalProcessingAgentCreateService {
	s.groupId = &value
	return s
}

func (s *LocalProcessingAgentCreateService) DisplayName(value string) *LocalProcessingAgentCreateService {
	s.displayName = &value
	return s
}

func (s *LocalProcessingAgentCreateService) EnvType(value string) *LocalProcessingAgentCreateService {
	s.envType = &value
	return s
}

func (s *LocalProcessingAgentCreateService) AcceptTerms(value bool) *LocalProcessingAgentCreateService {
	s.acceptTerms = &value
	return s
}

func (s *LocalProcessingAgentCreateService) Do(ctx context.Context) (LocalProcessingAgentCreateResponse, error) {
	var response LocalProcessingAgentCreateResponse
	err := s.HttpService.Do(ctx, "POST", "/local-processing-agents", s.request(), nil, 201, &response)
	return response, err
}