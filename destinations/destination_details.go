package destinations

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type DestinationDetailsService struct {
	httputils.HttpService
	destinationID *string
}

func (s *DestinationDetailsService) DestinationID(value string) *DestinationDetailsService {
	s.destinationID = &value
	return s
}

func (s *DestinationDetailsService) Do(ctx context.Context) (DestinationDetailsResponse, error) {
	var response DestinationDetailsResponse

	if s.destinationID == nil {
		return response, fmt.Errorf("missing required destinationID")
	}

	url := fmt.Sprintf("/destinations/%v", *s.destinationID)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}

func (s *DestinationDetailsService) DoCustom(ctx context.Context) (DestinationDetailsCustomResponse, error) {
	var response DestinationDetailsCustomResponse

	if s.destinationID == nil {
		return response, fmt.Errorf("missing required destinationID")
	}

	url := fmt.Sprintf("/destinations/%v", *s.destinationID)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}
