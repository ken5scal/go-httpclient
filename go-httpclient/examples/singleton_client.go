package examples

import (
	"time"

	"github.com/ken5scal/go-httpclient/gohttp"
)

var httpClient = getHttpClient()

func getHttpClient() gohttp.Client {
	client := gohttp.NewBuilder().
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(5 * time.Second).
		Build()
	return client
}
