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

// defined here so it can be overridden for testing
var HttpClient = http.DefaultClient

type Client struct {
	token      string
	httpClient *http.Client
}

func NewClient(token string) *Client {
	return &Client{
		token:      token,
		httpClient: &http.Client{Transport: &transport{token: token}},
	}
}

func (client *Client) run(method, path string, params map[string]interface{}) ([]byte, error) {
	var err error

	values := make(url.Values)
	for k, v := range params {
		values.Set(k, fmt.Sprintf("%v", v))
	}

	var req *http.Request
	if method == "GET" {
		req, err = http.NewRequest(method, basePath+path+"?"+values.Encode(), nil)
		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequest(method, basePath+path, strings.NewReader(values.Encode()))
		if err != nil {
			return nil, err
		}
	}

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

type transport struct {
	http.Transport
	token string
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+t.token)
	return HttpClient.Do(req)
}
