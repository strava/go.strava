package strava

import (
	"testing"
)

func TestActivityStreamsGet(t *testing.T) {
	types := []StreamType{
		StreamTypes.Time,
		StreamTypes.Location,
		StreamTypes.Distance,
		StreamTypes.Elevation,
		StreamTypes.Speed,
		StreamTypes.HeartRate,
		StreamTypes.Cadence,
		StreamTypes.Power,
		StreamTypes.Temperature,
		StreamTypes.Moving,
		StreamTypes.Grade,
	}

	client := newCassetteClient(testToken, "activity_stream")
	streams, err := NewActivityStreamsService(client).
		Get(103221154, types).Resolution("medium").Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	// time
	if l := len(streams.Time.Data); l != 999 {
		t.Errorf("data not parsed: %d", l)
	}

	if streams.Time.SeriesType != "distance" ||
		streams.Time.OriginalSize != 2829 ||
		streams.Time.Resolution != "medium" {

		t.Error("meta not parsed")
	}

	if streams.Time.RawData[1] != nil {
		t.Errorf("value should be nil but got %v", streams.Time.RawData[0])
	}

	di := []int{0, 0, 8, 10, 13, 16}
	for i := 2; i < len(di); i++ {
		if streams.Time.Data[i] != di[i] {
			t.Errorf("values incorrect: %d", i)
		}

		if *streams.Time.RawData[i] != di[i] {
			t.Errorf("values incorrect: %d", i)
		}
	}

	// location
	if l := len(streams.Location.Data); l != 999 {
		t.Errorf("data not parsed: %d", l)
	}

	if streams.Location.SeriesType != "distance" ||
		streams.Location.OriginalSize != 2829 ||
		streams.Location.Resolution != "medium" {

		t.Error("meta not parsed")
	}

	dl := [][2]float64{[2]float64{0, 0}, [2]float64{38.546876, -121.817203}, [2]float64{38.546881, -121.817439}}
	for i := 0; i < len(dl); i++ {
		if streams.Location.Data[i] != dl[i] {
			t.Errorf("values incorrect: %d", i)
		}
	}

	// distance
	if l := len(streams.Distance.Data); l != 999 {
		t.Errorf("data not parsed: %d", l)
	}

	if streams.Distance.SeriesType != "distance" ||
		streams.Distance.OriginalSize != 2829 ||
		streams.Distance.Resolution != "medium" {

		t.Error("meta not parsed")
	}

	if streams.Distance.RawData[1] != nil {
		t.Errorf("value should be nil but got %v", streams.Distance.RawData[1])
	}

	df := []float64{0.6, 0, 64.2, 80.0, 102.9, 126.1, 149.5, 166.9, 190.5, 205.5}
	for i := 2; i < len(df); i++ {
		if streams.Distance.Data[i] != df[i] {
			t.Errorf("values incorrect: %d", i)
		}

		if *streams.Distance.RawData[i] != df[i] {
			t.Errorf("values incorrect: %d", i)
		}
	}

	// elevation
	if l := len(streams.Elevation.Data); l != 999 {
		t.Errorf("data not parsed: %d", l)
	}

	if streams.Elevation.SeriesType != "distance" ||
		streams.Elevation.OriginalSize != 2829 ||
		streams.Elevation.Resolution != "medium" {

		t.Error("meta not parsed")
	}

	if streams.Elevation.RawData[1] != nil {
		t.Errorf("value should be nil but got %v", streams.Elevation.RawData[1])
	}

	df = []float64{23.0, 0, 23.0, 23.0}
	for i := 2; i < len(df); i++ {
		if streams.Elevation.Data[i] != df[i] {
			t.Errorf("values incorrect: %d", i)
		}

		if *streams.Elevation.RawData[i] != df[i] {
			t.Errorf("values incorrect: %d", i)
		}
	}

	// speed
	if l := len(streams.Speed.Data); l != 999 {
		t.Errorf("data not parsed: %d", l)
	}

	if streams.Speed.SeriesType != "distance" ||
		streams.Speed.OriginalSize != 2829 ||
		streams.Speed.Resolution != "medium" {

		t.Error("meta not parsed")
	}

	df = []float64{0.0, 8.6, 8.0, 7.3, 7.7, 7.7, 7.8, 8.2, 8.2, 7.7}
	for i := 0; i < len(df); i++ {
		if streams.Speed.Data[i] != df[i] {
			t.Errorf("values incorrect: %d", i)
		}
	}

	// heart rate
	if l := len(streams.HeartRate.Data); l != 999 {
		t.Errorf("data not parsed: %d", l)
	}

	if streams.HeartRate.SeriesType != "distance" ||
		streams.HeartRate.OriginalSize != 2829 ||
		streams.HeartRate.Resolution != "medium" {

		t.Error("meta not parsed")
	}

	if streams.HeartRate.RawData[1] != nil {
		t.Errorf("value should be nil but got %v", streams.HeartRate.RawData[1])
	}

	di = []int{111, 0, 109, 108, 104, 103, 103, 103, 102, 101, 102, 102, 95, 91, 91, 93}
	for i := 2; i < len(di); i++ {
		if streams.HeartRate.Data[i] != di[i] {
			t.Errorf("values incorrect: %d", i)
		}

		if *streams.HeartRate.RawData[i] != di[i] {
			t.Errorf("values incorrect: %d", i)
		}
	}

	// cadence
	if l := len(streams.Cadence.Data); l != 999 {
		t.Errorf("data not parsed: %d", l)
	}

	if streams.Cadence.SeriesType != "distance" ||
		streams.Cadence.OriginalSize != 2829 ||
		streams.Cadence.Resolution != "medium" {

		t.Error("meta not parsed")
	}

	if streams.Cadence.RawData[0] != nil {
		t.Errorf("value should be nil but got %v", streams.Cadence.RawData[0])
	}

	di = []int{0, 81, 80, 78, 79, 80, 81, 80, 80, 81, 78, 78, 0, 0, 0, 60, 74, 52}
	for i := 1; i < len(di); i++ {
		if streams.Cadence.Data[i] != di[i] {
			t.Errorf("values incorrect: %d", i)
		}
	}

	// power
	if l := len(streams.Power.Data); l != 999 {
		t.Errorf("data not parsed: %d", l)
	}

	if streams.Power.SeriesType != "distance" ||
		streams.Power.OriginalSize != 2829 ||
		streams.Power.Resolution != "medium" {

		t.Error("meta not parsed")
	}

	di = []int{163, 168, 146, 134, 155, 118, 119, 112, 124, 125, 88, 0}
	for i := 0; i < len(di); i++ {
		if streams.Power.Data[i] != di[i] {
			t.Errorf("values incorrect: %d", i)
		}
	}

	// temperature
	if l := len(streams.Temperature.Data); l != 999 {
		t.Errorf("data not parsed: %d", l)
	}

	if streams.Temperature.SeriesType != "distance" ||
		streams.Temperature.OriginalSize != 2829 ||
		streams.Temperature.Resolution != "medium" {

		t.Error("meta not parsed")
	}

	di = []int{26, 26, 26, 26, 26}
	for i := 0; i < len(di); i++ {
		if streams.Temperature.Data[i] != di[i] {
			t.Errorf("values incorrect: %d", i)
		}
	}

	// moving
	if l := len(streams.Moving.Data); l != 999 {
		t.Errorf("data not parsed: %d", l)
	}

	if streams.Moving.SeriesType != "distance" ||
		streams.Moving.OriginalSize != 2829 ||
		streams.Moving.Resolution != "medium" {

		t.Error("meta not parsed")
	}

	db := []bool{false, true, true}
	for i := 0; i < len(db); i++ {
		if streams.Moving.Data[i] != db[i] {
			t.Errorf("values incorrect: %d", i)
		}
	}

	// grade
	if l := len(streams.Grade.Data); l != 999 {
		t.Errorf("data not parsed: %d", l)
	}

	if streams.Grade.SeriesType != "distance" ||
		streams.Grade.OriginalSize != 2829 ||
		streams.Grade.Resolution != "medium" {

		t.Error("meta not parsed")
	}

	df = []float64{0.7, 2.6, 1.3, 1.3, 1.3, 0.7, 0.7, 0.0, 0.7, 0.0}
	for i := 60; i < len(df); i++ {
		if streams.Grade.Data[i] != df[i] {
			t.Errorf("values incorrect: %d", i)
		}
	}

	// from here on out just check the request parameters
	s := NewActivityStreamsService(newStoreRequestClient())

	// path
	s.Get(45255, []StreamType{StreamTypes.Location, StreamTypes.HeartRate}).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/activities/45255/streams/latlng,heartrate" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}

	// no types
	_, err = s.Get(45255, []StreamType{}).Do()
	if err == nil {
		t.Error("should return error if not types provided")
	}

	// parameters
	s.Get(123, []StreamType{StreamTypes.Location, StreamTypes.Power}).Resolution("medium").SeriesType("distance").Do()

	transport = s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/activities/123/streams/latlng,watts" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "resolution=medium&series_type=distance" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestSegmentStreamsGet(t *testing.T) {
	// from here on out just check the request parameters
	s := NewSegmentStreamsService(newStoreRequestClient())

	// path
	s.Get(229781, []StreamType{StreamTypes.Location, StreamTypes.Time}).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/segments/229781/streams/latlng,time" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}

	// parameters
	s.Get(123, []StreamType{StreamTypes.Location}).Resolution("low").SeriesType("distance").Do()

	transport = s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/segments/123/streams/latlng" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "resolution=low&series_type=distance" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestSegmentEffortStreamsGet(t *testing.T) {
	// from here on out just check the request parameters
	s := NewSegmentEffortStreamsService(newStoreRequestClient())

	// path
	s.Get(123123, []StreamType{StreamTypes.Location, StreamTypes.Cadence}).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/segment_efforts/123123/streams/latlng,cadence" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}

	// parameters
	s.Get(123, []StreamType{StreamTypes.Distance}).Resolution("low").SeriesType("time").Do()

	transport = s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/segment_efforts/123/streams/distance" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "resolution=low&series_type=time" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestStreamsBadSource(t *testing.T) {
	s := NewSegmentEffortStreamsService(newStoreRequestClient())
	s.parentType = 123123

	_, err := s.Get(123, []StreamType{StreamTypes.Location, StreamTypes.Cadence}).Do()

	if err == nil {
		t.Error("should have returned error")
	}
}
