package v1

import (
	"github.com/hq0101/go-jenkins/rest"
)

type NodeV1Interface interface {
	RESTClient() rest.Interface
	NodesGetter
}

type NodeV1Client struct {
	restClient rest.Interface
}

func (c *NodeV1Client) Nodes() NodesInterface {
	return newNodes(c)
}

func New(c rest.Interface) *NodeV1Client {
	return &NodeV1Client{
		restClient: c,
	}
}

func NewForConfig(c *rest.Config) (*NodeV1Client, error) {
	config := *c

	restClient, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &NodeV1Client{
		restClient: restClient,
	}, nil
}

func NewForConfigOrDie(c *rest.Config) *NodeV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

func (c *NodeV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
