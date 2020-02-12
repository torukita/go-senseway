package mqtt

import (
	"crypto/tls"
	"fmt"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/pkg/errors"
)

// Server: forrmat is tcp://localhost:1883 ssl://localhost:8883
type Config struct {
	Server   string
	Username string
	Password string
}

type Client struct {
	config     Config
	MQTTClient paho.Client
}

func NewClientWithConfig(config Config) *Client {
	server := config.Server
	username := config.Username
	password := config.Password

	opts := paho.NewClientOptions().AddBroker(server).SetCleanSession(true)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetConnectTimeout(3 * time.Second)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	opts.SetTLSConfig(tlsConfig)

	client := &Client{
		config:     config,
		MQTTClient: paho.NewClient(opts),
	}
	return client
}

func defaultOnConnect(c paho.Client) {
	fmt.Printf("connected to the mqtt server\n")
}

func defaultReceiveMessage(c paho.Client, m paho.Message) {
	fmt.Printf("received message\n")
}

func (c *Client) Connect() error {
	if token := c.MQTTClient.Connect(); token.Wait() && token.Error() != nil {
		return errors.Wrap(token.Error(), "failed to connect server")
	}
	return nil
}

// helper to subscribe
func (c *Client) SubscribeTopic(topic string, handler paho.MessageHandler) error {
	if token := c.MQTTClient.Subscribe(topic, 0, handler); token.Wait() && token.Error() != nil {
		return errors.Wrap(token.Error(), "failed to subscribe")
	}
	return nil
}

// helper to publish
func (c *Client) PublishTopic(topic string, bytes []byte) error {
	if token := c.MQTTClient.Publish(topic, 0, false, bytes); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (c *Client) SubscribeRx(deveui string, handler paho.MessageHandler) error {
	username := c.config.Username
	topic := fmt.Sprintf("lora/%s/%s/rx", username, deveui)
	return c.SubscribeTopic(topic, handler)
}

func (c *Client) SubscribeAck(deveui string, handler paho.MessageHandler) error {
	username := c.config.Username
	topic := fmt.Sprintf("lora/%s/%s/ack", username, deveui)
	return c.SubscribeTopic(topic, handler)
}

func (c *Client) SubscribeTx(deveui string, handler paho.MessageHandler) error {
	username := c.config.Username
	topic := fmt.Sprintf("lora/%s/%s/tx", username, deveui)
	return c.SubscribeTopic(topic, handler)
}

func (c *Client) SubscribeServerTx(deveui string, handler paho.MessageHandler) error {
	username := c.config.Username
	topic := fmt.Sprintf("lora/%s/%s/tx_send", username, deveui)
	return c.SubscribeTopic(topic, handler)
}

func (c *Client) SubscribeAll(deveui string, handler paho.MessageHandler) error {
	username := c.config.Username
	topic := fmt.Sprintf("lora/%s/%s/#", username, deveui)
	return c.SubscribeTopic(topic, handler)
}

func (c *Client) PublishTx(deveui string, bytes []byte) error {
	username := c.config.Username
	topic := fmt.Sprintf("lora/%s/%s/tx", username, deveui)
	return c.PublishTopic(topic, bytes)
}

// helper to subscribe with paho client
func SubscribeTopic(client paho.Client, topic string, handler paho.MessageHandler) error {
	if token := client.Subscribe(topic, 0, handler); token.Wait() && token.Error() != nil {
		return errors.Wrap(token.Error(), "failed to subscribe")
	}
	return nil
}

// helper to publish with paho client
func PublishTopic(client paho.Client, topic string, bytes []byte) error {
	if token := client.Publish(topic, 0, false, bytes); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
