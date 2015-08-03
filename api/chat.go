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

type TextMessage struct {
	Receiver Receiver `json: "receiver"`
	Sender   string   `json: "sender"`
	Msgtype  string   `json: "msgtype"`
	Text     Text     `json: "text"`
}

type Receiver struct {
	Type string `json: "type"`
	Id   string `json: "id"`
}

type Text struct {
	Content string `json: "content"`
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
	data, err := json.Marshal(Chat{
		ChatId:   "1",
		Name:     "测试会话",
		Owner:    "xiaoxi___0525",
		UserList: []string{"xiaoxi___0525,uio257918"},
	})

	_, err = a.Client.PostJSON(url, data)
	return err
}

func (a *API) SendTextMessage() error {
	token, err := a.Tokener.Token()
	if err != nil {
		return err
	}

	qs := make(url.Values)
	qs.Add("access_token", token)

	url := sendChatURL + "?" + qs.Encode()

	data, err := json.Marshal(TextMessage{
		Receiver: Receiver{
			Type: "single",
			Id:   "uio257918",
		},
		Sender:  "xiaoxi___0525",
		Msgtype: "text",
		Text: Text{
			Content: "测试聊天",
		},
	})

	_, err = a.Client.PostJSON(url, data)
	return err
}
