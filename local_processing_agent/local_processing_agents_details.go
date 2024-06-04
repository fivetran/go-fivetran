package localprocessingagent

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// LocalProcessingAgentDetailsService Retrieves the details of the specified local processing agent.
// Ref. https://fivetran.com/docs/rest-api/local-processing-agent-management#retrievelocalprocessingagentdetails
type LocalProcessingAgentDetailsService struct {
	httputils.HttpService
	agentId *string
}

func (s *LocalProcessingAgentDetailsService) AgentId(value string) *LocalProcessingAgentDetailsService {
	s.agentId = &value
	return s
}

func (s *LocalProcessingAgentDetailsService) Do(ctx context.Context) (LocalProcessingAgentDetailsResponse, error) {
	var response LocalProcessingAgentDetailsResponse

	if s.agentId == nil {
		return response, fmt.Errorf("missing required agentId")
	}

	url := fmt.Sprintf("/local-processing-agents/%v", *s.agentId)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}
