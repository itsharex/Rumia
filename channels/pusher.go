package channels

import (
	"Rumia/components/cache"
	"Rumia/models"
	"fmt"
	"strings"
)

var channelMapper = map[string]func(data *models.Data) error{
	"lark":     SendLarkMessage,
	"ding":     SendDingMessage,
	"wxwork":   SendWxWorkMessage,
	"telegram": SendTelegramMessage,
	"bark":     SendBarkMessage,
	"test":     TestPushMethod,
}

func PushMessage() {
	for {
		select {
		case c1 := <-cache.PushChannel:
			go func() {
				keys := strings.Split(c1.Message.Channel, ",")
				for i := 0; i < len(keys); i++ {
					callback := channelMapper[keys[i]]
					if callback == nil {
						// TODO: 处理不支持的channel推送
						msg := fmt.Sprintf("不支持%schannel推送", c1.Message.Channel)
						if c1.Message.Feedback {
							cache.FeedbackResult.Store(c1.Ids, msg)
						} else {
							panic(msg)
						}
						continue
					}

					err := callback(c1)
					if err != nil {
						cache.RetryChannel <- c1
					} else {
						if c1.Message.Feedback {
							cache.FeedbackResult.Store(c1.Ids, "ok")
						}
					}
				}
			}()

		case c2 := <-cache.RetryChannel:
			go func() {
				if c2.RepushNumber > 2 {
					// TODO: 处理推送失败的消息
					//panic("重新推送失败次数超过上限")
					cache.FeedbackResult.Store(c2.Ids, "重新推送失败次数超过上限")
				} else if c2.RetryNumber < 5 {
					cache.RetryChannel <- c2
					c2.RetryNumber++
					fmt.Println("正在重新推送：", c2)
				} else {
					cache.PushChannel <- c2
					c2.RetryNumber = 0
					c2.RepushNumber++
					fmt.Println("正在重新推送：", c2)
				}
			}()
		}
	}
}
