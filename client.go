package mirai

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/starudream/mirai-go/internal/query"
)

const (
	DefaultUserAgent = "starudream-mirai-go"
)

type Client struct {
	client *resty.Client

	baseURL *url.URL

	userAgent string

	verifyKey string
}

func NewClient(addr, verifyKey string, options ...ClientOptionFunc) (*Client, error) {
	c := &Client{
		client:    newRestyClient(),
		userAgent: DefaultUserAgent,
		verifyKey: verifyKey,
	}

	if err := c.setBaseURL(addr); err != nil {
		return nil, err
	}

	for i := 0; i < len(options); i++ {
		option := options[i]
		if option == nil {
			continue
		}
		if err := option(c); err != nil {
			return nil, err
		}
	}

	c.client.SetHeader("User-Agent", c.userAgent)

	return c, nil
}

func (c *Client) setBaseURL(urlStr string) error {
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}

	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	c.baseURL = baseURL

	return nil
}

func (c *Client) Pre(method, path string, opt any) (string, map[string]string, any, error) {
	u := *c.baseURL
	p, err := url.PathUnescape(path)
	if err != nil {
		return "", nil, nil, err
	}

	u.RawPath = c.baseURL.Path + path
	u.Path = c.baseURL.Path + p

	headers := map[string]string{}

	if v, ok := opt.(BaseReqInterface); ok && v.GetSessionKey() != "" {
		headers["Authorization"] = "sessionKey " + v.GetSessionKey()
	}

	var body any

	switch method {
	case http.MethodGet:
		if opt != nil {
			qvs, qe := query.Values(opt)
			if qe != nil {
				return "", nil, nil, qe
			}
			u.RawQuery = qvs.Encode()
		}
	case http.MethodPost:
		headers["Content-Type"] = "application/json"
		if opt != nil {
			body = opt
		}
	}

	return u.String(), headers, body, nil
}

func (c *Client) Exec(method, path string, opt, res any) (*resty.Response, error) {
	u, hds, body, err := c.Pre(method, path, opt)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.
		NewRequest().
		SetHeaders(hds).
		SetBody(body).
		SetResult(res).
		SetError(&BaseResp{}).
		Execute(method, u)
	if err != nil {
		return nil, err
	}

	if v, ok := resp.Error().(*BaseResp); ok {
		return resp, fmt.Errorf("status: %d, code: %d, msg: %s", resp.StatusCode(), v.Code, v.Msg)
	}

	if v, ok := res.(BaseRespInterface); ok && v.GetCode() != CodeOK {
		return resp, fmt.Errorf("code: %d, msg: %s", v.GetCode(), v.GetMsg())
	}

	return resp, nil
}

func newRestyClient() *resty.Client {
	c := resty.New()
	c.SetDebug(Debug())
	c.SetTimeout(5 * time.Minute)
	c.SetDisableWarn(true)
	c.SetLogger(&restyLogger{})

	return c
}
