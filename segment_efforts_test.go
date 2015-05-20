package strava

import (
	"reflect"
	"testing"
	"time"
)

func TestSegmentEffortsGet(t *testing.T) {
	client := newCassetteClient(testToken, "segment_effort_get")
	effort, err := NewSegmentEffortsService(client).Get(801006623).Do()

	expected := &SegmentEffortDetailed{}
	expected.Id = 801006623
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
	expected.Segment.Starred = false

	expected.Activity.Id = 46320211
	expected.Athlete.Id = 123529

	expected.AverageCadence = 0
	expected.AveragePower = 460.8
	expected.AverageHeartrate = 190.5
	expected.MaximumHeartrate = 199
	expected.KOMRank = 1
	expected.PRRank = 1
	expected.ElapsedTime = 360
	expected.MovingTime = 360
	expected.StartDate, _ = time.Parse(timeFormat, "2013-03-29T13:49:35Z")
	expected.StartDateLocal, _ = time.Parse(timeFormat, "2013-03-29T06:49:35Z")

	expected.Distance = 2659.89
	expected.StartIndex = 1992
	expected.EndIndex = 2310

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if !reflect.DeepEqual(effort, expected) {
		t.Errorf("should match\n%v\n%v", effort, expected)
	}

	// from here on out just check the request parameters
	s := NewSegmentEffortsService(newStoreRequestClient())

	// path
	s.Get(321).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/segment_efforts/321" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestSegmentEffortsBadJSON(t *testing.T) {
	var err error
	s := NewSegmentEffortsService(NewStubResponseClient("bad json"))

	_, err = s.Get(123).Do()
	if err == nil {
		t.Error("should return a bad json error")
	}
}
