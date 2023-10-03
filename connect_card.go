package fivetran

import (
    "context"
    "encoding/json"
    "fmt"

    httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ConnectCardService implements the https://fivetran.com/docs/rest-api/getting-started/connect-card#connectcards
// Ref.https://fivetran.com/docs/rest-api/getting-started/connect-card
type ConnectCardService struct {
    c                  *Client
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

func (c *Client) NewConnectCard() *ConnectCardService {
    return &ConnectCardService{c: c}
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

    url := fmt.Sprintf("%v/connectors/%v/connect-card", s.c.baseURL, *s.connectorId)
    expectedStatus := 200

    headers := s.c.commonHeaders()
    headers["Content-Type"] = "application/json"
    headers["Accept"] = restAPIv2

    reqBody, err := json.Marshal(s.request())
    if err != nil {
        return response, err
    }

    r := httputils.Request{
        Method:  "POST",
        Url:     url,
        Body:    reqBody,
        Queries: nil,
        Headers: headers,
        Client:  s.c.httpClient,
    }

    respBody, respStatus, err := r.Do(ctx)
    if err != nil {
        return response, err
    }

    if err := json.Unmarshal(respBody, &response); err != nil {
        return response, err
    }

    if respStatus != expectedStatus {
        err := fmt.Errorf("status code: %v; expected: %v", respStatus, expectedStatus)
        return response, err
    }

    return response, nil
}