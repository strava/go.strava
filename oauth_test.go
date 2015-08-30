package strava

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOAuthAuthenticatorCallbackHandler(t *testing.T) {
	// http client failure
	auth := OAuthAuthenticator{
		RequestClientGenerator: func(r *http.Request) *http.Client { return &http.Client{Transport: &storeRequestTransport{}} },
	}

	f := auth.HandlerFunc(func(auth *AuthorizationResponse, w http.ResponseWriter, r *http.Request) {
		t.Error("should handle request failure")
	}, func(err error, w http.ResponseWriter, r *http.Request) {
		if err == nil {
			t.Error("error should not be nil")
		}
	})

	req, _ := http.NewRequest("GET", "", nil)
	f(httptest.NewRecorder(), req)

	// http client doesn't exist
	auth = OAuthAuthenticator{
		RequestClientGenerator: func(r *http.Request) *http.Client { return nil },
	}

	f = auth.HandlerFunc(func(auth *AuthorizationResponse, w http.ResponseWriter, r *http.Request) {
		t.Error("should handle request failure")
	}, func(err error, w http.ResponseWriter, r *http.Request) {
		if err == nil {
			t.Error("error should not be nil")
		}
	})

	req, _ = http.NewRequest("GET", "", nil)
	f(httptest.NewRecorder(), req)

	// access denied
	auth = OAuthAuthenticator{}
	f = auth.HandlerFunc(func(auth *AuthorizationResponse, w http.ResponseWriter, r *http.Request) {
		t.Error("access denied should be failure")
	}, func(err error, w http.ResponseWriter, r *http.Request) {
		if err != OAuthAuthorizationDeniedErr {
			t.Errorf("returned incorrect error, got %v", err)
		}
	})

	req, _ = http.NewRequest("GET", "?error=access_denied", nil)
	f(httptest.NewRecorder(), req)

	// strava error
	auth = OAuthAuthenticator{
		RequestClientGenerator: func(r *http.Request) *http.Client {
			return NewStubResponseClient("{}", http.StatusInternalServerError).httpClient
		},
	}

	f = auth.HandlerFunc(func(auth *AuthorizationResponse, w http.ResponseWriter, r *http.Request) {
		t.Error("should return error when strava returned error")
	}, func(err error, w http.ResponseWriter, r *http.Request) {
		if err != OAuthServerErr {
			t.Errorf("returned incorrect error, got %v", err)
		}
	})

	req, _ = http.NewRequest("GET", "?code=75e251e3ff8fff", nil)
	f(httptest.NewRecorder(), req)

	// strava error
	auth = OAuthAuthenticator{
		RequestClientGenerator: func(r *http.Request) *http.Client {
			return NewStubResponseClient(`{"message":"bad","errors":[]}`, http.StatusBadRequest).httpClient
		},
	}

	f = auth.HandlerFunc(func(auth *AuthorizationResponse, w http.ResponseWriter, r *http.Request) {
		t.Error("should return error when strava returned error")
	}, func(err error, w http.ResponseWriter, r *http.Request) {
		if err != OAuthServerErr {
			t.Errorf("returned incorrect error, got %v", err)
		}
	})

	req, _ = http.NewRequest("GET", "?code=75e251e3ff8fff", nil)
	f(httptest.NewRecorder(), req)

	// strava invalid credentials error
	auth = OAuthAuthenticator{
		RequestClientGenerator: func(r *http.Request) *http.Client {
			return NewStubResponseClient(`{"message":"bad","errors":[{"resource":"Application","field":"","code":""}]}`, http.StatusBadRequest).httpClient
		},
	}

	f = auth.HandlerFunc(func(auth *AuthorizationResponse, w http.ResponseWriter, r *http.Request) {
		t.Error("should return error when strava returned error")
	}, func(err error, w http.ResponseWriter, r *http.Request) {
		if err != OAuthInvalidCredentialsErr {
			t.Errorf("returned incorrect error, got %v", err)
		}
	})

	req, _ = http.NewRequest("GET", "?code=75e251e3ff8fff", nil)
	f(httptest.NewRecorder(), req)

	// strava invalid code error
	auth = OAuthAuthenticator{
		RequestClientGenerator: func(r *http.Request) *http.Client {
			return NewStubResponseClient(`{"message":"bad","errors":[{"resource":"RequestToken","field":"","code":""}]}`, http.StatusBadRequest).httpClient
		},
	}

	f = auth.HandlerFunc(func(auth *AuthorizationResponse, w http.ResponseWriter, r *http.Request) {
		t.Error("should return error when strava returned error")
	}, func(err error, w http.ResponseWriter, r *http.Request) {
		if err != OAuthInvalidCodeErr {
			t.Errorf("returned incorrect error, got %v", err)
		}
	})

	req, _ = http.NewRequest("GET", "", nil)
	f(httptest.NewRecorder(), req)

	// other strava error
	auth = OAuthAuthenticator{
		RequestClientGenerator: func(r *http.Request) *http.Client {
			return NewStubResponseClient(`{"message":"bad","errors":[{"resource":"Other","field":"","code":""}]}`, http.StatusBadRequest).httpClient
		},
	}

	f = auth.HandlerFunc(func(auth *AuthorizationResponse, w http.ResponseWriter, r *http.Request) {
		t.Error("should return error when strava returned error")
	}, func(err error, w http.ResponseWriter, r *http.Request) {
		if _, ok := err.(*Error); !ok {
			t.Errorf("returned incorrect error, got %v", err)
		}
	})

	req, _ = http.NewRequest("GET", "?code=75e251e3ff8fff", nil)
	f(httptest.NewRecorder(), req)

	// bad json
	auth = OAuthAuthenticator{
		RequestClientGenerator: func(r *http.Request) *http.Client {
			return NewStubResponseClient(`bad json`, http.StatusOK).httpClient
		},
	}

	f = auth.HandlerFunc(func(auth *AuthorizationResponse, w http.ResponseWriter, r *http.Request) {
		t.Error("should return error when strava returned error")
	}, func(err error, w http.ResponseWriter, r *http.Request) {
		if err == nil {
			t.Error("error should not be nil")
		}
	})

	req, _ = http.NewRequest("GET", "", nil)
	f(httptest.NewRecorder(), req)

	// success!
	auth = OAuthAuthenticator{
		RequestClientGenerator: func(r *http.Request) *http.Client {
			return NewStubResponseClient(`{}`, http.StatusOK).httpClient
		},
	}

	f = auth.HandlerFunc(func(auth *AuthorizationResponse, w http.ResponseWriter, r *http.Request) {
	}, func(err error, w http.ResponseWriter, r *http.Request) {
		t.Error("should be success")
	})

	req, _ = http.NewRequest("GET", "?code=75e251e3ff8fff", nil)
	f(httptest.NewRecorder(), req)
}

