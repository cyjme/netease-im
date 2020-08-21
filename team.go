package netease

import (
	"encoding/json"
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"strconv"
)

const (
	teamCreatePoint          = neteaseBaseURL + "/team/create.action"
	teamAddPoint             = neteaseBaseURL + "/team/add.action"
	teamKickPoint            = neteaseBaseURL + "/team/kick.action"
	teamRemovePoint          = neteaseBaseURL + "/team/remove.action"
	teamUpdatePoint          = neteaseBaseURL + "/team/update.action"
	teamQueryPoint           = neteaseBaseURL + "/team/query.action"
	teamQueryDetailPoint     = neteaseBaseURL + "/team/queryDetail.action"
	teamGetMarkReadInfoPoint = neteaseBaseURL + "/team/getMarkReadInfo.action"
	teamChangeOwnerPoint     = neteaseBaseURL + "/team/changeOwner.action"
	teamAddManagerPoint      = neteaseBaseURL + "/team/addManager.action"
	teamRemoveManagerPoint   = neteaseBaseURL + "/team/removeManager.action"
	teamJoinTeamsPointPoint  = neteaseBaseURL + "/team/joinTeams.action"
	teamUpdateTeamNickPoint  = neteaseBaseURL + "/team/updateTeamNick.action"
	teamMuteTeamPoint        = neteaseBaseURL + "/team/muteTeamPoint.action"
	teamMuteTlistPoint       = neteaseBaseURL + "/team/muteTlist.action"
	teamLeavePoint           = neteaseBaseURL + "/team/leave.action"
	teamMuteTlistAllPoint    = neteaseBaseURL + "/team/muteTlistAll.action"
	teamListTeamMutePoint    = neteaseBaseURL + "/team/listTeamMute.action"
)

type CreateTeamReq struct {
	Tname           string `json:"tname"`
	Owner           string `json:"owner"`
	Members         string `json:"members"`
	Announcement    string `json:"announcement"`
	Intro           string `json:"intro"`
	Msg             string `json:"msg"`
	Magree          int    `json:"magree"`
	Joinmode        int    `json:"joinmode"`
	Custom          string `json:"custom"`
	Icon            string `json:"icon"`
	Beinvitemode    int    `json:"beinvitemode"`
	Invitemode      int    `json:"invitemode"`
	Uptinfomode     int    `json:"uptinfomode"`
	Upcustommode    int    `json:"upcustommode"`
	TeamMemberLimit int    `json:"teamMemberLimit"`
}

// teamCreatePoint 创建Team
/*
tname	String	是	群名称，最大长度64字符
owner	String	是	群主用户帐号，最大长度32字符
members	String	是	邀请的群成员列表。["aaa","bbb"](JSONArray对应的accid，如果解析出错会报414)，members与owner总和上限为200。members中无需再加owner自己的账号。
announcement	String	否	群公告，最大长度1024字符
intro	String	否	群描述，最大长度512字符
msg	String	是	邀请发送的文字，最大长度150字符
magree	int	是	管理后台建群时，0不需要被邀请人同意加入群，1需要被邀请人同意才可以加入群。其它会返回414
joinmode	int	是	群建好后，sdk操作时，0不用验证，1需要验证,2不允许任何人加入。其它返回414
custom	String	否	自定义高级群扩展属性，第三方可以跟据此属性自定义扩展自己的群属性。（建议为json）,最大长度1024字符
icon	String	否	群头像，最大长度1024字符
beinvitemode	int	否	被邀请人同意方式，0-需要同意(默认),1-不需要同意。其它返回414
invitemode	int	否	谁可以邀请他人入群，0-管理员(默认),1-所有人。其它返回414
uptinfomode	int	否	谁可以修改群资料，0-管理员(默认),1-所有人。其它返回414
upcustommode	int	否	谁可以更新群自定义属性，0-管理员(默认),1-所有人。其它返回414
teamMemberLimit	int	否	该群最大人数(包含群主)，范围：2至应用定义的最大群人数(默认:200)。其它返回414
*/
func (c *ImClient) CreateTeam(t *CreateTeamReq) (string, error) {
	param := map[string]string{}
	param["tname"] = t.Tname
	param["tname"] = t.Tname
	param["owner"] = t.Owner
	param["members"] = t.Members
	param["announcement"] = t.Announcement
	param["intro"] = t.Intro
	param["msg"] = t.Msg
	param["magree"] = strconv.Itoa(t.Magree)
	param["joinmode"] = strconv.Itoa(t.Joinmode)
	param["custom"] = t.Custom
	param["icon"] = t.Icon
	param["beinvitemode"] = strconv.Itoa(t.Beinvitemode)
	param["invitemode"] = strconv.Itoa(t.Invitemode)
	param["uptinfomode"] = strconv.Itoa(t.Uptinfomode)
	param["upcustommode"] = strconv.Itoa(t.Upcustommode)
	param["teamMemberLimit"] = strconv.Itoa(t.TeamMemberLimit)

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(teamCreatePoint)
	if err != nil {
		return "", err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return "", err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return "", err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return "", errors.New(msg)
	}

	var tid string
	err = json.Unmarshal(*jsonRes["tid"], &tid)
	if err != nil {
		return "", err
	}

	return tid, nil
}

