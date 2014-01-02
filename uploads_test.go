package strava

import (
	"testing"
)

func TestUploadsGet(t *testing.T) {
	// if you need to change this you should also update tests below
	if c := structAttributeCount(&UploadDetailed{}); c != 5 {
		t.Fatalf("incorrect number of detailed attributes, %d != 5", c)
	}

	client := newCassetteClient(testToken, "upload_get")
	upload, err := NewUploadsService(client).Get(46440854).Do()

	expected := &UploadDetailed{}
	expected.Id = 46440854
	expected.ExternalId = "25FA60D8-15CF-472E-8C86-228B16320F41"
	expected.Error = ""
	expected.Status = "The created activity has been deleted."
	expected.ActivityId = 0

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	for _, prob := range structCompare(t, upload, expected) {
		t.Error(prob)
	}

	// from here on out just check the request parameters
	s := NewUploadsService(newStoreRequestClient())

	// path
	s.Get(123).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/uploads/123" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestUploadsBadJSON(t *testing.T) {
	var err error
	s := NewUploadsService(newStubResponseClient("bad json"))

	_, err = s.Get(123).Do()
	if err == nil {
		t.Error("should return a bad json error")
	}
}
