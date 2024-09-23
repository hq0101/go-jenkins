package v1

import (
	"github.com/hq0101/go-jenkins/rest"
)

type QueueV1Interface interface {
	RESTClient() rest.Interface
	QueueGetter
}

type QueueV1Client struct {
	restClient rest.Interface
}

func (c *QueueV1Client) Queue() QueueInterface {
	return newQueue(c)
}

func New(c rest.Interface) *QueueV1Client {
	return &QueueV1Client{
		restClient: c,
	}
}

func NewForConfig(c *rest.Config) (*QueueV1Client, error) {
	config := *c

	restClient, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &QueueV1Client{
		restClient: restClient,
	}, nil
}

func NewForConfigOrDie(c *rest.Config) *QueueV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

func (c *QueueV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
