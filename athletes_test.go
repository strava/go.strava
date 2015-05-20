package strava

import (
	"reflect"
	"testing"
	"time"
)

func TestAthletesGet(t *testing.T) {
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
	expected.Country = "United States"
	expected.Gender = "M"
	expected.CreatedAt, _ = time.Parse(timeFormat, "2013-12-26T19:19:36Z")
	expected.UpdatedAt, _ = time.Parse(timeFormat, "2014-01-12T00:20:58Z")

	if !reflect.DeepEqual(athlete, expected) {
		t.Errorf("should match\n%v\n%v", athlete, expected)
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

func TestCAthletesListStarredSegments(t *testing.T) {
	client := newCassetteClient(testToken, "athlete_list_starred_segments")
	segments, err := NewAthletesService(client).ListStarredSegments(3545423).Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	expected := &PersonalSegmentSummary{}

	expected.Id = 229781
	expected.Name = "Hawk Hill"
	expected.ActivityType = ActivityTypes.Ride
	expected.Distance = 2684.82
	expected.AverageGrade = 5.7
	expected.MaximumGrade = 14.2
	expected.ElevationHigh = 245.3
	expected.ElevationLow = 92.4
	expected.StartLocation = Location{37.8331119, -122.4834356}
	expected.EndLocation = Location{37.8280722, -122.4981393}
	expected.ClimbCategory = ClimbCategories.Category4
	expected.City = "San Francisco"
	expected.State = "CA"
	expected.Country = "United States"
	expected.Private = false
	expected.Starred = true

	expected.AthletePR.Id = 3439333050
	expected.AthletePR.ElapsedTime = 550
	expected.AthletePR.Distance = 2713.4

	expected.AthletePR.StartDate, _ = time.Parse(timeFormat, "2013-01-21T19:05:07Z")
	expected.AthletePR.StartDateLocal, _ = time.Parse(timeFormat, "2013-01-21T11:05:07Z")

	expected.StarredDate, _ = time.Parse(timeFormat, "2014-07-24T23:23:24Z")

	if !reflect.DeepEqual(segments[0], expected) {
		t.Errorf("should match\n%v\n%v", segments[0], expected)
	}

	// from here on out just check the request parameters
	s := NewAthletesService(newStoreRequestClient())

	// path
	s.ListStarredSegments(123).Page(2).PerPage(3).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/athletes/123/segments/starred" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "page=2&per_page=3" {
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

func TestAthletesStats(t *testing.T) {
	client := newCassetteClient(testToken, "athlete_stats")
	stats, err := NewAthletesService(client).Stats(123).Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	expected := &AthleteStats{}
	expected.BiggestRideDistance = 205481
	expected.BiggestClimbElevationGain = 1224

	expected.RecentRideTotals.Count = 29
	expected.RecentRideTotals.Distance = 975679.00390625
	expected.RecentRideTotals.MovingTime = 155868
	expected.RecentRideTotals.ElapsedTime = 171286
	expected.RecentRideTotals.ElevationGain = 12460.902572631836
	expected.RecentRideTotals.AchievementCount = 336

	expected.RecentRunTotals.Count = 7
	expected.RecentRunTotals.Distance = 104608.0009765625
	expected.RecentRunTotals.MovingTime = 31383
	expected.RecentRunTotals.ElapsedTime = 32453
	expected.RecentRunTotals.ElevationGain = 2107.405592918396
	expected.RecentRunTotals.AchievementCount = 42

	expected.YTDRideTotals.Count = 45
	expected.YTDRideTotals.Distance = 1.58619e+06
	expected.YTDRideTotals.MovingTime = 257281
	expected.YTDRideTotals.ElapsedTime = 285315
	expected.YTDRideTotals.ElevationGain = 22430

	expected.YTDRunTotals.Count = 7
	expected.YTDRunTotals.Distance = 104608
	expected.YTDRunTotals.MovingTime = 31383
	expected.YTDRunTotals.ElapsedTime = 32453
	expected.YTDRunTotals.ElevationGain = 2107

	expected.AllRideTotals.Count = 765
	expected.AllRideTotals.Distance = 4.2918079e+07
	expected.AllRideTotals.MovingTime = 6386345
	expected.AllRideTotals.ElapsedTime = 7228437
	expected.AllRideTotals.ElevationGain = 550886

	expected.AllRunTotals.Count = 76
	expected.AllRunTotals.Distance = 937914
	expected.AllRunTotals.MovingTime = 256194
	expected.AllRunTotals.ElapsedTime = 268123
	expected.AllRunTotals.ElevationGain = 8062

	if !reflect.DeepEqual(stats, expected) {
		t.Errorf("should match\n%v\n%v", stats, expected)
	}
}

func TestAthletesListKOMs(t *testing.T) {
	client := newCassetteClient(testToken, "athlete_list_koms")
	efforts, err := NewAthletesService(client).ListKOMs(3776).Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if len(efforts) == 0 {
		t.Fatal("efforts not parsed")
	}

	if efforts[0].StartDate.IsZero() || efforts[0].StartDateLocal.IsZero() {
		t.Error("dates not parsed")
	}

	// from here on out just check the request parameters
	s := NewAthletesService(newStoreRequestClient())

	// parameters
	s.ListKOMs(123).PerPage(9).Page(8).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/athletes/123/koms" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "page=8&per_page=9" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestAthletesListActivities(t *testing.T) {
	client := newCassetteClient(testToken, "athlete_list_activies")
	activities, err := NewAthletesService(client).ListActivities(14507).Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if len(activities) == 0 {
		t.Fatal("efforts not parsed")
	}

	if activities[0].StartDate.IsZero() || activities[0].StartDateLocal.IsZero() {
		t.Error("dates not parsed")
	}

	// from here on out just check the request parameters
	s := NewAthletesService(newStoreRequestClient())

	// path
	s.ListActivities(123).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/athletes/123/activities" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}

	// parameters
	s.ListActivities(123).PerPage(9).Page(8).Do()

	transport = s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.RawQuery != "page=8&per_page=9" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}

	// parameters2
	s.ListActivities(123).Before(1391020072).Do()

	transport = s.client.httpClient.Transport.(*storeRequestTransport)

	if transport.request.URL.RawQuery != "before=1391020072" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}

	// parameters3
	s.ListActivities(123).After(0).Do()

	transport = s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.RawQuery != "after=0" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestAthletesBadJSON(t *testing.T) {
	var err error
	s := NewAthletesService(NewStubResponseClient("bad json"))

	_, err = s.Get(123).Do()
	if err == nil {
		t.Error("should return a bad json error")
	}

	_, err = s.ListStarredSegments(123).Do()
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

	_, err = s.ListKOMs(123).Do()
	if err == nil {
		t.Error("should return a bad json error")
	}

	_, err = s.ListActivities(123).Do()
	if err == nil {
		t.Error("should return a bad json error")
	}
}
