package mirai

import (
	"github.com/go-resty/resty/v2"
)

type ClientOptionFunc func(c *Client) error

func WithClient(client *resty.Client) ClientOptionFunc {
	return func(c *Client) error {
		c.client = client
		return nil
	}
}

func WithUserAgent(userAgent string) ClientOptionFunc {
	return func(c *Client) error {
		c.userAgent = userAgent
		return nil
	}
}
