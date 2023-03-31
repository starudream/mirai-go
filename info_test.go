package mirai

import (
	"testing"
)

func TestClient_About(t *testing.T) {
	res, _, err := client.About()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", res)
}

func TestClient_BotList(t *testing.T) {
	res, _, err := client.BotList()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", res)
}
