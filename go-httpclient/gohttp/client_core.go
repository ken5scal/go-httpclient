package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"strings"
)

func (c *httpClient) do(method, url string, headers http.Header, body interface{}) (*http.Response, error) {
	client := http.Client{}

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

	return client.Do(request)
}

func (c *httpClient) getRequestHeders(requestHeaders http.Header) http.Header {
	result := make(http.Header)

	// add common headers to the request
	for header, value := range c.Headers {
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
