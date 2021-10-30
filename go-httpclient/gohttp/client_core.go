package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	defaultMaxIdleConnections = 5
	defaultResponseTimeout    = 5 * time.Second
	defaultConnectionTimeout  = 1 * time.Second
)

func (c *httpClient) do(method, url string, headers http.Header, body interface{}) (*http.Response, error) {
	fullHeaders := c.getRequestHeders(headers)

	requestBody, err := c.getRequestBody(fullHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, errors.New("Error creating requestBody: " + err.Error())
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("Error creating request: " + err.Error())
	}

	request.Header = fullHeaders
	// return c.client.Do(request)

	client := c.getHttpClient()
	return client.Do(request)
}

func (c *httpClient) getHttpClient() *http.Client {
	if c.client != nil {
		return c.client
	}

	dialer := net.Dialer{Timeout: c.getConnectionTimeout()}
	c.client = &http.Client{
		Timeout: c.getResponseTimeout() + c.getConnectionTimeout(),
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   c.getMaXIdleCOnnections(),
			DialContext:           dialer.DialContext,
			ResponseHeaderTimeout: c.getResponseTimeout(),
		},
	}

	return c.client
}

func (c *httpClient) getMaXIdleCOnnections() int {
	if c.builder.maxIdleConnections > 0 {
		return c.builder.maxIdleConnections
	}
	return defaultMaxIdleConnections
}

func (c *httpClient) getResponseTimeout() time.Duration {
	if c.builder.responseTimeout > 0 {
		return c.builder.responseTimeout
	}

	if c.builder.disableTimeout {
		return 0
	}
	return defaultResponseTimeout
}

func (c *httpClient) getConnectionTimeout() time.Duration {
	if c.builder.connectionTimeout > 0 {
		return c.builder.connectionTimeout
	}

	if c.builder.disableTimeout {
		return 0
	}

	return defaultConnectionTimeout
}

func (c *httpClient) getRequestHeders(requestHeaders http.Header) http.Header {
	result := make(http.Header)

	// add common headers to the request
	for header, value := range c.builder.headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	// add custom headers to the requests:
	for header, value := range requestHeaders {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}
	return result
}

func (c *httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(body)
	case "application/xml":
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}
}
