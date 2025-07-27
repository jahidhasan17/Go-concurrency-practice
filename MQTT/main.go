package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
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
	err := godotenv.Load()
    if err != nil {
        fmt.Println("Warning: .env file not found, using system environment variables")
    }

    opts := mqtt.NewClientOptions()
    var brokerAddress string = os.Getenv("MQTT_BROKER_ADDRESS")
    brokerPortStr := os.Getenv("MQTT_BROKER_PORT")
    if brokerPortStr == "" {
        panic("MQTT_BROKER_PORT environment variable not set")
    }
    brokerPort, err := strconv.Atoi(brokerPortStr)
    if err != nil {
        panic(fmt.Sprintf("Invalid MQTT_BROKER_PORT: %v", err))
    }
    opts.AddBroker(fmt.Sprintf("ssl://%s:%d", brokerAddress, brokerPort))
	opts.SetClientID("go_mqtt_concurrent_client")
	opts.SetUsername(os.Getenv("MQTT_USERNAME"))
	opts.SetPassword(os.Getenv("MQTT_PASSWORD"))
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