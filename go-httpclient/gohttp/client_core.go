package gohttp

import (
	"net/http"
)

func (c *httpClient) do(method, url string, headers http.Header, body interface{}) (*http.Response, error) {

	client := http.Client{}
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Accept", "application/xml")

	return client.Do(request)
}
