package httputils

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

type Request struct {
	Method  string
	Url     string
	Body    []byte
	Queries map[string]string
	Headers map[string]string
	Client  HttpClient

	HandleRateLimits bool
	MaxRetryAttempts int
}

func (req *Request) Do(ctx context.Context) ([]byte, int, error) {
	return req.do(ctx, 0)
}

func (req *Request) do(ctx context.Context, attempt int) ([]byte, int, error) {
	if req.Client == nil {
		return nil, 0, errors.New("HTTP client is not provided")
	}

	newReq, err := http.NewRequestWithContext(ctx, req.Method, req.Url, bytes.NewReader(req.Body))
	if err != nil {
		return nil, 0, err
	}

	if len(req.Queries) > 0 {
		query := newReq.URL.Query()
		for k, v := range req.Queries {
			query.Add(k, v)
		}
		newReq.URL.RawQuery = query.Encode()
	}

	for k, v := range req.Headers {
		newReq.Header.Add(k, v)
	}

	if debug.enable {
		fmt.Printf("---\nDebug:\n  - HTTP Request:\n")
		fmt.Printf("    - Method: %v\n", req.Method)
		fmt.Printf("    - URL: %v\n", newReq.URL.String())
		fmt.Printf("    - Body: %s\n", req.Body)
		fmt.Printf("    - Headers:\n")
		for k, v := range req.Headers {
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

	resp, err := req.Client.Do(newReq)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 429 && req.HandleRateLimits {
		if attempt > req.MaxRetryAttempts {
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
		err = contextDelay(ctx, time.Duration(retryAfterSeconds)*time.Second)
		if err != nil {
			return nil, 0, err
		}
		if debug.enable {
			fmt.Printf("\n\t- Retry attempt: %v", attempt)
		}
		return req.do(ctx, attempt+1)
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

func contextDelay(ctx context.Context, d time.Duration) error {
	t := time.NewTimer(d)
	select {
	case <-ctx.Done():
		t.Stop()
		return fmt.Errorf("interrupted: context deadline exceeded")
	case <-t.C:
	}
	return nil
}
