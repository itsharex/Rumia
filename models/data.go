package models

type Data struct {
	Ids           string
	Message       *Message
	ChannelConfig *ChannelConfig
	RetryNumber   int
	RepushNumber  int
}
