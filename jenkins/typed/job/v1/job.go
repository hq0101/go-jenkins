package v1

import (
	"context"
	"fmt"
	v1 "github.com/hq0101/go-jenkins/api/job/v1"
	"github.com/hq0101/go-jenkins/rest"
	"net/http"
)

type JobsGetter interface {
	Jobs() JobsInterface
}

type JobsInterface interface {
	GetConfigXML(ctx context.Context, name string) (string, error)
	AllBuilds(ctx context.Context, name string) ([]v1.Build, error)
}

type jobs struct {
	client rest.Interface
}

func newJobs(c *JobV1Client) *jobs {
	return &jobs{
		client: c.restClient,
	}
}

func (c *jobs) BuildJob(ctx context.Context, name string, sec int) error {
	var statusCode int
	err := c.client.Get().
		AbsPath(fmt.Sprintf("/job/%s/build?delay=%dsec", name, sec)).
		Do(ctx).StatusCode(&statusCode).Error()
	if err != nil {
		return err
	}

	if statusCode != http.StatusOK {
		return fmt.Errorf("status code %d", statusCode)
	}
	return err
}

func (c *jobs) BuildJobWithParameters(ctx context.Context, name string) error {
	var statusCode int
	err := c.client.Get().
		AbsPath(fmt.Sprintf("/job/%s/buildWithParameters", name)).
		Do(ctx).StatusCode(&statusCode).Error()
	if err != nil {
		return err
	}

	if statusCode != http.StatusOK {
		return fmt.Errorf("status code %d", statusCode)
	}
	return nil
}

func (c *jobs) GetConfigXML(ctx context.Context, name string) (string, error) {
	var ret string
	var statusCode int
	err := c.client.Get().AbsPath(fmt.Sprintf("/job/%s/config.xml", name)).Do(ctx).StatusCode(&statusCode).Into(&ret)
	if err != nil {
		return "", err
	}

	if statusCode != http.StatusOK {
		return "", fmt.Errorf("status code %d", statusCode)
	}

	return ret, nil
}

func (c *jobs) AllBuilds(ctx context.Context, name string) ([]v1.Build, error) {
	var ret []v1.Build
	var statusCode int

	err := c.client.
		Get().
		AbsPath(fmt.Sprintf("/job/%s/api/json", name)).
		Param("tree", "allBuilds").
		Do(ctx).StatusCode(&statusCode).Into(&ret)
	if statusCode != http.StatusOK {
		return ret, fmt.Errorf("status code %d", statusCode)
	}
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func (c *jobs) ConfirmRename(ctx context.Context, name, newName, crumb string) error {
	url := fmt.Sprintf("/job/%s/confirmRename", name)
	var statusCode int

	err := c.client.Post().AbsPath(url).SetHeader("Content-Type", "application/xml").Do(ctx).StatusCode(&statusCode).Error()

	if err != nil {
		return err
	}

	if statusCode != http.StatusOK {
		return fmt.Errorf("status code %d", statusCode)
	}
	return nil
}

func (c *jobs) GetConsoleText(ctx context.Context, name string, pipelineID int) (string, error) {
	url := fmt.Sprintf("/job/%s/%d/consoleText", name, pipelineID)
	var statusCode int
	var result string

	err := c.client.Get().
		AbsPath(url).
		SetHeader("Content-Type", "application/xml").
		Do(ctx).
		StatusCode(&statusCode).Into(&result)

	if err != nil {
		return "", err
	}

	if statusCode != http.StatusOK {
		return "", fmt.Errorf("status code %d", statusCode)
	}

	return result, nil
}
