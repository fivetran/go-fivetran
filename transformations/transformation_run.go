package transformations

import (
	"context"
	"fmt"
	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type TransformationRunService struct {
	httputils.HttpService
	transformationId *string
}

func (s *TransformationRunService) TransformationId(value string) *TransformationRunService {
	s.transformationId = &value
	return s
}

func (s *TransformationRunService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse

	if s.transformationId == nil {
		return response, fmt.Errorf("missing required transformationId")
	}

	url := fmt.Sprintf("/transformations/%v/run", *s.transformationId)
	err := s.HttpService.Do(ctx, "POST", url, nil, nil, 200, &response)
	return response, err
}