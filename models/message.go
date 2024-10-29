package models

type Message struct {
	Sender  string `json:"sender"`
	Content string `json:"content"`
}

func NewMessage(sender string, content string) *Message {
	return &Message{
		Sender:  sender,
		Content: content,
	}
}
