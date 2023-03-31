package mirai

import (
	"os"
	"strconv"
	"testing"
)

var (
	client *Client

	botQQ       int64
	sessionKey  string
	targetQQ    int64
	targetGroup int64
)

func TestMain(m *testing.M) {
	addr := envStr("MIRAI_ADDR", true)
	verifyKey := envStr("MIRAI_VERIFY_KEY", true)
	botQQ = envInt("MIRAI_BOT_QQ", true)
	sessionKey = envStr("MIRAI_SESSION_KEY", false)
	targetQQ = envInt("MIRAI_TARGET_QQ", false)
	targetGroup = envInt("MIRAI_TARGET_GROUP", false)

	var err error

	client, err = NewClient(addr, verifyKey)
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func envStr(name string, required bool) string {
	v := os.Getenv(name)
	if required && v == "" {
		panic("'" + name + "' is not set")
	}
	return v
}

func envInt(name string, required bool) int64 {
	v := os.Getenv(name)
	if v == "" {
		if required {
			panic("'" + name + "' is not set")
		} else {
			return 0
		}
	}
	i, e := strconv.ParseInt(v, 10, 64)
	if e != nil {
		panic("'" + name + "' is not valid")
	}
	return i
}
