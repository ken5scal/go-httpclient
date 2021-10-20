package gohttp

type httpClient struct {
}

func New() HttpClient {
	client := &httpClient{}
	return client
}

type HttpClient interface {
	Get()
	Post()
	Delete()
	Put()
	Patch()
}

func (c *httpClient) Get() {
	c.do()
}

func (c *httpClient) Post() {
}

func (c *httpClient) Put() {
}

func (c *httpClient) Patch() {
}

func (c *httpClient) Delete() {
}
