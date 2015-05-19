package strava

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"testing"
)

var testToken string

func init() {
	testToken = os.Getenv("access_token")
	if testToken == "" {
		// please don't deauthorize this token. It offers no special access.
		testToken = "21b4fe41a815dd7de4f0cae7f04bbbf9aa0f9507"
	}
}

var cassetteDirectory = "cassettes"

func newCassetteClient(token, cassette string) *Client {
	c := NewClient(token)
	c.httpClient = &http.Client{
		Transport: &cassetteTransport{
			token:     token,
			directory: cassetteDirectory,
			cassette:  cassette},
	}

	return c
}

type cassetteTransport struct {
	http.Transport
	token     string
	directory string
	cassette  string
}

// if directory/cassette.json exists, return it
// else run, save and return
func (t *cassetteTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	filename := t.directory + "/" + t.cassette
	if _, e := os.Stat(filename + ".resp"); e == nil {
		content, _ := ioutil.ReadFile(filename + ".resp")
		var resp http.Response
		json.Unmarshal(content, &resp)

		// and the body
		resp.Body, _ = os.Open(filename + ".body")

		// check for the error too
		var err error = nil
		if _, e = os.Stat(filename + ".err"); err == nil {
			content, _ := ioutil.ReadFile(filename + ".err")

			json.Unmarshal(content, &err)
		}

		return &resp, err
	}

	// need to fetch the data from the web, so use the default transport
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return resp, err
	}

	// save the cassette to disk
	j, _ := json.Marshal(resp)
	ioutil.WriteFile(filename+".resp", j, 0644)

	body, _ := ioutil.ReadAll(resp.Body)
	ioutil.WriteFile(filename+".body", body, 0644)
	resp.Body, _ = os.Open(filename + ".body")

	if err != nil {
		j, _ = json.Marshal(err)
		ioutil.WriteFile(filename+".err", j, 0644)
	}

	return resp, err
}

/*********************************************************/

func newStoreRequestClient() *Client {
	c := NewClient("")
	c.httpClient = &http.Client{Transport: &storeRequestTransport{}}

	return c
}

type storeRequestTransport struct {
	http.Transport
	request *http.Request
}

func (t *storeRequestTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.request = req

	return nil, errors.New("for testing, no request made")
}

/*********************************************************/

func TestClient(t *testing.T) {
	c := NewClient("token")
	if c.token != "token" {
		t.Errorf("token not set correctly")
	}

	httpClient := &http.Client{}
	c = NewClient("token", httpClient)
	if c.httpClient != httpClient {
		t.Errorf("http client not set correctly")
	}
}

func TestRun(t *testing.T) {
	var err error
	c := newStoreRequestClient()

	_, err = c.run("GET", "pa%@th", nil)
	if err == nil {
		t.Error("should return error due to invalid path")
	}

	_, err = c.run("POST", "pa%@th", nil)
	if err == nil {
		t.Error("should return error due to invalid path")
	}
}

func TestCheckResponseForErrors(t *testing.T) {
	var err error
	var resp http.Response

	resp.StatusCode = 300
	_, err = checkResponseForErrors(&resp)
	if err == nil {
		t.Error("should have returned error")
	}

	resp.StatusCode = 503
	_, err = checkResponseForErrors(&resp)
	if err == nil {
		t.Error("should have returned error")
	}

	resp.StatusCode = 404
	resp.Body = ioutil.NopCloser(strings.NewReader(`{"message":"Record Not Found","errors":[{"resource":"Activity","field":"id","code":"invalid"}]}`))
	_, err = checkResponseForErrors(&resp)
	if err == nil {
		t.Error("should have returned error")
	}

	if se, ok := err.(Error); ok {
		if len(se.Errors) == 0 {
			t.Error("Detailed errors not parsed")
		}
	} else {
		t.Error("Should have returned strava error")
	}
}

func TestCheckResponseForErrorsWithErrorHandler(t *testing.T) {
	var err error
	var resp http.Response

	erorrHandler := func(response *http.Response) error {
		contents, _ := ioutil.ReadAll(resp.Body)
		var data interface{}
		json.Unmarshal(contents, &data)
		errorData := data.(map[string]interface{})["error"].(string)
		return errors.New(errorData)
	}

	resp.StatusCode = 400
	resp.Body = ioutil.NopCloser(strings.NewReader(`{"error":"bad request found"}`))
	_, err = checkResponseForErrorsWithErrorHandler(&resp, erorrHandler)

	if err == nil {
		t.Error("should have returned error")
	}
	if err.Error() != "bad request found" {
		t.Error("should have returned expected error message")
	}
}

func TestTransport(t *testing.T) {
	c := newStoreRequestClient()
	c.token = "token"
	NewClubsService(c).Get(122).Do()

	transport := c.httpClient.Transport.(*storeRequestTransport)
	if h := transport.request.Header.Get("Authorization"); h != "Bearer token" {
		t.Errorf("request header incorrect, got %v", h)
	}
}