type AddMemberToTeamReq struct {
	Tid     string `json:"tid"`
	Owner   string `json:"owner"`
	Members string `json:"members"`
	Magree  int    `json:"magree"`
	Msg     string `json:"msg"`
	Attach  string `json:"attach"`
}

//teamAddPoint 拉人入群
/*
tid	String	是	网易云通信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
owner	String	是	用户帐号，最大长度32字符，按照群属性invitemode传入
members	String	是	["aaa","bbb"](JSONArray对应的accid，如果解析出错会报414)，一次最多拉200个成员
magree	int	是	管理后台建群时，0不需要被邀请人同意加入群，1需要被邀请人同意才可以加入群。其它会返回414
msg	String	是	邀请发送的文字，最大长度150字符
attach	String	否	自定义扩展字段，最大长度512
*/
func (c *ImClient) AddMemberToTeam(tid string, t *AddMemberToTeamReq) (string, error) {
	param := map[string]string{}
	param["tid"] = tid

	param["owner"] = t.Owner
	param["members"] = t.Members
	param["magree"] = strconv.Itoa(t.Magree)
	param["msg"] = t.Msg
	if t.Attach != "" {
		param["attach"] = t.Attach
	}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(teamAddPoint)
	if err != nil {
		return "", err
	}

	fmt.Println("resp content", string(resp.Body()))
	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return "", err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return "", err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return "", errors.New(msg)
	}

	//faccid 可能不存在。
	//faccid := struct {
	//	Accid []string `json:"accid"`
	//	Msg   string   `json:"msg"`
	//}{
	//}
	//err = json.Unmarshal(*jsonRes["faccid"], &faccid)
	//if err != nil {
	//	return "", err
	//}
	//if len(faccid.Accid) > 0 {
	//	return "", errors.New(faccid.Msg + ":" + strings.Join(faccid.Accid, ","))
	//}

	return tid, nil
}

type RemoveMemberFromTeamReq struct {
	Tid     string `json:"tid"`
	Owner   string `json:"owner"`
	Member  string `json:"member"`
	Members string `json:"members"`
	Attach  string `json:"attach"`
}

//teamKickPoint 踢人出群
/*
tid 	String	是	网易云通信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
owner	String	是	管理员的accid，用户帐号，最大长度32字符
member	String	否	被移除人的accid，用户账号，最大长度32字符;注：member或members任意提供一个，优先使用member参数
members	String	否	["aaa","bbb"]（JSONArray对应的accid，如果解析出错，会报414）一次最多操作200个accid; 注：member或members任意提供一个，优先使用member参数
attach	String	否	自定义扩展字段，最大长度512
*/
func (c *ImClient) RemoveMemberFromTeam(tid string, t *RemoveMemberFromTeamReq) (string, error) {
	param := map[string]string{}
	param["tid"] = tid

	param["tid"] = t.Tid
	param["owner"] = t.Owner
	param["member"] = t.Member
	param["members"] = t.Members
	if t.Attach != "" {
		param["attach"] = t.Attach
	}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(teamKickPoint)
	if err != nil {
		return "", err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return "", err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return "", err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return "", errors.New(msg)
	}

	//faccid := struct {
	//	Accid []string `json:"accid"`
	//	Msg   string   `json:"msg"`
	//}{
	//}
	//err = json.Unmarshal(*jsonRes["faccid"], &faccid)
	//if err != nil {
	//	return "", err
	//}
	//if len(faccid.Accid) > 0 {
	//	return "", errors.New(faccid.Msg + ":" + strings.Join(faccid.Accid, ","))
	//}

	return tid, nil
}

//teamDeleteTeamPoint 解散群
/*
tid  	String	是	网易云通信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
owner	String	是	群主用户帐号，最大长度32字符
*/
func (c *ImClient) DeleteTeam(tid string, owner string) (string, error) {
	param := map[string]string{}
	param["tid"] = tid
	param["owner"] = owner
	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(teamRemovePoint)
	if err != nil {
		return "", err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return "", err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return "", err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return "", errors.New(msg)
	}

	return tid, nil
}

type UpdateTeamReq struct {
	Tid             string `json:"tid"`
	Tname           string `json:"tname"`
	Owner           string `json:"owner"`
	Announcement    string `json:"announcement"`
	Intro           string `json:"intro"`
	Joinmode        int    `json:"joinmode"`
	Custom          string `json:"custom"`
	Icon            string `json:"icon"`
	Beinvitemode    int    `json:"beinvitemode"`
	Invitemode      int    `json:"invitemode"`
	Uptinfomode     int    `json:"uptinfomode"`
	Upcustommode    int    `json:"upcustommode"`
	TeamMemberLimit int    `json:"teamMemberLimit"`
}

