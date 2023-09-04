package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

type DbtProjectCreateService struct {
	c               *Client
	groupID         *string
	dbtVersion      *string
	defaultSchema   *string
	targetName      *string
	threads         *int
	projectType     *string
	environmentVars *[]string
	projectConfig   *DbtProjectConfig
}

type dbtProjectCreateRequest struct {
	GroupID         *string                  `json:"group_id,omitempty"`
	DbtVersion      *string                  `json:"dbt_version,omitempty"`
	DefaultSchema   *string                  `json:"default_schema,omitempty"`
	TargetName      *string                  `json:"target_name,omitempty"`
	Threads         *int                     `json:"threads,omitempty"`
	EnvironmentVars *[]string                `json:"environment_vars,omitempty"`
	Type            *string                  `json:"type,omitempty"`
	ProjectConfig   *dbtProjectConfigRequest `json:"project_config,omitempty"`
}

func (c *Client) NewDbtProjectCreate() *DbtProjectCreateService {
	return &DbtProjectCreateService{c: c}
}

func (s *DbtProjectCreateService) request() *dbtProjectCreateRequest {
	var config *dbtProjectConfigRequest

	if s.projectConfig != nil {
		config = s.projectConfig.request()
	}

	return &dbtProjectCreateRequest{
		GroupID:         s.groupID,
		DbtVersion:      s.dbtVersion,
		DefaultSchema:   s.defaultSchema,
		TargetName:      s.targetName,
		Threads:         s.threads,
		Type:            s.projectType,
		EnvironmentVars: s.environmentVars,
		ProjectConfig:   config,
	}
}

func (s *DbtProjectCreateService) GroupID(value string) *DbtProjectCreateService {
	s.groupID = &value
	return s
}

func (s *DbtProjectCreateService) DbtVersion(value string) *DbtProjectCreateService {
	s.dbtVersion = &value
	return s
}

func (s *DbtProjectCreateService) DefaultSchema(value string) *DbtProjectCreateService {
	s.defaultSchema = &value
	return s
}

func (s *DbtProjectCreateService) TargetName(value string) *DbtProjectCreateService {
	s.targetName = &value
	return s
}

func (s *DbtProjectCreateService) Threads(value int) *DbtProjectCreateService {
	s.threads = &value
	return s
}

func (s *DbtProjectCreateService) EnvironmentVars(value []string) *DbtProjectCreateService {
	s.environmentVars = &value
	return s
}

func (s *DbtProjectCreateService) ProjectConfig(value *DbtProjectConfig) *DbtProjectCreateService {
	s.projectConfig = value
	return s
}

func (s *DbtProjectCreateService) Type(value string) *DbtProjectCreateService {
	s.projectType = &value
	return s
}

func (s *DbtProjectCreateService) Do(ctx context.Context) (DbtProjectDetailsResponse, error) {
	var response DbtProjectDetailsResponse
	url := fmt.Sprintf("%v/dbt/projects", s.c.baseURL)
	expectedStatus := 201

	headers := s.c.commonHeaders()
	headers["Content-Type"] = "application/json"

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return response, err
	}

	r := request{
		method:           "POST",
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
