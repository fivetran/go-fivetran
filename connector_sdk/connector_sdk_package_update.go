package connectorsdk

import (
	"context"
	"fmt"
	"io"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectorSdkPackageUpdateService struct {
	httputils.HttpService
	packageID   *string
	fileContent io.Reader
	fileName    *string
}

func (s *ConnectorSdkPackageUpdateService) PackageID(value string) *ConnectorSdkPackageUpdateService {
	s.packageID = &value
	return s
}

func (s *ConnectorSdkPackageUpdateService) FileContent(value io.Reader) *ConnectorSdkPackageUpdateService {
	s.fileContent = value
	return s
}

func (s *ConnectorSdkPackageUpdateService) FileName(value string) *ConnectorSdkPackageUpdateService {
	s.fileName = &value
	return s
}

func (s *ConnectorSdkPackageUpdateService) Do(ctx context.Context) (ConnectorSdkPackageResponse, error) {
	var response ConnectorSdkPackageResponse
	if s.packageID == nil {
		return response, fmt.Errorf("missing required packageID")
	}
	if s.fileContent == nil {
		return response, fmt.Errorf("missing required fileContent")
	}

	fileName := "code.zip"
	if s.fileName != nil {
		fileName = *s.fileName
	}

	url := fmt.Sprintf("/connector-sdk/packages/%v", *s.packageID)
	err := s.HttpService.DoMultipart(ctx, "PATCH", url, "file", fileName, s.fileContent, nil, 200, &response)
	return response, err
}
