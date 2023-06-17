package client

import (
	"context"
	"fmt"
	"github.com/BAN1ce/gardener/metric"
	"github.com/eclipse/paho.golang/autopaho"
	"github.com/eclipse/paho.golang/paho"
	"github.com/google/uuid"
	"net/url"
)

type Client struct {
	cm *autopaho.ConnectionManager
}

func (c *Client) Start(ctx context.Context) error {
	var (
		clientID = uuid.NewString()
		err      error
	)

	cliCfg := autopaho.ClientConfig{
		BrokerUrls:     []*url.URL{{Scheme: "tcp", Host: "localhost:1883"}},
		KeepAlive:      180,
		OnConnectionUp: func(*autopaho.ConnectionManager, *paho.Connack) { fmt.Println("mqtt connection up") },
		OnConnectError: func(err error) { fmt.Printf("error whilst attempting connection: %s\n", err) },
		Debug:          paho.NOOPLogger{},
		ClientConfig: paho.ClientConfig{
			ClientID:      clientID,
			Router:        &Handler{},
			OnClientError: func(err error) { fmt.Printf("server requested disconnect: %s\n", err) },
			OnServerDisconnect: func(d *paho.Disconnect) {
				if d.Properties != nil {
					fmt.Printf("server requested disconnect: %s\n", d.Properties.ReasonString)
				} else {
					fmt.Printf("server requested disconnect; reason code: %d\n", d.ReasonCode)
				}
			},
		},
	}
	c.cm, err = autopaho.NewConnection(ctx, cliCfg)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Publish(total int, topic string, payload []byte) error {
	for i := 0; i < total; i++ {
		_, err := c.cm.Publish(context.Background(), &paho.Publish{
			Topic:   topic,
			Payload: payload,
			QoS:     0,
		})
		if err != nil {
			continue
		}
		metric.PublishedCount.WithLabelValues(topic).Inc()
	}
	return nil
}

func (c *Client) Subscribe(topic string) error {
	_, err := c.cm.Subscribe(context.Background(), &paho.Subscribe{
		Properties:    nil,
		Subscriptions: map[string]paho.SubscribeOptions{topic: {QoS: 0, RetainHandling: 0}},
	})
	return err
}
