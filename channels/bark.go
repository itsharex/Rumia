package channels

import (
	"Rumia/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type barkMessageResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func SendBarkMessage(data *models.Data) error {
	if data.ChannelConfig.BarkServer == "" || data.ChannelConfig.BarkSecret == "" {
		return errors.New("未配置 Bark 消息推送")
	}

	url := ""
	if data.Message.Title != "" {
		url = fmt.Sprintf(
			"%s/%s/%s/%s",
			data.ChannelConfig.BarkServer,
			data.ChannelConfig.BarkSecret,
			data.Message.Title,
			data.Message.Content)

	} else {
		url = fmt.Sprintf(
			"%s/%s/%s",
			data.ChannelConfig.BarkServer,
			data.ChannelConfig.BarkSecret,
			data.Message.Content)
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	var res barkMessageResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return err
	}

	if res.Code != 200 {
		return errors.New(res.Message)
	}

	return nil
}
