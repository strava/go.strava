// oauth_example.go provides a simple example implementing Strava OAuth
// using the go.strava library.
//
// usage:
//   > go get github.com/strava/go.strava
//   > cd $GOPATH/github.com/strava/go.strava/examples
//   > go run oauth_example.go -id=youappsid -secret=yourappsecret
//
//   Visit http://localhost:8080 in your webbrowser
//
//   Application id and secret can be found at https://www.strava.com/settings/api
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/strava/go.strava"
)

var PORT = 8080 // port of local demo server

func main() {
	// setup the credentials for your app
	// These need to be set to reflect your application
	// and can be found at https://www.strava.com/settings/api
	flag.IntVar(&strava.ClientId, "id", 0, "Strava Client ID")
	flag.StringVar(&strava.ClientSecret, "secret", "", "Strava Client Secret")

	flag.Parse()

	if strava.ClientId == 0 || strava.ClientSecret == "" {
		fmt.Println("\nPlease provide your application's client_id and client_secret.")
		fmt.Println("For example: go run oauth_example.go -id=9 -secret=longrandomsecret\n")

		flag.PrintDefaults()
		os.Exit(1)
	}

	// define a callbackURL that includes full path of the callback,
	// this will be used in several places. Strava will redirect to this location
	// after the users grants access to your application
	strava.OAuthCallbackURL = fmt.Sprintf("http://localhost:%d/exchange_token", PORT)

	http.HandleFunc("/", indexHandler)

	path, err := strava.OAuthCallbackPath()
	if err != nil {
		// possibly strava.OAuthCallbackURL was not set, or was not a valid URL
		fmt.Println(err)
		os.Exit(1)
	}
	http.HandleFunc(path, strava.OAuthCallbackHandler(OAuthSuccess, OAuthFailure))

	// start the server
	fmt.Printf("Visit http://localhost:%d/ to view the demo\n", PORT)
	fmt.Printf("ctrl-c to exit")
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// you should make this a template in your real application
	fmt.Fprintf(w, `<a href="%s">`, strava.OAuthAuthorizationURL("state1", strava.Permissions.Public, true))
	fmt.Fprint(w, `<img src="http://strava.github.io/api/images/ConnectWithStrava.png" />`)
	fmt.Fprint(w, `</a>`)
}

func OAuthSuccess(auth *strava.AuthorizationResponse, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "SUCCESS:\nAt this point you can use this information to create a new user or link the account to one of your existing users\n")
	fmt.Fprintf(w, "State: %s\n\n", auth.State)
	fmt.Fprintf(w, "Access Token: %s\n\n", auth.AccessToken)

	fmt.Fprintf(w, "The Authenticated Athlete (you):\n")
	content, _ := json.MarshalIndent(auth.Athlete, "", " ")
	fmt.Fprint(w, string(content))
}

func OAuthFailure(err error, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Authorization Failure:\n")

	// some standard error checking
	if err == strava.OAuthAuthorizationDeniedErr {
		fmt.Fprint(w, "The user clicked the 'Do not Authorize' button on the previous page.\n")
		fmt.Fprint(w, "This is the main error your application should handle.")
	} else if err == strava.OAuthInvalidCredentialsErr {
		fmt.Fprint(w, "You provided an incorrect client_id or client_secret.\nDid you remember to set them at the begininng of this file?")
	} else if err == strava.OAuthInvalidCodeErr {
		fmt.Fprint(w, "The temporary token was not recognized, this shouldn't happen normally")
	} else if err == strava.OAuthServerErr {
		fmt.Fprint(w, "There was some sort of server error, try again to see if the problem continues")
	} else {
		fmt.Fprint(w, err)
	}
}
