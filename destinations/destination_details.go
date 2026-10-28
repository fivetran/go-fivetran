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

func (s *DestinationDetailsService) do(ctx context.Context, response any) error {
	if s.destinationID == nil {
		return fmt.Errorf("missing required destinationID")
	}

	url := fmt.Sprintf("/destinations/%v", *s.destinationID)
	return s.HttpService.Do(ctx, "GET", url, nil, nil, 200, response)
}

func (s *DestinationDetailsService) doWithUserAgentSuffix(ctx context.Context, userAgentSuffix string, response any) error {
	if s.destinationID == nil {
		return fmt.Errorf("missing required destinationID")
	}

	url := fmt.Sprintf("/destinations/%v", *s.destinationID)
	return s.HttpService.DoWithUserAgentSuffix(ctx, userAgentSuffix, "GET", url, nil, nil, 200, response)
}

func (s *DestinationDetailsService) Do(ctx context.Context) (DestinationDetailsResponse, error) {
	var response DestinationDetailsResponse

	err := s.do(ctx, &response)

	return response, err
}

func (s *DestinationDetailsService) DoCustom(ctx context.Context) (DestinationDetailsCustomResponse, error) {
	var response DestinationDetailsCustomResponse

	err := s.do(ctx, &response)

	return response, err
}

func (s *DestinationDetailsService) DoCustomWithUserAgentSuffix(ctx context.Context, userAgentSuffix string) (DestinationDetailsCustomResponse, error) {
	var response DestinationDetailsCustomResponse

	err := s.doWithUserAgentSuffix(ctx, userAgentSuffix, &response)

	return response, err
}
