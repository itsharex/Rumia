package channels

import (
	"Rumia/models"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type larkMessageRequest struct {
	Timestamp string `json:"timestamp"`
	Sign      string `json:"sign"`
	MsgType   string `json:"msg_type"`
	Content   struct {
		Text string `json:"text"`
	} `json:"content"`
}

type larkMessageResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func larkSign(secret string, timestamp int64) (string, error) {
	// https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN?lang=zh-CN
	//timestamp + key 做sha256, 再进行base64 encode

	stringToSign := fmt.Sprintf("%v", timestamp) + "\n" + secret

	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	_, err := h.Write(data)
	if err != nil {
		return "", err
	}

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature, nil
}

func SendLarkMessage(data *models.Data) error {
	if data.ChannelConfig.LarkWebhookURL == "" {
		return errors.New("未配置飞书群机器人消息推送")
	}

	timestamp := time.Now().Unix()
	sign, err := larkSign(data.ChannelConfig.LarkWebhookSecret, timestamp)
	if err != nil {
		return err
	}

	messageRequest := larkMessageRequest{
		Timestamp: strconv.FormatInt(timestamp, 10),
		Sign:      sign,
		MsgType:   "text",
		Content: struct {
			Text string `json:"text"`
		}(struct{ Text string }{Text: data.Message.Content}),
	}

	jsonData, err := json.Marshal(messageRequest)
	if err != nil {
		return err
	}

	resp, err := http.Post(
		data.ChannelConfig.LarkWebhookURL,
		"application/json",
		bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	var res larkMessageResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return err
	}

	if res.Code != 0 {
		return errors.New(res.Msg)
	}
	return nil
}
