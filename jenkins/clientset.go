package jenkins

import (
	jobv1 "github.com/hq0101/go-jenkins/jenkins/typed/job/v1"
	nodev1 "github.com/hq0101/go-jenkins/jenkins/typed/node/v1"
	pipelineModev1 "github.com/hq0101/go-jenkins/jenkins/typed/pipeline-model-definition/v1"
	queuev1 "github.com/hq0101/go-jenkins/jenkins/typed/queue/v1"
	viewv1 "github.com/hq0101/go-jenkins/jenkins/typed/view/v1"
	"github.com/hq0101/go-jenkins/rest"
)

type Interface interface {
	NodesV1() nodev1.NodeV1Interface
	JobsV1() jobv1.JobV1Client
	ViewV1() viewv1.ViewV1Client
	PipelineModeV1() pipelineModev1.PipelineModeV1Client
	QueueV1() queuev1.QueueV1Client
}

type Clientset struct {
	nodeV1         *nodev1.NodeV1Client
	jobV1          *jobv1.JobV1Client
	viewV1         *viewv1.ViewV1Client
	pipelineModeV1 *pipelineModev1.PipelineModeV1Client
	queueV1        *queuev1.QueueV1Client
}

func (c *Clientset) NodeV1() nodev1.NodeV1Interface {
	return c.nodeV1
}

func (c *Clientset) JobV1() jobv1.JobV1Interface {
	return c.jobV1
}

func (c *Clientset) ViewV1() viewv1.ViewV1Interface {
	return c.viewV1
}

func (c *Clientset) QueueV1() queuev1.QueueV1Interface {
	return c.queueV1
}

func (c *Clientset) PipelineModeV1() pipelineModev1.PipelineModeV1Interface {
	return c.pipelineModeV1
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

	client.viewV1, err = viewv1.NewForConfig(&config)
	if err != nil {
		return nil, err
	}

	client.pipelineModeV1, err = pipelineModev1.NewForConfig(&config)
	if err != nil {
		return nil, err
	}

	client.queueV1, err = queuev1.NewForConfig(&config)
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
	client.viewV1 = viewv1.New(c)
	client.pipelineModeV1 = pipelineModev1.New(c)
	client.queueV1 = queuev1.New(c)
	return &client
}
