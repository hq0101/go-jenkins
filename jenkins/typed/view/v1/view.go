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
		return nil, fmt.Errorf("status code: %d", statusCode)
	}

	return categories, err
}

func (c *view) CheckIncludeRegex(ctx context.Context, name string) error {
	url := fmt.Sprintf("/view/%s/descriptorByName/hudson.model.ListView/checkIncludeRegex", name)
	var statusCode int

	err := c.client.
		Post().
		AbsPath(url).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		Do(ctx).
		StatusCode(&statusCode).Error()

	if statusCode != http.StatusOK {
		return fmt.Errorf("status code: %d", statusCode)
	}

	return err
}

// CreateJobToView 列表试图
func (c *view) CreateJobToView(ctx context.Context, viewName, jobName, crumb string) error {
	url := fmt.Sprintf("/view/%s/addJobToView", viewName)
	var statusCode int

	err := c.client.
		Post().
		AbsPath(url).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		Body(map[string]string{
			"name":          jobName,
			"Submit":        "",
			"Jenkins-Crumb": crumb,
			"json":          `{"name":"` + jobName + `","submit":"` + "" + `","Jenkins-Crumb":"` + crumb + `"}`,
		}).
		Do(ctx).
		StatusCode(&statusCode).Error()

	if statusCode != http.StatusOK {
		return fmt.Errorf("status code: %d", statusCode)
	}

	return err
}

// CreateView 我的视图
func (c *view) CreateView(ctx context.Context, name, submit, crumb string) error {
	url := "/createView"
	var statusCode int

	err := c.client.
		Post().
		AbsPath(url).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		Body(map[string]string{
			"name":          name,
			"mode":          v1.ModeMyView,
			"Submit":        submit,
			"Jenkins-Crumb": crumb,
			"json":          `{"name":"` + name + `","mode":"` + v1.ModeMyView + `","submit":"` + "" + `","Jenkins-Crumb":"` + crumb + `"}`,
		}).
		Do(ctx).
		StatusCode(&statusCode).Error()

	if statusCode != http.StatusOK {
		return fmt.Errorf("status code: %d", statusCode)
	}

	return err
}
