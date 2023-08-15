package fivetran

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

const (
	DEBUG_BACKOFF_DELAY = 60
)

type request struct {
	method  string
	url     string
	body    []byte
	queries map[string]string
	headers map[string]string
	client  HttpClient

	handleRateLimits bool
	maxRetryAttempts int
}

func (req *request) httpRequest(ctx context.Context) ([]byte, int, error) {
	return req.httpRequestImpl(ctx, 0)
}

func (req *request) httpRequestImpl(ctx context.Context, attempt int) ([]byte, int, error) {

	if req.client == nil {
		return nil, 0, errors.New("HTTP client is not provided")
	}

	newReq, err := http.NewRequestWithContext(ctx, req.method, req.url, bytes.NewReader(req.body))
	if err != nil {
		return nil, 0, err
	}

	if len(req.queries) > 0 {
		query := newReq.URL.Query()
		for k, v := range req.queries {
			query.Add(k, v)
		}
		newReq.URL.RawQuery = query.Encode()
	}

	for k, v := range req.headers {
		newReq.Header.Add(k, v)
	}

	if debug.enable {
		fmt.Printf("---\nDebug:\n  - HTTP Request:\n")
		fmt.Printf("    - Method: %v\n", req.method)
		fmt.Printf("    - URL: %v\n", newReq.URL.String())
		fmt.Printf("    - Body: %s\n", req.body)
		fmt.Printf("    - Headers:\n")
		for k, v := range req.headers {
			if k == "Authorization" {
				if debug.authEnable {
					fmt.Printf("      - %v: %v\n", k, v)
				}
				if !debug.authEnable {
					fmt.Printf("      - %v: <omitted>\n", k)
				}
				continue
			}
			fmt.Printf("      - %v: %v\n", k, v)
		}
	}

	resp, err := req.client.Do(newReq)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 429 && req.handleRateLimits {
		if attempt > req.maxRetryAttempts {
			return nil, resp.StatusCode, fmt.Errorf("rate limit retry max attempts count reached")
		}
		fmt.Println(resp)
		retryAfterSeconds, err := strconv.Atoi(resp.Header.Get("Retry-After"))
		if err != nil {
			return nil, 0, err
		}
		if debug.enable {
			fmt.Printf("\n\t- Waiting for retry: %v seconds left", retryAfterSeconds)
		}
		time.Sleep(time.Duration(retryAfterSeconds) * time.Second)
		if debug.enable {
			fmt.Printf("\n\t- Retry attempt: %v", attempt)
		}
		return req.httpRequestImpl(ctx, attempt+1)
	}

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, resp.StatusCode, err
	}

	if debug.enable {
		fmt.Printf("  - HTTP Response:\n")
		fmt.Printf("    - Status Code: %v\n", resp.StatusCode)
		fmt.Printf("    - Body: %s\n---\n", respBody)
	}

	return respBody, resp.StatusCode, nil
}
