package connectcard

import (
    "context"
    "fmt"

    httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ConnectCardService implements the https://fivetran.com/docs/rest-api/getting-started/connect-card#connectcards
// Ref.https://fivetran.com/docs/rest-api/getting-started/connect-card
type ConnectCardService struct {
    httputils.HttpService
    connectorId        *string
    config             *ConnectCardConfig
}

type connectCardRequest struct {
    ConnectCardConfig  *connectCardConfigRequest `json:"connect_card_config"`
}

type ConnectCardResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    Data    struct {
        ConnectorId  string `json:"connector_id"`
        ConnectCard  struct {
            Token    string `json:"token"`
            Uri      string `json:"uri"`
        } `json:"connect_card"`
        ConnectCardConfig    ConnectCardConfigResponse   `json:"connect_card_config"`
    } `json:"data"`
}

func (s *ConnectCardService) request() *connectCardRequest {
    var config *connectCardConfigRequest

    if s.config != nil {
        config = s.config.request()
    }

    return &connectCardRequest{
        ConnectCardConfig: config,
    }
}

func (s *ConnectCardService) ConnectorId(value string) *ConnectCardService {
    s.connectorId = &value
    return s
}

func (s *ConnectCardService) Config(value *ConnectCardConfig) *ConnectCardService {
    s.config = value
    return s
}

func (s *ConnectCardService) Do(ctx context.Context) (ConnectCardResponse, error) {
    var response ConnectCardResponse

    if s.connectorId == nil {
        return response, fmt.Errorf("missing required connectorId")
    }

    if s.config == nil {
        return response, fmt.Errorf("missing required config")
    }

    url := fmt.Sprintf("/connectors/%v/connect-card", *s.connectorId)
    err := s.HttpService.Do(ctx, "POST", url, s.request(), nil, 200, &response)
    return response, err
}