func TestOAuthAuthenticatorAuthorize(t *testing.T) {
	auth := OAuthAuthenticator{}

	_, err := auth.Authorize("", nil)
	if err != OAuthInvalidCodeErr {
		t.Errorf("returned incorrect error, got %v", err)
	}
}

func TestOAuthAuthenticatorCallbackPath(t *testing.T) {
	auth := OAuthAuthenticator{}

	_, err := auth.CallbackPath()
	if err == nil {
		t.Error("should return error since callback url is not set")
	}

	auth = OAuthAuthenticator{
		CallbackURL: "http://www.strava.c%om/",
	}
	_, err = auth.CallbackPath()
	if err == nil {
		t.Error("should return error since not a callback url")
	}

	auth = OAuthAuthenticator{
		CallbackURL: "http://abc.com/strava/oauth",
	}
	s, _ := auth.CallbackPath()
	if s != "/strava/oauth" {
		t.Error("incorrect path")
	}
}

func TestOAuthAuthenticatorAuthorizationURL(t *testing.T) {
	auth := OAuthAuthenticator{
		CallbackURL: "http://abc.com/strava/oauth",
	}

	url := auth.AuthorizationURL("state", Permissions.Public, false)
	if url != basePath+"/oauth/authorize?client_id=0&response_type=code&redirect_uri=http://abc.com/strava/oauth&scope=public&state=state" {
		t.Errorf("incorrect oauth url, got %v", url)
	}

	url = auth.AuthorizationURL("state", Permissions.Public, true)
	if url != basePath+"/oauth/authorize?client_id=0&response_type=code&redirect_uri=http://abc.com/strava/oauth&scope=public&state=state&approval_prompt=force" {
		t.Errorf("incorrect oauth url, got %v", url)
	}

	url = auth.AuthorizationURL("state", Permissions.ViewPrivate, false)
	if url != basePath+"/oauth/authorize?client_id=0&response_type=code&redirect_uri=http://abc.com/strava/oauth&scope=view_private&state=state" {
		t.Errorf("incorrect oauth url, got %v", url)
	}

	url = auth.AuthorizationURL("", Permissions.Public, false)
	if url != basePath+"/oauth/authorize?client_id=0&response_type=code&redirect_uri=http://abc.com/strava/oauth&scope=public" {
		t.Errorf("incorrect oauth url, got %v", url)
	}
}

func TestOAuthErrorError(t *testing.T) {
	err := OAuthAuthorizationDeniedErr
	if err.Error() != err.message {
		t.Error("should simply print message")
	}
}

func TestErrorString(t *testing.T) {
	err := Error{
		Message: "bad bad bad",
		Errors:  []*ErrorDetailed{&ErrorDetailed{"auth", "code", "missing"}},
	}
	if err.Error() != `{"message":"bad bad bad","errors":[{"resource":"auth","field":"code","code":"missing"}]}` {
		t.Errorf("should simply print message, got %v", err.Error())
	}
}

func TestOAuthDeauthorize(t *testing.T) {
	client := newCassetteClient("f774a6dd01b16de401ba1531e770c951dcd7523f", "oauth_deauthorize")
	err := NewOAuthService(client).Deauthorize().Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	// from here on out just check the request parameters
	s := NewOAuthService(newStoreRequestClient())

	// path
	s.Deauthorize().Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/oauth/deauthorize" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.Method != "POST" {
		t.Errorf("request method incorrect, got %v", transport.request.Method)
	}

	if transport.request.URL.RawQuery != "" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}
