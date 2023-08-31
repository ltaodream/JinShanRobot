# 金山协作群机器人消息推送封装
### 使用格式
```go
package main

import bot "github.com/ltaodream/JinShanRobot"

func main() {
	bot := bot.NewBot("your_webhook_url")
	err := bot.SendText("Hello, World!")
	if err != nil {
		return
	}
	err = bot.SendMarkdown("## Hello, World!")
	if err != nil {
		return
	}
	err = bot.SendLink("Title", "Text", "https://kdocs.cn", "查看详情")
	if err != nil {
		return
	}
}
```