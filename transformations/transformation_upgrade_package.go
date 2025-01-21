package transformations

import (
	"context"
	"fmt"
	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type TransformationUpgradePackageService struct {
	httputils.HttpService
	transformationId *string
}

func (s *TransformationUpgradePackageService) TransformationId(value string) *TransformationUpgradePackageService {
	s.transformationId = &value
	return s
}

func (s *TransformationUpgradePackageService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse

	if s.transformationId == nil {
		return response, fmt.Errorf("missing required transformationId")
	}

	url := fmt.Sprintf("/transformations/%v/upgrade", *s.transformationId)
	err := s.HttpService.Do(ctx, "POST", url, nil, nil, 200, &response)
	return response, err
}