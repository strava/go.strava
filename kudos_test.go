package strava

import (
	"testing"
)

func TestActivitiesKudosList(t *testing.T) {
	client := newCassetteClient(testToken, "activity_kudos_list")
	athletes, err := NewActivityKudosService(client, 103221154).List().Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if len(athletes) == 0 {
		t.Fatal("kudoers not parsed")
	}

	if athletes[0].CreatedAt.IsZero() || athletes[0].UpdatedAt.IsZero() {
		t.Error("athlete dates not parsed")
	}

	// from here on out just check the request parameters
	s := NewActivityKudosService(newStoreRequestClient(), 123)

	// path
	s.List().Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/activities/123/kudos" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}

	// parameters
	s.List().Page(2).PerPage(3).Do()

	transport = s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/activities/123/kudos" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "page=2&per_page=3" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestActivityKudosCreate(t *testing.T) {
	client := newCassetteClient(testToken, "activity_kudos_post")
	err := NewActivityKudosService(client, 118229063).Create().Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	// from here on out just check the request parameters
	s := NewActivityKudosService(newStoreRequestClient(), 123)

	// path
	s.Create().Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/activities/123/kudos" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.Method != "POST" {
		t.Errorf("request method incorrect, got %v", transport.request.Method)
	}

	if transport.request.URL.RawQuery != "" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestActivityKudosDelete(t *testing.T) {
	client := newCassetteClient(testToken, "activity_kudos_delete")
	err := NewActivityKudosService(client, 118229063).Delete().Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	// from here on out just check the request parameters
	s := NewActivityKudosService(newStoreRequestClient(), 123)

	// path
	s.Delete().Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/activities/123/kudos" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.Method != "DELETE" {
		t.Errorf("request method incorrect, got %v", transport.request.Method)
	}

	if transport.request.URL.RawQuery != "" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestActivityKudosBadJSON(t *testing.T) {
	var err error
	s := NewActivityKudosService(NewStubResponseClient("bad json"), 123)

	_, err = s.List().Do()
	if err == nil {
		t.Error("should return a bad json error")
	}
}
