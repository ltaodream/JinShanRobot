package JinShanRobot

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type Bot struct {
	WebhookURL string
	Timeout    int //毫秒
	Retry      int
}

type Message struct {
	MsgType  string      `json:"msgtype"`
	Text     interface{} `json:"text,omitempty"`
	Markdown interface{} `json:"markdown,omitempty"`
	Link     interface{} `json:"link,omitempty"`
}

func NewBot(webhookURL string, timeout int, retry int) *Bot {
	return &Bot{WebhookURL: webhookURL, Timeout: timeout, Retry: retry}
}

func (b *Bot) SendText(content string) error {
	msg := Message{
		MsgType: "text",
		Text:    map[string]string{"content": content},
	}
	return b.send(msg)
}

func (b *Bot) SendMarkdown(text string) error {
	msg := Message{
		MsgType:  "markdown",
		Markdown: map[string]string{"text": text},
	}
	return b.send(msg)
}

func (b *Bot) SendLink(title, text, messageUrl, btnTitle string) error {
	msg := Message{
		MsgType: "link",
		Link: map[string]interface{}{
			"title":      title,
			"text":       text,
			"messageUrl": messageUrl,
			"btnTitle":   btnTitle,
		},
	}
	return b.send(msg)
}

func (b *Bot) send(msg Message) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	client := &http.Client{Timeout: time.Duration(b.Timeout) * time.Millisecond}

	for i := 0; i < b.Retry; i++ {
		resp, err := client.Post(b.WebhookURL, "application/json", bytes.NewBuffer(data))
		if err != nil {
			continue
		}
		resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			return nil
		}
	}

	return err
}
