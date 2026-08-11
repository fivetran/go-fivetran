package httputils

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
)

type capturingClient struct {
	request *http.Request
}

func (c *capturingClient) Do(req *http.Request) (*http.Response, error) {
	c.request = req
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(`{"code":"Success"}`)),
	}, nil
}

func TestDoWithUserAgentSuffixAppendsSuffixWithoutMutatingCommonHeaders(t *testing.T) {
	client := &capturingClient{}
	service := HttpService{
		BaseUrl: "https://api.example.com/v1",
		CommonHeaders: map[string]string{
			"Authorization": "Basic token",
			"User-Agent":    "Go-Fivetran/1.3.9 terraform-provider-fivetran/1.9.37",
		},
		Client: client,
	}

	var response map[string]string
	err := service.DoWithUserAgentSuffix(
		context.Background(),
		"fivetran_connection_v2",
		http.MethodGet,
		"/connections/connection_id",
		nil,
		nil,
		http.StatusOK,
		&response,
	)
	if err != nil {
		t.Fatalf("DoWithUserAgentSuffix returned error: %v", err)
	}

	if got, want := client.request.Header.Get("User-Agent"), "Go-Fivetran/1.3.9 terraform-provider-fivetran/1.9.37 fivetran_connection_v2"; got != want {
		t.Fatalf("request User-Agent = %q, want %q", got, want)
	}
	if got, want := service.CommonHeaders["User-Agent"], "Go-Fivetran/1.3.9 terraform-provider-fivetran/1.9.37"; got != want {
		t.Fatalf("common User-Agent = %q, want %q", got, want)
	}
}
