package strava

import (
	"io/ioutil"
	"reflect"
	"testing"
	"time"
)

func TestActivitiesGet(t *testing.T) {
	client := newCassetteClient(testToken, "activity_get")
	activity, err := NewActivitiesService(client).Get(103221154).Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	expected := &ActivityDetailed{}

	expected.Id = 103221154
	expected.ExternalId = "2010-08-15-11-04-29.fit"
	expected.UploadId = 112859609
	expected.Athlete.Id = 227615
	expected.Athlete.FirstName = "John"
	expected.Athlete.LastName = "Applestrava"
	expected.Athlete.ProfileMedium = "http://dgalywyr863hv.cloudfront.net/pictures/athletes/227615/41555/3/medium.jpg"
	expected.Athlete.Profile = "http://dgalywyr863hv.cloudfront.net/pictures/athletes/227615/41555/3/large.jpg"
	expected.Athlete.City = "San Francisco"
	expected.Athlete.State = "CA"
	expected.Athlete.Country = "United States"
	expected.Athlete.Gender = "M"
	expected.Athlete.Friend = "accepted"
	expected.Athlete.Follower = "accepted"
	expected.Athlete.Premium = true
	expected.Athlete.CreatedAt, _ = time.Parse(timeFormat, "2012-01-18T18:20:37Z")
	expected.Athlete.UpdatedAt, _ = time.Parse(timeFormat, "2014-01-21T06:23:32Z")

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
	expected.Country = "United States"
	expected.Private = false

	expected.StartDate, _ = time.Parse(timeFormat, "2010-08-15T18:04:29Z")
	expected.StartDateLocal, _ = time.Parse(timeFormat, "2010-08-15T11:04:29Z")

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
	expected.Gear.Id = "b77076"
	expected.Gear.Name = "burrito burner"
	expected.Gear.Primary = false
	expected.Gear.Distance = 536292

	expected.AverageSpeed = 7.313
	expected.MaximunSpeed = 13.7
	expected.AverageCadence = 73.2
	expected.AverageTemperature = 27.0
	expected.AveragePower = 140.2
	expected.WeightedAveragePower = 202
	expected.Kilojoules = 397.5
	expected.DeviceWatts = true
	expected.AverageHeartrate = 104.4
	expected.MaximumHeartrate = 147.0
	expected.Calories = 443.2
	expected.Truncated = 0

	expected.SegmentEfforts = make([]*SegmentEffortSummary, 1)
	expected.SegmentEfforts[0] = new(SegmentEffortSummary)
	expected.SegmentEfforts[0].Id = 2226314143
	expected.SegmentEfforts[0].Name = "Russell Sprint "
	expected.SegmentEfforts[0].Segment.Id = 5858222
	expected.SegmentEfforts[0].Segment.Name = "Russell Sprint "
	expected.SegmentEfforts[0].Segment.ActivityType = ActivityTypes.Ride
	expected.SegmentEfforts[0].Segment.Distance = 780.5
	expected.SegmentEfforts[0].Segment.AverageGrade = 0.2
	expected.SegmentEfforts[0].Segment.MaximumGrade = 1.3
	expected.SegmentEfforts[0].Segment.ElevationHigh = 24.2
	expected.SegmentEfforts[0].Segment.ElevationLow = 22.4
	expected.SegmentEfforts[0].Segment.StartLocation = Location{38.547070, -121.823156}
	expected.SegmentEfforts[0].Segment.EndLocation = Location{38.547150, -121.832495}
	expected.SegmentEfforts[0].Segment.ClimbCategory = ClimbCategories.NotCategorized
	expected.SegmentEfforts[0].Segment.City = "Davis"
	expected.SegmentEfforts[0].Segment.State = "CA"
	expected.SegmentEfforts[0].Segment.Country = "United States"
	expected.SegmentEfforts[0].Segment.Private = false
	expected.SegmentEfforts[0].Segment.Starred = false
	expected.SegmentEfforts[0].AverageCadence = 78.7
	expected.SegmentEfforts[0].AveragePower = 153
	expected.SegmentEfforts[0].AverageHeartrate = 107.9
	expected.SegmentEfforts[0].MaximumHeartrate = 119
	expected.SegmentEfforts[0].Activity.Id = 103221154
	expected.SegmentEfforts[0].Athlete.Id = 227615
	expected.SegmentEfforts[0].KOMRank = 0
	expected.SegmentEfforts[0].PRRank = 0
	expected.SegmentEfforts[0].ElapsedTime = 112
	expected.SegmentEfforts[0].MovingTime = 112
	expected.SegmentEfforts[0].StartDate, _ = time.Parse(timeFormat, "2010-08-15T18:05:56Z")
	expected.SegmentEfforts[0].StartDateLocal, _ = time.Parse(timeFormat, "2010-08-15T11:05:56Z")
	expected.SegmentEfforts[0].Distance = 812.6
	expected.SegmentEfforts[0].StartIndex = 83
	expected.SegmentEfforts[0].EndIndex = 194
	expected.SegmentEfforts[0].Hidden = false

	expected.SplitsMetric = []*Split{}
	expected.SplitsStandard = []*Split{}
	expected.BestEfforts = []*BestEffort{}

	if len(activity.SegmentEfforts) == 0 {
		t.Fatal("no segment efforts!?!?!")
	}

	if !reflect.DeepEqual(activity.SegmentEfforts[0], expected.SegmentEfforts[0]) {
		t.Errorf("should match\n%v\n%v", activity.SegmentEfforts[0], expected.SegmentEfforts[0])
	}

	// not comparing these here
	activity.SegmentEfforts = expected.SegmentEfforts
	activity.SplitsMetric = expected.SplitsMetric
	activity.SplitsStandard = expected.SplitsStandard
	activity.BestEfforts = expected.BestEfforts

	if !reflect.DeepEqual(activity, expected) {
		t.Errorf("should match\n%v\n%v", activity, expected)
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

	if !reflect.DeepEqual(activity.SplitsMetric[0], split) {
		t.Errorf("should match\n%v\n%v", activity.SplitsMetric[0], split)
	}

	split = &Split{
		Distance:            1612.0,
		ElapsedTime:         509,
		ElevationDifference: 12.6,
		MovingTime:          454,
		Split:               1,
	}

	if !reflect.DeepEqual(activity.SplitsStandard[0], split) {
		t.Errorf("should match\n%v\n%v", activity.SplitsStandard[0], split)
	}

	bestEffort := &BestEffort{}

	bestEffort.Id = 474685446
	bestEffort.Name = "400m"
	bestEffort.ElapsedTime = 111
	bestEffort.MovingTime = 112

	bestEffort.StartDate, _ = time.Parse(timeFormat, "2013-09-23T00:15:15Z")
	bestEffort.StartDateLocal, _ = time.Parse(timeFormat, "2013-09-22T17:15:15Z")

	bestEffort.Distance = 400
	bestEffort.StartIndex = 1
	bestEffort.EndIndex = 109

	bestEffort.Activity.Id = 103359122
	bestEffort.Athlete.Id = 227615

	if !reflect.DeepEqual(activity.BestEfforts[0], bestEffort) {
		t.Errorf("should match\n%v\n%v", activity.BestEfforts[0], bestEffort)
	}

	// hidden efforts
	client = newCassetteClient(testToken, "activity_get_ride_all_efforts")
	activity, err = NewActivitiesService(client).Get(103221154).IncludeAllEfforts().Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if activity.SegmentEfforts[0].Hidden == false {
		t.Errorf("effort should be hidden")
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

func TestActivitiesDelete(t *testing.T) {
	// from here on out just check the request parameters
	s := NewActivitiesService(newStoreRequestClient())

	// path
	s.Delete(123).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/activities/123" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.Method != "DELETE" {
		t.Errorf("request method incorrect, got %v", transport.request.Method)
	}
}

func TestActivitiesCreate(t *testing.T) {
	client := newCassetteClient(testToken, "activity_post")
	activity, err := NewActivitiesService(client).Create("name", ActivityTypes.Ride, time.Now(), 100).Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if activity.StartDate.IsZero() || activity.StartDateLocal.IsZero() {
		t.Error("dates not parsed")
	}

	// from here on out just check the request parameters
	s := NewActivitiesService(newStoreRequestClient())
	start := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)

	// path
	s.Create("name", ActivityTypes.Ride, start, 100).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/activities" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.Method != "POST" {
		t.Errorf("request method incorrect, got %v", transport.request.Method)
	}

	body, _ := ioutil.ReadAll(transport.request.Body)
	if string(body) != "elapsed_time=100&name=name&start_date_local=2009-11-10T23%3A00%3A00Z&type=Ride" {
		t.Errorf("request body incorrect, got %s", body)
	}

	// parameters1
	s.Create("name", ActivityTypes.Ride, start, 100).Distance(100.0).Do()

	body, _ = ioutil.ReadAll(transport.request.Body)
	if string(body) != "distance=100&elapsed_time=100&name=name&start_date_local=2009-11-10T23%3A00%3A00Z&type=Ride" {
		t.Errorf("request body incorrect, got %s", body)
	}

	// parameters2
	s.Create("name", ActivityTypes.Ride, start, 100).Description("description").Do()

	body, _ = ioutil.ReadAll(transport.request.Body)
	if string(body) != "description=description&elapsed_time=100&name=name&start_date_local=2009-11-10T23%3A00%3A00Z&type=Ride" {
		t.Errorf("request body incorrect, got %s", body)
	}
}

