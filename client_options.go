package mirai

type ClientOptionFunc func(c *Client) error

func WithUserAgent(userAgent string) ClientOptionFunc {
	return func(c *Client) error {
		c.userAgent = userAgent
		return nil
	}
}
