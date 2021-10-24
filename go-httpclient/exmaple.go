package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ken5scal/go-httpclient/gohttp"
)

var (
	httpClient       = gohttp.New()
	githubHttpClient = getGithubClient()
)

func getGithubClient() gohttp.HttpClient {
	client := gohttp.New()

	client.SetConnectionTimeout(2 * time.Second)
	client.SetResponseTimeout(1000 * time.Millisecond)
	client.DisableTimeout(true)

	commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "Bearer ABC-123")
	client.SetHeaders(commonHeaders)

	return client
}

func main() {
	getUrls()
}

func getUrls() {
	resp, err := githubHttpClient.Get("http://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.StatusCode)
	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func createUser(user User) {
	resp, err := githubHttpClient.Post("http://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.StatusCode)
	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
}
