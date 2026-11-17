package transformations

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
	"github.com/fivetran/go-fivetran/utils"
)

type TransformationProjectDetailsService struct {
	httputils.HttpService
	projectId *string
}

func (s *TransformationProjectDetailsService) ProjectId(value string) *TransformationProjectDetailsService {
	s.projectId = &value
	return s
}

func (s *TransformationProjectDetailsService) do(ctx context.Context, response any) error {
	if s.projectId == nil {
		return fmt.Errorf("missing required projectId")
	}

	url := fmt.Sprintf("/transformation-projects/%v", *s.projectId)
	return s.HttpService.Do(ctx, "GET", url, nil, nil, 200, response)
}

func (s *TransformationProjectDetailsService) Do(ctx context.Context) (TransformationProjectResponse, error) {
	var response TransformationProjectResponse
	err := s.do(ctx, &response)
	return response, err
}

func (s *TransformationProjectDetailsService) DoCustomMerged(ctx context.Context) (TransformationProjectCustomMergedResponse, error) {
	var response TransformationProjectCustomMergedResponse
	err := s.do(ctx, &response)
	if err == nil {
		err = utils.FetchFromMap(&response.Data.ProjectConfigCustom, &response.Data.ProjectConfig)
	}
	return response, err
}
