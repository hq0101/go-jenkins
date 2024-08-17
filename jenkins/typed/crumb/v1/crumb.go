package v1

import (
	"context"
	"fmt"
	v1 "github.com/hq0101/go-jenkins/api/crumb/v1"
	"github.com/hq0101/go-jenkins/rest"
	"net/http"
)

type CrumbGetter interface {
	Crumb() CrumbInterface
}

type CrumbInterface interface {
	GetCrumb(ctx context.Context) (*v1.CrumbIssuer, error)
}

type crumb struct {
	client rest.Interface
}

func newCrumb(c *CrumbV1Client) *crumb {
	return &crumb{
		client: c.RESTClient(),
	}
}

func (c *crumb) GetCrumb(ctx context.Context) (*v1.CrumbIssuer, error) {
	crumbIssuer := &v1.CrumbIssuer{}
	var statusCode int
	err := c.client.Get().AbsPath("/crumbIssuer/api/json").Do(ctx).StatusCode(&statusCode).Into(crumbIssuer)
	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("get crumb failed, status code: %d", statusCode)
	}
	if err != nil {
		return nil, err
	}
	return crumbIssuer, nil
}
