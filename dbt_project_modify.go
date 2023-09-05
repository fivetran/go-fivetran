package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

type DbtProjectModifyService struct {
	c         *Client
	projectId *string

	dbtVersion      *string
	targetName      *string
	threads         *int
	environmentVars *[]string
	projectConfig   *DbtProjectConfig
}

type dbtProjectModifyRequest struct {
	DbtVersion      *string                        `json:"dbt_version,omitempty"`
	TargetName      *string                        `json:"target_name,omitempty"`
	Threads         *int                           `json:"threads,omitempty"`
	EnvironmentVars *[]string                      `json:"environment_vars,omitempty"`
	ProjectConfig   *updateDbtProjectConfigRequest `json:"project_config,omitempty"`
}

func (c *Client) NewDbtProjectModify() *DbtProjectModifyService {
	return &DbtProjectModifyService{c: c}
}

func (s *DbtProjectModifyService) ProjectId(value string) *DbtProjectModifyService {
	s.projectId = &value
	return s
}

func (s *DbtProjectModifyService) TargetName(value string) *DbtProjectModifyService {
	s.targetName = &value
	return s
}

func (s *DbtProjectModifyService) Threads(value int) *DbtProjectModifyService {
	s.threads = &value
	return s
}

func (s *DbtProjectModifyService) EnvironmentVars(value []string) *DbtProjectModifyService {
	s.environmentVars = &value
	return s
}

func (s *DbtProjectModifyService) ProjectConfig(value *DbtProjectConfig) *DbtProjectModifyService {
	s.projectConfig = value
	return s
}

func (s *DbtProjectModifyService) DbtVersion(value string) *DbtProjectModifyService {
	s.dbtVersion = &value
	return s
}

func (s *DbtProjectModifyService) request() *dbtProjectModifyRequest {
	var config *updateDbtProjectConfigRequest

	if s.projectConfig != nil {
		config = s.projectConfig.updateRequest()
	}

	return &dbtProjectModifyRequest{
		DbtVersion:      s.dbtVersion,
		TargetName:      s.targetName,
		Threads:         s.threads,
		EnvironmentVars: s.environmentVars,
		ProjectConfig:   config,
	}
}

func (s *DbtProjectModifyService) Do(ctx context.Context) (DbtProjectDetailsResponse, error) {
	var response DbtProjectDetailsResponse

	if s.projectId == nil {
		return response, fmt.Errorf("missing required dbt project ID")
	}
	url := fmt.Sprintf("%v/dbt/projects/%v", s.c.baseURL, *s.projectId)
	expectedStatus := 200

	headers := s.c.commonHeaders()
	headers["Content-Type"] = "application/json"

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return response, err
	}

	r := request{
		method:           "PATCH",
		url:              url,
		body:             reqBody,
		queries:          nil,
		headers:          headers,
		client:           s.c.httpClient,
		handleRateLimits: s.c.handleRateLimits,
		maxRetryAttempts: s.c.maxRetryAttempts,
	}

	respBody, respStatus, err := r.httpRequest(ctx)
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return response, err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected: %v", respStatus, expectedStatus)
		return response, err
	}

	return response, nil
}
