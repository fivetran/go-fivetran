package connectorsdk

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectorSdkPackageListService struct {
	httputils.HttpService
	limit  *int
	cursor *string
}

func (s *ConnectorSdkPackageListService) Limit(value int) *ConnectorSdkPackageListService {
	s.limit = &value
	return s
}

func (s *ConnectorSdkPackageListService) Cursor(value string) *ConnectorSdkPackageListService {
	s.cursor = &value
	return s
}

func (s *ConnectorSdkPackageListService) Do(ctx context.Context) (ConnectorSdkPackagesListResponse, error) {
	var response ConnectorSdkPackagesListResponse
	var queries map[string]string = nil
	if s.cursor != nil || s.limit != nil {
		queries = make(map[string]string)
		if s.cursor != nil {
			queries["cursor"] = *s.cursor
		}
		if s.limit != nil {
			queries["limit"] = fmt.Sprintf("%v", *s.limit)
		}
	}
	err := s.HttpService.Do(ctx, "GET", "/connector-sdk/packages", nil, queries, 200, &response)
	return response, err
}
