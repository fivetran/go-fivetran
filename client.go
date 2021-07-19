package fivetran

import (
	"encoding/base64"
	"fmt"
)

// Client holds client configuration
type Client struct {
	baseURL       string
	authorization string
}

const defaultBaseURL = "https://api.fivetran.com/v1"
const restAPIv2 = "application/json;version=2"

// New receives API Key and API Secret, and returns a new Client
func New(apiKey, apiSecret string) *Client {
	return &Client{
		baseURL:       defaultBaseURL,
		authorization: fmt.Sprintf("Basic %v", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v:%v", apiKey, apiSecret)))),
	}
}

// BaseURL changes Client base REST API endpoint URL
func (c *Client) BaseURL(baseURL string) {
	c.baseURL = baseURL
}
