# 金山协作群机器人消息推送封装
### 使用格式
```go
package main

import bot "github.com/ltaodream/JinShanRobot"

func main() {
	bot := bot.NewBot("your_webhook_url")
	bot.SendText("Hello, World!")
	bot.SendMarkdown("## Hello, World!")
	bot.SendLink("Title", "Text", "https://kdocs.cn", "查看详情")
}
```