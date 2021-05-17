package fivetran

import (
	"context"
	"io"
	"net/http"
)

// Request is a type used to pass http requests to func httpRequest.
type Request struct {
	method  string
	url     string
	body    io.Reader
	queries map[string]string
	headers map[string]string
}

// httpRequest receives a Request and a context.Context and calls func http.NewRequestWithContext.
// If the request succeeds, it returns the response body, the response status code, and a nil error.
// It returns an error if the request fails.
func httpRequest(req Request, ctx context.Context) ([]byte, int, error) {
	client := &http.Client{}

	newReq, err := http.NewRequestWithContext(ctx, req.method, req.url, req.body)
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

	resp, err := client.Do(newReq)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return respBody, resp.StatusCode, nil
}
