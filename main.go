package main

import (
	"Rumia/channels"
	"Rumia/components/cache"
	"Rumia/components/server"
	"Rumia/models"
	"math/rand"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

func pushHandle(c server.Context) {
	msg := models.Message{
		Token:    "",
		Channel:  "test",
		Topic:    "",
		Title:    "",
		Tag:      "",
		Content:  "",
		Feedback: false,
	}
	err := c.DecodeJson(&msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, server.H{"msg": "解析json数据失败"})
		return
	}

	if msg.Token == "" || msg.Content == "" {
		c.JSON(http.StatusBadRequest, server.H{"msg": "Bad Request"})
		return
	}

	if msg.Token != testToken {
		c.JSON(http.StatusUnauthorized, server.H{"msg": "token错误或者没有权限推送消息"})
		return
	}

	currentTime := strconv.Itoa(int(time.Now().Unix())/rand.Intn(1024) + rand.Intn(1024))
	data := models.Data{
		Ids:           currentTime,
		Message:       &msg,
		ChannelConfig: user.ChannelConfig,
		RetryNumber:   0,
		RepushNumber:  0,
	}
	cache.PushChannel <- &data

	if msg.Feedback {
		var ok bool
		var result any
		for {
			result, ok = cache.FeedbackResult.LoadAndDelete(currentTime)
			if ok {
				break
			}
			runtime.Gosched()
		}
		c.JSON(http.StatusOK, server.H{"msg": result.(string)})
	} else {
		c.JSON(http.StatusOK, server.H{"msg": "ok"})
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())

	go channels.PushMessage()

	svr := server.NewServer()
	svr.POST("/push", pushHandle)
	svr.Run(":8080")
}
