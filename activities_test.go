package strava

import (
	"testing"
)

func TestActivitiesGet(t *testing.T) {
	// if you need to change this you should also update tests below
	if c := structAttributeCount(&ActivityDetailed{}); c != 61 {
		t.Fatalf("incorrect number of detailed attributes, %d != 61", c)
	}

	client := newCassetteClient(testToken, "activity_get")
	activity, err := NewActivitiesService(client).Get(103221154).Do()

	expected := &ActivityDetailed{}

	expected.Id = 103221154
	expected.ResourceState = 3
	expected.ExternalId = "2010-08-15-11-04-29.fit"
	expected.UploadId = 112859609
	expected.Athlete.Id = 227615
	expected.Name = "08/15/2010 Davis, CA"
	expected.Description = "Something Special"
	expected.Type = ActivityTypes.Ride
	expected.Distance = 20739.1
	expected.MovingTime = 2836
	expected.ElapsedTime = 3935
	expected.TotalElevationGain = 22.0
	expected.StartLocation = Location{38.55, -121.82}
	expected.EndLocation = Location{38.56, -121.78}
	expected.City = "Davis"
	expected.State = "CA"
	expected.Private = false

	expected.StartDateString = "2010-08-15T18:04:29Z"
	expected.StartDateLocalString = "2010-08-15T11:04:29Z"

	expected.AchievementCount = 0
	expected.KudosCount = 1
	expected.CommentCount = 1
	expected.AthleteCount = 2
	expected.PhotoCount = 0

	expected.Map.Id = "a103221154"
	expected.Map.Polyline = "_ugjFpiofV?dSOp@@BF@DD@PCbA?|AFzAGZ[JITDp@ArD@jHEtF@vECvH@vB?bHC|D@zCGvD@JPT@VEj@T~@D^EvB?tIOpDCxAGnAAnCP`IFnA@`AF`BNhBBbCPzDB~BAv@B|@C|NBjWChE@bAAv@@rEEbBC^KZm@h@]f@WTOREPIDY?C@MNW`@k@`@sAl@oBl@UL[`@Uj@S~@Ab@Br@VjAB`@G\\ELOLGD_@Fw@IYA]By@Cg@De@@gACyB@oACaLBqO?y@BiAEyCBwFAqEBsBA{AFa@CqA@q@Eo@BqAEaF@oACiH@}DC}C@w@CIEEICMD_A@cBCmK@_DCmABaAA}GMm@Mg@Yg@a@[s@QoD@{FCiD?oFEwD@yJImAD_ACeE@uAGc@BmE@_CCuD@aBE}A@q@A_A@s@AqB@s@Fo@Le@Ng@XsAdAq@|@q@fA{@z@aAp@_A\\cFdAo@Rq@LkAXyATiDLmFEECCECIHsA?yF@kBCoBBoA@sJ@i@CiE@aHCkG@}LAqB@oKCoDBiFCmCBeDAuA?aa@?eDDsBAyMDcN@kVAuCBaMCoCYiCGgAAa@Ba@VaBFu@Bw@CeDDc@To@z@uAPs@Bc@IkE@iDAw@Kw@Gc@UcA_@sASsAGw@BkB@wD?oPDc@N[XSXKVOZ_@J[HuAGw@}BaIG_@Cc@BwAAcADiH?mHDwJ?_JPc[AaPAq@F{CAs@@eCAgCDwBCmCBmGDyBA{AFgFEyIBeMDgFHkD@yC?kDGiC@]FILEn@C\\KRQf@k@N[jAeETk@f@}@v@gAh@a@`Ae@nA]fAGlE@jBIjDAnFBvH^dFPzBLfFPvHZpIFnAAjCF|B?~@AhHDdB?bB@`AAdBBtA?hFJbAGjGFxG?f@Ft@Nz@XpF|C~@\\`@HNEXcALS`@_@CDQEITOHGFKv@Mh@BPHLbA`@|A|@~@b@bATbADPBJLDNDhCPrA?nCDVX^LZ?HIXCRKdM@TDD^F@JGlAI`@Fv@ALK^A^@VFf@BBTAHCD]FKp@FBJAJAA"
	expected.Map.SummaryPolyline = "cugjFjiofVg@bfBnA~eCkCfE{JzEg@fJotBRQod@mC_DwwASoPnKsN~CoK?SknFf@wLbB{EScLcBoKRgYlCcGmCwLz@{~D~CcBfEoKbGwBjeCjC~MbGbBwB{@~CzJfEpAvL?fY`BR"

	expected.Trainer = false
	expected.Commute = true
	expected.Manual = false
	expected.Private = false
	expected.Flagged = false
	expected.HasKudoed = false

	expected.GearId = "b77076"

	expected.AverageSpeed = 7.3
	expected.MaximunSpeed = 13.7
	expected.AverageCadence = 73.2
	expected.AverageTemperature = 27.0
	expected.AveragePower = 140.2
	expected.Kilojoules = 397.5
	expected.AverageHeartrate = 104.4
	expected.MaximumHeartrate = 147.0
	expected.Calories = 443.2
	expected.Truncated = 0

	expected.SegmentEfforts = make([]*SegmentEffortSummary, 1)
	expected.SegmentEfforts[0] = new(SegmentEffortSummary)
	expected.SegmentEfforts[0].Id = 2226314149
	expected.SegmentEfforts[0].Name = "DBC Junior Airport loop sprint orange flag"
	expected.SegmentEfforts[0].Segment.Id = 5775164
	expected.SegmentEfforts[0].Segment.Name = "DBC Junior Airport loop sprint orange flag"
	expected.SegmentEfforts[0].Segment.ActivityType = ActivityTypes.Ride
	expected.SegmentEfforts[0].Segment.Distance = 809.9
	expected.SegmentEfforts[0].Segment.AverageGrade = 0
	expected.SegmentEfforts[0].Segment.MaximumGrade = 0.6
	expected.SegmentEfforts[0].Segment.ElevationHigh = 28.3
	expected.SegmentEfforts[0].Segment.ElevationLow = 27.7
	expected.SegmentEfforts[0].Segment.StartLocation = Location{38.58258, -121.851906}
	expected.SegmentEfforts[0].Segment.EndLocation = Location{38.589291, -121.854774}
	expected.SegmentEfforts[0].Segment.ClimbCategory = ClimbCategories.NotCategorized
	expected.SegmentEfforts[0].Segment.City = "Davis"
	expected.SegmentEfforts[0].Segment.State = "CA"
	expected.SegmentEfforts[0].Segment.Private = false
	expected.SegmentEfforts[0].Segment.PRTime = 113
	expected.SegmentEfforts[0].Segment.PRDistance = 805.6
	expected.SegmentEfforts[0].Activity.Id = 103221154
	expected.SegmentEfforts[0].Athlete.Id = 227615
	expected.SegmentEfforts[0].KOMRank = 0
	expected.SegmentEfforts[0].PRRank = 0
	expected.SegmentEfforts[0].ElapsedTime = 113
	expected.SegmentEfforts[0].MovingTime = 113
	expected.SegmentEfforts[0].StartDateString = "2010-08-15T18:23:07Z"
	expected.SegmentEfforts[0].StartDateLocalString = "2010-08-15T11:23:07Z"
	expected.SegmentEfforts[0].Distance = 805.6
	expected.SegmentEfforts[0].StartIndex = 1112
	expected.SegmentEfforts[0].EndIndex = 1225

	expected.SplitsMetric = make([]*Split, 0)
	expected.SplitsStandard = make([]*Split, 0)
	expected.BestEfforts = make([]*BestEffort, 0)

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if activity.StartDate.IsZero() || activity.StartDateLocal.IsZero() {
		t.Error("dates are not parsed")
	}

	if len(activity.SegmentEfforts) == 0 {
		t.Fatal("no segment efforts!?!?!")
	}

	if activity.SegmentEfforts[0].StartDate.IsZero() || activity.SegmentEfforts[0].StartDateLocal.IsZero() {
		t.Error("segment effort dates are not parsed")
	}

	for _, prob := range structCompare(t, activity, expected) {
		t.Error(prob)
	}

	for _, prob := range structCompare(t, activity.SegmentEfforts[0], expected.SegmentEfforts[0]) {
		t.Error(prob)
	}

	// run
	client = newCassetteClient(testToken, "activity_get_run")
	activity, err = NewActivitiesService(client).Get(103359122).Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if activity.Type != ActivityTypes.Run {
		t.Error("activity type should be run")
	}

	if len(activity.SplitsMetric) == 0 {
		t.Fatal("no metric splits")
	}

	if len(activity.SplitsStandard) == 0 {
		t.Fatal("no standard splits")
	}

	if len(activity.BestEfforts) == 0 {
		t.Fatal("no best efforts")
	}

	split := &Split{
		Distance:            1000.0,
		ElapsedTime:         327,
		ElevationDifference: 14.4,
		MovingTime:          272,
		Split:               1,
	}

	for _, prob := range structCompare(t, split, activity.SplitsMetric[0]) {
		t.Error(prob)
	}

	split = &Split{
		Distance:            1612.0,
		ElapsedTime:         509,
		ElevationDifference: 12.6,
		MovingTime:          454,
		Split:               1,
	}

	for _, prob := range structCompare(t, split, activity.SplitsStandard[0]) {
		t.Error(prob)
	}

	bestEffort := &BestEffort{
		Id:                   474685446,
		Name:                 "400m",
		ElapsedTime:          111,
		MovingTime:           112,
		StartDateString:      "2013-09-23T00:15:15Z",
		StartDateLocalString: "2013-09-22T17:15:15Z",
		Distance:             400,
		StartIndex:           1,
		EndIndex:             109,
	}

	bestEffort.Activity.Id = 103359122
	bestEffort.Athlete.Id = 227615

	for _, prob := range structCompare(t, bestEffort, activity.BestEfforts[0]) {
		t.Error(prob)
	}

	// from here on out just check the request parameters
	s := NewActivitiesService(newStoreRequestClient())

	// path
	s.Get(123).IncludeAllEfforts().Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/activities/123" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "include_all_efforts=true" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestActivitiesListComments(t *testing.T) {
	// if you need to change this you should also update tests below
	if c := structAttributeCount(&CommentSummary{}); c != 18 {
		t.Fatalf("incorrect number of detailed attributes, %d != 18", c)
	}

	client := newCassetteClient(testToken, "activity_list_comments")
	comments, err := NewActivitiesService(client).ListComments(103221154).Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if len(comments) == 0 {
		t.Fatal("comments not parsed")
	}

	if v := comments[0].Id; v != 19035182 {
		t.Errorf("value incorrect, got %v", v)
	}

	if v := comments[0].ActivityId; v != 103221154 {
		t.Errorf("value incorrect, got %v", v)
	}

	if v := comments[0].Text; v != "Testing!!!" {
		t.Errorf("value incorrect, got %v", v)
	}

	if comments[0].CreatedAt.IsZero() {
		t.Error("dates not parsed")
	}

	if comments[0].Athlete.CreatedAt.IsZero() || comments[0].Athlete.UpdatedAt.IsZero() {
		t.Error("athlete dates not parsed")
	}

	// from here on out just check the request parameters
	s := NewActivitiesService(newStoreRequestClient())

	// path
	s.ListComments(123).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/activities/123/comments" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}

	// parameters
	s.ListComments(123).IncludeMarkdown().Do()

	transport = s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/activities/123/comments" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "markdown=true" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}

	// parameters2
	s.ListComments(123).Page(1).PerPage(10).Do()

	transport = s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/activities/123/comments" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "page=1&per_page=10" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestActivitiesListKudoers(t *testing.T) {
	client := newCassetteClient(testToken, "activity_list_kudoers")
	athletes, err := NewActivitiesService(client).ListKudoers(103221154).Do()

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
	s := NewActivitiesService(newStoreRequestClient())

	// path
	s.ListKudoers(123).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/activities/123/kudos" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}

	// parameters
	s.ListKudoers(123).Page(2).PerPage(3).Do()

	transport = s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/activities/123/kudos" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "page=2&per_page=3" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestActivitiesListPhotos(t *testing.T) {
	// if you need to change this you should also update tests below
	if c := structAttributeCount(&PhotoSummary{}); c != 9 {
		t.Fatalf("incorrect number of detailed attributes, %d != 9", c)
	}

	// token for 3545423, I wasn't able to post a test photo for the other account
	client := newCassetteClient("f578367dbb2288fb9f91090fa676111fdc5e8698", "activity_list_photos")
	photos, err := NewActivitiesService(client).ListPhotos(103374194).Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if len(photos) == 0 {
		t.Fatal("photos not parsed")
	}

	expected := &PhotoSummary{}

	expected.Id = 19219017
	expected.ActivityId = 103374194
	expected.Reference = "http://instagram.com/p/ipv-OOyd3a/"
	expected.UID = "624241007441599962_905799726"
	expected.Caption = "Yest"
	expected.Type = "InstagramPhoto"
	expected.UploadedAtString = "2014-01-02T04:02:28Z"
	expected.CreatedAtString = "2014-01-02T04:04:00Z"

	if photos[0].CreatedAt.IsZero() || photos[0].UploadedAt.IsZero() {
		t.Error("dates are not parsed")
	}

	for _, prob := range structCompare(t, photos[0], expected) {
		t.Error(prob)
	}

	// from here on out just check the request parameters
	s := NewActivitiesService(newStoreRequestClient())

	// path
	s.ListPhotos(321).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/activities/321/photos" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestActivitiesListZones(t *testing.T) {
	// if you need to change this you should also update tests below
	if c := structAttributeCount(&ZonesSummary{}); c != 9 {
		t.Fatalf("incorrect number of detailed attributes, %d != 9", c)
	}

	// token for 3545423, I wasn't able to post a test photo for the other account
	client := newCassetteClient(testToken, "activity_list_zones")
	zones, err := NewActivitiesService(client).ListZones(103221154).Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if len(zones) == 0 {
		t.Fatal("zones not parsed")
	}

	if v := zones[0].Score; v != 12 {
		t.Errorf("value incorrect, got %v", v)
	}

	if v := zones[0].Type; v != "heartrate" {
		t.Errorf("value incorrect, got %v", v)
	}

	if v := zones[0].SensorBased; v != true {
		t.Errorf("value incorrect, got %v", v)
	}

	if v := zones[0].CustonZones; v != false {
		t.Errorf("value incorrect, got %v", v)
	}

	if v := zones[0].Max; v != 184 {
		t.Errorf("value incorrect, got %v", v)
	}

	if v := zones[1].BikeWeight; v != 10 {
		t.Errorf("value incorrect, got %v", v)
	}

	if v := zones[1].AthleteWeight; v != 75 {
		t.Errorf("value incorrect, got %v", v)
	}

	if len(zones[0].Buckets) == 0 {
		t.Fatal("Buckets not parsed")
	}

	expected := &ZoneBucket{Max: 143, Min: 108, Time: 910}

	for _, prob := range structCompare(t, zones[0].Buckets[1], expected) {
		t.Error(prob)
	}

	// from here on out just check the request parameters
	s := NewActivitiesService(newStoreRequestClient())

	// path
	s.ListZones(321).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/activities/321/zones" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestActivitiesBadJSON(t *testing.T) {
	var err error
	s := NewActivitiesService(newStubResponseClient("bad json"))

	_, err = s.Get(123).Do()
	if err == nil {
		t.Error("should return a bad json error")
	}

	_, err = s.ListComments(123).Do()
	if err == nil {
		t.Error("should return a bad json error")
	}

	_, err = s.ListKudoers(123).Do()
	if err == nil {
		t.Error("should return a bad json error")
	}

	_, err = s.ListPhotos(123).Do()
	if err == nil {
		t.Error("should return a bad json error")
	}

	_, err = s.ListZones(123).Do()
	if err == nil {
		t.Error("should return a bad json error")
	}
}

