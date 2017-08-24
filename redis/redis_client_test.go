package redis

import (
	"fmt"
	"testing"
	"time"
)

func Test_Pub(t *testing.T) {
	client := NewRedisClient()
	pubsub := client.Subscribe("test")
	defer pubsub.Close()

	subscr, err := pubsub.ReceiveTimeout(time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Println(subscr)

	err = client.Publish("test", "hello").Err()
	if err != nil {
		panic(err)
	}

	msg, err := pubsub.ReceiveMessage()
	if err != nil {
		panic(err)
	}

	fmt.Println(msg.Channel, msg.Payload)
}

type AuthenticationEvent struct {
}
