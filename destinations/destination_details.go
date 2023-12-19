package destinations

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// DestinationDetailsService implements the Destination Management, Retrieve destination details API.
// Ref. https://fivetran.com/docs/rest-api/destinations#retrievedestinationdetails
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
