package channels

import (
	"Rumia/models"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type wxWorkMessageRequest struct {
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

type wxWorkMessageResponse struct {
	Code    int    `json:"errcode"`
	Message string `json:"errmsg"`
}

func SendWxWorkMessage(data *models.Data) error {
	if data.ChannelConfig.WxWorkWebhookURL == "" {
		return errors.New("未配置企业微信群机器人消息推送")
	}

	messageRequest := wxWorkMessageRequest{
		Msgtype: "text",
		Text: struct {
			Content string `json:"content"`
		}(struct{ Content string }{Content: data.Message.Content}),
	}

	jsonData, err := json.Marshal(messageRequest)
	if err != nil {
		return err
	}

	resp, err := http.Post(
		fmt.Sprintf("%s", data.ChannelConfig.WxWorkWebhookURL), "application/json",
		bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	var res wxWorkMessageResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return err
	}

	if res.Code != 0 {
		return errors.New(res.Message)
	}

	return nil
}
