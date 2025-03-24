package destinations

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type DestinationDeleteService struct {
	httputils.HttpService
	destinationID *string
}

func (s *DestinationDeleteService) DestinationID(value string) *DestinationDeleteService {
	s.destinationID = &value
	return s
}

func (s *DestinationDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	
	if s.destinationID == nil {
		return response, fmt.Errorf("missing required destinationID")
	}

	url := fmt.Sprintf("/destinations/%v", *s.destinationID)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}