package localprocessingagent

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// LocalProcessingAgentDeleteService Deletes the specified local processing agent from your Fivetran account.
// Ref. https://fivetran.com/docs/rest-api/local-processing-agent-management#deletealocalprocessingagent
type LocalProcessingAgentDeleteService struct {
	httputils.HttpService
	agentId   *string
}

func (s *LocalProcessingAgentDeleteService) AgentId(value string) *LocalProcessingAgentDeleteService {
	s.agentId = &value
	return s
}

func (s *LocalProcessingAgentDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	
	if s.agentId == nil {
		return response, fmt.Errorf("missing required agentId")
	}

	url := fmt.Sprintf("/local-processing-agents/%v", *s.agentId)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}
