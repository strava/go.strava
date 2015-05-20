package strava

import (
	"reflect"
	"testing"
	"time"
)

func TestSegmentsGet(t *testing.T) {
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
	expected.Starred = false

	expected.CreatedAt, _ = time.Parse(timeFormat, "2009-09-21T20:29:41Z")
	expected.UpdatedAt, _ = time.Parse(timeFormat, "2014-06-18T13:01:35Z")

	expected.TotalElevationGain = 155.733

	expected.Map.Id = "s229781"
	expected.Map.Polyline = "}g|eFnpqjVl@En@Md@HbAd@d@^h@Xx@VbARjBDh@OPQf@w@d@k@XKXDFPH\\EbGT`AV`@v@|@NTNb@?XOb@cAxAWLuE@eAFMBoAv@eBt@q@b@}@tAeAt@i@dAC`AFZj@dB?~@[h@MbAVn@b@b@\\d@Eh@Qb@_@d@eB|@c@h@WfBK|AMpA?VF\\\\t@f@t@h@j@|@b@hCb@b@XTd@Bl@GtA?jAL`ALp@Tr@RXd@Rx@Pn@^Zh@Tx@Zf@`@FTCzDy@f@Yx@m@n@Op@VJr@"

	expected.EffortCount = 83035
	expected.AthleteCount = 10506
	expected.Hazardous = false
	expected.StarCount = 405

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if !reflect.DeepEqual(segment, expected) {
		t.Errorf("should match\n%v\n%v", segment, expected)
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

func TestSegmentsListEfforts(t *testing.T) {
	client := newCassetteClient(testToken, "segment_list_efforts")
	efforts, err := NewSegmentsService(client).ListEfforts(229781).Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if len(efforts) == 0 {
		t.Fatal("efforts not parsed")
	}

	if len(efforts) != 30 {
		t.Fatal("wrong number of efforts returned")
	}

	expected := &SegmentEffortSummary{}

	expected.Id = 1323785488
	expected.Name = "Hawk Hill"

	expected.Segment.Id = 229781
	expected.Segment.Name = "Hawk Hill"
	expected.Segment.ActivityType = ActivityTypes.Ride
	expected.Segment.Distance = 2684.82
	expected.Segment.AverageGrade = 5.7
	expected.Segment.MaximumGrade = 14.2
	expected.Segment.ElevationHigh = 245.3
	expected.Segment.ElevationLow = 92.4
	expected.Segment.StartLocation = Location{37.8331119, -122.4834356}
	expected.Segment.EndLocation = Location{37.8280722, -122.4981393}
	expected.Segment.ClimbCategory = ClimbCategories.Category4
	expected.Segment.City = "San Francisco"
	expected.Segment.State = "CA"
	expected.Segment.Country = "United States"
	expected.Segment.Private = false

	expected.Activity.Id = 67124336
	expected.Athlete.Id = 118571

	expected.ElapsedTime = 769
	expected.MovingTime = 769

	expected.StartDate, _ = time.Parse(timeFormat, "1970-01-01T00:29:39Z")
	expected.StartDateLocal, _ = time.Parse(timeFormat, "1969-12-31T16:29:39Z")

	expected.Distance = 2697.7
	expected.StartIndex = 1623
	expected.EndIndex = 2239

	if !reflect.DeepEqual(efforts[0], expected) {
		t.Errorf("should match\n%v\n%v", efforts[0], expected)
	}

	// from here on out just check the request parameters
	s := NewSegmentsService(newStoreRequestClient())

	// path
	s.ListEfforts(321).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/segments/321/all_efforts" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}

	// parameters
	s.ListEfforts(123).AthleteId(321).Do()

	transport = s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.RawQuery != "athlete_id=321" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}

	// parameters2
	sTime, _ := time.Parse(timeFormat, "1969-12-31T16:29:39Z")
	eTime, _ := time.Parse(timeFormat, "1970-01-01T00:29:39Z")
	s.ListEfforts(123).DateRange(sTime, eTime).Do()

	transport = s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.RawQuery != "end_date_local=1970-01-01T00%3A29%3A39Z&start_date_local=1969-12-31T16%3A29%3A39Z" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}

	// parameters3
	s.ListEfforts(123).Page(1).PerPage(2).Do()

	transport = s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.RawQuery != "page=1&per_page=2" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestSegmentsGetLeaderboard(t *testing.T) {
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
	expected.AthleteGender = Genders.Male
	expected.AthleteId = 123529
	expected.AverageHeartrate = 190.5
	expected.AveragePower = 460.8
	expected.Distance = 2659.9
	expected.ElapsedTime = 360
	expected.MovingTime = 360
	expected.ActivityId = 46320211
	expected.EffortId = 801006623
	expected.Rank = 1

	expected.StartDate, _ = time.Parse(timeFormat, "2013-03-29T13:49:35Z")
	expected.StartDateLocal, _ = time.Parse(timeFormat, "2013-03-29T06:49:35Z")

	if !reflect.DeepEqual(leaderboard.Entries[0], expected) {
		t.Errorf("should match\n%v\n%v", leaderboard.Entries[0], expected)
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
	s.GetLeaderboard(123).Gender(Genders.Male).Following().AgeGroup(AgeGroups.From25to34).Do()

	transport = s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.RawQuery != "age_group=25_34&following=true&gender=M" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}

	// parameters2
	s.GetLeaderboard(123).Gender(Genders.Female).ClubId(2).WeightClass(WeightClasses.From200PlusPounds).Do()

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

	// parameters5
	s.GetLeaderboard(123).ContextEntries(3).Do()

	transport = s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.RawQuery != "context_entries=3" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestSegmentsExplore(t *testing.T) {
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

	if !reflect.DeepEqual(segments[0], expected) {
		t.Errorf("should match\n%v\n%v", segments[0], expected)
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

	_, err = s.ListEfforts(123).Do()
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
