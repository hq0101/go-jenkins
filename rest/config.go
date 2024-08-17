package rest

import (
	"fmt"
	"github.com/hq0101/go-jenkins/pkg/version"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	gruntime "runtime"
	"strings"
	"time"
)

type Config struct {
	Host        string
	UserName    string
	Password    string
	BearerToken string
	UserAgent   string
	Timeout     time.Duration
}

func RESTClientFor(config *Config) (*RESTClient, error) {
	if config.Host == "" {
		return nil, fmt.Errorf("host must be a URL or a host:port pair")
	}

	clientContent := ClientContentConfig{
		Username:  config.UserName,
		Password:  config.Password,
		Token:     config.BearerToken,
		UserAgent: config.UserAgent,
	}

	httpClient := HTTPClientFor(config)

	hostURL, err := url.Parse(config.Host)
	if err != nil || hostURL.Scheme == "" || hostURL.Host == "" {
		scheme := "http://"
		hostURL, err = url.Parse(scheme + config.Host)
		if err != nil {
			return nil, err
		}
		if hostURL.Path != "" && hostURL.Path != "/" {
			return nil, fmt.Errorf("host must be a URL or a host:port pair: %q", config.Host)
		}
		return nil, err
	}

	restClient, err := NewRESTClient(hostURL, clientContent, httpClient)
	if err != nil {
		return nil, err
	}

	return restClient, nil
}

func HTTPClientFor(config *Config) *http.Client {
	var httpClient *http.Client

	if config.Timeout > 0 {
		httpClient = &http.Client{
			Timeout: config.Timeout,
		}
	} else {
		httpClient = http.DefaultClient
	}

	return httpClient
}

// adjustCommit returns sufficient significant figures of the commit's git hash.
func adjustCommit(c string) string {
	if len(c) == 0 {
		return "unknown"
	}
	if len(c) > 7 {
		return c[:7]
	}
	return c
}

// adjustVersion strips "alpha", "beta", etc. from version in form
// major.minor.patch-[alpha|beta|etc].
func adjustVersion(v string) string {
	if len(v) == 0 {
		return "unknown"
	}
	seg := strings.SplitN(v, "-", 2)
	return seg[0]
}

// adjustCommand returns the last component of the
// OS-specific command path for use in User-Agent.
func adjustCommand(p string) string {
	// Unlikely, but better than returning "".
	if len(p) == 0 {
		return "unknown"
	}
	return filepath.Base(p)
}

// buildUserAgent builds a User-Agent string from given args.
func buildUserAgent(command, version, os, arch, commit string) string {
	return fmt.Sprintf(
		"%s/%s (%s/%s) go-jenkins/%s", command, version, os, arch, commit)
}

// DefaultUserAgent returns a User-Agent string built from static global vars.
func DefaultUserAgent() string {
	return buildUserAgent(
		adjustCommand(os.Args[0]),
		adjustVersion(version.Get().GitVersion),
		gruntime.GOOS,
		gruntime.GOARCH,
		adjustCommit(version.Get().GitCommit))
}
