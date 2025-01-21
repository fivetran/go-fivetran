package transformations

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type TransformationProjectDeleteService struct {
	httputils.HttpService
	projectId *string
}

func (s *TransformationProjectDeleteService) ProjectId(value string) *TransformationProjectDeleteService {
	s.projectId = &value
	return s
}

func (s *TransformationProjectDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	
	if s.projectId == nil {
		return response, fmt.Errorf("missing required projectId")
	}

	url := fmt.Sprintf("/transformation-projects/%v", *s.projectId)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}
