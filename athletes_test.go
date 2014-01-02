package strava

import (
	"testing"
)

func TestAthletesGet(t *testing.T) {
	// if you need to change this you should also update tests below
	if c := structAttributeCount(&AthleteSummary{}); c != 14 {
		t.Fatalf("incorrect number of detailed attributes, %d != 14", c)
	}

	client := newCassetteClient(testToken, "athlete_get")
	athlete, err := NewAthletesService(client).Get(3545423).Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	expected := &AthleteSummary{}

	expected.Id = 3545423
	expected.FirstName = "Strava"
	expected.LastName = "Testing"
	expected.Friend = "accepted"
	expected.Follower = "accepted"
	expected.Profile = "avatar/athlete/large.png"
	expected.ProfileMedium = "avatar/athlete/medium.png"
	expected.City = "Palo Alto"
	expected.State = "CA"
	expected.Gender = "M"
	expected.CreatedAtString = "2013-12-26T19:19:36Z"
	expected.UpdatedAtString = "2014-01-02T04:42:17Z"

	if athlete.CreatedAt.IsZero() || athlete.UpdatedAt.IsZero() {
		t.Error("segment effort dates are not parsed")
	}

	for _, prob := range structCompare(t, athlete, expected) {
		t.Error(prob)
	}

	// from here on out just check the request parameters
	s := NewAthletesService(newStoreRequestClient())

	// path
	s.Get(111).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/athletes/111" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestAthletesListFriends(t *testing.T) {
	client := newCassetteClient(testToken, "athlete_list_friends")
	friends, err := NewAthletesService(client).ListFriends(3545423).Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if len(friends) == 0 {
		t.Fatal("friends not parsed")
	}

	if friends[0].CreatedAt.IsZero() || friends[0].UpdatedAt.IsZero() {
		t.Error("dates not parsed")
	}

	// from here on out just check the request parameters
	s := NewAthletesService(newStoreRequestClient())

	// parameters
	s.ListFriends(123).Page(2).PerPage(3).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/athletes/123/friends" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "page=2&per_page=3" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestAthletesListFollowers(t *testing.T) {
	client := newCassetteClient(testToken, "athlete_list_followers")
	followers, err := NewAthletesService(client).ListFollowers(3545423).Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if len(followers) == 0 {
		t.Fatal("followers not parsed")
	}

	if followers[0].CreatedAt.IsZero() || followers[0].UpdatedAt.IsZero() {
		t.Error("dates not parsed")
	}

	// from here on out just check the request parameters
	s := NewAthletesService(newStoreRequestClient())

	// parameters
	s.ListFollowers(123).Page(2).PerPage(3).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/athletes/123/followers" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "page=2&per_page=3" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestAthletesListBothFollowing(t *testing.T) {
	client := newCassetteClient(testToken, "athlete_list_both_following")
	followers, err := NewAthletesService(client).ListBothFollowing(3545423).Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if len(followers) == 0 {
		t.Fatal("followers not parsed")
	}

	if followers[0].CreatedAt.IsZero() || followers[0].UpdatedAt.IsZero() {
		t.Error("dates not parsed")
	}

	// from here on out just check the request parameters
	s := NewAthletesService(newStoreRequestClient())

	// parameters
	s.ListBothFollowing(123).PerPage(7).Page(8).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/athletes/123/both-following" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "page=8&per_page=7" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestAthletesBadJSON(t *testing.T) {
	var err error
	s := NewAthletesService(newStubResponseClient("bad json"))

	_, err = s.Get(123).Do()
	if err == nil {
		t.Error("should return a bad json error")
	}

	_, err = s.ListFriends(123).Do()
	if err == nil {
		t.Error("should return a bad json error")
	}

	_, err = s.ListFollowers(123).Do()
	if err == nil {
		t.Error("should return a bad json error")
	}

	_, err = s.ListBothFollowing(123).Do()
	if err == nil {
		t.Error("should return a bad json error")
	}
}
