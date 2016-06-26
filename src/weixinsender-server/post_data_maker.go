package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

const (
	MSGTYPE_TEXT   = "text"
	MSGTYPE_IMAGE  = "image"
	MSGTYPE_VOICE  = "voice"
	MSGTYPE_VIDEO  = "video"
	MSGTYPE_MUSIC  = "music"
	MSGTYPE_NEWS   = "news"
	MSGTYPE_MPNEWS = "mpnews"
)

/*
POST数据示例：
{
    "kf_account" : "test1@test",
    "nickname" : "客服1",
    "password" : "pswmd5",
}
*/
func MakeKFAccountBody(kfaccount, nickname, password string) (body []byte, err error) {
	mBody := map[string]string{
		"kf_account": kfaccount,
		"nickname":   nickname,
		"password":   password,
	}
	return json.Marshal(mBody)
}

/*
各消息类型所需的JSON数据包如下：

1. 发送文本消息
{
    "touser":"OPENID",
    "msgtype":"text",
    "text":
    {
        "content":"Hello World"
    }
}

2. 发送图片消息
{
    "touser":"OPENID",
    "msgtype":"image",
    "image":
    {
        "media_id":"MEDIA_ID"
    }
}

3. 发送语音消息
{
    "touser":"OPENID",
    "msgtype":"voice",
    "voice":
    {
        "media_id":"MEDIA_ID"
    }
}

4. 发送视频消息
{
    "touser":"OPENID",
    "msgtype":"video",
    "video":
    {
        "media_id":"MEDIA_ID",
        "thumb_media_id":"MEDIA_ID",
        "title":"TITLE",
        "description":"DESCRIPTION"
    }
}

5. 发送音乐消息
{
    "touser":"OPENID",
    "msgtype":"music",
    "music":
    {
        "title":"MUSIC_TITLE",
        "description":"MUSIC_DESCRIPTION",
        "musicurl":"MUSIC_URL",
        "hqmusicurl":"HQ_MUSIC_URL",
        "thumb_media_id":"THUMB_MEDIA_ID"
    }
}

6. 发送图文消息 图文消息条数限制在10条以内，注意，如果图文数超过10，则将会无响应。
{
    "touser":"OPENID",
    "msgtype":"news",
    "news":{
        "articles": [
        {
            "title":"Happy Day",
            "description":"Is Really A Happy Day",
            "url":"URL",
            "picurl":"PIC_URL"
        },
        {
            "title":"Happy Day",
            "description":"Is Really A Happy Day",
            "url":"URL",
            "picurl":"PIC_URL"
        }
        ]
    }
}

7. 如果需要以某个客服帐号来发消息（在微信6.0.2及以上版本中显示自定义头像），则需在JSON数据包的后半部分加入customservice参数，例如发送文本消息则改为：
{
    "touser":"OPENID",
    "msgtype":"text",
    "text":
    {
        "content":"Hello World"
    },
    "customservice":
    {
        "kf_account": "test1@kftest"
    }
}
*/

func MakeSendMessageBody(touser, type_a1, data, kf_account string) (body []byte, err error) {
	mBody := map[string]interface{}{
		"touser":  touser,
		"msgtype": type_a1,
	}

	var v interface{}
	err = json.Unmarshal([]byte(data), &v)
	if err != nil {
		return nil, err
	}

	switch type_a1 {
	case MSGTYPE_TEXT:
		mBody[MSGTYPE_TEXT] = v
	case MSGTYPE_IMAGE:
		mBody[MSGTYPE_IMAGE] = v
	case MSGTYPE_VOICE:
		mBody[MSGTYPE_VOICE] = v
	case MSGTYPE_VIDEO:
		mBody[MSGTYPE_VIDEO] = v
	case MSGTYPE_MUSIC:
		mBody[MSGTYPE_MUSIC] = v
	case MSGTYPE_NEWS:
		mBody[MSGTYPE_NEWS] = v
	default:
		return nil, fmt.Errorf("not support madia type: %v", type_a1)
	}

	if kf_account != "" {
		mBody["customservice"] = map[string]string{
			"kf_account": kf_account,
		}
	}

	return json.Marshal(mBody)
}

func MakeMediaFileBody(media_file string) (body []byte, err error) {

	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)

	fw, err := w.CreateFormFile("media", media_file)
	if err != nil {
		return nil, err
	}

	if strings.HasPrefix(media_file, "http://") || strings.HasPrefix(media_file, "https://") {

		resp, err := http.Get(media_file)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		pix, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		_, err = io.Copy(fw, bytes.NewReader(pix))
		if err != nil {
			return nil, err
		}
	} else {

		fd, err := os.Open(media_file)
		if err != nil {
			return nil, err
		}
		defer fd.Close()

		_, err = io.Copy(fw, fd)
		if err != nil {
			return nil, err
		}

	}

	w.Close()

	return buf.Bytes(), nil
}

