package main

import "Rumia/models"

var testToken = "eFNb02fjKZ6vyGzpMcxIoc"

var channelConfig = models.ChannelConfig{
	LarkWebhookURL:    "",
	LarkWebhookSecret: "",
	DingWebhookURL:    "",
	DingWebhookSecret: "",
	WxWorkWebhookURL:  "",
}

var user = models.Data{
	Message:       nil,
	ChannelConfig: &channelConfig,
	RetryNumber:   0,
	RepushNumber:  0,
}
