package models

type User struct {
	Username      string         `json:"username"`
	Password      string         `json:"password"`
	ChannelConfig *ChannelConfig `json:"channel_config"`
}
