package v1

import (
	"context"
	"fmt"
	v1 "github.com/hq0101/go-jenkins/api/view/v1"
	"github.com/hq0101/go-jenkins/rest"
	"net/http"
)

type ViewGetter interface {
	View() ViewInterface
}

type ViewInterface interface {
	GetCategories(ctx context.Context, name string, depth int) (*v1.Categories, error)
}

type view struct {
	client rest.Interface
}

func newView(c *ViewV1Client) *view {
	return &view{
		client: c.RESTClient(),
	}
}

func (c *view) GetCategories(ctx context.Context, name string, depth int) (*v1.Categories, error) {
	url := fmt.Sprintf("/view/%s/itemCategories", name)

	categories := &v1.Categories{}
	var statusCode int
	err := c.client.
		Get().
		AbsPath(url).
		Param("depth", fmt.Sprintf("%d", depth)).
		Do(ctx).
		StatusCode(&statusCode).Into(categories)
	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("get crumb failed, status code: %d", statusCode)
	}
	if err != nil {
		return nil, err
	}
	return categories, nil
}
