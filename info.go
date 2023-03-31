package mirai

import (
	"net/http"
)

type AboutResp struct {
	BaseResp
	Data AboutInfo `json:"data"`
}

type AboutInfo struct {
	Version string `json:"version"`
}

func (c *Client) About() (AboutInfo, *http.Response, error) {
	res := AboutResp{}
	resp, err := c.Exec(http.MethodGet, "about", nil, &res)
	if err != nil {
		return res.Data, nil, err
	}
	return res.Data, resp.RawResponse, nil
}

type BotListResp struct {
	BaseResp
	Data []int64 `json:"data"`
}

func (c *Client) BotList() ([]int64, *http.Response, error) {
	res := BotListResp{}
	resp, err := c.Exec(http.MethodGet, "botList", nil, &res)
	if err != nil {
		return res.Data, nil, err
	}
	return res.Data, resp.RawResponse, nil
}
