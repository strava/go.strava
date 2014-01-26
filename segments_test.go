package strava

import (
	"testing"
)

func TestSegmentsGet(t *testing.T) {
	// if you need to change this you should also update tests below
	if c := structAttributeCount(&SegmentDetailed{}); c != 28 {
		t.Fatalf("incorrect number of detailed attributes, %d != 28", c)
	}

	client := newCassetteClient(testToken, "segment_get")
	segment, err := NewSegmentsService(client).Get(229781).Do()

	expected := &SegmentDetailed{}

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

	expected.CreatedAtString = "2009-09-21T20:29:41Z"
	expected.UpdatedAtString = "2014-01-20T14:02:10Z"

	expected.TotalElevationGain = 155.733

	expected.Map.Id = "s229781"
	expected.Map.Polyline = "}g|eFnpqjVl@En@Md@HbAd@d@^h@Xx@VbARjBDh@OPQf@w@d@k@XKXDFPH\\EbGT`AV`@v@|@NTNb@?XOb@cAxAWLuE@eAFMBoAv@eBt@q@b@}@tAeAt@i@dAC`AFZj@dB?~@[h@MbAVn@b@b@\\d@Eh@Qb@_@d@eB|@c@h@WfBK|AMpA?VF\\\\t@f@t@h@j@|@b@hCb@b@XTd@Bl@GtA?jAL`ALp@Tr@RXd@Rx@Pn@^Zh@Tx@Zf@`@FTCzDy@f@Yx@m@n@Op@VJr@"

	expected.EffortCount = 64277
	expected.AthleteCount = 8604
	expected.Hazardous = false
	expected.Starred = false

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if segment.CreatedAt.IsZero() || segment.UpdatedAt.IsZero() {
		t.Error("dates are not parsed")
	}

	for _, prob := range structCompare(t, segment, expected) {
		t.Error(prob)
	}

	// from here on out just check the request parameters
	s := NewSegmentsService(newStoreRequestClient())

	// path
	s.Get(321).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/segments/321" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestSegmentsGetLeaderboard(t *testing.T) {
	// if you need to change this you should also update tests below
	if c := structAttributeCount(&SegmentLeaderboard{}); c != 3 {
		t.Fatalf("incorrect number of attributes, %d != 3", c)
	}

	if c := structAttributeCount(&SegmentLeaderboardEntry{}); c != 14 {
		t.Fatalf("incorrect number of attributes, %d != 14", c)
	}

	client := newCassetteClient(testToken, "segment_get_leaderboard")
	leaderboard, err := NewSegmentsService(client).GetLeaderboard(229781).Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if len(leaderboard.Entries) == 0 {
		t.Fatal("entries not parsed")
	}

	expected := &SegmentLeaderboardEntry{}

	expected.AthleteName = "Jim Whimpey"
	expected.AthleteProfile = "http://dgalywyr863hv.cloudfront.net/pictures/athletes/123529/15953/2/large.jpg"
	expected.AthleteGender = Male
	expected.AthleteId = 123529
	expected.AverageHeartrate = 190.5
	expected.AveragePower = 460.8
	expected.Distance = 2659.9
	expected.ElapsedTime = 360
	expected.MovingTime = 360
	expected.ActivityId = 46320211
	expected.EffortId = 801006623
	expected.Rank = 1

	expected.StartDateString = "2013-03-29T13:49:35Z"
	expected.StartDateLocalString = "2013-03-29T06:49:35Z"

	if leaderboard.Entries[0].StartDate.IsZero() || leaderboard.Entries[0].StartDateLocal.IsZero() {
		t.Error("dates are not parsed")
	}

	for _, prob := range structCompare(t, leaderboard.Entries[0], expected) {
		t.Error(prob)
	}

	// from here on out just check the request parameters
	s := NewSegmentsService(newStoreRequestClient())

	// path
	s.GetLeaderboard(123).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/segments/123/leaderboard" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}

	// parameters
	s.GetLeaderboard(123).Gender(Male).Following().AgeGroup(AgeGroups.From25to34).Do()

	transport = s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.RawQuery != "age_group=25_34&following=true&gender=M" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}

	// parameters2
	s.GetLeaderboard(123).Gender(Female).ClubId(2).WeightClass(WeightClasses.From200PlusPounds).Do()

	transport = s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.RawQuery != "club_id=2&gender=F&weight_class=200_plus" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}

	// parameters3
	s.GetLeaderboard(123).PerPage(3).DateRange(DateRanges.ThisWeek).Do()

	transport = s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.RawQuery != "date_range=this_week&per_page=3" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}

	// parameters4
	s.GetLeaderboard(123).PerPage(3).Page(1).DateRange(DateRanges.ThisWeek).Do()

	transport = s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.RawQuery != "date_range=this_week&page=1&per_page=3" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestSegmentsExplore(t *testing.T) {
	// if you need to change this you should also update tests below
	if c := structAttributeCount(&SegmentExplorerSegment{}); c != 9 {
		t.Fatalf("incorrect number of attributes, %d != 9", c)
	}

	client := newCassetteClient(testToken, "segment_explore")
	segments, err := NewSegmentsService(client).Explore(37.674887, -122.595185, 37.840461, -122.280015).Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if len(segments) == 0 {
		t.Fatal("segments not parsed")
	}

	expected := &SegmentExplorerSegment{}

	expected.Id = 229781
	expected.Name = "Hawk Hill"
	expected.ClimbCategory = ClimbCategories.Category4
	expected.AverageGrade = 5.7
	expected.StartLocation = Location{37.8331119, -122.4834356}
	expected.EndLocation = Location{37.8280722, -122.4981393}
	expected.ElevationDifference = 152.8
	expected.Distance = 2684.8
	expected.Polyline = "}g|eFnpqjVl@En@Md@HbAd@d@^h@Xx@VbARjBDh@OPQf@w@d@k@XKXDFPH\\EbGT`AV`@v@|@NTNb@?XOb@cAxAWLuE@eAFMBoAv@eBt@q@b@}@tAeAt@i@dAC`AFZj@dB?~@[h@MbAVn@b@b@\\d@Eh@Qb@_@d@eB|@c@h@WfBK|AMpA?VF\\\\t@f@t@h@j@|@b@hCb@b@XTd@Bl@GtA?jAL`ALp@Tr@RXd@Rx@Pn@^Zh@Tx@Zf@`@FTCzDy@f@Yx@m@n@Op@VJr@"

	for _, prob := range structCompare(t, segments[0], expected) {
		t.Error(prob)
	}

	// from here on out just check the request parameters
	s := NewSegmentsService(newStoreRequestClient())

	// path
	s.Explore(1, 2, 3, 4).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/segments/explore" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "bounds=1.000000%2C2.000000%2C3.000000%2C4.000000" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}

	// parameters
	s.Explore(4, 3, 2, 1).ActivityType("running").MinimumCategory(1).MaximumCategory(2).Do()

	transport = s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.RawQuery != "activity_type=running&bounds=4.000000%2C3.000000%2C2.000000%2C1.000000&max_cat=2&min_cat=1" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestSegmentsBadJSON(t *testing.T) {
	var err error
	s := NewSegmentsService(NewStubResponseClient("bad json"))

	_, err = s.Get(123).Do()
	if err == nil {
		t.Error("should return a bad json error")
	}

	_, err = s.GetLeaderboard(123).Do()
	if err == nil {
		t.Error("should return a bad json error")
	}

	_, err = s.Explore(1, 2, 3, 4).Do()
	if err == nil {
		t.Error("should return a bad json error")
	}
}

func TestClimbCategory(t *testing.T) {
	if id := ClimbCategories.NotCategorized.Id(); id != 0 {
		t.Errorf("climb category id incorrect, got %v", id)
	}

	if s := ClimbCategories.NotCategorized.String(); s != "Not Categorized" {
		t.Errorf("climb category string incorrect, got %v", s)
	}

	if id := ClimbCategories.Category4.Id(); id != 1 {
		t.Errorf("climb category id incorrect, got %v", id)
	}

	if s := ClimbCategories.Category4.String(); s != "Category 4" {
		t.Errorf("climb category string incorrect, got %v", s)
	}

	if id := ClimbCategories.Category3.Id(); id != 2 {
		t.Errorf("climb category id incorrect, got %v", id)
	}

	if s := ClimbCategories.Category3.String(); s != "Category 3" {
		t.Errorf("climb category string incorrect, got %v", s)
	}

	if id := ClimbCategories.Category2.Id(); id != 3 {
		t.Errorf("climb category id incorrect, got %v", id)
	}

	if s := ClimbCategories.Category2.String(); s != "Category 2" {
		t.Errorf("climb category string incorrect, got %v", s)
	}

	if id := ClimbCategories.Category1.Id(); id != 4 {
		t.Errorf("climb category id incorrect, got %v", id)
	}

	if s := ClimbCategories.Category1.String(); s != "Category 1" {
		t.Errorf("climb category string incorrect, got %v", s)
	}

	if id := ClimbCategories.HorsCategorie.Id(); id != 5 {
		t.Errorf("climb category id incorrect, got %v", id)
	}

	if s := ClimbCategories.HorsCategorie.String(); s != "Hors Categorie" {
		t.Errorf("climb category string incorrect, got %v", s)
	}
}
