package certificates

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectionCertificateRevokeService struct {
	httputils.HttpService
	connectionID *string
	hash        *string
}

func (s *ConnectionCertificateRevokeService) ConnectionID(value string) *ConnectionCertificateRevokeService {
	s.connectionID = &value
	return s
}

func (s *ConnectionCertificateRevokeService) Hash(value string) *ConnectionCertificateRevokeService {
	s.hash = &value
	return s
}

func (s *ConnectionCertificateRevokeService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	url := fmt.Sprintf("/connections/%v/certificates/%v", *s.connectionID, *s.hash)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}
