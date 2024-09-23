package v1

import (
	"context"
	"fmt"
	v1 "github.com/hq0101/go-jenkins/api/pipeline-model-definition/v1"
	"github.com/hq0101/go-jenkins/rest"
	"net/http"
)

type PipelineModeGetter interface {
	PipelineMode() PipelineModeInterface
}

type PipelineModeInterface interface {
	Validate(ctx context.Context, jenkinsFile string) (string, error)
	ValidateJenkinsFile(ctx context.Context, jenkinsFile string) (*v1.Result, error)
	ValidateJson(ctx context.Context, jenkinsFileJson string) (*v1.Result, error)
	ToJson(ctx context.Context, jenkinsFile string) (*v1.Result, error)
	ToJenkinsFile(ctx context.Context, jenkinsFileJson string) (*v1.Result, error)
	StepsToJson(ctx context.Context, jenkinsFile string) (*v1.Result, error)
	StepsToJenkinsFile(ctx context.Context, jenkinsFileJson string) (*v1.Result, error)
}

type pipelineMode struct {
	client rest.Interface
}

func newPipelineMode(c *PipelineModeV1Client) *pipelineMode {
	return &pipelineMode{
		client: c.RESTClient(),
	}
}

func (c *pipelineMode) Validate(ctx context.Context, jenkinsFile string) (string, error) {
	var ret string
	var statusCode int
	err := c.client.
		Post().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		AbsPath("/pipeline-model-converter/validate").
		Param("jenkinsfile", jenkinsFile).
		Do(ctx).
		StatusCode(&statusCode).Into(&ret)
	if statusCode != http.StatusOK {
		return "", fmt.Errorf("get crumb failed, status code: %d", statusCode)
	}

	return ret, err
}

// ValidateJenkinsFile Validation of Jenkinsfile
func (c *pipelineMode) ValidateJenkinsFile(ctx context.Context, jenkinsFile string) (*v1.Result, error) {

	ret := &v1.Result{}
	var statusCode int
	err := c.client.
		Post().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		AbsPath("/pipeline-model-converter/validateJenkinsfile").
		Param("jenkinsfile", jenkinsFile).
		Do(ctx).
		StatusCode(&statusCode).Into(ret)
	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("get crumb failed, status code: %d", statusCode)
	}

	return ret, err
}

// ValidateJson Validation of JSON representation
func (c *pipelineMode) ValidateJson(ctx context.Context, jenkinsFileJson string) (*v1.Result, error) {
	ret := &v1.Result{}
	var statusCode int
	err := c.client.
		Post().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		AbsPath("/pipeline-model-converter/validateJson").
		Param("json", jenkinsFileJson).
		Do(ctx).
		StatusCode(&statusCode).Into(ret)
	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("get crumb failed, status code: %d", statusCode)
	}

	return ret, err
}

func (c *pipelineMode) ToJson(ctx context.Context, jenkinsFile string) (*v1.Result, error) {
	ret := &v1.Result{}
	var statusCode int
	err := c.client.
		Post().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		AbsPath("/pipeline-model-converter/toJson").
		Param("jenkinsfile", jenkinsFile).
		Do(ctx).
		StatusCode(&statusCode).Into(ret)
	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("get crumb failed, status code: %d", statusCode)
	}

	return ret, err
}

func (c *pipelineMode) ToJenkinsFile(ctx context.Context, jenkinsFileJson string) (*v1.Result, error) {
	ret := &v1.Result{}
	var statusCode int
	err := c.client.
		Post().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		AbsPath("/pipeline-model-converter/toJenkinsfile").
		Param("json", jenkinsFileJson).
		Do(ctx).
		StatusCode(&statusCode).Into(ret)
	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("get crumb failed, status code: %d", statusCode)
	}

	return ret, err
}

func (c *pipelineMode) StepsToJson(ctx context.Context, jenkinsFile string) (*v1.Result, error) {
	ret := &v1.Result{}
	var statusCode int
	err := c.client.
		Post().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		AbsPath("/pipeline-model-converter/stepsToJson").
		Param("jenkinsfile", jenkinsFile).
		Do(ctx).
		StatusCode(&statusCode).Into(ret)
	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("get crumb failed, status code: %d", statusCode)
	}

	return ret, err
}

func (c *pipelineMode) StepsToJenkinsFile(ctx context.Context, jenkinsFileJson string) (*v1.Result, error) {
	ret := &v1.Result{}
	var statusCode int
	err := c.client.
		Post().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		AbsPath("/pipeline-model-converter/stepsToJenkinsfile").
		Param("json", jenkinsFileJson).
		Do(ctx).
		StatusCode(&statusCode).Into(ret)
	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("get crumb failed, status code: %d", statusCode)
	}

	return ret, err
}
