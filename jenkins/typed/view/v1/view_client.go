package v1

import "github.com/hq0101/go-jenkins/rest"

type ViewV1Interface interface {
	RESTClient() rest.Interface
	ViewGetter
}

type ViewV1Client struct {
	restClient rest.Interface
}

func (c *ViewV1Client) View() ViewInterface {
	return newView(c)
}

func NewForConfig(c *rest.Config) (*ViewV1Client, error) {
	config := *c
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ViewV1Client{client}, nil
}

func NewForConfigOrDie(c *rest.Config) *ViewV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

func New(c rest.Interface) *ViewV1Client {
	return &ViewV1Client{c}
}

func (c *ViewV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
