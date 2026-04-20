package connectorsdk

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectorSdkPackageDeleteService struct {
	httputils.HttpService
	packageID *string
}

func (s *ConnectorSdkPackageDeleteService) PackageID(value string) *ConnectorSdkPackageDeleteService {
	s.packageID = &value
	return s
}

func (s *ConnectorSdkPackageDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	if s.packageID == nil {
		return response, fmt.Errorf("missing required packageID")
	}

	url := fmt.Sprintf("/connector-sdk/packages/%v", *s.packageID)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}
