package mirai

import (
	"encoding/json"
	"testing"
)

func TestMessageUnmarshal(t *testing.T) {
	data := `
{
   "code": 0,
   "msg": "",
   "data": [
      {
         "type": "FriendMessage",
         "messageChain": [
            {
               "type": "Source",
               "id": 0,
               "time": 1680242400
            },
            {
               "type": "Plain",
               "text": "Hello World"
            }
         ],
         "sender": {
            "id": 0,
            "nickname": "Bot",
            "remark": ""
         }
      }
   ]
}`
	v := &GetMessageResp{}
	e := json.Unmarshal([]byte(data), v)
	if e != nil {
		t.Fatal(e)
	}
	bs, _ := json.Marshal(v)
	t.Log(string(bs))

	if v.Data.Empty() {
		return
	}

	t.Log(v.Data[0].Type)
}
