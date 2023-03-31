package mirai

import (
	"net/http"
)

type Sex string

const (
	SexUnknown = "UNKNOWN"
	SexMale    = "MALE"
	SexFemale  = "FEMALE"
)

type BotProfileReq struct {
	BaseReq
}

type ProfileResp struct {
	BaseResp
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Age      int64  `json:"age"`
	Level    int64  `json:"level"`
	Sign     string `json:"sign"`
	Sex      Sex    `json:"sex"`
}

func (c *Client) BotProfile(req BotProfileReq) (ProfileResp, *http.Response, error) {
	res := ProfileResp{}
	resp, err := c.Exec(http.MethodGet, "botProfile", req, &res)
	if err != nil {
		return res, nil, err
	}
	return res, resp.RawResponse, nil
}

type FriendProfileReq struct {
	BaseReq
	Target int64 `url:"target,omitempty" json:"target,omitempty"`
}

func (c *Client) FriendProfile(req FriendProfileReq) (ProfileResp, *http.Response, error) {
	res := ProfileResp{}
	resp, err := c.Exec(http.MethodGet, "friendProfile", req, &res)
	if err != nil {
		return res, nil, err
	}
	return res, resp.RawResponse, nil
}
