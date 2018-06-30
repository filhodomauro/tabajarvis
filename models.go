package main

type Event struct {
	Object  string  `json:"object,omitempty"`
	Entries []Entry `json:"entry,omitempty"`
}

type Entry struct {
	ID        string      `json:"id,omitempty"`
	Time      int         `json:"time,omitempty"`
	Messaging []Messaging `messaging:"messaging,omitempty"`
}

type Messaging struct {
	Sender    User    `json:"sender,omitempty"`
	Recipient User    `json:"recipient,omitempty"`
	Timestamp int     `json:"timestamp,omitempty"`
	Message   Message `json:"message,omitempty"`
}

type User struct {
	ID string `json:"id,omitempty"`
}

type Message struct {
	MID  string `json:"mid,omitempty"`
	Text string `json:"text,omitempty"`
}