func TestLocationString(t *testing.T) {
	l := Location{1, 2}

	if l.String() != "[1.000000, 2.000000]" {
		t.Errorf("location string has changed, got %v", l.String())
	}
}

func TestActivityType(t *testing.T) {
	if id := ActivityTypes.Ride.Id(); id != 1 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.Ride.String(); s != "Ride" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	if id := ActivityTypes.AlpineSki.Id(); id != 2 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.AlpineSki.String(); s != "Alpine Ski" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	if id := ActivityTypes.BackcountrySki.Id(); id != 3 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.BackcountrySki.String(); s != "Backcountry Ski" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	if id := ActivityTypes.Hike.Id(); id != 4 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.Hike.String(); s != "Hike" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	if id := ActivityTypes.IceSkate.Id(); id != 5 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.IceSkate.String(); s != "Ice Skate" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	if id := ActivityTypes.InlineSkate.Id(); id != 6 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.InlineSkate.String(); s != "Inline Skate" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	if id := ActivityTypes.NordicSki.Id(); id != 7 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.NordicSki.String(); s != "Nordic Ski" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	if id := ActivityTypes.RollerSki.Id(); id != 8 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.RollerSki.String(); s != "Roller Ski" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	if id := ActivityTypes.Run.Id(); id != 9 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.Run.String(); s != "Run" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	if id := ActivityTypes.Walk.Id(); id != 10 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.Walk.String(); s != "Walk" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	if id := ActivityTypes.Workout.Id(); id != 11 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.Workout.String(); s != "Workout" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	if id := ActivityTypes.Snowboard.Id(); id != 12 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.Snowboard.String(); s != "Snowboard" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	if id := ActivityTypes.Snowshoe.Id(); id != 13 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.Snowshoe.String(); s != "Snowshoe" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	if id := ActivityTypes.Kitesurf.Id(); id != 14 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.Kitesurf.String(); s != "Kitesurf" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	if id := ActivityTypes.Windsurf.Id(); id != 15 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.Windsurf.String(); s != "Windsurf" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	if id := ActivityTypes.Swim.Id(); id != 16 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.Swim.String(); s != "Swim" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	// other
	ty := ActivityType(30)
	if id := ty.Id(); id != 0 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ty.String(); s != "Activity" {
		t.Errorf("activity type string incorrect, got %v", s)
	}
}
