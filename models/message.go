package models

type Message struct {
	Token    string `json:"token"`
	Channel  string `json:"channel"`
	Topic    string `json:"topic"`
	Title    string `json:"title"`
	Tag      string `json:"tag"`
	Content  string `json:"content"`
	Feedback bool   `json:"feedback"`
}
