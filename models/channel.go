package models

type ChannelConfig struct {
	LarkWebhookURL    string `json:"lark_webhook_url"`
	LarkWebhookSecret string `json:"lark_webhook_secret"`
	DingWebhookURL    string `json:"ding_webhook_url"`
	DingWebhookSecret string `json:"ding_webhook_secret"`
	WxWorkWebhookURL  string `json:"wx_work_webhook_url"`
	TelegramBotToken  string `json:"telegram_bot_token"`
	TelegramChatId    string `json:"telegram_chat_id"`
	BarkServer        string `json:"bark_server"`
	BarkSecret        string `json:"bark_secret"`
}
