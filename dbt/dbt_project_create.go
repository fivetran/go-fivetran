package dbt

import (
	"context"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type DbtProjectCreateService struct {
	httputils.HttpService
	groupID         *string
	dbtVersion      *string
	defaultSchema   *string
	targetName      *string
	threads         *int
	projectType     *string
	environmentVars *[]string
	projectConfig   *DbtProjectConfig
}

func (s *DbtProjectCreateService) request() *dbtProjectCreateRequest {
	var config interface{}

	if s.projectConfig != nil {
		config = s.projectConfig.Request()
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
	err := s.HttpService.Do(ctx, "POST", "/dbt/projects", s.request(), nil, 201, &response)
	return response, err
}
