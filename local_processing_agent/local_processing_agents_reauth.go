package localprocessingagent

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// LocalProcessingAgentReAuthService Regenerates authentication keys for the specified local processing agent.
// Ref. https://fivetran.com/docs/rest-api/local-processing-agent-management#regeneratekeys
type LocalProcessingAgentReAuthService struct {
	httputils.HttpService
	agentId *string
}

func (s *LocalProcessingAgentReAuthService) AgentId(value string) *LocalProcessingAgentReAuthService {
	s.agentId = &value
	return s
}

func (s *LocalProcessingAgentReAuthService) Do(ctx context.Context) (LocalProcessingAgentCreateResponse, error) {
	var response LocalProcessingAgentCreateResponse

	if s.agentId == nil {
		return response, fmt.Errorf("missing required agentId")
	}

	url := fmt.Sprintf("/local-processing-agents/%v/re-auth", *s.agentId)
	err := s.HttpService.Do(ctx, "POST", url, nil, nil, 200, &response)
	return response, err
}