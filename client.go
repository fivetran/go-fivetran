package fivetran

import (
	"encoding/base64"
	"fmt"
)

// Client holds client configuration
type Client struct {
	baseURL       string
	authorization string
	userAgent     string
}

const defaultBaseURL = "https://api.fivetran.com/v1"
const restAPIv2 = "application/json;version=2"

// WARNING: Update Agent version on each release!
const defaultUserAgent = "Go-Fivetran/0.2.3"

// New receives API Key and API Secret, and returns a new Client
func New(apiKey, apiSecret string) *Client {
	return &Client{
		baseURL:       defaultBaseURL,
		authorization: fmt.Sprintf("Basic %v", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v:%v", apiKey, apiSecret)))),
		userAgent:     defaultUserAgent,
	}
}

// BaseURL changes Client base REST API endpoint URL
func (c *Client) BaseURL(baseURL string) {
	c.baseURL = baseURL
}

// UserAgent sets custom User-Agent header in Client requests
func (c *Client) UserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c *Client) commonHeaders() map[string]string {
	return map[string]string{
		"Authorization": c.authorization,
		"User-Agent":    c.userAgent,
	}
}
