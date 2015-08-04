package api

import (
	"encoding/json"
	"net/url"
	"time"
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
func (a *API) CreateChat(chatTitle string, chatOwner string, chatUserlist []string) error {
	token, err := a.Tokener.Token()
	if err != nil {
		return err
	}

	qs := make(url.Values)
	qs.Add("access_token", token)

	url := createChatURL + "?" + qs.Encode()

	chatId := time.Now().UnixNano().(string) + chatOwner

	data, err := json.Marshal(Chat{
		ChatId:   chatId,
		Name:     chatTitle,
		Owner:    chatOwner,
		UserList: chatUserlist,
	})

	_, err = a.Client.PostJSON(url, data)
	return err
}

func (a *API) SendTextMessage(toUser string, fromUser string, content string) error {
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
			Id:   toUser,
		},
		Sender:  fromUser,
		Msgtype: "text",
		Text: &ChatText{
			Content: content,
		},
	})

	_, err = a.Client.PostJSON(url, data)
	return err
}
