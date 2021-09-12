package models

type Message struct {
	Name string
}

func (t *Message) CollectionName() string {
	return "messages"
}
