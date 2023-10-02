package httputils

import (
	"context"
	"encoding/json"
	"fmt"
)

type HttpParams struct {
	Method         string
	ExpectedStatus int
}

type HttpService struct {
	HttpParams
	BaseUrl          string
	CommonHeaders    map[string]string
	Client           HttpClient
	HandleRateLimits bool
	MaxRetryAttempts int
}

func (s HttpService) Do(
	ctx context.Context,
	url string,
	requestBody any,
	queries map[string]string,
	response any) error {

	var body []byte = nil
	var err error = nil

	if requestBody != nil {
		body, err = json.Marshal(requestBody)
		if err != nil {
			return err
		}
	}

	r := Request{
		Method:           s.Method,
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

	if respStatus != s.ExpectedStatus {
		err := fmt.Errorf("status code: %v; expected: %v", respStatus, s.ExpectedStatus)
		return err
	}
	return nil
}
