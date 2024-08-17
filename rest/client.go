package rest

import (
	"net/http"
	"net/url"
	"strings"
)

type Interface interface {
	Verb(verb string) *Request
	Post() *Request
	Put() *Request
	Get() *Request
	Delete() *Request
}

type ClientContentConfig struct {
	ContentType string
	Username    string
	Password    string
	Token       string
	UserAgent   string
}

// HasBasicAuth returns whether the configuration has basic authentication or not.
func (c *ClientContentConfig) HasBasicAuth() bool {
	return len(c.Username) != 0
}

// HasTokenAuth returns whether the configuration has token authentication or not.
func (c *ClientContentConfig) HasTokenAuth() bool {
	return len(c.Token) != 0
}

type RESTClient struct {
	base           *url.URL
	versionAPIPath string
	content        ClientContentConfig
	Client         *http.Client
}

func NewRESTClient(baseURL *url.URL, content ClientContentConfig, client *http.Client) (*RESTClient, error) {
	if len(content.ContentType) == 0 {
		content.ContentType = "application/json"
	}

	if len(content.UserAgent) == 0 {
		content.UserAgent = DefaultUserAgent()
	}

	base := *baseURL
	if !strings.HasSuffix(base.Path, "/") {
		base.Path += "/"
	}
	base.RawQuery = ""
	base.Fragment = ""

	return &RESTClient{
		base:    baseURL,
		content: content,
		Client:  client,
	}, nil
}

func (c *RESTClient) Verb(verb string) *Request {
	return NewRequest(c).Verb(verb)
}

func (c *RESTClient) Post() *Request {
	return c.Verb("POST")
}

func (c *RESTClient) Get() *Request {
	return c.Verb("GET")
}

func (c *RESTClient) Put() *Request {
	return c.Verb("PUT")
}

func (c *RESTClient) Delete() *Request {
	return c.Verb("DELETE")
}
