package v1

import (
	"context"
	"fmt"
	v1 "github.com/hq0101/go-jenkins/api/job/v1"
	"github.com/hq0101/go-jenkins/rest"
	"net/http"
	"net/url"
)

type JobsGetter interface {
	Jobs() JobsInterface
}

type JobsInterface interface {
	BuildJob(ctx context.Context, name string, sec int) error
	BuildJobWithParameters(ctx context.Context, name string) error
	GetConfigXML(ctx context.Context, name string) (string, error)
	AllBuilds(ctx context.Context, name string) ([]v1.Build, error)
	ConfirmRename(ctx context.Context, name, newName, crumb string) error
	GetConsoleText(ctx context.Context, name string, pipelineID int) (string, error)
	CreateWorkflowMultiBranchProject(ctx context.Context, viewName, name, jenkinsCrumb string) error
	CreateWorkflowJob(ctx context.Context, viewName, name, jenkinsCrumb string) error
	CopyJob(ctx context.Context, viewName, newJobName, oldJobName, jenkinsCrumb string) error
	DeleteJob(ctx context.Context, name, crumb string) error
	RemoveJobFromView(ctx context.Context, viewName, jobName, crumb string) error
	GetPipelineRuns(ctx context.Context, name string) ([]v1.PipelineRun, error)
	DescribeJobRun(ctx context.Context, name string, number int) (*v1.PipelineRun, error)
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
		Do(ctx).
		StatusCode(&statusCode).
		Error()

	if statusCode != http.StatusOK {
		return fmt.Errorf("status code %d", statusCode)
	}
	return err
}

func (c *jobs) BuildJobWithParameters(ctx context.Context, name string) error {
	var statusCode int
	err := c.client.Get().
		AbsPath(fmt.Sprintf("/job/%s/buildWithParameters", name)).
		Do(ctx).
		StatusCode(&statusCode).
		Error()

	if statusCode != http.StatusOK {
		return fmt.Errorf("status code %d", statusCode)
	}
	return err
}

func (c *jobs) GetConfigXML(ctx context.Context, name string) (string, error) {
	var ret string
	var statusCode int
	err := c.client.
		Get().
		AbsPath(fmt.Sprintf("/job/%s/config.xml", name)).
		Do(ctx).
		StatusCode(&statusCode).
		Into(&ret)

	if statusCode != http.StatusOK {
		return "", fmt.Errorf("status code %d", statusCode)
	}

	return ret, err
}

func (c *jobs) AllBuilds(ctx context.Context, name string) ([]v1.Build, error) {
	var ret []v1.Build
	var statusCode int

	err := c.client.
		Get().
		AbsPath(fmt.Sprintf("/job/%s/api/json", name)).
		Param("tree", "allBuilds").
		Do(ctx).StatusCode(&statusCode).
		Into(&ret)
	if statusCode != http.StatusOK {
		return ret, fmt.Errorf("status code %d", statusCode)
	}

	return ret, err
}

func (c *jobs) ConfirmRename(ctx context.Context, name, newName, crumb string) error {
	var statusCode int

	err := c.client.
		Post().
		AbsPath(fmt.Sprintf("/job/%s/confirmRename", name)).
		SetHeader("Content-Type", "application/xml").
		Do(ctx).
		StatusCode(&statusCode).
		Error()

	if statusCode != http.StatusOK {
		return fmt.Errorf("status code %d", statusCode)
	}

	return err
}

func (c *jobs) GetConsoleText(ctx context.Context, name string, pipelineID int) (string, error) {
	var statusCode int
	var result string

	err := c.client.Get().
		AbsPath(fmt.Sprintf("/job/%s/%d/consoleText", name, pipelineID)).
		SetHeader("Content-Type", "application/xml").
		Do(ctx).
		StatusCode(&statusCode).
		Into(&result)

	if statusCode != http.StatusOK {
		return "", fmt.Errorf("status code %d", statusCode)
	}

	return result, err
}

func (c *jobs) CreateWorkflowMultiBranchProject(ctx context.Context, viewName, name, jenkinsCrumb string) error {
	var statusCode int

	err := c.client.
		Post().
		AbsPath(fmt.Sprintf("/view/%s/createItem", viewName)).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		Body(url.Values{
			"name":          {name},
			"mode":          {v1.ModeWorkflowMultiBranchProject},
			"Jenkins-Crumb": {jenkinsCrumb},
			"Json":          {`{"name":"` + name + `","mode":"` + v1.ModeWorkflowMultiBranchProject + `","Jenkins-Crumb":"` + jenkinsCrumb + `"}`},
		}).
		Do(ctx).
		StatusCode(&statusCode).Error()

	if statusCode != http.StatusOK {
		return fmt.Errorf("status code %d", statusCode)
	}

	return err
}

