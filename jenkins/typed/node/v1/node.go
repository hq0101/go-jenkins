package v1

import (
	"context"
	"fmt"
	"github.com/hq0101/go-jenkins/api/node/v1"
	"github.com/hq0101/go-jenkins/rest"
	"net/http"
)

type NodesGetter interface {
	Nodes() NodesInterface
}

type NodesInterface interface {
	GetNodes(ctx context.Context) (*v1.Node, error)
}

type nodes struct {
	client rest.Interface
}

func newNodes(c *NodeV1Client) *nodes {
	return &nodes{
		client: c.RESTClient(),
	}
}

func (c *nodes) GetNodes(ctx context.Context) (*v1.Node, error) {
	result := &v1.Node{}
	var statusCode int
	err := c.client.Get().
		AbsPath("/manage/computer/api/json").
		Do(ctx).
		StatusCode(&statusCode).
		Into(result)
	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("status code %d", statusCode)
	}
	return result, err
}
