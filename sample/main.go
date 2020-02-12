package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/torukita/go-senseway/mqtt"
)

func MessageHandler(client MQTT.Client, message MQTT.Message) {
	fmt.Printf("%s\n", string(message.Payload()))
}

func main() {
	// sample setting
	// server definition forrmat is like this (tcp://localhost:1883 ssl://localhost:8883)
	server := os.Getenv("MQTT_SERVER")
	username := os.Getenv("MQTT_USERNAME")
	password := os.Getenv("MQTT_PASSWORD")
	deveui := os.Getenv("DEVEUI")

	cfg := mqtt.Config{
		Server:   server,
		Username: username,
		Password: password,
	}

	client := mqtt.NewClientWithConfig(cfg)
	if err := client.Connect(); err != nil {
		panic(err)
	}
	if err := client.SubscribeRx(deveui, MessageHandler); err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
