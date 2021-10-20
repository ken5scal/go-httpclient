package gohttp

import "net/http"

func (c *httpClient) do() {

	client := http.Client{}
	request, err := http.NewRequest(httpMethod, url, nil)
	request.Header.Set("Accept", "application/xml")

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

}
