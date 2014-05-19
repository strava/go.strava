// +build !appengine

package strava

import (
	"net/http"
)

var httpClient = func(r *http.Request) *http.Client {
	return http.DefaultClient
}
