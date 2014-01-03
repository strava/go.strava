package strava

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var OAuthCallbackURL string

type Permission string

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

type AuthorizationResponse struct {
	AccessToken string          `json:"access_token"`
	State       string          `json:"State"`
	Athlete     AthleteDetailed `json:"athlete"`
}

// returns the path portion of the callbackURL. A callbackURL should
// be defined in the application and used here and in OAuthAuthorizationURL.
// For example:
//		http.HandleFunc(strava.OAuthCallbackPath(callbackURL), strava.OAuthCallbackHandler(successCallback, failureCallback))
func OAuthCallbackPath() (string, error) {
	if OAuthCallbackURL == "" {
		return "", errors.New("You must set strava.OAuthCallbackURL to be the full path of the oauth response")
	}
	url, err := url.Parse(OAuthCallbackURL)
	if err != nil {
		return "", err
	}
	return url.Path, nil
}

// When a user authorizes an application on strava.com, the request is redirect
// back to the strava.OAuthCallbackURL where the provided code must be exchanged for an
// access token. This method handles the exchange and calls success or failure after
// it completes
func OAuthCallbackHandler(
	success func(auth *AuthorizationResponse, w http.ResponseWriter, r *http.Request),
	failure func(err error, w http.ResponseWriter, r *http.Request)) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// user denied authorization
		if r.FormValue("error") == "access_denied" {
			failure(OAuthAuthorizationDeniedErr, w, r)
			return
		}

		resp, err := httpClient.PostForm(basePath+"/oauth/token",
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

// OAuthAuthorizationURL constructs the url you should send the user to so they can authorize
func OAuthAuthorizationURL(state string, scope Permission, force bool) string {
	path := fmt.Sprintf("%s/oauth/authorize?client_id=%d&response_type=code&redirect_uri=%s&scope=%v", basePath, ClientId, OAuthCallbackURL, scope)

	if state != "" {
		path += "&state=" + state
	}

	if force {
		path += "&approval_prompt=force"
	}

	return path
}
