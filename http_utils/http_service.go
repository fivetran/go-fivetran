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
	return s.do(ctx, method, url, requestBody, queries, expectedStatus, response)
}

func (s HttpService) DoWithUserAgentSuffix(
	ctx context.Context,
	userAgentSuffix,
	method,
	url string,
	requestBody any,
	queries map[string]string,
	expectedStatus int,
	response any) error {
	return s.doWithUserAgentSuffix(ctx, userAgentSuffix, method, url, requestBody, queries, expectedStatus, response)
}

func (s HttpService) do(
	ctx context.Context,
	method,
	url string,
	requestBody any,
	queries map[string]string,
	expectedStatus int,
	response any) error {
	return s.send(ctx, nil, method, url, requestBody, queries, expectedStatus, response)
}

func (s HttpService) doWithUserAgentSuffix(
	ctx context.Context,
	userAgentSuffix,
	method,
	url string,
	requestBody any,
	queries map[string]string,
	expectedStatus int,
	response any) error {
	return s.send(ctx, &userAgentSuffix, method, url, requestBody, queries, expectedStatus, response)
}

func (s HttpService) send(
	ctx context.Context,
	userAgentSuffix *string,
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

	headers := copyHeaders(s.CommonHeaders)
	if userAgentSuffix != nil && *userAgentSuffix != "" {
		if headers["User-Agent"] == "" {
			headers["User-Agent"] = *userAgentSuffix
		} else {
			headers["User-Agent"] += " " + *userAgentSuffix
		}
	}

	if method == "POST" || method == "PATCH" {
		headers["Content-Type"] = "application/json"
	}

	r := Request{
		Method:           method,
		Url:              s.BaseUrl + url,
		Body:             body,
		Queries:          queries,
		Headers:          headers,
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

func copyHeaders(headers map[string]string) map[string]string {
	result := make(map[string]string, len(headers))
	for k, v := range headers {
		result[k] = v
	}
	return result
}
