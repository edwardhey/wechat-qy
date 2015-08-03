package api

import (
	"encoding/json"
	"net/url"
)

const (
	createChatURL  = "https://qyapi.weixin.qq.com/cgi-bin/chat/create"
	getChatURL     = "https://qyapi.weixin.qq.com/cgi-bin/chat/get"
	updateChatURL  = "https://qyapi.weixin.qq.com/cgi-bin/chat/update"
	quitChatURL    = "https://qyapi.weixin.qq.com/cgi-bin/chat/quit"
	clearnotifyURL = "https://qyapi.weixin.qq.com/cgi-bin/chat/clearnotify"
	sendChatURL    = "https://qyapi.weixin.qq.com/cgi-bin/chat/send"
	setmuteChatURL = "https://qyapi.weixin.qq.com/cgi-bin/chat/setmute"
)

type Chat struct {
	ChatId   string   `json: "chatid"`
	Name     string   `json: "name"`
	Owner    string   `json: "owner"`
	UserList []string `json: "userlist"`
}

// CreateChat 方法用于创建微信聊天
func (a *API) CreateChat(chat *Chat) error {
	token, err := a.Tokener.Token()
	if err != nil {
		return err
	}

	qs := make(url.Values)
	qs.Add("access_token", token)

	url := createChatURL + "?" + qs.Encode()

	data, err := json.Marshal(chat)
	if err != nil {
		return err
	}

	_, err = a.Client.PostJSON(url, data)
	return err
}
