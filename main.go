package main

import (
	"github.com/go-toast/toast"
	"log"
	"time"
)

func main() {
	//PushMessage()
	PullMessage()
}

func PushMessage() {
	message := RedMessage{ChannelDefault, "Test Title", "Test Content"}
	log.Println("Initializing Redis client")
	client := NewClient()
	Push(client, message)
}

func PullMessage() {
	log.Println("Initializing Redis client")
	client := NewClient()

	log.Println("Start pulling messages")
	for {
		time.Sleep(10 * time.Second)

		message := Pull(client, ChannelDefault)
		if message != nil {
			log.Println("Result Message: ", message)
			ShowNotify(message)
		}
	}
}

func ShowNotify(message *RedMessage) {
	notification := toast.Notification{
		AppID:   "Microsoft.Windows.Shell.RunDialog",
		Title:   message.Title,
		Message: message.Content,
		Actions: []toast.Action{},
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}
