package v1

import "github.com/hq0101/go-jenkins/rest"

type PipelineModeV1Interface interface {
	RESTClient() rest.Interface
	PipelineModeGetter
}

type PipelineModeV1Client struct {
	restClient rest.Interface
}

func (c *PipelineModeV1Client) PipelineMode() PipelineModeInterface {
	return newPipelineMode(c)
}

func NewForConfig(c *rest.Config) (*PipelineModeV1Client, error) {
	config := *c
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &PipelineModeV1Client{client}, nil
}

func NewForConfigOrDie(c *rest.Config) *PipelineModeV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

func New(c rest.Interface) *PipelineModeV1Client {
	return &PipelineModeV1Client{c}
}

func (c *PipelineModeV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
