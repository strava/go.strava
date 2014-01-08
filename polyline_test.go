package strava

import (
	"testing"
)

func TestPolylineDecode(t *testing.T) {
	var encoded Polyline = "_p~iF~ps|U_ulLnnqC_mqNvxq`@"
	expected := [][2]float64{{38.5, -120.2}, {40.7, -120.95}, {43.252, -126.453}}

	latlng := encoded.Decode()
	if len(latlng) != len(expected) {
		t.Fatal("Polyline, decode incorrect number of results")
	}

	for i, v := range latlng {
		if v[0] != expected[i][0] || v[1] != expected[i][1] {
			t.Errorf("Polyline, decode error on element %d, expected %v, got %v", i, expected[i], v)
		}
	}
}
