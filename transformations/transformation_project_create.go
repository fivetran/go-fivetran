package transformations

import (
	"context"

	httputils "github.com/fivetran/go-fivetran/http_utils"
	"github.com/fivetran/go-fivetran/utils"
)

type TransformationProjectCreateService struct {
	httputils.HttpService
	groupId      		*string
	projectType    		*string
	runTests     		*bool
	projectConfig       *TransformationProjectConfig
	projectConfigCustom *map[string]interface{}
}

func (s *TransformationProjectCreateService) requestBase() transformationProjectCreateRequestBase {
	return transformationProjectCreateRequestBase{
		GroupId: 			s.groupId,
		ProjectType: 		s.projectType,
		RunTests: 			s.runTests,
	}
}

func (s *TransformationProjectCreateService) request() *transformationProjectCreateRequest {
	var projectConfig interface{}
	if s.projectConfig != nil {
		projectConfig = s.projectConfig.CreateRequest()
	}

	r := &transformationProjectCreateRequest{
		transformationProjectCreateRequestBase: 		s.requestBase(),
		ProjectConfig:                         			projectConfig,
	}

	return r
}

func (s *TransformationProjectCreateService) requestCustom() *transformationProjectCustomCreateRequest {
	return &transformationProjectCustomCreateRequest{
		transformationProjectCreateRequestBase: s.requestBase(),
		ProjectConfig:                 		    s.projectConfigCustom,
	}
}

func (s *TransformationProjectCreateService) requestCustomMerged() (*transformationProjectCustomCreateRequest, error) {
	currentConfig := s.projectConfigCustom

	if s.projectConfig != nil {
		var err error
		currentConfig, err = s.projectConfig.MergeForCreate(currentConfig)
		if err != nil {
			return nil, err
		}
	}

	return &transformationProjectCustomCreateRequest{
		transformationProjectCreateRequestBase: s.requestBase(),
		ProjectConfig:                          currentConfig,
	}, nil
}

func (s *TransformationProjectCreateService) GroupId(value string) *TransformationProjectCreateService {
	s.groupId = &value
	return s
}

func (s *TransformationProjectCreateService) ProjectType(value string) *TransformationProjectCreateService {
	s.projectType = &value
	return s
}

func (s *TransformationProjectCreateService) RunTests(value bool) *TransformationProjectCreateService {
	s.runTests = &value
	return s
}

func (s *TransformationProjectCreateService) ProjectConfig(value *TransformationProjectConfig) *TransformationProjectCreateService {
	s.projectConfig = value
	return s
}

func (s *TransformationProjectCreateService) ProjectConfigCustom(value *map[string]interface{}) *TransformationProjectCreateService {
	s.projectConfigCustom = value
	return s
}

func (s *TransformationProjectCreateService) do(ctx context.Context, req, response any) error {
	err := s.HttpService.Do(ctx, "POST", "/transformation-projects", req, nil, 201, &response)
	return err
}

func (s *TransformationProjectCreateService) Do(ctx context.Context) (TransformationProjectResponse, error) {
	var response TransformationProjectResponse

	err := s.do(ctx, s.request(), &response)

	return response, err
}

func (s *TransformationProjectCreateService) DoCustom(ctx context.Context) (TransformationProjectCustomResponse, error) {
	var response TransformationProjectCustomResponse

	err := s.do(ctx, s.requestCustom(), &response)

	return response, err
}

func (s *TransformationProjectCreateService) DoCustomMerged(ctx context.Context) (TransformationProjectCustomMergedResponse, error) {
	var response TransformationProjectCustomMergedResponse

	req, err := s.requestCustomMerged()

	if err != nil {
		return response, err
	}

	err = s.do(ctx, req, &response)

	if err == nil {
		err = utils.FetchFromMap(&response.Data.ProjectConfigCustom, &response.Data.ProjectConfig)
	}

	return response, err
}
