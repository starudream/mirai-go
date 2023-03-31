package mirai

import (
	"testing"
)

func TestClient_BotProfile(t *testing.T) {
	res, _, err := client.BotProfile(BotProfileReq{
		BaseReq: BaseReq{SessionKey: sessionKey},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", res)
}

func TestClient_FriendProfile(t *testing.T) {
	res, _, err := client.FriendProfile(FriendProfileReq{
		BaseReq: BaseReq{SessionKey: sessionKey},
		Target:  targetQQ,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", res)
}
