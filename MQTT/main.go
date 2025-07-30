package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

var (
	connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
		fmt.Println("Connected")
	}
	connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
		fmt.Printf("Connection lost: %v\n", err)
	}
)

type MQTTMessage struct {
	Topic   string
	Payload []byte
}

func messageHandler(msgChan chan<- MQTTMessage) mqtt.MessageHandler {
	return func(client mqtt.Client, msg mqtt.Message) {
		msgChan <- MQTTMessage{
			Topic:   msg.Topic(),
			Payload: msg.Payload(),
		}
	}
}

func worker(id int, msgChan <-chan MQTTMessage) {
	for msg := range msgChan {
		fmt.Printf("Worker %d processing: %s - %s\n", id, msg.Topic, string(msg.Payload))

		time.Sleep(100 * time.Millisecond)
	}
}

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

	msgChan := make(chan MQTTMessage, 100)

	for i := 0; i < 5; i++ {
		go worker(i, msgChan)
	}

	opts.SetDefaultPublishHandler(messageHandler(msgChan))
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	opts.SetWill("topic/lwt", "Client disconnected unexpectedly", 1, true)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	defer client.Disconnect(250)

	if token := client.Subscribe("topic/#", 1, nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	go func() {
		publishConcurrently(client, "topic/test", 20)
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
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
		time.Sleep(100 * time.Millisecond)
	}
}