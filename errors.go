package strava

import (
	"encoding/json"
)

type Error struct {
	Message string           `json:"message"`
	Errors  []*ErrorDetailed `json:"errors"`
}

type ErrorDetailed struct {
	Resource string `json:"resource"`
	Field    string `json:"field"`
	Code     string `json:"code"`
}

func (e Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

// returned during oauth if there was a user caused problem
// such as user did not grant access or the id/secret was invalid
type OAuthError struct {
	message string
}

func (e *OAuthError) Error() string {
	return e.message
}

var (
	OAuthAuthorizationDeniedErr = &OAuthError{"authorization denied by user"}
	OAuthInvalidCredentialsErr  = &OAuthError{"invalid client_id or client_secret"}
	OAuthInvalidCodeErr         = &OAuthError{"unrecognized code"}
	OAuthServerErr              = &OAuthError{"server error"}
)
