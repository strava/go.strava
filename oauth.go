package strava

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// An OAuthAuthenticator holds state about how OAuth requests should be authenticated.
type OAuthAuthenticator struct {
	CallbackURL string // used to help generate the AuthorizationURL

	// The RequestClientGenerator builds the http.Client that will be used
	// to complete the token exchange. If nil, http.DefaultClient will be used.
	// On Google's App Engine http.DefaultClient is not available and this generator
	// can be used to create a client using the incoming request, for Example:
	//    func(r *http.Request) { return urlfetch.Client(appengine.NewContext(r)) }
	RequestClientGenerator func(r *http.Request) *http.Client
}

// Permission represents the access of an access_token.
// The permission type is requested during the token exchange.
type Permission string

// Permissions defines the available permissions
var Permissions = struct {
	Public           Permission
	ViewPrivate      Permission
	Write            Permission
	WriteViewPrivate Permission
}{
	"public",
	"view_private",
	"write",
	"write,view_private",
}

// AuthorizationResponse is returned as a result of the token exchange
type AuthorizationResponse struct {
	AccessToken string          `json:"access_token"`
	State       string          `json:"State"`
	Athlete     AthleteDetailed `json:"athlete"`
}

// CallbackPath returns the path portion of the CallbackURL.
// Useful when setting a http path handler, for example:
//		http.HandleFunc(stravaOAuth.CallbackURL(), stravaOAuth.HandlerFunc(successCallback, failureCallback))
func (auth OAuthAuthenticator) CallbackPath() (string, error) {
	if auth.CallbackURL == "" {
		return "", errors.New("callbackURL is empty")
	}
	url, err := url.Parse(auth.CallbackURL)
	if err != nil {
		return "", err
	}
	return url.Path, nil
}

// HandlerFunc builds a http.HandlerFunc that will complete the token exchange
// after a user authorizes an application on strava.com.
// This method handles the exchange and calls success or failure after it completes.
func (auth OAuthAuthenticator) HandlerFunc(
	success func(auth *AuthorizationResponse, w http.ResponseWriter, r *http.Request),
	failure func(err error, w http.ResponseWriter, r *http.Request)) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// user denied authorization
		if r.FormValue("error") == "access_denied" {
			failure(OAuthAuthorizationDeniedErr, w, r)
			return
		}

		// use the client generator if provided.
		client := http.DefaultClient
		if auth.RequestClientGenerator != nil {
			client = auth.RequestClientGenerator(r)
		}

		if client == nil {
			failure(OAuthHTTPClientUnavailable, w, r)
			return
		}

		resp, err := client.PostForm(basePath+"/oauth/token",
			url.Values{"client_id": {fmt.Sprintf("%d", ClientId)}, "client_secret": {ClientSecret}, "code": {r.FormValue("code")}})

		// this was a poor request, maybe strava servers down?
		if err != nil {
			failure(err, w, r)
			return
		}
		defer resp.Body.Close()

		// check status code, could be 500, or most likely the client_secret is incorrect
		if resp.StatusCode/100 == 5 {
			failure(OAuthServerErr, w, r)
			return
		}

		if resp.StatusCode/100 != 2 {
			var response Error
			contents, _ := ioutil.ReadAll(resp.Body)
			json.Unmarshal(contents, &response)

			if len(response.Errors) == 0 {
				failure(OAuthServerErr, w, r)
				return
			}

			if response.Errors[0].Resource == "Application" {
				failure(OAuthInvalidCredentialsErr, w, r)
				return
			}

			if response.Errors[0].Resource == "RequestToken" {
				failure(OAuthInvalidCodeErr, w, r)
				return
			}

			failure(&response, w, r)
			return
		}

		var response AuthorizationResponse
		contents, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(contents, &response)

		if err != nil {
			failure(err, w, r)
			return
		}

		response.Athlete.postProcessDetailed()
		response.State = r.FormValue("state")

		success(&response, w, r)
		return
	}
}

// AuthorizationURL constructs the url a user should use to authorize this specific application.
func (auth OAuthAuthenticator) AuthorizationURL(state string, scope Permission, force bool) string {
	path := fmt.Sprintf("%s/oauth/authorize?client_id=%d&response_type=code&redirect_uri=%s&scope=%v", basePath, ClientId, auth.CallbackURL, scope)

	if state != "" {
		path += "&state=" + state
	}

	if force {
		path += "&approval_prompt=force"
	}

	return path
}
