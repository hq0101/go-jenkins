package jenkins

import (
	jobv1 "github.com/hq0101/go-jenkins/jenkins/typed/job/v1"
	nodev1 "github.com/hq0101/go-jenkins/jenkins/typed/node/v1"
	"github.com/hq0101/go-jenkins/rest"
)

type Interface interface {
	NodesV1() nodev1.NodeV1Interface
	JobsV1() jobv1.JobV1Client
}

type Clientset struct {
	nodeV1 *nodev1.NodeV1Client
	jobV1  *jobv1.JobV1Client
}

func (c *Clientset) NodeV1() nodev1.NodeV1Interface {
	return c.nodeV1
}

func (c *Clientset) Jobs() jobv1.JobV1Interface {
	return c.jobV1
}

func NewForConfig(c *rest.Config) (*Clientset, error) {
	config := *c

	client := &Clientset{}
	var err error

	client.nodeV1, err = nodev1.NewForConfig(&config)
	if err != nil {
		return nil, err
	}
	client.jobV1, err = jobv1.NewForConfig(&config)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var client Clientset
	client.nodeV1 = nodev1.New(c)
	client.jobV1 = jobv1.New(c)
	return &client
}
