package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

type ConnectorDeleteService struct {
	c           *Client
	connectorID *string
}

type ConnectorDelete struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c *Client) NewConnectorDeleteService() *ConnectorDeleteService {
	return &ConnectorDeleteService{c: c}
}

func (s *ConnectorDeleteService) ConnectorID(connectorID string) *ConnectorDeleteService {
	s.connectorID = &connectorID
	return s
}

func (s *ConnectorDeleteService) Do(ctx context.Context) (ConnectorDelete, error) {
	if s.connectorID == nil { // we don't validate business rules (unless it is strictly necessary)
		err := fmt.Errorf("missing required ConnectorID")
		return ConnectorDelete{}, err
	}

	url := fmt.Sprintf("%v/connectors/%v", s.c.baseURL, *s.connectorID)
	expectedStatus := 200
	headers := make(map[string]string)

	headers["Authorization"] = s.c.authorization

	r := Request{
		method:  "DELETE",
		url:     url,
		body:    nil,
		queries: nil,
		headers: headers,
	}

	respBody, respStatus, err := httpRequest(r, ctx)
	if err != nil {
		return ConnectorDelete{}, err
	}

	var connectorDelete ConnectorDelete
	if err := json.Unmarshal(respBody, &connectorDelete); err != nil {
		return ConnectorDelete{}, err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected %v", respStatus, expectedStatus)
		return connectorDelete, err
	}

	return connectorDelete, nil
}