func TestActivitiesUpdate(t *testing.T) {
	client := newCassetteClient(testToken, "activity_put")
	activity, err := NewActivitiesService(client).Update(141818870).Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if activity.StartDate.IsZero() || activity.StartDateLocal.IsZero() {
		t.Error("dates not parsed")
	}

	// from here on out just check the request parameters
	s := NewActivitiesService(newStoreRequestClient())

	// path
	s.Update(123).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/activities/123" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.Method != "PUT" {
		t.Errorf("request method incorrect, got %v", transport.request.Method)
	}

	// parameters1
	s.Update(123).Name("name").Description("description").Do()

	if transport.request.URL.RawQuery != "description=description&name=name" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}

	// parameters2
	s.Update(123).Type(ActivityTypes.AlpineSki).Gear("g123").Do()

	if transport.request.URL.RawQuery != "gear_id=g123&type=AlpineSki" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}

	// parameters3
	s.Update(123).Private(false).Commute(true).Trainer(false).Do()

	if transport.request.URL.RawQuery != "commute=true&private=0&trainer=false" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}

	// parameters4
	s.Update(123).Private(true).Do()

	if transport.request.URL.RawQuery != "private=1" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestActivitiesListPhotos(t *testing.T) {
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
	expected.UploadedAt, _ = time.Parse(timeFormat, "2014-01-02T04:02:28Z")
	expected.CreatedAt, _ = time.Parse(timeFormat, "2014-01-02T04:04:00Z")

	if !reflect.DeepEqual(photos[0], expected) {
		t.Errorf("should match\n%v\n%v", photos[0], expected)
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

	if len(zones[0].Buckets) == 0 {
		t.Fatal("Buckets not parsed")
	}

	expected := &ZoneBucket{Max: 143, Min: 108, Time: 910}

	if !reflect.DeepEqual(zones[0].Buckets[1], expected) {
		t.Errorf("should match\n%v\n%v", zones[0].Buckets[1], expected)
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

func TestActivitiesListLaps(t *testing.T) {
	client := newCassetteClient(testToken, "activity_list_laps")
	laps, err := NewActivitiesService(client).ListLaps(103373338).Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if len(laps) == 0 {
		t.Fatal("laps not parsed")
	}

	expected := &LapEffortSummary{}

	expected.Id = 429913783
	expected.Activity.Id = 103373338
	expected.Athlete.Id = 227615

	expected.Name = "Lap 1"
	expected.ElapsedTime = 6219
	expected.MovingTime = 5118

	expected.StartDate, _ = time.Parse(timeFormat, "2013-09-28T17:27:59Z")
	expected.StartDateLocal, _ = time.Parse(timeFormat, "2013-09-28T10:27:59Z")

	expected.Distance = 25109.4
	expected.StartIndex = 0
	expected.EndIndex = 5087

	expected.TotalElevationGain = 90
	expected.AverageSpeed = 4
	expected.MaximunSpeed = 8.9
	expected.AveragePower = 70
	expected.LapIndex = 1

	if !reflect.DeepEqual(laps[0], expected) {
		t.Errorf("should match\n%v\n%v", laps[0], expected)
	}

	// from here on out just check the request parameters
	s := NewActivitiesService(newStoreRequestClient())

	// path
	s.ListLaps(321).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/activities/321/laps" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestActivitiesBadJSON(t *testing.T) {
	var err error
	s := NewActivitiesService(NewStubResponseClient("bad json"))

	_, err = s.Get(123).Do()
	if err == nil {
		t.Error("should return a bad json error")
	}

	_, err = s.Create("name", ActivityTypes.Ride, time.Now(), 123).Do()
	if err == nil {
		t.Error("should return a bad json error")
	}

	_, err = s.Update(123).Do()
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

	_, err = s.ListLaps(123).Do()
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
	if id := ActivityTypes.VirtualRide.Id(); id != 17 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.VirtualRide.String(); s != "VirtualRide" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	if id := ActivityTypes.EBikeRide.Id(); id != 18 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.EBikeRide.String(); s != "EBikeRide" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	if id := ActivityTypes.WaterSport.Id(); id != 20 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.WaterSport.String(); s != "WaterSport" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	if id := ActivityTypes.Canoeing.Id(); id != 21 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.Canoeing.String(); s != "Canoeing" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	if id := ActivityTypes.Kayaking.Id(); id != 22 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.Kayaking.String(); s != "Kayaking" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	if id := ActivityTypes.Rowing.Id(); id != 23 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.Rowing.String(); s != "Rowing" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	if id := ActivityTypes.StandUpPaddling.Id(); id != 24 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.StandUpPaddling.String(); s != "StandUpPaddling" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	if id := ActivityTypes.Surfing.Id(); id != 25 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.Surfing.String(); s != "Surfing" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	if id := ActivityTypes.Crossfit.Id(); id != 26 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.Crossfit.String(); s != "Crossfit" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	if id := ActivityTypes.Elliptical.Id(); id != 27 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.Elliptical.String(); s != "Elliptical" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	if id := ActivityTypes.RockClimbing.Id(); id != 28 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.RockClimbing.String(); s != "RockClimbing" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	if id := ActivityTypes.StairStepper.Id(); id != 29 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.StairStepper.String(); s != "StairStepper" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	if id := ActivityTypes.WeightTraining.Id(); id != 30 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.WeightTraining.String(); s != "WeightTraining" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	if id := ActivityTypes.Yoga.Id(); id != 31 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.Yoga.String(); s != "Yoga" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	if id := ActivityTypes.WinterSport.Id(); id != 40 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.WinterSport.String(); s != "WinterSport" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	if id := ActivityTypes.CrossCountrySkiing.Id(); id != 41 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ActivityTypes.CrossCountrySkiing.String(); s != "CrossCountrySkiing" {
		t.Errorf("activity type string incorrect, got %v", s)
	}

	// other
	ty := ActivityType(100)
	if id := ty.Id(); id != 0 {
		t.Errorf("activity type id incorrect, got %v", id)
	}

	if s := ty.String(); s != "Activity" {
		t.Errorf("activity type string incorrect, got %v", s)
	}
}
