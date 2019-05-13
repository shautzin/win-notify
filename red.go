package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/json-iterator/go"
	"log"
)

type RedMessage struct {
	Channel MessageChannel `json:"channel"`
	Title   string         `json:"title"`
	Content string         `json:"content"`
}

const (
	RedisAddr     = "redis.lt:6379"
	RedisPassword = "********"
)

type MessageChannel string

const (
	ChannelDefault = "00"
	ChannelMobile  = "10"
	ChannelPc      = "20"
)

const MessageListNamePrefix = "MESSAGE_LIST_"

func NewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Password: RedisPassword,
		Addr:     RedisAddr,
		DB:       0,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Client connection test:", pong)
	return client
}

func Push(client *redis.Client, message RedMessage) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	bytes, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}
	client.LPush(MessageListNamePrefix+string(message.Channel), bytes)
}

func Pull(client *redis.Client, ch MessageChannel) *RedMessage {
	bytes, err := client.RPop(MessageListNamePrefix + string(ch)).Bytes()
	if err != nil {
		log.Println(err)
		return nil
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	message := RedMessage{}
	err = json.Unmarshal(bytes, &message)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &message
}
