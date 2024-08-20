package v1

import "github.com/hq0101/go-jenkins/rest"

type JobV1Interface interface {
	RESTClient() rest.Interface
	JobsGetter
}

type JobV1Client struct {
	restClient rest.Interface
}

func (c *JobV1Client) Jobs() JobsInterface {
	return newJobs(c)
}

func NewForConfig(c *rest.Config) (*JobV1Client, error) {
	config := *c

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &JobV1Client{restClient: client}, nil
}

func NewForConfigOrDie(c *rest.Config) *JobV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

func New(c rest.Interface) *JobV1Client {
	return &JobV1Client{restClient: c}
}

func (c *JobV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
