package transformations

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type QuickstartPackageDetailsService struct {
	httputils.HttpService
	packageDefinitionId *string
}

func (s *QuickstartPackageDetailsService) PackageDefinitionId(value string) *QuickstartPackageDetailsService {
	s.packageDefinitionId = &value
	return s
}

func (s *QuickstartPackageDetailsService) Do(ctx context.Context) (QuickstartPackageResponse, error) {
	var response QuickstartPackageResponse

	if s.packageDefinitionId == nil {
		return response, fmt.Errorf("missing required packageDefinitionId")
	}

	url := fmt.Sprintf("/transformations/package-metadata/%v", *s.packageDefinitionId)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}
