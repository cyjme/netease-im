package tests

import (
	"github.com/MrSong0607/netease-im"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

var jsonTool = jsoniter.ConfigCompatibleWithStandardLibrary

func TestCreateTeam(t *testing.T) {
	members, err := jsonTool.MarshalToString([]string{"test2"})
	if err != nil {
		t.Error(err)
	}
	team := &netease.Team{
		Tname:           "testTeam",
		Owner:           "test1",
		Members:         members,
		Announcement:    "群公告测试内容",
		Intro:           "群描述测试内容",
		Msg:             "欢迎加入 testTeam 群聊",
		Magree:          0,
		Joinmode:        1,
		TeamMemberLimit: 200, //网易云信默认值为200,非必填。但因为 Team struct teamMemberLimit int 默认值为0, 所以需指定.
	}
	tid, err := client.CreateTeam(team)
	if err != nil {
		t.Error(err)
	}

	t.Log(tid)
}
