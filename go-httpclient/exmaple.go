package main

import (
	"fmt"

	"github.com/ken5scal/go-httpclient/gohttp"
)

func main() {
	client := gohttp.New()
	resp, err := client.Get("http://example.com", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
}
