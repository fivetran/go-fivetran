package transformations

import (
	"context"
	"fmt"
	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type TransformationCancelService struct {
	httputils.HttpService
	transformationId *string
}

func (s *TransformationCancelService) TransformationId(value string) *TransformationCancelService {
	s.transformationId = &value
	return s
}

func (s *TransformationCancelService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse

	if s.transformationId == nil {
		return response, fmt.Errorf("missing required transformationId")
	}

	url := fmt.Sprintf("/transformations/%v/cancel", *s.transformationId)
	err := s.HttpService.Do(ctx, "POST", url, nil, nil, 200, &response)
	return response, err
}