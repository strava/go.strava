package strava

import (
	"reflect"
	"testing"
)

func TestClubsGet(t *testing.T) {
	client := newCassetteClient(testToken, "club_get")
	club, err := NewClubsService(client).Get(45255).Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	expected := &ClubDetailed{}
	expected.Id = 45255
	expected.Name = "Test Club"
	expected.ProfileMedium = "avatar/club/medium.png"
	expected.Profile = "avatar/club/large.png"
	expected.Description = "test description"
	expected.Type = ClubTypes.CasualClub
	expected.SportType = SportTypes.Cycling
	expected.City = "San Francisco"
	expected.State = "California"
	expected.Country = "United States"
	expected.Private = true
	expected.MemberCount = 2

	if !reflect.DeepEqual(club, expected) {
		t.Errorf("should match\n%v\n%v", club, expected)
	}

	// from here on out just check the request parameters
	s := NewClubsService(newStoreRequestClient())

	// path
	s.Get(321).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/clubs/321" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestClubsListMembers(t *testing.T) {
	client := newCassetteClient(testToken, "club_list_members")
	members, err := NewClubsService(client).ListMembers(45255).Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if len(members) == 0 {
		t.Fatal("no members, how can I test?")
	}

	if members[0].CreatedAt.IsZero() || members[0].UpdatedAt.IsZero() {
		t.Error("dates are not parsed")
	}

	// from here on out just check the request parameters
	s := NewClubsService(newStoreRequestClient())

	// per_page
	s.ListMembers(45255).PerPage(1).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.RawQuery != "per_page=1" {
		t.Error("per_page request incorrect")
	}

	// page
	s.ListMembers(45255).Page(3).Do()

	transport = s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.RawQuery != "page=3" {
		t.Error("page request incorrect")
	}
}

func TestClubsListActivities(t *testing.T) {
	client := newCassetteClient(testToken, "club_list_activities")
	activities, err := NewClubsService(client).ListActivities(45255).Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if len(activities) == 0 {
		t.Fatal("no activities, how can I test")
	}

	if activities[0].StartDate.IsZero() || activities[0].StartDateLocal.IsZero() {
		t.Error("dates are not parsed")
	}

	// from here on out just check the request parameters
	s := NewClubsService(newStoreRequestClient())

	// per_page
	s.ListActivities(45255).PerPage(1).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.RawQuery != "per_page=1" {
		t.Error("per_page request incorrect")
	}

	// page
	s.ListActivities(45255).Page(21).Do()
	transport = s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.RawQuery != "page=21" {
		t.Error("page request incorrect")
	}
}

func TestClubsBadJSON(t *testing.T) {
	var err error
	s := NewClubsService(NewStubResponseClient("bad json"))

	_, err = s.Get(123).Do()
	if err == nil {
		t.Error("should return a bad json error")
	}

	_, err = s.ListMembers(123).Do()
	if err == nil {
		t.Error("should return a bad json error")
	}

	_, err = s.ListActivities(123).Do()
	if err == nil {
		t.Error("should return a bad json error")
	}
}
