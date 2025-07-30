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

const (
	workerPoolSize    = 5
	messageBufferSize = 100
	defaultWaitTime   = 100 * time.Millisecond
	lwtTopic          = "topic/lwt"
	subscribeTopic    = "topic/#"
	testTopic         = "topic/test"
	testMessageCount  = 20
)

type MQTTConfig struct {
	BrokerAddress string
	BrokerPort    int
	Username      string
	Password      string
}

type MQTTMessage struct {
	Topic   string
	Payload []byte
}

func main() {
	config, err := loadConfiguration()
	if err != nil {
		panic(fmt.Sprintf("Configuration error: %v", err))
	}

	msgChan := make(chan MQTTMessage, messageBufferSize)
	startWorkers(workerPoolSize, msgChan)

	client := setupMQTTClient(config, msgChan)
	defer client.Disconnect(250)

	subscribeToTopics(client)
	publishTestMessages(client)

	waitForShutdownSignal()
}

func loadConfiguration() (*MQTTConfig, error) {
	_ = godotenv.Load() // Optional .env file

	portStr := os.Getenv("MQTT_BROKER_PORT")
	if portStr == "" {
		return nil, fmt.Errorf("MQTT_BROKER_PORT environment variable not set")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("invalid MQTT_BROKER_PORT: %v", err)
	}

	return &MQTTConfig{
		BrokerAddress: os.Getenv("MQTT_BROKER_ADDRESS"),
		BrokerPort:    port,
		Username:      os.Getenv("MQTT_USERNAME"),
		Password:      os.Getenv("MQTT_PASSWORD"),
	}, nil
}

func startWorkers(workerCount int, msgChan <-chan MQTTMessage) {
	for i := 0; i < workerCount; i++ {
		go processMessages(i, msgChan)
	}
}

func processMessages(workerID int, msgChan <-chan MQTTMessage) {
	for msg := range msgChan {
		fmt.Printf("Worker %d processing: %s - %s\n", workerID, msg.Topic, string(msg.Payload))
		time.Sleep(defaultWaitTime)
	}
}

func setupMQTTClient(config *MQTTConfig, msgChan chan<- MQTTMessage) mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("ssl://%s:%d", config.BrokerAddress, config.BrokerPort))
	opts.SetClientID("go_mqtt_concurrent_client")
	opts.SetUsername(config.Username)
	opts.SetPassword(config.Password)
	opts.SetDefaultPublishHandler(createMessageHandler(msgChan))
	opts.OnConnect = onConnectHandler
	opts.OnConnectionLost = onConnectionLostHandler
	opts.SetWill(lwtTopic, "Client disconnected unexpectedly", 1, true)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(fmt.Sprintf("Connection error: %v", token.Error()))
	}

	return client
}

func createMessageHandler(msgChan chan<- MQTTMessage) mqtt.MessageHandler {
	return func(client mqtt.Client, msg mqtt.Message) {
		msgChan <- MQTTMessage{
			Topic:   msg.Topic(),
			Payload: msg.Payload(),
		}
	}
}

func onConnectHandler(client mqtt.Client) {
	fmt.Println("Connected")
}

func onConnectionLostHandler(client mqtt.Client, err error) {
	fmt.Printf("Connection lost: %v\n", err)
}

func subscribeToTopics(client mqtt.Client) {
	if token := client.Subscribe(subscribeTopic, 1, nil); token.Wait() && token.Error() != nil {
		panic(fmt.Sprintf("Subscription error: %v", token.Error()))
	}
}

func publishTestMessages(client mqtt.Client) {
	go func() {
		for i := 0; i < testMessageCount; i++ {
			msg := fmt.Sprintf("Message %d at %s", i, time.Now().Format(time.RFC3339))
			publishMessage(client, testTopic, msg)
			time.Sleep(defaultWaitTime)
		}
	}()
}

func publishMessage(client mqtt.Client, topic, message string) {
	token := client.Publish(topic, 0, false, message)
	if token.Wait() && token.Error() != nil {
		fmt.Printf("Publish error: %v\n", token.Error())
		return
	}
	fmt.Printf("Published: %s\n", message)
}

func waitForShutdownSignal() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	fmt.Println("Shutdown signal received")
}