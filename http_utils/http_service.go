package httputils

import (
	"context"
	"encoding/json"
	"fmt"
)

type HttpService struct {
	BaseUrl          string
	CommonHeaders    map[string]string
	Client           HttpClient
	HandleRateLimits bool
	MaxRetryAttempts int
}

func (s HttpService) Do(
	ctx context.Context,
	method,
	url string,
	requestBody any,
	queries map[string]string,
	expectedStatus int,
	response any) error {

	var body []byte = nil
	var err error = nil

	if requestBody != nil {
		body, err = json.Marshal(requestBody)
		if err != nil {
			return err
		}
	}

	if method == "POST" || method == "PATCH" {
		s.CommonHeaders["Content-Type"] = "application/json"
	}

	r := Request{
		Method:           method,
		Url:              s.BaseUrl + url,
		Body:             body,
		Queries:          queries,
		Headers:          s.CommonHeaders,
		Client:           s.Client,
		HandleRateLimits: s.HandleRateLimits,
		MaxRetryAttempts: s.MaxRetryAttempts,
	}

	respBody, respStatus, err := r.Do(ctx)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected: %v", respStatus, expectedStatus)
		return err
	}
	return nil
}
