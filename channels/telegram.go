package channels

import (
	"Rumia/models"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type telegramMessageRequest struct {
	ChatId    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

type telegramMessageResponse struct {
	Ok          bool   `json:"ok"`
	Description string `json:"description"`
}

func SendTelegramMessage(data *models.Data) error {
	if data.ChannelConfig.TelegramBotToken == "" || data.ChannelConfig.TelegramChatId == "" {
		return errors.New("未配置 Telegram 机器人消息推送")
	}
	messageRequest := telegramMessageRequest{
		ChatId:    data.ChannelConfig.TelegramChatId,
		Text:      data.Message.Content,
		ParseMode: "markdown",
	}

	jsonData, err := json.Marshal(messageRequest)
	if err != nil {
		return err
	}

	resp, err := http.Post(
		fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", data.ChannelConfig.TelegramBotToken),
		"application/json",
		bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	var res telegramMessageResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return err
	}

	if !res.Ok {
		return errors.New(res.Description)
	}

	return nil
}
