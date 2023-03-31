package mirai

import (
	"net/http"
)

type CountMessageReq struct {
	BaseReq
}

type CountMessageResp struct {
	BaseResp
	Data int `json:"data"`
}

func (c *Client) CountMessage(req CountMessageReq) (int, *http.Response, error) {
	res := CountMessageResp{}
	resp, err := c.Exec(http.MethodGet, "countMessage", req, &res)
	if err != nil {
		return res.Data, nil, err
	}
	return res.Data, resp.RawResponse, nil
}

type GetMessageReq struct {
	BaseReq
	Count int `url:"count,omitempty" json:"count,omitempty"`
}

type GetMessageResp struct {
	BaseResp
	Data Messages `json:"data"`
}

func (c *Client) getMessage(path string, req GetMessageReq) (Messages, *http.Response, error) {
	res := GetMessageResp{}
	resp, err := c.Exec(http.MethodGet, path, req, &res)
	if err != nil {
		return res.Data, nil, err
	}
	return res.Data, resp.RawResponse, nil
}

func (c *Client) FetchMessage(req GetMessageReq) (Messages, *http.Response, error) {
	return c.getMessage("fetchMessage", req)
}

func (c *Client) FetchLatestMessage(req GetMessageReq) (Messages, *http.Response, error) {
	return c.getMessage("fetchLatestMessage", req)
}

func (c *Client) PeakMessage(req GetMessageReq) (Messages, *http.Response, error) {
	return c.getMessage("peekMessage", req)
}

func (c *Client) PeakLatestMessage(req GetMessageReq) (Messages, *http.Response, error) {
	return c.getMessage("peekLatestMessage", req)
}

type SendMessageReq struct {
	BaseReq
	Target int64 `json:"target,omitempty"`
	QQ     int64 `json:"qq,omitempty"`
	Group  int64 `json:"group,omitempty"`
	Quote  int64 `json:"quote,omitempty"`

	MessageChain []MessageInfoInterface `json:"messageChain,omitempty"`
}

func (req SendMessageReq) fill() SendMessageReq {
	for i := range req.MessageChain {
		req.MessageChain[i].setType()
	}
	return req
}

type SendMessageResp struct {
	BaseResp
	MessageId int64 `json:"messageId"`
}

func (c *Client) SendFriendMessage(req SendMessageReq) (int64, *http.Response, error) {
	res := SendMessageResp{}
	resp, err := c.Exec(http.MethodPost, "sendFriendMessage", req.fill(), &res)
	if err != nil {
		return res.MessageId, nil, err
	}
	return res.MessageId, resp.RawResponse, nil
}

func (c *Client) SendGroupMessage(req SendMessageReq) (int64, *http.Response, error) {
	res := SendMessageResp{}
	resp, err := c.Exec(http.MethodPost, "sendGroupMessage", req.fill(), &res)
	if err != nil {
		return res.MessageId, nil, err
	}
	return res.MessageId, resp.RawResponse, nil
}

func (c *Client) SendTempMessage(req SendMessageReq) (int64, *http.Response, error) {
	res := SendMessageResp{}
	resp, err := c.Exec(http.MethodPost, "sendTempMessage", req.fill(), &res)
	if err != nil {
		return res.MessageId, nil, err
	}
	return res.MessageId, resp.RawResponse, nil
}