//teamUpdateTeamPoint 编辑群资料
/*
tid	String	是	网易云通信服务器产生，群唯一标识，创建群时会返回
tname	String	否	群名称，最大长度64字符
owner	String	是	群主用户帐号，最大长度32字符
announcement	String	否	群公告，最大长度1024字符
intro	String	否	群描述，最大长度512字符
joinmode	int	否	群建好后，sdk操作时，0不用验证，1需要验证,2不允许任何人加入。其它返回414
custom	String	否	自定义高级群扩展属性，第三方可以跟据此属性自定义扩展自己的群属性。（建议为json）,最大长度1024字符
icon	String	否	群头像，最大长度1024字符
beinvitemode	int	否	被邀请人同意方式，0-需要同意(默认),1-不需要同意。其它返回414
invitemode	int	否	谁可以邀请他人入群，0-管理员(默认),1-所有人。其它返回414
uptinfomode	int	否	谁可以修改群资料，0-管理员(默认),1-所有人。其它返回414
upcustommode	int	否	谁可以更新群自定义属性，0-管理员(默认),1-所有人。其它返回414
teamMemberLimit	int	否	该群最大人数(包含群主)，范围：2至应用定义的最大群人数(默认:200)。其它返回414
*/
func (c *ImClient) UpdateTeam(tid string, t *UpdateTeamReq) (string, error) {
	param := map[string]string{}
	param["tid"] = tid

	param["tname"] = t.Tname
	param["owner"] = t.Owner
	param["announcement"] = t.Announcement
	param["intro"] = t.Intro
	param["joinmode"] = strconv.Itoa(t.Joinmode)
	param["custom"] = t.Custom
	param["icon"] = t.Icon
	param["beinvitemode"] = strconv.Itoa(t.Beinvitemode)
	param["invitemode"] = strconv.Itoa(t.Invitemode)
	param["uptinfomode"] = strconv.Itoa(t.Uptinfomode)
	param["upcustommode"] = strconv.Itoa(t.Upcustommode)
	param["teamMemberLimit"] = strconv.Itoa(t.TeamMemberLimit)

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(teamUpdatePoint)
	if err != nil {
		return "", err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return "", err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return "", err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return "", errors.New(msg)
	}

	return tid, nil
}

//teamQueryTeamPoint 群信息与成员列表查询
/*
tids	String	是	群id列表，如["3083","3084"]
ope	int	是	1表示带上群成员列表，0表示不带群成员列表，只返回群信息
ignoreInvalid	Boolean	否	是否忽略无效的tid，默认为false。设置为true时将忽略无效tid，并在响应结果中返回无效的tid
*/
func (c *ImClient) QueryTeam(tids string, ope int, ignoreInvalid bool) ([]TeamInfoInQueryAction, error) {
	teamDetails := []TeamInfoInQueryAction{}
	param := map[string]string{}
	param["tids"] = tids
	param["ope"] = strconv.Itoa(ope)
	if ignoreInvalid {
		param["ignoreInvalid"] = "true"
	} else {
		param["ignoreInvalid"] = "false"
	}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(teamQueryPoint)
	if err != nil {
		return teamDetails, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return teamDetails, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return teamDetails, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return teamDetails, errors.New(msg)
	}

	err = json.Unmarshal(*jsonRes["tinfos"], &teamDetails)
	if err != nil {
		return teamDetails, err
	}

	return teamDetails, nil
}

//teamQueryDetailPoint 获取群组详细信息
/*
tid	long	是	群id，群唯一标识，创建群时会返回
*/
func (c *ImClient) QueryTeamDetail(tid string) (TeamDetailInfo, error) {
	param := map[string]string{}
	var teamDetail TeamDetailInfo
	param["tid"] = tid
	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(teamQueryDetailPoint)
	if err != nil {
		return teamDetail, err
	}
	fmt.Println(string(resp.Body()), "000000000")

	var jsonRes map[string]*json.RawMessage
	err = json.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return teamDetail, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return teamDetail, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return teamDetail, errors.New(msg)
	}

	err = json.Unmarshal(*jsonRes["tinfo"], &teamDetail)
	if err != nil {
		return teamDetail, err
	}

	fmt.Println(teamDetail.Owner)

	return teamDetail, nil
}

//teamGetMarkReadInfoPoint 获取群组已读消息的已读详情信息

//teamChangeOwnerPoint 移交群主

//teamAddManagerPoint 认命管理员

//teamRemoveManagerPoint 移除管理员

//teamJoinTeamsPointPoint 获取某用户所加入的群信息

//teamUpdateTeamNickPoint 修改群昵称

//teamMuteTeamPoint 修改消息提醒开关

//teamMuteTlistPoint 禁言群成员

//teamLeavePoint 主动退群

//teamMuteTlistAllPoint 将群组整体禁言

//teamListTeamMutePoint 获取群组禁言列表
