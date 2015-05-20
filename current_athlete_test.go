package strava

import (
	"reflect"
	"testing"
	"time"
)

func TestCurrentAthleteGet(t *testing.T) {
	client := newCassetteClient(testToken, "current_athlete_get")
	athlete, err := NewCurrentAthleteService(client).Get().Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	expected := &AthleteDetailed{}

	expected.Id = 227615
	expected.FirstName = "John"
	expected.LastName = "Applestrava"
	expected.Profile = "http://dgalywyr863hv.cloudfront.net/pictures/athletes/227615/41555/3/large.jpg"
	expected.ProfileMedium = "http://dgalywyr863hv.cloudfront.net/pictures/athletes/227615/41555/3/medium.jpg"
	expected.City = "San Francisco"
	expected.State = "CA"
	expected.Country = "United States"
	expected.Gender = "M"
	expected.CreatedAt, _ = time.Parse(timeFormat, "2012-01-18T18:20:37Z")
	expected.UpdatedAt, _ = time.Parse(timeFormat, "2014-01-21T06:23:32Z")
	expected.Premium = true

	expected.FollowerCount = 1
	expected.FriendCount = 35
	expected.MutualFriendCount = 0
	expected.DatePreference = "%m/%d/%Y"
	expected.MeasurementPreference = "feet"
	expected.FTP = 200
	expected.Weight = 70.1
	expected.Email = "mobiledemo@strava.com"

	expected.Clubs = make([]*ClubSummary, 1)
	expected.Clubs[0] = new(ClubSummary)
	expected.Clubs[0].Id = 45255
	expected.Clubs[0].Name = "Test Club"
	expected.Clubs[0].Profile = "avatar/club/large.png"
	expected.Clubs[0].ProfileMedium = "avatar/club/medium.png"

	expected.Bikes = make([]*GearSummary, 1)
	expected.Bikes[0] = new(GearSummary)
	expected.Bikes[0].Id = "b77076"
	expected.Bikes[0].Name = "burrito burner"
	expected.Bikes[0].Distance = 536292.3

	expected.Shoes = make([]*GearSummary, 1)
	expected.Shoes[0] = new(GearSummary)
	expected.Shoes[0].Id = "g5697"
	expected.Shoes[0].Name = "ASICS Kayano"
	expected.Shoes[0].Primary = true
	expected.Shoes[0].Distance = 17224.6

	if !reflect.DeepEqual(athlete, expected) {
		t.Errorf("should match\n%v\n%v", athlete, expected)
	}

	// from here on out just check the request parameters
	s := NewCurrentAthleteService(newStoreRequestClient())

	// path
	s.Get().Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/athlete" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestCurrentAthleteUpdate(t *testing.T) {
	client := newCassetteClient(testToken, "current_athlete_put")
	athlete, err := NewCurrentAthleteService(client).Update().Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if athlete.CreatedAt.IsZero() || athlete.UpdatedAt.IsZero() {
		t.Error("dates not parsed")
	}

	// from here on out just check the request parameters
	s := NewCurrentAthleteService(newStoreRequestClient())

	// parameters1
	s.Update().City("city").State("state").Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/athlete" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.Method != "PUT" {
		t.Errorf("request method incorrect, got %v", transport.request.Method)
	}

	if transport.request.URL.RawQuery != "city=city&state=state" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}

	// parameters2
	s.Update().Country("USA").Gender("M").Do()

	transport = s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.RawQuery != "country=USA&sex=M" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}

	// parameters3
	s.Update().Weight(100.0).Do()

	transport = s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.RawQuery != "weight=100" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestCurrentAthleteListActivities(t *testing.T) {
	client := newCassetteClient(testToken, "current_athlete_list_activities")
	activities, err := NewCurrentAthleteService(client).ListActivities().Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if len(activities) == 0 {
		t.Fatal("activities not parsed")
	}

	if activities[0].StartDate.IsZero() || activities[0].StartDateLocal.IsZero() {
		t.Error("dates not parsed")
	}

	// from here on out just check the request parameters
	s := NewCurrentAthleteService(newStoreRequestClient())

	// parameters1
	s.ListActivities().Page(2).PerPage(3).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/athlete/activities" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "page=2&per_page=3" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}

	// parameters2
	s.ListActivities().Before(10002).After(2003).Do()

	transport = s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/athlete/activities" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "after=2003&before=10002" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestCurrentAthleteListFriendsActivities(t *testing.T) {
	client := newCassetteClient(testToken, "current_athlete_list_friends_activities")
	activities, err := NewCurrentAthleteService(client).ListFriendsActivities().Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if len(activities) == 0 {
		t.Fatal("activities not parsed")
	}

	if activities[0].StartDate.IsZero() || activities[0].StartDateLocal.IsZero() {
		t.Error("dates not parsed")
	}

	// from here on out just check the request parameters
	s := NewCurrentAthleteService(newStoreRequestClient())

	// parameters1
	s.ListFriendsActivities().Page(2).PerPage(3).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/activities/following" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "page=2&per_page=3" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}

	// parameters2
	s.ListFriendsActivities().Before(10002).Do()

	transport = s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/activities/following" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "before=10002" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestCurrentAthleteListFriends(t *testing.T) {
	client := newCassetteClient(testToken, "current_athlete_list_friends")
	friends, err := NewCurrentAthleteService(client).ListFriends().Do()

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
	s := NewCurrentAthleteService(newStoreRequestClient())

	// parameters
	s.ListFriends().Page(2).PerPage(3).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/athlete/friends" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "page=2&per_page=3" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestCurrentAthleteListFollowers(t *testing.T) {
	client := newCassetteClient(testToken, "current_athlete_list_followers")
	followers, err := NewCurrentAthleteService(client).ListFollowers().Do()

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
	s := NewCurrentAthleteService(newStoreRequestClient())

	// parameters
	s.ListFollowers().Page(5).PerPage(3).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/athlete/followers" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "page=5&per_page=3" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestCurrentAthleteListClubs(t *testing.T) {
	client := newCassetteClient(testToken, "current_athlete_list_clubs")
	clubs, err := NewCurrentAthleteService(client).ListClubs().Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if len(clubs) == 0 {
		t.Fatal("clubs not parsed")
	}

	// from here on out just check the request parameters
	s := NewCurrentAthleteService(newStoreRequestClient())

	// path
	s.ListClubs().Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/athlete/clubs" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestCurrentAthleteListStarredSegments(t *testing.T) {
	client := newCassetteClient(testToken, "current_athlete_list_starred_segments")
	segments, err := NewCurrentAthleteService(client).ListStarredSegments().Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if len(segments) == 0 {
		t.Fatal("segments not parsed")
	}

	// from here on out just check the request parameters
	s := NewCurrentAthleteService(newStoreRequestClient())

	// path
	s.ListStarredSegments().Page(1).PerPage(2).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/segments/starred" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "page=1&per_page=2" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestCurrentAthleteBadJSON(t *testing.T) {
	var err error
	s := NewCurrentAthleteService(NewStubResponseClient("bad json"))

	_, err = s.Get().Do()
	if err == nil {
		t.Error("should return a bad json error")
	}

	_, err = s.Update().Do()
	if err == nil {
		t.Error("should return a bad json error")
	}

	_, err = s.ListActivities().Do()
	if err == nil {
		t.Error("should return a bad json error")
	}

	_, err = s.ListFriends().Do()
	if err == nil {
		t.Error("should return a bad json error")
	}

	_, err = s.ListFollowers().Do()
	if err == nil {
		t.Error("should return a bad json error")
	}

	_, err = s.ListClubs().Do()
	if err == nil {
		t.Error("should return a bad json error")
	}

	_, err = s.ListStarredSegments().Do()
	if err == nil {
		t.Error("should return a bad json error")
	}
}
