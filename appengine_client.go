// +build appengine

package strava

import (
	"appengine"
	"appengine/urlfetch"
	"net/http"
)

var httpClient = func(r *http.Request) *http.Client {
	return urlfetch.Client(appengine.NewContext(r))
}
