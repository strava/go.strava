package strava

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var ClientId int
var ClientSecret string

const basePath = "https://www.strava.com/api/v3"
const timeFormat = "2006-01-02T15:04:05Z"

type Client struct {
	token      string
	httpClient *http.Client
}

// NewClient builds a normal client for making requests to the strava api.
// a http.Client can be passed in if http.DefaultClient can not be used.
func NewClient(token string, client ...*http.Client) *Client {
	c := &Client{token: token}
	if len(client) != 0 {
		c.httpClient = client[0]
	} else {
		c.httpClient = http.DefaultClient
	}
	return c
}

// NewStubResponseClient can be used for testing
// TODO, stub out with an actual response
func NewStubResponseClient(content string, statusCode ...int) *Client {
	c := NewClient("")
	t := &stubResponseTransport{content: content}

	if len(statusCode) != 0 {
		t.statusCode = statusCode[0]
	}

	c.httpClient = &http.Client{Transport: t}

	return c
}

type stubResponseTransport struct {
	http.Transport
	content    string
	statusCode int
}

func (t *stubResponseTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	resp := &http.Response{
		Status:     http.StatusText(t.statusCode),
		StatusCode: t.statusCode,
	}
	resp.Body = ioutil.NopCloser(strings.NewReader(t.content))

	return resp, nil
}

func (client *Client) run(method, path string, params map[string]interface{}) ([]byte, error) {
	var err error

	values := make(url.Values)
	for k, v := range params {
		values.Set(k, fmt.Sprintf("%v", v))
	}

	var req *http.Request
	if method == "POST" {
		req, err = http.NewRequest("POST", basePath+path, strings.NewReader(values.Encode()))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req, err = http.NewRequest(method, basePath+path+"?"+values.Encode(), nil)
		if err != nil {
			return nil, err
		}
	}

	return client.runRequest(req)
}

func (client *Client) runRequest(req *http.Request) ([]byte, error) {
	req.Header.Add("Authorization", "Bearer "+client.token)
	resp, err := client.httpClient.Do(req)

	// this was a poor request, maybe strava servers down?
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return checkResponseForErrors(resp)
}

func checkResponseForErrors(resp *http.Response) ([]byte, error) {
	// check status code, could be 500, or most likely the client_secret is incorrect
	if resp.StatusCode/100 == 5 {
		return nil, errors.New("server error")
	}

	if resp.StatusCode/100 == 4 {
		var response Error
		contents, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(contents, &response)

		return nil, response
	}

	if resp.StatusCode/100 == 3 {
		return nil, errors.New("redirect error")
	}

	return ioutil.ReadAll(resp.Body)
}
