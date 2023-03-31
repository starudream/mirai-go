package mirai

import (
	"testing"
)

func TestClient_CountMessage(t *testing.T) {
	count, _, err := client.CountMessage(CountMessageReq{
		BaseReq: BaseReq{SessionKey: sessionKey},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("count: %d", count)
}

func TestClient_FetchMessage(t *testing.T) {
	res, _, err := client.FetchMessage(GetMessageReq{
		BaseReq: BaseReq{SessionKey: sessionKey},
		Count:   1,
	})
	if err != nil {
		t.Fatal(err)
	}
	if res.Empty() {
		return
	}
	t.Log(res[0].Type)
}

func TestClient_SendFriendMessage(t *testing.T) {
	id, _, err := client.SendFriendMessage(SendMessageReq{
		BaseReq: BaseReq{SessionKey: sessionKey},
		Target:  targetQQ,
		MessageChain: []MessageInfoInterface{
			&MessageInfoPlain{Text: "Hello"},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("id: %d", id)
}

func TestClient_SendGroupMessage(t *testing.T) {
	id, _, err := client.SendGroupMessage(SendMessageReq{
		BaseReq: BaseReq{SessionKey: sessionKey},
		Target:  targetGroup,
		MessageChain: []MessageInfoInterface{
			&MessageInfoPlain{Text: "Hello"},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("id: %d", id)
}
