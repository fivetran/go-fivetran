package transformations

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type TransformationProjectTestsService struct {
	httputils.HttpService
	projectId *string
}

func (s *TransformationProjectTestsService) ExternalLoggingId(value string) *TransformationProjectTestsService {
	s.projectId = &value
	return s
}

func (s *TransformationProjectTestsService) Do(ctx context.Context) (TransformationProjectResponse, error) {
	var response TransformationProjectResponse

	if s.projectId == nil {
		return response, fmt.Errorf("missing required projectId")
	}

	url := fmt.Sprintf("/transformation-projects/%v/test", *s.projectId)
	err := s.HttpService.Do(ctx, "POST", url, nil, nil, 200, &response)
	return response, err
}