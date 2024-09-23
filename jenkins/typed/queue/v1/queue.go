package v1

import (
	"context"
	"fmt"
	"github.com/hq0101/go-jenkins/api/queue/v1"
	"github.com/hq0101/go-jenkins/rest"
	"net/http"
)

type QueueGetter interface {
	Queue() QueueInterface
}

type QueueInterface interface {
	GetQueue(ctx context.Context) (*v1.Queue, error)
}

type queue struct {
	client rest.Interface
}

func newQueue(c *QueueV1Client) *queue {
	return &queue{
		client: c.RESTClient(),
	}
}

func (c *queue) GetQueue(ctx context.Context) (*v1.Queue, error) {
	result := &v1.Queue{}
	var statusCode int
	err := c.client.Get().
		AbsPath("/queue/api/json").
		Do(ctx).
		StatusCode(&statusCode).
		Into(result)
	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("status code %d", statusCode)
	}
	return result, err
}
