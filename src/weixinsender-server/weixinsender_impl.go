package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// WeixinSenderImpl implementaion
type WeixinSenderImpl struct {
}

func (this *WeixinSenderImpl) Ping() (r string, err error) {
	g_logger.Info("请求ping方法")
	return "pong", nil
}

func (this *WeixinSenderImpl) GetAccessToken(appid string, appsecret string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("获取公众号[%v]的访问凭证开始", appid)
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%v&secret=%v", appid, appsecret)

	outputStr = fmt.Sprintf("获取公众号[%v]的访问凭证的请求URL地址为: %v", appid, requestUrl)
	g_logger.Debug(outputStr)

	resp, err := http.Get(requestUrl)
	if err != nil {
		outputStr = fmt.Sprintf("请求获取公众号[%v]的访问凭证失败. 失败原因：%v", appid, err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(body)

	outputStr = fmt.Sprintf("获取公众号[%v]的访问凭证：%v", appid, r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("获取公众号[%v]的访问凭证失败。Http状态码：%v", appid, resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(body)
	if err != nil {
		outputStr = fmt.Sprintf("获取公众号[%v]的访问凭证失败，失败原因: %v", appid, err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("获取公众号[%v]的访问凭证成功", appid)
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) AddKFAccount(access_token string, kfaccount string, nickname string, password string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("添加客服账号[%v]开始", kfaccount)
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/customservice/kfaccount/add?access_token=%v", access_token)

	outputStr = fmt.Sprintf("添加客服账号[%v]的请求URL地址为: %v", kfaccount, requestUrl)
	g_logger.Debug(requestUrl)

	//Make Post Data
	reqBody, err := MakeKFAccountBody(kfaccount, nickname, password)
	if err != nil {
		outputStr = fmt.Sprintf("生成微信客服账号POST数据失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	resp, err := http.Post(requestUrl, "text/plain", bytes.NewReader(reqBody))
	if err != nil {
		outputStr = fmt.Sprintf("请求添加客服账号[%v]失败. 失败原因：%v", kfaccount, err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(resBody)

	outputStr = fmt.Sprintf("添加客服账号[%v]请求的状态：%v", kfaccount, r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("添加客服账号[%v]失败。Http状态码：%v", kfaccount, resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(resBody)
	if err != nil {
		outputStr = fmt.Sprintf("添加客服账号[%v]失败，失败原因: %v", kfaccount, err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("添加客服账号[%v]成功", kfaccount)
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) UpdateKFAccount(access_token string, kfaccount string, nickname string, password string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("修改客服账号[%v]开始", kfaccount)
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/customservice/kfaccount/update?access_token=%v", access_token)

	outputStr = fmt.Sprintf("修改客服账号[%v]的请求URL地址为: %v", kfaccount, requestUrl)
	g_logger.Debug(requestUrl)

	//Make Post Data
	reqBody, err := MakeKFAccountBody(kfaccount, nickname, password)
	if err != nil {
		outputStr = fmt.Sprintf("生成微信客服账号POST数据失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	resp, err := http.Post(requestUrl, "text/plain", bytes.NewReader(reqBody))
	if err != nil {
		outputStr = fmt.Sprintf("请求修改客服账号[%v]失败. 失败原因：%v", kfaccount, err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(resBody)

	outputStr = fmt.Sprintf("修改客服账号[%v]请求的状态：%v", kfaccount, r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("修改客服账号[%v]失败。Http状态码：%v", kfaccount, resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(resBody)
	if err != nil {
		outputStr = fmt.Sprintf("修改客服账号[%v]失败，失败原因: %v", kfaccount, err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("修改客服账号[%v]成功", kfaccount)
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) DeleteKFAccount(access_token string, kfaccount string, nickname string, password string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("删除客服账号[%v]开始", kfaccount)
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/customservice/kfaccount/del?access_token=%v", access_token)

	outputStr = fmt.Sprintf("删除客服账号[%v]的请求URL地址为: %v", kfaccount, requestUrl)
	g_logger.Debug(requestUrl)

	//Make Post Data
	reqBody, err := MakeKFAccountBody(kfaccount, nickname, password)
	if err != nil {
		outputStr = fmt.Sprintf("生成微信客服账号POST数据失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	resp, err := http.Post(requestUrl, "text/plain", bytes.NewReader(reqBody))
	if err != nil {
		outputStr = fmt.Sprintf("请求删除客服账号[%v]失败. 失败原因：%v", kfaccount, err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(resBody)

	outputStr = fmt.Sprintf("删除客服账号[%v]请求的状态：%v", kfaccount, r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("删除客服账号[%v]失败。Http状态码：%v", kfaccount, resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(resBody)
	if err != nil {
		outputStr = fmt.Sprintf("删除客服账号[%v]失败，失败原因: %v", kfaccount, err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("删除客服账号[%v]成功", kfaccount)
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) SetKFHeadImg(access_token string, kfaccount string, media_file string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("设置客服账号[%v]的头像开始", kfaccount)
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("http://api.weixin.qq.com/customservice/kfaccount/uploadheadimg?access_token=%v&kf_account=%v", access_token, kfaccount)

	outputStr = fmt.Sprintf("设置客服账号[%v]头像的请求URL地址为: %v", kfaccount, requestUrl)
	g_logger.Debug(requestUrl)

	//Make Post Data
	reqBody, err := MakeMediaFileBody(media_file)
	if err != nil {
		outputStr = fmt.Sprintf("生成客服头像的POST数据失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	req, err := http.NewRequest("POST", requestUrl, bytes.NewReader(reqBody))
	if err != nil {
		outputStr = fmt.Sprintf("创建HTTP POST请求失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}
	req.Header.Set("Content-Type", "multipart/form-data")
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		outputStr = fmt.Sprintf("请求设置客服账号[%v]的头像失败. 失败原因：%v", kfaccount, err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(resBody)

	outputStr = fmt.Sprintf("设置客服账号[%v]头像的状态：%v", kfaccount, r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("设置客服账号[%v]头像失败。Http状态码：%v", kfaccount, resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(resBody)
	if err != nil {
		outputStr = fmt.Sprintf("设置客服账号[%v]头像失败，失败原因: %v", kfaccount, err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("设置客服账号[%v]头像成功", kfaccount)
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) GetKFAccountList(access_token string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("获取客服账号列表开始")
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/customservice/getkflist?access_token=%v", access_token)

	outputStr = fmt.Sprintf("获取客服账号列表的请求URL地址为: %v", requestUrl)
	g_logger.Debug(outputStr)

	resp, err := http.Get(requestUrl)
	if err != nil {
		outputStr = fmt.Sprintf("请求获取客服账号列表失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(body)

	outputStr = fmt.Sprintf("获取客服账号列表：%v", r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("获取客服账号列表失败。Http状态码：%v", resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(body)
	if err != nil {
		outputStr = fmt.Sprintf("获取客服账号列表失败，失败原因: %v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("获取客服账号列表成功")
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) SendMessage(access_token string, touser string, type_a1 string, data string, kfaccount string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("给[%v]发送微信消息开始", touser)
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%v", access_token)

	outputStr = fmt.Sprintf("给[%v]发送微信消息的请求URL地址为: %v", touser, requestUrl)
	g_logger.Debug(requestUrl)

	//Make Post Data
	reqBody, err := MakeSendMessageBody(touser, type_a1, data, kfaccount)
	if err != nil {
		outputStr = fmt.Sprintf("生成发送微信消息的Body数据失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	resp, err := http.Post(requestUrl, "text/plain", bytes.NewReader(reqBody))
	if err != nil {
		outputStr = fmt.Sprintf("请求给[%v]发送微信消息失败. 失败原因：%v", touser, err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(resBody)

	outputStr = fmt.Sprintf("给[%v]发送微信消息的状态：%v", touser, r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("给[%v]发送微信消息失败。Http状态码：%v", touser, resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(resBody)
	if err != nil {
		outputStr = fmt.Sprintf("给[%v]发送微信消息失败，失败原因: %v", touser, err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("给[%v]发送微信消息成功", touser)
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) UploadTempMedia(access_token string, type_a1 string, media_file string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("上传临时媒体文件[%]开始", media_file)
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/media/upload?access_token=%v&type=%v", access_token, type_a1)

	outputStr = fmt.Sprintf("上传临时媒体文件[%v]的请求URL地址为: %v", media_file, requestUrl)
	g_logger.Debug(requestUrl)

	//Make Post Data
	reqBody, err := MakeMediaFileBody(media_file)
	if err != nil {
		outputStr = fmt.Sprintf("生成临时媒体的POST数据失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	req, err := http.NewRequest("POST", requestUrl, bytes.NewReader(reqBody))
	if err != nil {
		outputStr = fmt.Sprintf("创建HTTP POST请求失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}
	req.Header.Set("Content-Type", "multipart/form-data")
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		outputStr = fmt.Sprintf("请求上传临时媒体文件[%v]失败. 失败原因：%v", media_file, err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(resBody)

	outputStr = fmt.Sprintf("上传临时媒体文件[%v]的状态：%v", media_file, r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("上传临时媒体文件[%v]失败。Http状态码：%v", media_file, resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(resBody)
	if err != nil {
		outputStr = fmt.Sprintf("上传临时媒体文件[%v]失败，失败原因: %v", media_file, err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("上传临时媒体文件[%v]成功", media_file)
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) DownloadTempMedia(access_token string, media_id string) (r []byte, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("下载临时媒体素材[%v]开始", media_id)
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/media/get?access_token=%v&media_id=%v", access_token, media_id)

	outputStr = fmt.Sprintf("下载临时媒体素材[%v]的请求URL地址为: %v", media_id, requestUrl)
	g_logger.Debug(outputStr)

	resp, err := http.Get(requestUrl)
	if err != nil {
		outputStr = fmt.Sprintf("请求下载临时媒体素材[%v]失败. 失败原因：%v", media_id, err)
		g_logger.Error(outputStr)
		return nil, err
	}
	defer resp.Body.Close()

	r, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return nil, err
	}

	outputStr = fmt.Sprintf("下载的临时媒体素材[%v]的头信息：%v", media_id, resp.Header)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("下载临时媒体素材[%v]失败。Http状态码：%v", media_id, resp.StatusCode)
		g_logger.Error(outputStr)
		return nil, fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(r)
	if err != nil {
		outputStr = fmt.Sprintf("下载临时媒体素材[%v]失败，失败原因: %v", media_id, err)
		g_logger.Error(outputStr)
		return nil, err
	}

	outputStr = fmt.Sprintf("下载临时媒体素材[%v]成功", media_id)
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) UploadPermanentMedia(access_token string, type_a1 string, media_file string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("上传永久媒体文件[%]开始", media_file)
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("http://api.weixin.qq.com/cgi-bin/material/add_material?access_token=%v&type=%v", access_token, type_a1)

	outputStr = fmt.Sprintf("上传永久媒体文件[%v]的请求URL地址为: %v", media_file, requestUrl)
	g_logger.Debug(requestUrl)

	//Make Post Data
	reqBody, err := MakeMediaFileBody(media_file)
	if err != nil {
		outputStr = fmt.Sprintf("生成永久媒体的POST数据失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	req, err := http.NewRequest("POST", requestUrl, bytes.NewReader(reqBody))
	if err != nil {
		outputStr = fmt.Sprintf("创建HTTP POST请求失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}
	req.Header.Set("Content-Type", "multipart/form-data")
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		outputStr = fmt.Sprintf("请求上传永久媒体文件[%v]失败. 失败原因：%v", media_file, err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(resBody)

	outputStr = fmt.Sprintf("上传永久媒体文件[%v]的状态：%v", media_file, r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("上传永久媒体文件[%v]失败。Http状态码：%v", media_file, resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(resBody)
	if err != nil {
		outputStr = fmt.Sprintf("上传永久媒体文件[%v]失败，失败原因: %v", media_file, err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("上传永久媒体文件[%v]成功", media_file)
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) DownloadPermanentMedia(access_token string, media_id string) (r []byte, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("下载永久媒体素材[%v]开始", media_id)
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/material/get_material?access_token=%v", access_token)

	outputStr = fmt.Sprintf("下载永久媒体素材[%v]的请求URL地址为: %v", media_id, requestUrl)
	g_logger.Debug(outputStr)

	//Make Post Data
	reqBody, err := MakeMediaIDBody(media_id)
	if err != nil {
		outputStr = fmt.Sprintf("生成下载永久媒体素材[%v]的POST数据失败. 失败原因：%v", media_id, err)
		g_logger.Error(outputStr)
		return nil, err
	}

	resp, err := http.Post(requestUrl, "text/plain", bytes.NewReader(reqBody))
	if err != nil {
		outputStr = fmt.Sprintf("请求下载永久媒体素材[%v]失败. 失败原因：%v", media_id, err)
		g_logger.Error(outputStr)
		return nil, err
	}
	defer resp.Body.Close()

	r, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return nil, err
	}

	outputStr = fmt.Sprintf("下载的永久媒体素材[%v]的头信息：%v", media_id, resp.Header)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("下载永久媒体素材[%v]失败。Http状态码：%v", media_id, resp.StatusCode)
		g_logger.Error(outputStr)
		return nil, fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(r)
	if err != nil {
		outputStr = fmt.Sprintf("下载永久媒体素材[%v]失败，失败原因: %v", media_id, err)
		g_logger.Error(outputStr)
		return nil, err
	}

	outputStr = fmt.Sprintf("下载永久媒体素材[%v]成功", media_id)
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) DeletePermanentMedia(access_token string, media_id string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("删除永久媒体素材[%v]开始", media_id)
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/material/del_material?access_token=%v", access_token)

	outputStr = fmt.Sprintf("删除永久媒体素材[%v]的请求URL地址为: %v", media_id, requestUrl)
	g_logger.Debug(outputStr)

	//Make Post Data
	reqBody, err := MakeMediaIDBody(media_id)
	if err != nil {
		outputStr = fmt.Sprintf("生成删除永久媒体素材[%v]的POST数据失败. 失败原因：%v", media_id, err)
		g_logger.Error(outputStr)
		return "", err
	}

	resp, err := http.Post(requestUrl, "text/plain", bytes.NewReader(reqBody))
	if err != nil {
		outputStr = fmt.Sprintf("请求删除永久媒体素材[%v]失败. 失败原因：%v", media_id, err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(body)

	outputStr = fmt.Sprintf("删除的永久媒体素材[%v]的状态：%v", media_id, r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("删除永久媒体素材[%v]失败。Http状态码：%v", media_id, resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(body)
	if err != nil {
		outputStr = fmt.Sprintf("删除永久媒体素材[%v]失败，失败原因: %v", media_id, err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("删除永久媒体素材[%v]成功", media_id)
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) UploadNews(access_token string, news []byte) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("上传图文消息开始")
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/media/uploadnews?access_token=%v", access_token)

	outputStr = fmt.Sprintf("上传图文消息的请求URL地址为: %v", requestUrl)
	g_logger.Debug(requestUrl)

	resp, err := http.Post(requestUrl, "text/plain", bytes.NewReader(news))
	if err != nil {
		outputStr = fmt.Sprintf("请求上传图文消息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(resBody)

	outputStr = fmt.Sprintf("上传图文消息的状态：%v", r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("上传图文消息失败。Http状态码：%v", resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(resBody)
	if err != nil {
		outputStr = fmt.Sprintf("上传图文消息失败，失败原因: %v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("上传图文消息成功")
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) SendNews(access_token string, is_to_all bool, group_id string, msg_type string, content string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("发送图文消息开始")
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token=%v", access_token)

	outputStr = fmt.Sprintf("发送图文消息的请求URL地址为: %v", requestUrl)
	g_logger.Debug(outputStr)

	//Make Post Data
	reqBody, err := MakeSendNewsBody(is_to_all, group_id, msg_type, content)
	if err != nil {
		outputStr = fmt.Sprintf("生成图文消息的POST数据失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	resp, err := http.Post(requestUrl, "text/plain", bytes.NewReader(reqBody))
	if err != nil {
		outputStr = fmt.Sprintf("请求发送图文消息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(body)

	outputStr = fmt.Sprintf("发送图文消息的状态：%v", r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("发送图文消息失败。Http状态码：%v", resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(body)
	if err != nil {
		outputStr = fmt.Sprintf("发送图文消息失败，失败原因: %v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("发送图文消息成功")
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) DeleteNews(access_token string, msg_id string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("删除图文消息[%v]开始", msg_id)
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/mass/delete?access_token=%v", access_token)

	outputStr = fmt.Sprintf("删除图文消息[%v]的请求URL地址为: %v", msg_id, requestUrl)
	g_logger.Debug(outputStr)

	//Make Post Data
	reqBody, err := MakeMsgIDBody(msg_id)
	if err != nil {
		outputStr = fmt.Sprintf("生成删除图文消息[%v]的POST数据失败. 失败原因：%v", msg_id, err)
		g_logger.Error(outputStr)
		return "", err
	}

	resp, err := http.Post(requestUrl, "text/plain", bytes.NewReader(reqBody))
	if err != nil {
		outputStr = fmt.Sprintf("请求删除图文消息[%v]失败. 失败原因：%v", msg_id, err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(body)

	outputStr = fmt.Sprintf("删除图文消息[%v]的状态：%v", msg_id, r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("删除图文消息[%v]失败。Http状态码：%v", msg_id, resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(body)
	if err != nil {
		outputStr = fmt.Sprintf("删除图文消息[%v]失败，失败原因: %v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("删除图文消息[%v]成功", msg_id)
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) PreviewNewsByOpenId(access_token string, touser string, msg_type string, content string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("预览图文消息开始")
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/mass/preview?access_token=%v", access_token)

	outputStr = fmt.Sprintf("预览图文消息的请求URL地址为: %v", requestUrl)
	g_logger.Debug(outputStr)

	//Make Post Data
	reqBody, err := MakePreviewNewsBodyByOpenId(touser, msg_type, content)
	if err != nil {
		outputStr = fmt.Sprintf("生成图文消息的POST数据失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	resp, err := http.Post(requestUrl, "text/plain", bytes.NewReader(reqBody))
	if err != nil {
		outputStr = fmt.Sprintf("请求预览图文消息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(body)

	outputStr = fmt.Sprintf("预览图文消息的状态：%v", r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("预览图文消息失败。Http状态码：%v", resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(body)
	if err != nil {
		outputStr = fmt.Sprintf("预览图文消息失败，失败原因: %v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("预览图文消息成功")
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) PreviewNewsByWeixinName(access_token string, towxname string, msg_type string, content string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("发送给[%v]预览图文消息开始", towxname)
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/mass/preview?access_token=%v", access_token)

	outputStr = fmt.Sprintf("发送给[%v]预览图文消息的请求URL地址为: %v", towxname, requestUrl)
	g_logger.Debug(outputStr)

	//Make Post Data
	reqBody, err := MakePreviewNewsBodyByWeixinName(towxname, msg_type, content)
	if err != nil {
		outputStr = fmt.Sprintf("生成发送给[%v]的图文消息的POST数据失败. 失败原因：%v", towxname, err)
		g_logger.Error(outputStr)
		return "", err
	}

	resp, err := http.Post(requestUrl, "text/plain", bytes.NewReader(reqBody))
	if err != nil {
		outputStr = fmt.Sprintf("请求发送给[%v]预览图文消息失败. 失败原因：%v", towxname, err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(body)

	outputStr = fmt.Sprintf("发送给[%v]预览图文消息的状态：%v", towxname, r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("发送给[%v]预览图文消息失败。Http状态码：%v", towxname, resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(body)
	if err != nil {
		outputStr = fmt.Sprintf("发送给[%v]预览图文消息失败，失败原因: %v", towxname, err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("发送给[%v]预览图文消息成功", towxname)
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) GetNewsStatus(access_token string, msg_id string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("获取图文消息[%v]的状态开始", msg_id)
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/mass/get?access_token=%v", access_token)

	outputStr = fmt.Sprintf("获取图文消息[%v]的状态的请求URL地址为: %v", msg_id, requestUrl)
	g_logger.Debug(outputStr)

	//Make Post Data
	reqBody, err := MakeMsgIDBody(msg_id)
	if err != nil {
		outputStr = fmt.Sprintf("生成获取图文消息[%v]的状态的POST数据失败. 失败原因：%v", msg_id, err)
		g_logger.Error(outputStr)
		return "", err
	}

	resp, err := http.Post(requestUrl, "text/plain", bytes.NewReader(reqBody))
	if err != nil {
		outputStr = fmt.Sprintf("请求获取图文消息[%v]的状态失败. 失败原因：%v", msg_id, err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(body)

	outputStr = fmt.Sprintf("获取图文消息[%v]的状态：%v", msg_id, r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("获取图文消息[%v]的状态失败。Http状态码：%v", msg_id, resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(body)
	if err != nil {
		outputStr = fmt.Sprintf("获取图文消息[%v]的状态失败，失败原因: %v", msg_id, err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("获取图文消息[%v]的状态成功", msg_id)
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) CreateUserGroup(access_token string, group_name string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("创建用户分组[%v]开始", group_name)
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/groups/create?access_token=%v", access_token)

	outputStr = fmt.Sprintf("创建用户分组[%v]的请求URL地址为: %v", group_name, requestUrl)
	g_logger.Debug(outputStr)

	//Make Post Data
	reqBody, err := MakeCreateUserGroupBody(group_name)
	if err != nil {
		outputStr = fmt.Sprintf("生成创建用户分组[%v]的POST数据失败. 失败原因：%v", group_name, err)
		g_logger.Error(outputStr)
		return "", err
	}

	resp, err := http.Post(requestUrl, "text/plain", bytes.NewReader(reqBody))
	if err != nil {
		outputStr = fmt.Sprintf("请求创建用户分组[%v]失败. 失败原因：%v", group_name, err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(body)

	outputStr = fmt.Sprintf("创建用户分组[%v]：%v", group_name, r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("创建用户分组[%v]失败。Http状态码：%v", group_name, resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(body)
	if err != nil {
		outputStr = fmt.Sprintf("创建用户分组[%v]失败，失败原因: %v", group_name, err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("创建用户分组[%v]成功", group_name)
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) UpdateUserGroup(access_token string, group_id string, new_group_name string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("修改用户分组[%v]开始", group_id)
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/groups/update?access_token=%v", access_token)

	outputStr = fmt.Sprintf("修改用户分组[%v]的请求URL地址为: %v", group_id, requestUrl)
	g_logger.Debug(outputStr)

	//Make Post Data
	reqBody, err := MakeUpdateUserGroupBody(group_id, new_group_name)
	if err != nil {
		outputStr = fmt.Sprintf("生成修改用户分组[%v]的POST数据失败. 失败原因：%v", group_id, err)
		g_logger.Error(outputStr)
		return "", err
	}

	resp, err := http.Post(requestUrl, "text/plain", bytes.NewReader(reqBody))
	if err != nil {
		outputStr = fmt.Sprintf("请求修改用户分组[%v]失败. 失败原因：%v", group_id, err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(body)

	outputStr = fmt.Sprintf("修改用户分组[%v]：%v", group_id, r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("修改用户分组[%v]失败。Http状态码：%v", group_id, resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(body)
	if err != nil {
		outputStr = fmt.Sprintf("修改用户分组[%v]失败，失败原因: %v", group_id, err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("修改用户分组[%v]成功", group_id)
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) GetUserGroupList(access_token string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("获取用户分组列表开始")
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/groups/get?access_token=%v", access_token)

	outputStr = fmt.Sprintf("获取用户分组列表的请求URL地址为: %v", requestUrl)
	g_logger.Debug(outputStr)

	resp, err := http.Get(requestUrl)
	if err != nil {
		outputStr = fmt.Sprintf("请求获取用户分组列表失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(body)

	outputStr = fmt.Sprintf("获取用户分组列表：%v", r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("获取用户分组列表失败。Http状态码：%v", resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(body)
	if err != nil {
		outputStr = fmt.Sprintf("获取用户分组列表失败。失败原因: %v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("获取用户分组列表成功")
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) GetUserGroupByOpenID(access_token string, openid string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("查询用户[%v]的分组开始", openid)
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/groups/getid?access_token=%v", access_token)

	outputStr = fmt.Sprintf("查询用户[%v]的分组的请求URL地址为: %v", openid, requestUrl)
	g_logger.Debug(outputStr)

	//Make Post Data
	reqBody, err := MakeOpenIDBody(openid)
	if err != nil {
		outputStr = fmt.Sprintf("生成查询用户[%v]的分组的POST数据失败. 失败原因：%v", openid, err)
		g_logger.Error(outputStr)
		return "", err
	}

	resp, err := http.Post(requestUrl, "text/plain", bytes.NewReader(reqBody))
	if err != nil {
		outputStr = fmt.Sprintf("请求查询用户[%v]的分组失败. 失败原因：%v", openid, err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(body)

	outputStr = fmt.Sprintf("查询用户[%v]的分组：%v", openid, r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("查询用户[%v]的分组失败。Http状态码：%v", openid, resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(body)
	if err != nil {
		outputStr = fmt.Sprintf("查询用户[%v]的分组失败。失败原因: %v", openid, err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("查询用户[%v]的分组成功", openid)
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) MoveUserToGroup(access_token string, openid_list []string, to_groupid string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("移动用户[%v]到分组[%v]开始", openid_list, to_groupid)
	g_logger.Info(outputStr)

	var requestUrl string
	nCount := len(openid_list)
	switch {
	case nCount == 1:
		requestUrl = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/groups/members/update?access_token=%v", access_token)
	case nCount > 1:
		requestUrl = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/groups/members/batchupdate?access_token=%v", access_token)
	default:
		outputStr = fmt.Sprintf("用户列表不能空")
		g_logger.Error(outputStr)
		return "", fmt.Errorf(outputStr)
	}

	outputStr = fmt.Sprintf("移动用户[%v]到分组[%v]的请求URL地址为: %v", openid_list, to_groupid, requestUrl)
	g_logger.Debug(outputStr)

	//Make Post Data
	reqBody, err := MakeMoveUserToGroupBody(openid_list, to_groupid)
	if err != nil {
		outputStr = fmt.Sprintf("生成移动用户[%v]到分组[%v]的POST数据失败. 失败原因：%v", openid_list, to_groupid, err)
		g_logger.Error(outputStr)
		return "", err
	}

	resp, err := http.Post(requestUrl, "text/plain", bytes.NewReader(reqBody))
	if err != nil {
		outputStr = fmt.Sprintf("请求移动用户[%v]到分组[%v]失败. 失败原因：%v", openid_list, to_groupid, err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(body)

	outputStr = fmt.Sprintf("移动用户[%v]到分组[%v]：%v", openid_list, to_groupid, r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("移动用户[%v]到分组[%v]失败。Http状态码：%v", openid_list, to_groupid, resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(body)
	if err != nil {
		outputStr = fmt.Sprintf("移动用户[%v]到分组[%v]失败。失败原因: %v", openid_list, to_groupid, err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("移动用户[%v]到分组[%v]成功", openid_list, to_groupid)
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) RemarkUser(access_token string, openid string, remark string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("设置用户[%v]的备注名[%v]开始", openid, remark)
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/user/info/updateremark?access_token=%v", access_token)

	outputStr = fmt.Sprintf("设置用户[%v]的备注名[%v]的请求URL地址为: %v", openid, remark, requestUrl)
	g_logger.Debug(outputStr)

	//Make Post Data
	reqBody, err := MakeRemarkUserBody(openid, remark)
	if err != nil {
		outputStr = fmt.Sprintf("生成设置用户[%v]的备注名[%v]的POST数据失败. 失败原因：%v", openid, remark, err)
		g_logger.Error(outputStr)
		return "", err
	}

	resp, err := http.Post(requestUrl, "text/plain", bytes.NewReader(reqBody))
	if err != nil {
		outputStr = fmt.Sprintf("请求设置用户[%v]的备注名[%v]失败. 失败原因：%v", openid, remark, err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(body)

	outputStr = fmt.Sprintf("设置用户[%v]的备注名[%v]：%v", openid, remark, r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("设置用户[%v]的备注名[%v]失败。Http状态码：%v", openid, remark, resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(body)
	if err != nil {
		outputStr = fmt.Sprintf("设置用户[%v]的备注名[%v]失败。失败原因: %v", openid, remark, err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("设置用户[%v]的备注名[%v]成功", openid, remark)
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) GetUserInfo(access_token string, openid string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("获取用户[%v]的基本信息开始", openid)
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/user/info?access_token=%v&openid=%v&lang=zh_CN", access_token, openid)

	outputStr = fmt.Sprintf("获取用户[%v]的基本信息的请求URL地址为: %v", openid, requestUrl)
	g_logger.Debug(outputStr)

	resp, err := http.Get(requestUrl)
	if err != nil {
		outputStr = fmt.Sprintf("请求获取用户[%v]的基本信息失败. 失败原因：%v", openid, err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(body)

	outputStr = fmt.Sprintf("获取用户[%v]的基本信息：%v", openid, r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("获取用户[%v]的基本信息失败。Http状态码：%v", openid, resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(body)
	if err != nil {
		outputStr = fmt.Sprintf("获取用户[%v]的基本信息失败。失败原因: %v", openid, err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("获取用户[%v]的基本信息成功", openid)
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) GetUserList(access_token string, next_openid string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("获取用户列表开始")
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/user/get?access_token=%v&next_openid=%v", access_token, next_openid)

	outputStr = fmt.Sprintf("获取用户列表的请求URL地址为: %v", requestUrl)
	g_logger.Debug(outputStr)

	resp, err := http.Get(requestUrl)
	if err != nil {
		outputStr = fmt.Sprintf("请求获取用户列表失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(body)

	outputStr = fmt.Sprintf("获取用户列表：%v", r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("获取用户列表失败。Http状态码：%v", resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(body)
	if err != nil {
		outputStr = fmt.Sprintf("获取用户列表失败。失败原因: %v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("获取用户列表成功")
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) CreateMenu(access_token string, menu_data []byte) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("创建菜单开始")
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf(" https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%v", access_token)

	outputStr = fmt.Sprintf("创建菜单的请求URL地址为: %v", requestUrl)
	g_logger.Debug(outputStr)

	resp, err := http.Post(requestUrl, "text/plain", bytes.NewReader(menu_data))
	if err != nil {
		outputStr = fmt.Sprintf("请求创建菜单失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(body)

	outputStr = fmt.Sprintf("创建菜单的状态：%v", r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("创建菜单失败。Http状态码：%v", resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(body)
	if err != nil {
		outputStr = fmt.Sprintf("创建菜单失败。失败原因: %v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("创建菜单成功")
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) DeleteMenu(access_token string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("删除菜单开始")
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=%v", access_token)

	outputStr = fmt.Sprintf("删除菜单的请求URL地址为: %v", requestUrl)
	g_logger.Debug(outputStr)

	resp, err := http.Get(requestUrl)
	if err != nil {
		outputStr = fmt.Sprintf("请求删除菜单失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(body)

	outputStr = fmt.Sprintf("删除菜单的状态：%v", r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("删除菜单失败。Http状态码：%v", resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(body)
	if err != nil {
		outputStr = fmt.Sprintf("删除菜单失败。失败原因: %v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("删除菜单成功")
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) GetMenu(access_token string) (r []byte, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("获取菜单开始")
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/menu/get?access_token=%v", access_token)

	outputStr = fmt.Sprintf("获取菜单的请求URL地址为: %v", requestUrl)
	g_logger.Debug(outputStr)

	resp, err := http.Get(requestUrl)
	if err != nil {
		outputStr = fmt.Sprintf("请求获取菜单失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return nil, err
	}
	defer resp.Body.Close()

	r, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return nil, err
	}

	outputStr = fmt.Sprintf("获取菜单的状态：%v", string(r))
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("获取菜单失败。Http状态码：%v", resp.StatusCode)
		g_logger.Error(outputStr)
		return nil, fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(r)
	if err != nil {
		outputStr = fmt.Sprintf("获取菜单失败。失败原因: %v", err)
		g_logger.Error(outputStr)
		return nil, err
	}

	outputStr = fmt.Sprintf("获取菜单成功")
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) GetArticleTotal(access_token string, begin_date string, end_date string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("获取[%v]发送的图文统计数据开始", begin_date)
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/datacube/getarticletotal?access_token=%v", access_token)

	outputStr = fmt.Sprintf("获取[%v]发送的图文统计数据的请求URL地址为: %v", begin_date, requestUrl)
	g_logger.Debug(outputStr)

	//Make Post Data
	reqBody, err := MakeDateRangeBody(begin_date, end_date)
	if err != nil {
		outputStr = fmt.Sprintf("生成获取[%v]发送的图文统计数据的POST数据失败. 失败原因：%v", begin_date, err)
		g_logger.Error(outputStr)
		return "", err
	}

	resp, err := http.Post(requestUrl, "text/plain", bytes.NewReader(reqBody))
	if err != nil {
		outputStr = fmt.Sprintf("请求获取[%v]发送的图文统计数据失败. 失败原因：%v", begin_date, err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(body)

	outputStr = fmt.Sprintf("获取[%v]发送的图文统计数据的状态：%v", begin_date, r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("获取[%v]发送的图文统计数据失败。Http状态码：%v", begin_date, resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(body)
	if err != nil {
		outputStr = fmt.Sprintf("获取[%v]发送的图文统计数据失败，失败原因: %v", begin_date, err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("获取[%v]发送的图文统计数据成功", begin_date)
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) GetUserRead(access_token, begin_date, end_date string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("获取截止[%v]发送的图文统计数据开始", begin_date)
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/datacube/getuserread?access_token=%v", access_token)

	outputStr = fmt.Sprintf("获取截止[%v]发送的图文统计数据的请求URL地址为: %v", begin_date, requestUrl)
	g_logger.Debug(outputStr)

	//Make Post Data
	reqBody, err := MakeDateRangeBody(begin_date, end_date)
	if err != nil {
		outputStr = fmt.Sprintf("生成获取截止[%v]发送的图文统计数据的POST数据失败. 失败原因：%v", begin_date, err)
		g_logger.Error(outputStr)
		return "", err
	}

	resp, err := http.Post(requestUrl, "text/plain", bytes.NewReader(reqBody))
	if err != nil {
		outputStr = fmt.Sprintf("请求获取截止[%v]发送的图文统计数据失败. 失败原因：%v", begin_date, err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(body)

	outputStr = fmt.Sprintf("获取截止[%v]发送的图文统计数据的状态：%v", begin_date, r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("获取截止[%v]发送的图文统计数据失败。Http状态码：%v", begin_date, resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(body)
	if err != nil {
		outputStr = fmt.Sprintf("获取截止[%v]发送的图文统计数据失败，失败原因: %v", begin_date, err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("获取截止[%v]发送的图文统计数据成功", begin_date)
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) GetArticleSummary(access_token, begin_date, end_date string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("获取当天[%v]发送的图文统计数据开始", begin_date)
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/datacube/getarticlesummary?access_token=%v", access_token)

	outputStr = fmt.Sprintf("获取当天[%v]发送的图文统计数据的请求URL地址为: %v", begin_date, requestUrl)
	g_logger.Debug(outputStr)

	//Make Post Data
	reqBody, err := MakeDateRangeBody(begin_date, end_date)
	if err != nil {
		outputStr = fmt.Sprintf("生成获取当天[%v]发送的图文统计数据的POST数据失败. 失败原因：%v", begin_date, err)
		g_logger.Error(outputStr)
		return "", err
	}
	resp, err := http.Post(requestUrl, "text/plain", bytes.NewReader(reqBody))
	if err != nil {
		outputStr = fmt.Sprintf("请求获取当天[%v]发送的图文统计数据失败. 失败原因：%v", begin_date, err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(body)

	outputStr = fmt.Sprintf("获取当天[%v]发送的图文统计数据的状态：%v", begin_date, r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("获取当天[%v]发送的图文统计数据失败。Http状态码：%v", begin_date, resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(body)
	if err != nil {
		outputStr = fmt.Sprintf("获取当天[%v]发送的图文统计数据失败，失败原因: %v", begin_date, err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("获取当天[%v]发送的图文统计数据成功", begin_date)
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) GetUserCumulate(access_token, begin_date, end_date string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("获取累计用户数据开始")
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/datacube/getusercumulate?access_token=%v", access_token)

	//Make Post Data
	reqBody, err := MakeDateRangeBody(begin_date, end_date)
	if err != nil {
		outputStr = fmt.Sprintf("生成获取当天[%v]累计用户数据的POST数据失败. 失败原因：%v", begin_date, err)
		g_logger.Error(outputStr)
		return "", err
	}
	resp, err := http.Post(requestUrl, "text/plain", bytes.NewReader(reqBody))
	if err != nil {
		outputStr = fmt.Sprintf("请求获取当天[%v]累计用户数据失败. 失败原因：%v", begin_date, err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(body)

	outputStr = fmt.Sprintf("获取累计用户数据：%v", r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("获取累计用户数据失败。Http状态码：%v", resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(body)
	if err != nil {
		outputStr = fmt.Sprintf("获取累计用户数据失败。失败原因: %v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("获取累计用户数据成功")
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) GetUserSummary(access_token, begin_date, end_date string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("获取用户增减数据开始")
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/datacube/getusersummary?access_token=%v", access_token)

	//Make Post Data
	reqBody, err := MakeDateRangeBody(begin_date, end_date)
	if err != nil {
		outputStr = fmt.Sprintf("生成获取当天[%v]用户增减数据的POST数据失败. 失败原因：%v", begin_date, err)
		g_logger.Error(outputStr)
		return "", err
	}
	resp, err := http.Post(requestUrl, "text/plain", bytes.NewReader(reqBody))
	if err != nil {
		outputStr = fmt.Sprintf("请求获取当天[%v]用户增减数据失败. 失败原因：%v", begin_date, err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(body)

	outputStr = fmt.Sprintf("获取用户增减数据：%v", r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("获取用户增减数据失败。Http状态码：%v", resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(body)
	if err != nil {
		outputStr = fmt.Sprintf("获取用户增减数据失败。失败原因: %v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("获取用户增减数据成功")
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) GetUserReadHour(access_token, begin_date, end_date string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("获取图文统计分时数据开始")
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/datacube/getuserreadhour?access_token=%v", access_token)

	//Make Post Data
	reqBody, err := MakeDateRangeBody(begin_date, end_date)
	if err != nil {
		outputStr = fmt.Sprintf("生成获取当天[%v]图文统计分时数据的POST数据失败. 失败原因：%v", begin_date, err)
		g_logger.Error(outputStr)
		return "", err
	}
	resp, err := http.Post(requestUrl, "text/plain", bytes.NewReader(reqBody))
	if err != nil {
		outputStr = fmt.Sprintf("请求获取当天[%v]图文统计分时数据失败. 失败原因：%v", begin_date, err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(body)

	outputStr = fmt.Sprintf("获取图文统计分时数据：%v", r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("获取图文统计分时数据失败。Http状态码：%v", resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(body)
	if err != nil {
		outputStr = fmt.Sprintf("获取图文统计分时数据失败。失败原因: %v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("获取图文统计分时数据成功")
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) GetUserShare(access_token, begin_date, end_date string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("获取图文分享转发数据开始")
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/datacube/getusershare?access_token=%v", access_token)

	//Make Post Data
	reqBody, err := MakeDateRangeBody(begin_date, end_date)
	if err != nil {
		outputStr = fmt.Sprintf("生成获取当天[%v]图文分享转发数据的POST数据失败. 失败原因：%v", begin_date, err)
		g_logger.Error(outputStr)
		return "", err
	}
	resp, err := http.Post(requestUrl, "text/plain", bytes.NewReader(reqBody))
	if err != nil {
		outputStr = fmt.Sprintf("请求获取当天[%v]图文分享转发数据失败. 失败原因：%v", begin_date, err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(body)

	outputStr = fmt.Sprintf("获取图文分享转发数据：%v", r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("获取图文分享转发数据失败。Http状态码：%v", resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(body)
	if err != nil {
		outputStr = fmt.Sprintf("获取图文分享转发数据失败。失败原因: %v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("获取图文分享转发数据成功")
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) GetUserShareHour(access_token, begin_date, end_date string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("获取图文分享转发分时数据开始")
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/datacube/getusersharehour?access_token=%v", access_token)

	//Make Post Data
	reqBody, err := MakeDateRangeBody(begin_date, end_date)
	if err != nil {
		outputStr = fmt.Sprintf("生成获取当天[%v]图文分享转发分时数据的POST数据失败. 失败原因：%v", begin_date, err)
		g_logger.Error(outputStr)
		return "", err
	}
	resp, err := http.Post(requestUrl, "text/plain", bytes.NewReader(reqBody))
	if err != nil {
		outputStr = fmt.Sprintf("请求获取当天[%v]图文分享转发分时数据失败. 失败原因：%v", begin_date, err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(body)

	outputStr = fmt.Sprintf("获取图文分享转发分时数据：%v", r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("获取图文分享转发分时数据失败。Http状态码：%v", resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(body)
	if err != nil {
		outputStr = fmt.Sprintf("获取图文分享转发分时数据失败。失败原因: %v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("获取图文分享转发分时数据成功")
	g_logger.Info(outputStr)

	return
}

func (this *WeixinSenderImpl) UploadNewsImg(access_token string, media_file string) (r string, err error) {
	var outputStr string

	outputStr = fmt.Sprintf("上传临时图文图片[%v]开始", media_file)
	g_logger.Info(outputStr)

	var requestUrl string
	requestUrl = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=%v", access_token)

	outputStr = fmt.Sprintf("上传临时图文图片[%v]的请求URL地址为: %v", media_file, requestUrl)
	g_logger.Debug(requestUrl)

	//Make Post Data
	reqBody, err := MakeMediaFileBody(media_file)
	if err != nil {
		outputStr = fmt.Sprintf("生成临时媒体的POST数据失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	req, err := http.NewRequest("POST", requestUrl, bytes.NewReader(reqBody))
	if err != nil {
		outputStr = fmt.Sprintf("创建HTTP POST请求失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}
	req.Header.Set("Content-Type", "multipart/form-data")
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		outputStr = fmt.Sprintf("请求上传临时图文图片[%v]失败. 失败原因：%v", media_file, err)
		g_logger.Error(outputStr)
		return "", err
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		outputStr = fmt.Sprintf("读取HTTP响应信息失败. 失败原因：%v", err)
		g_logger.Error(outputStr)
		return "", err
	}

	r = string(resBody)

	outputStr = fmt.Sprintf("上传临时图文图片[%v]的状态：%v", media_file, r)
	g_logger.Debug(outputStr)

	if resp.StatusCode != 200 {
		outputStr = fmt.Sprintf("上传临时图文图片[%v]失败。Http状态码：%v", media_file, resp.StatusCode)
		g_logger.Error(outputStr)
		return "", fmt.Errorf("Http Status Code: %v", resp.StatusCode)
	}

	err = ParseErrMsg(resBody)
	if err != nil {
		outputStr = fmt.Sprintf("上传临时图文图片[%v]失败，失败原因: %v", media_file, err)
		g_logger.Error(outputStr)
		return "", err
	}

	outputStr = fmt.Sprintf("上传临时图文图片[%v]成功", media_file)
	g_logger.Info(outputStr)

	return
}
