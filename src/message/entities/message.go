package entities

type Message struct {
	Name string
}

func (t *Message) TableName() string {
	return "messages"
}
