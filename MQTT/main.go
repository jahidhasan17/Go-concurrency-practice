package main

import (
	"fmt"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	}
	connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
		fmt.Println("Connected")
	}
	connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
		fmt.Printf("Connection lost: %v\n", err)
	}
)

func main() {
	opts := mqtt.NewClientOptions()
	var brokerAddress string = "p063f046.ala.asia-southeast1.emqxsl.com";
    var brokerPort int = 8883;
    opts.AddBroker(fmt.Sprintf("ssl://%s:%d", brokerAddress, brokerPort))
	opts.SetClientID("go_mqtt_concurrent_client")
	opts.SetUsername("jahid123")
	opts.SetPassword("jahid123")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	opts.SetWill("topic/lwt", "Client disconnected unexpectedly", 1, true)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		subscribeConcurrently(client, "topic/test")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		publishConcurrently(client, "topic/test", 5)
	}()

	wg.Wait()

	client.Disconnect(10000)
	fmt.Println("Disconnected")
}

func subscribeConcurrently(client mqtt.Client, topic string) {
	token := client.Subscribe(topic, 1, nil)
	if token.Wait() && token.Error() != nil {
		fmt.Printf("Subscribe error: %v\n", token.Error())
		return
	}
	fmt.Printf("Subscribed to topic: %s\n", topic)
}

func publishConcurrently(client mqtt.Client, topic string, count int) {
	for i := 0; i < count; i++ {
		msg := fmt.Sprintf("Message %d at %s", i, time.Now().Format(time.RFC3339))
		token := client.Publish(topic, 0, false, msg)
		if token.Wait() && token.Error() != nil {
			fmt.Printf("Publish error: %v\n", token.Error())
			continue
		}
		fmt.Printf("Published: %s\n", msg)
		time.Sleep(1 * time.Second)
	}
}