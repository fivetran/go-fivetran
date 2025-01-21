package transformations

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type TransformationProjectDetailsService struct {
	httputils.HttpService
	projectId *string
}

func (s *TransformationProjectDetailsService) ProjectId(value string) *TransformationProjectDetailsService {
	s.projectId = &value
	return s
}

func (s *TransformationProjectDetailsService) Do(ctx context.Context) (TransformationProjectResponse, error) {
	var response TransformationProjectResponse

	if s.projectId == nil {
		return response, fmt.Errorf("missing required projectId")
	}

	url := fmt.Sprintf("/transformation-projects/%v", *s.projectId)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}
