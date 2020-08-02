package tests

import (
	"encoding/json"
	"os"
	"testing"

	netease "github.com/MrSong0607/netease-im"
)

var (
	appKey    = "41f5431e11d5dd372fc58fbef33eeae1"
	appSecret = "e792ded9b0bc"
)
var client = netease.CreateImClient(appKey, appSecret, "")

//var client = netease.CreateImClient("", "", "")

func init() {
	os.Setenv("GOCACHE", "off")
}

func TestToken(t *testing.T) {
	user := &netease.ImUser{ID: "test2", Name: "test2", Gender: 1}
	tk, err := client.CreateImUser(user)
	if err != nil {
		t.Error(err)
	}
	t.Log(tk)
}

func TestRefreshToken(t *testing.T) {
	tk, err := client.RefreshToken("7")
	if err != nil {
		t.Error(err)
	}
	b, err := json.Marshal(tk)
	t.Log(string(b), err)
}

func Benchmark_SyncMap(b *testing.B) {
	netease.CreateImClient("", "", "")
}
