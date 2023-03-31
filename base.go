package mirai

type BaseReqInterface interface {
	GetSessionKey() string
}

type BaseReq struct {
	SessionKey string `url:"sessionKey,omitempty" json:"sessionKey,omitempty"`
}

var _ BaseReqInterface = (*BaseReq)(nil)

func (req BaseReq) GetSessionKey() string {
	return req.SessionKey
}

// --------------------------------------------------------------------------------

type BaseRespInterface interface {
	GetCode() Code
	GetMsg() string
}

type BaseResp struct {
	Code Code   `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

var _ BaseRespInterface = (*BaseResp)(nil)

func (resp BaseResp) GetCode() Code {
	return resp.Code
}

func (resp BaseResp) GetMsg() string {
	return resp.Msg
}
