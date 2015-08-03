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
	ChatId   string   `json:"chatid"`
	Name     string   `json:"name"`
	Owner    string   `json:"owner,omitempty"`
	UserList []string `json:"userlist,omitempty"`
}

type ChatTextMessage struct {
	Receiver *ChatReceiver `json:"receiver"`
	Sender   string        `json:"sender,omitempty"`
	Msgtype  string        `json:"msgtype"`
	Text     *ChatText     `json:"text"`
}

type ChatReceiver struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

type ChatText struct {
	Content string `json:"content"`
}

// CreateChat 方法用于创建微信聊天
func (a *API) CreateChat() error {
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
		UserList: []string{"xiaoxi___0525", "uio257918", "yankebin001"},
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

	data, err := json.Marshal(ChatTextMessage{
		Receiver: &ChatReceiver{
			Type: "single",
			Id:   "xiaoxi___0525",
		},
		Sender:  "xiaoxi___0525",
		Msgtype: "text",
		Text: &ChatText{
			Content: "测试聊天",
		},
	})

	_, err = a.Client.PostJSON(url, data)
	return err
}
