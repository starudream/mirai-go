package mirai

import (
	"testing"
)

func TestClient_VerifyAndBind(t *testing.T) {
	sk, _, err := client.Verify()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("sessionKey: %s", sk)

	_, err = client.Bind(BindReq{
		BaseReq: BaseReq{SessionKey: sk},
		QQ:      botQQ,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("bind success")

	res, _, err := client.SessionInfo(SessionInfoReq{
		BaseReq: BaseReq{SessionKey: sk},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("sessionInfo: %#v", res)

}

func TestClient_Release(t *testing.T) {
	_, err := client.Release(ReleaseReq{
		BaseReq: BaseReq{SessionKey: sessionKey},
		QQ:      botQQ,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("release success")
}
