package mirai

import (
	"net/http"
)

type VerifyReq struct {
	VerifyKey string `json:"verifyKey,omitempty"`
}

type VerifyResp struct {
	BaseResp
	Session string `json:"session"`
}

func (c *Client) Verify() (string, *http.Response, error) {
	res := VerifyResp{}
	resp, err := c.Exec(http.MethodPost, "verify", VerifyReq{VerifyKey: c.verifyKey}, &res)
	if err != nil {
		return "", nil, err
	}
	return res.Session, resp.RawResponse, nil
}

type BindReq struct {
	BaseReq
	QQ int64 `json:"qq"`
}

func (c *Client) Bind(req BindReq) (*http.Response, error) {
	res := BaseResp{}
	resp, err := c.Exec(http.MethodPost, "bind", req, &res)
	if err != nil {
		return nil, err
	}
	return resp.RawResponse, nil
}

type SessionInfoReq struct {
	BaseReq
}

type SessionInfoResp struct {
	BaseResp
	Data SessionInfoData `json:"data"`
}

type SessionInfoData struct {
	SessionKey string        `json:"sessionKey"`
	QQ         SessionInfoQQ `json:"qq"`
}

type SessionInfoQQ struct {
	Id       int64  `json:"id"`
	Nickname string `json:"nickname"`
	Remark   string `json:"remark"`
}

func (c *Client) SessionInfo(req SessionInfoReq) (SessionInfoQQ, *http.Response, error) {
	res := SessionInfoResp{}
	resp, err := c.Exec(http.MethodGet, "sessionInfo", req, &res)
	if err != nil {
		return res.Data.QQ, nil, err
	}
	return res.Data.QQ, resp.RawResponse, nil
}

type ReleaseReq struct {
	BaseReq
	QQ int64 `json:"qq"`
}

func (c *Client) Release(req ReleaseReq) (*http.Response, error) {
	res := BaseResp{}
	resp, err := c.Exec(http.MethodPost, "release", req, &res)
	if err != nil {
		return nil, err
	}
	return resp.RawResponse, nil
}
