package connectorsdk

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectorSdkPackageDetailsService struct {
	httputils.HttpService
	packageID *string
}

func (s *ConnectorSdkPackageDetailsService) PackageID(value string) *ConnectorSdkPackageDetailsService {
	s.packageID = &value
	return s
}

func (s *ConnectorSdkPackageDetailsService) Do(ctx context.Context) (ConnectorSdkPackageResponse, error) {
	var response ConnectorSdkPackageResponse
	if s.packageID == nil {
		return response, fmt.Errorf("missing required packageID")
	}

	url := fmt.Sprintf("/connector-sdk/packages/%v", *s.packageID)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}
