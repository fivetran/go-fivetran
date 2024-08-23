package dbt

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type DbtProjectModifyService struct {
	httputils.HttpService
	dbtProjectID *string

	dbtVersion      *string
	targetName      *string
	threads         *int
	environmentVars *[]string
	projectConfig   *DbtProjectConfig
}

func (s *DbtProjectModifyService) DbtProjectID(value string) *DbtProjectModifyService {
	s.dbtProjectID = &value
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
	var config interface{}

	if s.projectConfig != nil {
		config = s.projectConfig.Request()
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
	if s.dbtProjectID == nil {
		return response, fmt.Errorf("missing required dbtProjectID")
	}
	url := fmt.Sprintf("/dbt/projects/%v", *s.dbtProjectID)
	err := s.HttpService.Do(ctx, "PATCH", url, s.request(), nil, 200, &response)
	return response, err
}
