package transformations

import (
    "context"
    "fmt"

    httputils "github.com/fivetran/go-fivetran/http_utils"
    "github.com/fivetran/go-fivetran/utils"
)

type TransformationProjectUpdateService struct {
    httputils.HttpService
    projectId               *string
    runTests                *bool
    projectConfig           *TransformationProjectConfig
    projectConfigCustom     *map[string]interface{}
}

func (s *TransformationProjectUpdateService) requestBase() transformationProjectUpdateRequestBase {
    return transformationProjectUpdateRequestBase{
        RunTests: s.runTests,
    }
}

func (s *TransformationProjectUpdateService) request() *transformationProjectUpdateRequest {
    var projectConfig interface{}

    if s.projectConfig != nil {
        projectConfig = s.projectConfig.UpdateRequest()
    }

    return &transformationProjectUpdateRequest{
        transformationProjectUpdateRequestBase: s.requestBase(),
        ProjectConfig:                          projectConfig,
    }
}

func (s *TransformationProjectUpdateService) requestCustom() *transformationProjectCustomUpdateRequest {
    return &transformationProjectCustomUpdateRequest{
        transformationProjectUpdateRequestBase:  s.requestBase(),
        ProjectConfig:                           s.projectConfigCustom,
    }
}

func (s *TransformationProjectUpdateService) requestCustomMerged() (*transformationProjectCustomUpdateRequest, error) {
    currentConfig := s.projectConfigCustom

    if s.projectConfig != nil {
        var err error
        currentConfig, err = s.projectConfig.MergeForUpdate(currentConfig)
        if err != nil {
            return nil, err
        }
    }

    return &transformationProjectCustomUpdateRequest{
        transformationProjectUpdateRequestBase: s.requestBase(),
        ProjectConfig:                          currentConfig,
    }, nil
}

func (s *TransformationProjectUpdateService) ProjectId(value string) *TransformationProjectUpdateService {
    s.projectId = &value
    return s
}

func (s *TransformationProjectUpdateService) RunTests(value bool) *TransformationProjectUpdateService {
    s.runTests = &value
    return s
}

func (s *TransformationProjectUpdateService) ProjectConfig(value *TransformationProjectConfig) *TransformationProjectUpdateService {
    s.projectConfig = value
    return s
}

func (s *TransformationProjectUpdateService) ProjectConfigCustom(value *map[string]interface{}) *TransformationProjectUpdateService {
    s.projectConfigCustom = value
    return s
}

func (s *TransformationProjectUpdateService) do(ctx context.Context, req, response any) error {
    if s.projectId == nil {
        return fmt.Errorf("missing required projectId")
    }

    url := fmt.Sprintf("/transformation-projects/%v", *s.projectId)
    err := s.HttpService.Do(ctx, "PATCH", url, req, nil, 200, &response)
    return err
}

func (s *TransformationProjectUpdateService) Do(ctx context.Context) (TransformationProjectResponse, error) {
    var response TransformationProjectResponse

    err := s.do(ctx, s.request(), &response)

    return response, err
}

func (s *TransformationProjectUpdateService) DoCustom(ctx context.Context) (TransformationProjectCustomResponse, error) {
    var response TransformationProjectCustomResponse

    err := s.do(ctx, s.requestCustom(), &response)

    return response, err
}

func (s *TransformationProjectUpdateService) DoCustomMerged(ctx context.Context) (TransformationProjectCustomMergedResponse, error) {
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
