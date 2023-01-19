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
	"net/url"
	"time"
)

type dingMessageRequest struct {
	MessageType string `json:"msgtype"`
	Text        struct {
		Content string `json:"content"`
	} `json:"text"`
}

type dingMessageResponse struct {
	Code    int    `json:"errcode"`
	Message string `json:"errmsg"`
}

func dingSign(secret string, timestamp int64) (string, error) {
	// https://open.dingtalk.com/document/robots/customize-robot-security-settings
	// 使用HmacSHA256算法计算签名，然后进行Base64 encode
	// "timestamp" + \n + "secret"
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, secret)
	h := hmac.New(sha256.New, []byte(secret))
	_, err := h.Write([]byte(stringToSign))
	if err != nil {
		return "", err
	}

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	signature = url.QueryEscape(signature)
	return signature, nil
}

func SendDingMessage(data *models.Data) error {
	if data.ChannelConfig.DingWebhookURL == "" {
		return errors.New("未配置钉钉群机器人消息推送")
	}

	timestamp := time.Now().UnixMilli()
	sign, err := dingSign(data.ChannelConfig.DingWebhookSecret, timestamp)
	if err != nil {
		return err
	}

	messageRequest := dingMessageRequest{
		MessageType: "text",
		Text: struct {
			Content string `json:"content"`
		}(struct{ Content string }{Content: data.Message.Content}),
	}

	jsonData, err := json.Marshal(messageRequest)
	if err != nil {
		return err
	}

	resp, err := http.Post(
		fmt.Sprintf("%s&timestamp=%d&sign=%s", data.ChannelConfig.DingWebhookURL, timestamp, sign),
		"application/json",
		bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	var res dingMessageResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return err
	}

	if res.Code != 0 {
		return errors.New(res.Message)
	}
	return nil
}
