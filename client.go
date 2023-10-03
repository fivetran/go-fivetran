package fivetran

import (
	"encoding/base64"
	"fmt"
	"net/http"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// Client holds client configuration
type Client struct {
	baseURL          string
	authorization    string
	customUserAgent  string
	httpClient       httputils.HttpClient
	handleRateLimits bool
	maxRetryAttempts int
}

const defaultBaseURL = "https://api.fivetran.com/v1"
const restAPIv2 = "application/json;version=2"

// WARNING: Update Agent version on each release!
const defaultUserAgent = "Go-Fivetran/0.7.8"

// New receives API Key and API Secret, and returns a new Client with the
// default HTTP client
func New(apiKey, apiSecret string) *Client {
	credentials := fmt.Sprintf("Basic %v", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v:%v", apiKey, apiSecret))))
	return &Client{
		baseURL:          defaultBaseURL,
		authorization:    credentials,
		httpClient:       &http.Client{},
		maxRetryAttempts: 2,
		handleRateLimits: true,
	}
}

// BaseURL changes Client base REST API endpoint URL
func (c *Client) BaseURL(baseURL string) {
	c.baseURL = baseURL
}

// CustomUserAgent sets custom User-Agent header in Client requests
func (c *Client) CustomUserAgent(customUserAgent string) {
	c.customUserAgent = customUserAgent
}

// SetHttpClient sets custom HTTP client to perform requests with
func (c *Client) SetHttpClient(httpClient httputils.HttpClient) {
	c.httpClient = httpClient
}

// SetHandleRateLimits sets custom HTTP client to handle rate limits automatically
func (c *Client) SetHandleRateLimits(handleRateLimits bool) {
	c.handleRateLimits = handleRateLimits
}

// SetMaxRetryAttempts sets custom HTTP client maximum retry attempts count
func (c *Client) SetMaxRetryAttempts(maxRetryAttempts int) {
	c.maxRetryAttempts = maxRetryAttempts
}

func (c *Client) NewHttpService() httputils.HttpService {
	return httputils.HttpService{
		CommonHeaders:    c.commonHeaders(),
		BaseUrl:          c.baseURL,
		MaxRetryAttempts: c.maxRetryAttempts,
		HandleRateLimits: c.handleRateLimits,
		Client:           c.httpClient,
	}
}

func (c *Client) commonHeaders() map[string]string {
	userAgent := defaultUserAgent

	if c.customUserAgent != "" {
		userAgent += " " + c.customUserAgent
	}

	return map[string]string{
		"Authorization": c.authorization,
		"User-Agent":    userAgent,
	}
}
