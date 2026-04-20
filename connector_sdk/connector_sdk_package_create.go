package connectorsdk

import (
	"context"
	"fmt"
	"io"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectorSdkPackageCreateService struct {
	httputils.HttpService
	fileContent io.Reader
	fileName    *string
}

func (s *ConnectorSdkPackageCreateService) FileContent(value io.Reader) *ConnectorSdkPackageCreateService {
	s.fileContent = value
	return s
}

func (s *ConnectorSdkPackageCreateService) FileName(value string) *ConnectorSdkPackageCreateService {
	s.fileName = &value
	return s
}

func (s *ConnectorSdkPackageCreateService) Do(ctx context.Context) (ConnectorSdkPackageResponse, error) {
	var response ConnectorSdkPackageResponse
	if s.fileContent == nil {
		return response, fmt.Errorf("missing required fileContent")
	}

	fileName := "code.zip"
	if s.fileName != nil {
		fileName = *s.fileName
	}

	err := s.HttpService.DoMultipart(ctx, "POST", "/connector-sdk/packages", "file", fileName, s.fileContent, nil, 201, &response)
	return response, err
}