func (c *jobs) CreateWorkflowJob(ctx context.Context, viewName, name, jenkinsCrumb string) error {
	var statusCode int

	err := c.client.
		Post().
		AbsPath(fmt.Sprintf("/view/%s/createItem", viewName)).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		Body(url.Values{
			"name":          {name},
			"mode":          {v1.ModeWorkflowJob},
			"Jenkins-Crumb": {jenkinsCrumb},
			"json":          {`{"name":"` + name + `","mode":"` + v1.ModeWorkflowJob + `","Jenkins-Crumb":"` + jenkinsCrumb + `"}`},
		}).
		Do(ctx).
		StatusCode(&statusCode).Error()

	if statusCode != http.StatusOK {
		return fmt.Errorf("status code %d", statusCode)
	}

	return err
}

// CopyJob 复制现有的 Jenkins Job
func (c *jobs) CopyJob(ctx context.Context, viewName, newJobName, oldJobName, jenkinsCrumb string) error {
	var statusCode int

	err := c.client.
		Post().
		AbsPath(fmt.Sprintf("/view/%s/createItem", viewName)).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		Body(map[string]string{
			"name":          newJobName,
			"mode":          v1.ModeCopy,
			"from":          oldJobName,
			"Jenkins-Crumb": jenkinsCrumb,
			"Json":          `{"name":"` + newJobName + `","mode":"` + v1.ModeCopy + `","Jenkins-Crumb":"` + jenkinsCrumb + `"}`,
		}).
		Do(ctx).
		StatusCode(&statusCode).Error()

	if statusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d", statusCode)
	}

	return err
}

func (c *jobs) DeleteJob(ctx context.Context, name, crumb string) error {
	var statusCode int

	err := c.client.
		Post().
		AbsPath(fmt.Sprintf("/job/%s/doDelete", name)).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		Body(map[string]string{
			"Jenkins-Crumb": crumb,
		}).
		Do(ctx).
		StatusCode(&statusCode).Error()

	if statusCode != http.StatusFound {
		return fmt.Errorf("unexpected status code %d", statusCode)
	}

	return err
}

// RemoveJobFromView 只适用"列表视图"
func (c *jobs) RemoveJobFromView(ctx context.Context, viewName, jobName, crumb string) error {
	var statusCode int

	err := c.client.
		Post().
		AbsPath(fmt.Sprintf("/view/%s/removeJobFromView", viewName)).
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

func (c *jobs) Description(ctx context.Context, name string) error {
	var statusCode int

	err := c.client.
		Get().
		AbsPath(fmt.Sprintf("/job/%s/description", name)).
		SetHeader("Content-Type", "application/xml").
		Do(ctx).
		StatusCode(&statusCode).
		Error()
	if statusCode != http.StatusOK {
		return fmt.Errorf("status code: %d", statusCode)
	}

	return err
}

func (c *jobs) GetPipelineRuns(ctx context.Context, name string) ([]v1.PipelineRun, error) {
	var statusCode int

	var ret []v1.PipelineRun
	err := c.client.
		Get().
		AbsPath(fmt.Sprintf("/job/%s/wfapi/runs", name)).
		Do(ctx).
		StatusCode(&statusCode).
		Into(&ret)
	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", statusCode)
	}

	return ret, err
}

func (c *jobs) DescribeJobRun(ctx context.Context, name string, number int) (*v1.PipelineRun, error) {
	var statusCode int

	var ret v1.PipelineRun
	err := c.client.
		Get().
		AbsPath(fmt.Sprintf("/job/%s/%d/wfapi/describe", name, number)).
		Do(ctx).
		StatusCode(&statusCode).
		Into(&ret)
	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", statusCode)
	}
	return &ret, err
}

func (c *jobs) PauseJob(ctx context.Context, name string, number int) error {
	var statusCode int

	err := c.client.
		Get().
		AbsPath(fmt.Sprintf("/job/%s/%d/pause/toggle", name, number)).
		Do(ctx).
		StatusCode(&statusCode).
		Error()
	if statusCode != http.StatusOK {
		return fmt.Errorf("status code: %d", statusCode)
	}
	return err
}

func (c *jobs) StopJob(ctx context.Context, name string, number int) error {
	var statusCode int

	err := c.client.
		Get().
		AbsPath(fmt.Sprintf("/job/%s/%d/stop", name, number)).
		Do(ctx).
		StatusCode(&statusCode).
		Error()
	if statusCode != http.StatusOK {
		return fmt.Errorf("status code: %d", statusCode)
	}
	return err
}

func (c *jobs) GetContextMenu(ctx context.Context, name string, number int) (*v1.ContextMenu, error) {
	var statusCode int
	var ret v1.ContextMenu
	err := c.client.
		Get().
		AbsPath(fmt.Sprintf("/job/%s/%d/wfapi/contextMenu", name, number)).
		Do(ctx).
		StatusCode(&statusCode).
		Into(&ret)
	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", statusCode)
	}
	return &ret, err
}
