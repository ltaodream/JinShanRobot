package JinShanRobot

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Bot struct {
	WebhookURL string
}

type Message struct {
	MsgType  string      `json:"msgtype"`
	Text     interface{} `json:"text,omitempty"`
	Markdown interface{} `json:"markdown,omitempty"`
	Link     interface{} `json:"link,omitempty"`
}

func NewBot(webhookURL string) *Bot {
	return &Bot{WebhookURL: webhookURL}
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

	resp, err := http.Post(b.WebhookURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
