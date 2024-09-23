package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

type Request struct {
	c          *RESTClient
	verb       string
	pathPrefix string
	timeout    time.Duration
	params     url.Values
	headers    http.Header
	err        error
	body       io.Reader
}

func NewRequest(c *RESTClient) *Request {
	var timeout time.Duration
	if c.Client != nil {
		timeout = c.Client.Timeout
	}

	var pathPrefix string
	if c.base != nil {
		pathPrefix = path.Join("/", c.base.Path)
	} else {
		pathPrefix = path.Join("/", c.versionAPIPath)
	}

	r := &Request{
		c:          c,
		timeout:    timeout,
		pathPrefix: pathPrefix,
		params:     url.Values{},
		headers:    http.Header{},
	}

	if len(c.content.ContentType) > 0 {
		r.SetHeader("Content-Type", c.content.ContentType)
	}

	if len(c.content.UserAgent) > 0 {
		r.SetHeader("User-Agent", c.content.UserAgent)
	}

	return r
}

func (r *Request) Verb(verb string) *Request {
	r.verb = verb
	return r
}

func (r *Request) Param(paramName, s string) *Request {
	if r.err != nil {
		return r
	}
	return r.setParam(paramName, s)
}

func (r *Request) setParam(paramName, value string) *Request {
	if r.params == nil {
		r.params = make(url.Values)
	}
	r.params[paramName] = append(r.params[paramName], value)
	return r
}

func (r *Request) Error() error {
	return r.err
}

func (r *Request) SetHeader(key string, values ...string) *Request {
	if r.headers == nil {
		r.headers = http.Header{}
	}
	r.headers.Del(key)
	for _, value := range values {
		r.headers.Add(key, value)
	}
	return r
}

// AbsPath overwrites an existing path with the segments provided. Trailing slashes are preserved
// when a single segment is passed.
func (r *Request) AbsPath(segments ...string) *Request {
	if r.err != nil {
		return r
	}

	r.pathPrefix = path.Join(r.c.base.Path, path.Join(segments...))

	if len(segments) == 1 && (len(r.c.base.Path) > 1 || len(segments[0]) > 1) && strings.HasSuffix(segments[0], "/") {
		r.pathPrefix += "/"
	}

	return r
}

// Timeout makes the request use the given duration as an overall timeout for the
// request. Additionally, if set passes the value as "timeout" parameter in URL.
func (r *Request) Timeout(d time.Duration) *Request {
	if r.err != nil {
		return r
	}
	r.timeout = d
	return r
}

// RequestURI overwrites existing path and parameters with the value of the provided server relative
// URI.
func (r *Request) RequestURI(uri string) *Request {
	if r.err != nil {
		return r
	}

	locator, err := url.Parse(uri)
	if err != nil {
		r.err = err
		return r
	}
	r.pathPrefix = locator.Path
	if len(locator.Query()) > 0 {
		if r.params == nil {
			r.params = make(url.Values)
		}
		for k, v := range locator.Query() {
			r.params[k] = v
		}
	}

	return r
}

func (r *Request) Body(obj interface{}) *Request {
	if r.err != nil {
		return r
	}

	switch t := obj.(type) {
	case io.Reader:
		r.body = t
	case []byte:
		r.body = bytes.NewReader(t)
	case string:
		r.body = bytes.NewReader([]byte(t))
	case url.Values:
		r.body = bytes.NewBufferString(t.Encode())
	default:
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		if err := enc.Encode(obj); err != nil {
			r.err = err
			fmt.Printf("Failed to encode object to JSON: %v", err)
			return r
		}

		r.body = &buf
	}

	return r
}

func (r *Request) URL() *url.URL {
	p := r.pathPrefix

	finalURL := &url.URL{}
	if r.c.base != nil {
		*finalURL = *r.c.base
	}

	finalURL.Path = p

	query := url.Values{}
	for key, values := range r.params {
		for _, value := range values {
			query.Add(key, value)
		}
	}

	if r.timeout != 0 {
		query.Set("timeout", r.timeout.String())
	}
	finalURL.RawQuery = query.Encode()
	return finalURL
}

type Result struct {
	body        []byte
	err         error
	statusCode  int
	contentType string
}

func (r Result) StatusCode(statusCode *int) Result {
	*statusCode = r.statusCode
	return r
}

func (r Result) ContentType(contentType *string) Result {
	*contentType = r.contentType
	return r
}

func (r Result) Raw() ([]byte, error) {
	return r.body, r.err
}

func (r Result) Error() error {
	return r.err
}

func (r *Request) Do(ctx context.Context) Result {
	var result Result
	if r.err != nil {
		result.err = r.err
		return result
	}

	client := r.c.Client
	if client == nil {
		client = http.DefaultClient
	}

	if r.timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, r.timeout)
		defer cancel()
	}

	req, err := http.NewRequest(r.verb, r.URL().String(), r.body)
	if err != nil {
		result.err = err
		return result
	}

	req.Header = r.headers

	switch {
	case r.c.content.HasTokenAuth():
		r.SetHeader("Authorization", fmt.Sprintf("Bearer %s", r.c.content.Token))
	case r.c.content.HasBasicAuth():
		req.SetBasicAuth(r.c.content.Username, r.c.content.Password)
	}

	resp, err := client.Do(req)
	if err != nil {
		result.err = err
		return result
	}
	if resp == nil {
		return result
	}

	defer resp.Body.Close()

	result.body, result.err = io.ReadAll(resp.Body)
	result.statusCode = resp.StatusCode
	return result
}

func (r Result) Into(v interface{}) error {
	if r.err != nil {
		return r.Error()
	}

	if obj, ok := v.(*string); ok {
		*obj = string(r.body)
		return nil
	}

	if err := json.Unmarshal(r.body, v); err != nil {
		return err
	}

	return nil
}
