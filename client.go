package fivetran

import (
	"encoding/base64"
	"fmt"
)

type Client struct {
	baseURL       string
	authorization string
}

const defaultBaseURL = "https://api.fivetran.com/v1"

func NewClient(apiKey string, apiSecret string) *Client {
	return &Client{
		baseURL:       defaultBaseURL,
		authorization: fmt.Sprintf("Basic %v", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v:%v", apiKey, apiSecret)))),
	}
}

func (c *Client) BaseURL(baseURL string) {
	c.baseURL = baseURL
}