func MakeMediaIDBody(media_id string) (body []byte, err error) {
	mBody := map[string]string{
		"media_id": media_id,
	}
	return json.Marshal(mBody)
}

type NewsFilter struct {
	is_to_all bool
	group_id  string
}

func MakeSendNewsBody(is_to_all bool, group_id string, msg_type string, content string) (body []byte, err error) {
	mBody := map[string]interface{}{
		"filter": NewsFilter{
			is_to_all: is_to_all,
			group_id:  group_id,
		},
		"msgtype": msg_type,
	}

	switch msg_type {
	case MSGTYPE_TEXT:
		mBody[MSGTYPE_TEXT] = map[string]string{"content": content}
	case MSGTYPE_IMAGE:
		mBody[MSGTYPE_IMAGE] = map[string]string{"media_id": content}
	case MSGTYPE_VOICE:
		mBody[MSGTYPE_VOICE] = map[string]string{"media_id": content}
	case MSGTYPE_MPNEWS:
		mBody[MSGTYPE_MPNEWS] = map[string]string{"media_id": content}
	default:
		return nil, fmt.Errorf("not support madia type: %v", msg_type)
	}

	return json.Marshal(mBody)
}

func MakeMsgIDBody(msg_id string) (body []byte, err error) {
	mBody := map[string]string{
		"msg_id": msg_id,
	}
	return json.Marshal(mBody)
}

func MakePreviewNewsBodyByOpenId(touser string, msg_type string, content string) (body []byte, err error) {
	mBody := map[string]interface{}{
		"touser":  touser,
		"msgtype": msg_type,
	}

	switch msg_type {
	case MSGTYPE_TEXT:
		mBody[MSGTYPE_TEXT] = map[string]string{"content": content}
	case MSGTYPE_IMAGE:
		mBody[MSGTYPE_IMAGE] = map[string]string{"media_id": content}
	case MSGTYPE_VOICE:
		mBody[MSGTYPE_VOICE] = map[string]string{"media_id": content}
	case MSGTYPE_MPNEWS:
		mBody[MSGTYPE_MPNEWS] = map[string]string{"media_id": content}
	default:
		return nil, fmt.Errorf("not support madia type: %v", msg_type)
	}

	return json.Marshal(mBody)
}

func MakePreviewNewsBodyByWeixinName(towxname string, msg_type string, content string) (body []byte, err error) {
	mBody := map[string]interface{}{
		"towxname": towxname,
		"msgtype":  msg_type,
	}

	switch msg_type {
	case MSGTYPE_TEXT:
		mBody[MSGTYPE_TEXT] = map[string]string{"content": content}
	case MSGTYPE_IMAGE:
		mBody[MSGTYPE_IMAGE] = map[string]string{"media_id": content}
	case MSGTYPE_VOICE:
		mBody[MSGTYPE_VOICE] = map[string]string{"media_id": content}
	case MSGTYPE_MPNEWS:
		mBody[MSGTYPE_MPNEWS] = map[string]string{"media_id": content}
	default:
		return nil, fmt.Errorf("not support madia type: %v", msg_type)
	}

	return json.Marshal(mBody)
}

func MakeCreateUserGroupBody(group_name string) (body []byte, err error) {
	mBody := map[string]interface{}{
		"group": map[string]string{"name": group_name},
	}
	return json.Marshal(mBody)
}

func MakeUpdateUserGroupBody(group_id, group_name string) (body []byte, err error) {
	mBody := map[string]interface{}{
		"group": map[string]string{
			"id":   group_id,
			"name": group_name,
		},
	}
	return json.Marshal(mBody)
}

func MakeOpenIDBody(openid string) (body []byte, err error) {
	mBody := map[string]string{
		"openid": openid,
	}
	return json.Marshal(mBody)
}

func MakeMoveUserToGroupBody(openid_list []string, to_groupid string) (body []byte, err error) {
	var mBody interface{}
	nCount := len(openid_list)
	switch {
	case nCount == 1:
		mBody = map[string]string{
			"openid":     openid_list[0],
			"to_groupid": to_groupid,
		}
	case nCount > 1:
		mBody = map[string]interface{}{
			"openid_list": openid_list,
			"to_groupid":  to_groupid,
		}
	}
	return json.Marshal(mBody)
}

func MakeRemarkUserBody(openid, remark string) (body []byte, err error) {
	mBody := map[string]string{
		"openid": openid,
		"remark": remark,
	}
	return json.Marshal(mBody)
}

type WeixinErrorInfo struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func ParseErrMsg(result []byte) error {
	var info WeixinErrorInfo
	err := json.Unmarshal(result, &info)
	if err == nil {
		if info.Errcode != 0 {
			return fmt.Errorf("%v: %v", info.Errcode, info.Errmsg)
		}
	}
	return nil
}

func MakeDateRangeBody(begin_date, end_date string) (body []byte, err error) {
	mBody := map[string]string{
		"begin_date": begin_date,
		"end_date":   end_date,
	}
	return json.Marshal(mBody)
}
