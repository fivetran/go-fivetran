package connections

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectionReSyncTableService struct {
    httputils.HttpService
	connectionID *string
	schema      *string
	table       *string
}

func (s *ConnectionReSyncTableService) ConnectionID(value string) *ConnectionReSyncTableService {
	s.connectionID = &value
	return s
}

func (s *ConnectionReSyncTableService) Schema(value string) *ConnectionReSyncTableService {
	s.schema = &value
	return s
}

func (s *ConnectionReSyncTableService) Table(value string) *ConnectionReSyncTableService {
	s.table = &value
	return s
}

func (s *ConnectionReSyncTableService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse

	if s.connectionID == nil {
		return response, fmt.Errorf("missing required ConnectionID")
	}
	if s.schema == nil {
		return response, fmt.Errorf("missing required Schema")
	}
	if s.table == nil {
		return response, fmt.Errorf("missing required Table")
	}
	
	url := fmt.Sprintf("/connections/%v/schemas/%v/tables/%v/resync", *s.connectionID, *s.schema, *s.table)
	err := s.HttpService.Do(ctx, "POST", url, nil, nil, 200, &response)
	return response, err
}