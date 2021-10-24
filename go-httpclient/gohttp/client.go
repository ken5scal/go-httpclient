package gohttp

import (
	"net/http"
	"time"
)

type HttpClient interface {
	// Configure the client
	SetHeaders(headers http.Header)
	SetConnectionTimeout(timeout time.Duration)
	SetResponseTimeout(timeout time.Duration)
	SetMaxIdleConnections(maxIdleConnections int)
	DisableTimeout(disable bool)

	// Execute the CLient
	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	Put(url string, headers http.Header) (*http.Response, error)
	Patch(url string, headers http.Header) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
}

type httpClient struct {
	client             *http.Client
	Headers            http.Header
	maxIdleConnections int
	disableTimeout     bool
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
}

func New() HttpClient {
	// dialer := net.Dialer{Timeout: 1 * time.Second}
	// client := http.Client{
	// 	Transport: &http.Transport{
	// 		MaxIdleConnsPerHost:   5,
	// 		DialContext:           dialer.DialContext,
	// 		ResponseHeaderTimeout: 5 * time.Second,
	// 	},
	// }
	// return &httpClient{client: &client}
	return &httpClient{}
}

func (c *httpClient) SetHeaders(headers http.Header) {
	c.Headers = headers
}

func (c *httpClient) SetConnectionTimeout(timeout time.Duration) {
	c.connectionTimeout = timeout
}

func (c *httpClient) SetResponseTimeout(timeout time.Duration) {
	c.responseTimeout = timeout
}

func (c *httpClient) SetMaxIdleConnections(maxIdleConnections int) {
	c.maxIdleConnections = maxIdleConnections
}

func (c *httpClient) DisableTimeout(disable bool) {
	c.disableTimeout = disable
}

func (c *httpClient) Get(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

func (c *httpClient) Post(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}

func (c *httpClient) Put(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodPut, url, headers, nil)
}

func (c *httpClient) Patch(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodPatch, url, headers, nil)
}

func (c *httpClient) Delete(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}
