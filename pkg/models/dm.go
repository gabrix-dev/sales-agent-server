package models

type User struct {
	ID string `json:"id,omitempty"`
}

type Message struct {
	ID        string `json:"mid,omitempty"`
	IsDeleted bool   `json:"is_deleted,omitempty"`
	Text      string `json:"text,omitempty"`
}

type Messaging struct {
	Sender    User    `json:"sender,omitempty"`
	Recipient User    `json:"recipient,omitempty"`
	Timestamp int64   `json:"timestamp,omitempty"`
	Message   Message `json:"message,omitempty"`
}

type Entry struct {
	Time      int64       `json:"time,omitempty"`
	ID        string      `json:"id,omitempty"`
	Messaging []Messaging `json:"messaging,omitempty"`
}

type DMNotification struct {
	Object string  `json:"object,omitempty"`
	Entry  []Entry `json:"entry,omitempty"`
}
