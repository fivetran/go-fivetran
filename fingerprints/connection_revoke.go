package fingerprints

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectionFingerprintRevokeService struct {
	httputils.HttpService
	connectionID *string
	hash        *string
}

func (s *ConnectionFingerprintRevokeService) ConnectionID(value string) *ConnectionFingerprintRevokeService {
	s.connectionID = &value
	return s
}

func (s *ConnectionFingerprintRevokeService) Hash(value string) *ConnectionFingerprintRevokeService {
	s.hash = &value
	return s
}

func (s *ConnectionFingerprintRevokeService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	url := fmt.Sprintf("/connections/%v/fingerprints/%v", *s.connectionID, *s.hash)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}
