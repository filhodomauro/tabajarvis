package main

type Event struct {
	Object  string  `json:"object"`
	Entries []Entry `json:"entry"`
}

type Entry struct {
	ID        string      `json:"id"`
	Time      int         `json:"time"`
	Messaging []Messaging `messaging:"messaging"`
}

type Messaging struct {
	Sender    User    `json:"sender"`
	Recipient User    `json:"recipient"`
	Timestamp int     `json:"timestamp"`
	Message   Message `json:"message"`
}

type User struct {
	ID string `json:"id"`
}

type Message struct {
	MID  string `json:"mid"`
	Text string `json:"text"`
}
