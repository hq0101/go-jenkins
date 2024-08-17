package v1

import "github.com/hq0101/go-jenkins/rest"

type CrumbV1Interface interface {
	RESTClient() rest.Interface
}

type CrumbV1Client struct {
	restClient rest.Interface
}

func (c *CrumbV1Client) Crumb() CrumbInterface {
	return newCrumb(c)
}

func NewForConfig(c *rest.Config) (*CrumbV1Client, error) {
	config := *c
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &CrumbV1Client{client}, nil
}

func NewForConfigOrDie(c *rest.Config) *CrumbV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

func New(c rest.Interface) *CrumbV1Client {
	return &CrumbV1Client{c}
}

func (c *CrumbV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
