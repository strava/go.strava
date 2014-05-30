package strava

import (
	"reflect"
	"testing"
)

func TestGearGet(t *testing.T) {
	// bike
	client := newCassetteClient(testToken, "gear_get_bike")
	gear, err := NewGearService(client).Get("b77076").Do()

	expected := &GearDetailed{}
	expected.Id = "b77076"
	expected.Primary = false
	expected.Name = "burrito burner"
	expected.Distance = 536292.3
	expected.BrandName = "Schwinn"
	expected.ModelName = ""
	expected.FrameType = FrameTypes.Road
	expected.Description = ""

	if err != nil {
		t.Fatalf("Gear service error: %v", err)
	}

	if !reflect.DeepEqual(gear, expected) {
		t.Errorf("should match\n%v\n%v", gear, expected)
	}

	// shoe
	client = newCassetteClient(testToken, "gear_get_shoe")
	gear, err = NewGearService(client).Get("g5697").Do()

	expected = &GearDetailed{}
	expected.Id = "g5697"
	expected.Primary = true
	expected.Name = "ASICS Kayano"
	expected.Distance = 17224.6
	expected.BrandName = "ASICS"
	expected.ModelName = "Kayano"
	expected.Description = ""

	if err != nil {
		t.Fatalf("Gear service error: %v", err)
	}

	if !reflect.DeepEqual(gear, expected) {
		t.Errorf("should match\n%v\n%v", gear, expected)
	}

	// from here on out just check the request parameters
	s := NewGearService(newStoreRequestClient())

	// path
	s.Get("b123").Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/gear/b123" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}

}

func TestGearBadJSON(t *testing.T) {
	var err error
	s := NewGearService(NewStubResponseClient("bad json"))

	_, err = s.Get("b123").Do()
	if err == nil {
		t.Error("should return a bad json error")
	}
}

func TestFrameType(t *testing.T) {
	if FrameTypes.MountainBike.Id() != int(FrameTypes.MountainBike) {
		t.Error("frame type id should just be int val")
	}

	if s := FrameTypes.MountainBike.String(); s != "Mountain Bike" {
		t.Errorf("frame name incorrect, got %v", s)
	}

	if s := FrameTypes.Cyclocross.String(); s != "Cyclocross" {
		t.Errorf("frame name incorrect, got %v", s)
	}

	if s := FrameTypes.Road.String(); s != "Road" {
		t.Errorf("frame name incorrect, got %v", s)
	}

	if s := FrameTypes.TimeTrial.String(); s != "Time Trial" {
		t.Errorf("frame name incorrect, got %v", s)
	}

	if s := FrameType(0).String(); s != "Bike Frame" {
		t.Errorf("frame name incorrect, got %v", s)
	}
}
